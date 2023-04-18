package service

import (
	"k8s-client/pkg/resource/common"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Service struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	ClusterIP  string `json:"clusterIP"`
	ExternalIP string `json:"externalIP"`
	Ports      string `json:"ports"`
	Age        int64  `json:"age"`
}

func GetServiceList(client k8sClient.Interface, namespace string) ([]Service, error) {

	channel := common.GetServiceListChannelWithOptions(client, namespace, metav1.ListOptions{})

	services := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var serviceList []Service

	for _, service := range services.Items {
		serviceInfo := toService(&service)
		serviceList = append(serviceList, serviceInfo)
	}

	return serviceList, nil
}

func toService(service *corev1.Service) Service {

	return Service{
		Name:       service.Name,
		Type:       string(service.Spec.Type),
		ClusterIP:  service.Spec.ClusterIP,
		ExternalIP: strings.Join(GetExternalEndpoints(service), ","),
		Ports:      strings.Join(GetPorts(service), ","),
		Age:        service.CreationTimestamp.Time.Unix(),
	}
}

func ToInterface(ServiceList []Service) []interface{} {
	var result []interface{}
	for _, Service := range ServiceList {
		result = append(result, Service)
	}
	return result
}
