package pod

import (
	"k8s-client/pkg/resource/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Pod struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Age       int64    `json:"age"` // unix timestamp
	Restarts  int32    `json:"restarts"`
	Images    []string `json:"images"`
	Node      string   `json:"node"` // node name
	Status    string   `json:"status"`
}

func GetPodList(client k8sClient.Interface, namespace string) ([]Pod, error) {

	channel := common.GetPodListChannelWithOptions(client, namespace, metav1.ListOptions{})

	pods := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var podList []Pod

	for _, pod := range pods.Items {
		podInfo := toPod(&pod)
		podList = append(podList, podInfo)
	}

	return podList, nil
}

func toPod(pod *corev1.Pod) Pod {
	return Pod{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		Age:       pod.CreationTimestamp.Unix(),
		Restarts:  getRestartCount(pod),
		Images:    getImages(pod),
		Node:      pod.Spec.NodeName,
		Status:    getPodStatus(pod),
	}
}

func ToInterface(pods []Pod) []interface{} {
	var result []interface{}
	for _, pod := range pods {
		result = append(result, pod)
	}
	return result
}
