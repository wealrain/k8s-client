package pod

import (
	"context"
	"fmt"
	"k8s-client/pkg/resource/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type PodDetail struct {
	Name             string              `json:"name"`
	Namespace        string              `json:"namespace"`
	Labels           []string            `json:"labels"`
	Created          string              `json:"created"`
	ControlledBy     ControllerBy        `json:"controlledBy"`
	Status           string              `json:"status"` // Running, Pending, Succeeded, Failed, Unknown
	Node             string              `json:"node"`
	PodIP            string              `json:"podIP"`
	ServiceAccount   string              `json:"serviceAccount"`
	QoSClass         string              `json:"qosClass"` // Guaranteed, Burstable, BestEffort
	Tolerations      []corev1.Toleration `json:"tolerations"`
	Affinity         string              `json:"affinity"` // TODO: add affinity struct
	InitContainers   []Container         `json:"initContainers"`
	Containers       []Container         `json:"containers"`
	ImagePullSecrets []string            `json:"imagePullSecrets"`
	// Metrics          []Metric          `json:"metrics"`
	// Volumes []string `json:"volumes"` // TODO: add volume struct
	// Events           []Event           `json:"events"`
	// Annotations      map[string]string `json:"annotations"`
}

type ControllerBy struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type Container struct {
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	Ports    []string `json:"ports"`
	Commands []string `json:"commands"`
	Args     []string `json:"args"`
	Status   string   `json:"status"`
	// VolumeMounts []VolumeMount `json:"volumeMounts"`
	// LivenessProbe  *corev1.Probe                `json:"livenessProbe"`
	// ReadinessProbe *corev1.Probe                `json:"readinessProbe"`
	// StartupProbe   *corev1.Probe                `json:"startupProbe"`
	// Env            []EnvVar                     `json:"env"`
}

type EnvVar struct {
	Name      string               `json:"name"`
	Value     string               `json:"value"`
	ValueFrom *corev1.EnvVarSource `json:"valueFrom"`
}

type VolumeMount struct {
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readOnly"`
	MountPath string `json:"mountPath"`
	SubPath   string `json:"subPath"`
}

type Metric struct {
	Name      string `json:"name"`
	Value     uint64 `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

func GetPodDetail(client k8sClient.Interface, namespace, name string) (*PodDetail, error) {
	pod, err := client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// controllerBy
	controller := getPodController(client, namespace, pod)

	return toPodDetail(pod, controller), nil
}

func toPodDetail(pod *corev1.Pod, controller ControllerBy) *PodDetail {
	return &PodDetail{
		Name:             pod.Name,
		Namespace:        pod.Namespace,
		Created:          pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		Labels:           getPodLabels(pod.Labels),
		ControlledBy:     controller,
		Status:           getPodStatus(pod),
		Node:             pod.Spec.NodeName,
		PodIP:            pod.Status.PodIP,
		ServiceAccount:   pod.Spec.ServiceAccountName,
		QoSClass:         string(pod.Status.QOSClass),
		Tolerations:      pod.Spec.Tolerations,
		Affinity:         common.ToYaml(pod.Spec.Affinity),
		InitContainers:   getContainers(pod.Spec.InitContainers, pod),
		Containers:       getContainers(pod.Spec.Containers, pod),
		ImagePullSecrets: getImagePullSecrets(pod),
		// Volumes:        getVolumes(pod.Spec.Volumes),
		// Events:         getEvents(pod),
		// Metrics:        getMetrics(pod),
		// Annotations:    pod.Annotations,

	}
}

func getPodLabels(labels map[string]string) []string {
	var result []string
	for k, v := range labels {
		result = append(result, fmt.Sprintf("%s=%s", k, v))
	}

	return result
}

func getPodController(client k8sClient.Interface, namespace string, pod *corev1.Pod) ControllerBy {
	ownerRef := metav1.GetControllerOf(pod)
	if ownerRef == nil {
		return ControllerBy{}
	}

	return ControllerBy{
		Kind: ownerRef.Kind,
		Name: ownerRef.Name,
	}
}

func getContainers(containers []corev1.Container, pod *corev1.Pod) []Container {
	var result []Container
	for _, container := range containers {
		result = append(result, Container{
			Name:     container.Name,
			Image:    container.Image,
			Commands: container.Command,
			Args:     container.Args,
			Ports:    getContainerPorts(container.Ports),
			Status:   getContainerStatus(&container, pod),
		})
	}

	return result
}

func getContainerStatus(container *corev1.Container, pod *corev1.Pod) string {
	for _, status := range pod.Status.ContainerStatuses {
		if status.Name == container.Name {
			// 获取容器状态
			if status.State.Running != nil {
				return "Running"
			}
			if status.State.Waiting != nil {
				return "Waiting"
			}
			if status.State.Terminated != nil {
				return "Terminated"
			}
		}
	}

	return "Unknown"
}

func getContainerPorts(ports []corev1.ContainerPort) []string {
	var result []string
	for _, port := range ports {
		port := fmt.Sprintf("%d/%s", port.ContainerPort, port.Protocol)
		result = append(result, port)
	}

	return result
}

func getImagePullSecrets(pod *corev1.Pod) []string {
	var imagePullSecrets []string

	for _, secret := range pod.Spec.ImagePullSecrets {
		imagePullSecrets = append(imagePullSecrets, secret.Name)
	}

	return imagePullSecrets
}
