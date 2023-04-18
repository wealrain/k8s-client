package deployment

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type DeploymentDetail struct {
	Name         string   `json:"name"`
	Namespace    string   `json:"namespace"`
	Created      string   `json:"created"`
	Annotations  []string `json:"annotations"`
	Replicas     int32    `json:"replicas"`
	Selector     []string `json:"selector"`
	StrategyType string   `json:"strategyType"`
	// Conditions   []string  `json:"conditions"`
	Pods []PodInfo `json:"pods"`
	// Affinitys	[]string `json:"affinitys"`
	// Event []string `json:"event"`
}

type PodInfo struct {
	Name   string `json:"name"`
	Ready  string `json:"ready"`
	Status string `json:"status"`
}

func GetDeploymentDetail(client k8sClient.Interface, namespace, name string) (*DeploymentDetail, error) {
	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	podInfo := getPods(client, namespace, deployment.Spec.Selector)
	return toDeploymentDetail(deployment, podInfo), nil
}

func toDeploymentDetail(deployment *appsv1.Deployment, podInfo []PodInfo) *DeploymentDetail {
	return &DeploymentDetail{
		Name:         deployment.Name,
		Namespace:    deployment.Namespace,
		Created:      deployment.CreationTimestamp.Format("2006-01-02 15:04:05"),
		Annotations:  getAnnotations(deployment.Annotations),
		Replicas:     deployment.Status.Replicas,
		Selector:     getSelector(deployment.Spec.Selector.MatchLabels),
		StrategyType: string(deployment.Spec.Strategy.Type),
		// Conditions:   getConditions(deployment.Status.Conditions),
		Pods: podInfo,
		// Affinitys: getAffinitys(deployment.Spec.Template.Spec.Affinity),
		// Event: getEvent(deployment),
	}
}

func getAnnotations(annotations map[string]string) []string {
	var result []string
	for k, v := range annotations {
		result = append(result, k+"="+v)
	}

	return result
}

func getSelector(labels map[string]string) []string {
	var result []string
	for k, v := range labels {
		result = append(result, k+"="+v)
	}

	return result
}

func getPods(client k8sClient.Interface, namespace string, selector *metav1.LabelSelector) []PodInfo {
	labelSelector := metav1.FormatLabelSelector(selector)
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})

	if err != nil {
		return nil
	}

	var result []PodInfo
	for _, pod := range podList.Items {
		result = append(result, PodInfo{
			Name:   pod.Name,
			Ready:  getPodReady(pod.Status),
			Status: string(pod.Status.Phase),
		})
	}

	return result
}

func getPodReady(status corev1.PodStatus) string {
	ready := 0
	for _, containerStatus := range status.ContainerStatuses {
		if containerStatus.Ready {
			ready++
		}
	}

	return fmt.Sprintf("%d/%d", ready, len(status.ContainerStatuses))
}
