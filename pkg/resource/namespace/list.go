package namespace

import (
	"k8s-client/pkg/resource/common"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Namespace struct {
	Name  string `json:"name"`
	Phase string `json:"phase"`
}

func GetNamespaceList(client k8sClient.Interface) ([]Namespace, error) {

	channel := common.GetNamespaceListChannelWithOptions(client, metav1.ListOptions{})

	namespaces := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var namespaceList []Namespace

	for _, namespace := range namespaces.Items {
		namespaceInfo := toNamespace(&namespace)
		namespaceList = append(namespaceList, namespaceInfo)
	}

	return namespaceList, nil
}

func toNamespace(namespace *corev1.Namespace) Namespace {

	return Namespace{
		Name:  namespace.Name,
		Phase: strings.ToLower(string(namespace.Status.Phase)),
	}
}

func ToInterface(namespaceList []Namespace) []interface{} {
	var result []interface{}
	for _, namespace := range namespaceList {
		result = append(result, namespace)
	}
	return result
}
