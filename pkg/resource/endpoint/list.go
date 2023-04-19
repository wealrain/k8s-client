package endpoint

import (
	"k8s-client/pkg/resource/common"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Endpoint struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Endpoints string `json:"endpoints"`
	Age       int64  `json:"age"`
}

func GetEndpointList(client k8sClient.Interface, namespace string) ([]Endpoint, error) {

	channel := common.GetEndpointListChannelWithOptions(client, namespace, metav1.ListOptions{})

	endpoints := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var endpointList []Endpoint

	for _, endpoint := range endpoints.Items {
		endpointInfo := toEndpoint(&endpoint)
		endpointList = append(endpointList, endpointInfo)
	}

	return endpointList, nil
}

func toEndpoint(endpoint *corev1.Endpoints) Endpoint {

	return Endpoint{
		Name:      endpoint.Name,
		Namespace: endpoint.Namespace,
		Endpoints: strings.Join(GetEndpointAddresses(endpoint), ","),
		Age:       endpoint.CreationTimestamp.Time.Unix(),
	}
}

func ToInterface(endpointList []Endpoint) []interface{} {
	var result []interface{}
	for _, endpoint := range endpointList {
		result = append(result, endpoint)
	}
	return result
}
