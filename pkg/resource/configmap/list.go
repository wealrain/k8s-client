package configmap

import (
	"context"
	"k8s-client/pkg/errors"
	"k8s-client/pkg/resource/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Configmap struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"` // 命名空间
	Age       int64    `json:"age"`       // unix timestamp
	Keys      []string `json:"keys"`
}

func GetConfigmapList(client k8sClient.Interface, namespace string) ([]Configmap, error) {

	channel := common.GetConfigMapListChannelWithOptions(client, namespace, metav1.ListOptions{})

	configmaps := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var configmapList []Configmap

	for _, configmap := range configmaps.Items {
		configmapInfo := toConfigmap(&configmap)
		configmapList = append(configmapList, configmapInfo)
	}

	return configmapList, nil
}

func toConfigmap(configmap *corev1.ConfigMap) Configmap {
	return Configmap{
		Name:      configmap.Name,
		Namespace: configmap.Namespace,
		Age:       configmap.CreationTimestamp.Unix(),
		Keys:      getKeys(configmap.Data),
	}
}

func getKeys(data map[string]string) []string {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	return keys
}

func GetConfigmapData(client k8sClient.Interface, namespace, name string) map[string]string {
	configmap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	errors.HandleError(err)
	return configmap.Data
}

func ToInterface(configMapList []Configmap) []interface{} {
	var result []interface{}
	for _, configmap := range configMapList {
		result = append(result, configmap)
	}
	return result
}
