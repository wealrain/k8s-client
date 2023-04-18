package ingress

import (
	"k8s-client/pkg/resource/common"
	"strings"

	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Ingress struct {
	Name         string `json:"name"`
	LoadBalancer string `json:"loadBalancer"`
	Age          int64  `json:"age"`
}

func GetIngressList(client k8sClient.Interface, namespace string) ([]Ingress, error) {

	channel := common.GetIngressListChannelWithOptions(client, namespace, metav1.ListOptions{})

	ingresss := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var ingressList []Ingress

	for _, ingress := range ingresss.Items {
		ingressInfo := toIngress(&ingress)
		ingressList = append(ingressList, ingressInfo)
	}

	return ingressList, nil
}

func toIngress(ingress *netv1.Ingress) Ingress {

	return Ingress{
		Name:         ingress.Name,
		LoadBalancer: strings.Join(GetLoadBalancer(ingress), ","),
		Age:          ingress.CreationTimestamp.Time.Unix(),
	}
}

func ToInterface(ingressList []Ingress) []interface{} {
	var result []interface{}
	for _, ingress := range ingressList {
		result = append(result, ingress)
	}
	return result
}
