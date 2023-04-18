package configmap

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type ConfigmapDetail struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Created     string            `json:"created"`
	Annotations []string          `json:"annotations"`
	Data        map[string]string `json:"data"`
}

func GetConfigMapDetail(client k8sClient.Interface, namespace, name string) (*ConfigmapDetail, error) {
	configmap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &ConfigmapDetail{
		Name:        configmap.Name,
		Namespace:   configmap.Namespace,
		Created:     configmap.CreationTimestamp.Format(time.RFC3339),
		Annotations: getAnnotations(configmap.Annotations),
		Data:        configmap.Data,
	}, nil
}

func getAnnotations(annotations map[string]string) []string {
	var result []string
	for k, v := range annotations {
		result = append(result, k+"="+v)
	}
	return result
}
