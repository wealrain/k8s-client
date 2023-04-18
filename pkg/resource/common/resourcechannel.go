package common

/// 异步获取资源列表

import (
	"context"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
)

type PodListChannel struct {
	List  chan *corev1.PodList
	Error chan error
}

func GetPodListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) PodListChannel {
	channel := PodListChannel{
		List:  make(chan *corev1.PodList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.CoreV1().Pods(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type DeploymentListChannel struct {
	List  chan *appv1.DeploymentList
	Error chan error
}

func GetDeploymentListChannel(client client.Interface, namespace string) DeploymentListChannel {
	return GetDeploymentListChannelWithOptions(client, namespace, ListEverything)
}

func GetDeploymentListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) DeploymentListChannel {
	channel := DeploymentListChannel{
		List:  make(chan *appv1.DeploymentList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.AppsV1().Deployments(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type StatefulSetListChannel struct {
	List  chan *appv1.StatefulSetList
	Error chan error
}

func GetStatefulSetListChannel(client client.Interface, namespace string) StatefulSetListChannel {
	return GetStatefulSetListChannelWithOptions(client, namespace, ListEverything)
}

func GetStatefulSetListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) StatefulSetListChannel {
	channel := StatefulSetListChannel{
		List:  make(chan *appv1.StatefulSetList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.AppsV1().StatefulSets(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type ReplicaSetListChannel struct {
	List  chan *appv1.ReplicaSetList
	Error chan error
}

func GetReplicaSetListChannel(client client.Interface, namespace string) ReplicaSetListChannel {
	return GetReplicaSetListChannelWithOptions(client, namespace, ListEverything)
}

func GetReplicaSetListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) ReplicaSetListChannel {
	channel := ReplicaSetListChannel{
		List:  make(chan *appv1.ReplicaSetList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.AppsV1().ReplicaSets(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type ServiceListChannel struct {
	List  chan *corev1.ServiceList
	Error chan error
}

func GetServiceListChannel(client client.Interface, namespace string) ServiceListChannel {
	return GetServiceListChannelWithOptions(client, namespace, ListEverything)
}

func GetServiceListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) ServiceListChannel {
	channel := ServiceListChannel{
		List:  make(chan *corev1.ServiceList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.CoreV1().Services(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type IngressListChannel struct {
	List  chan *netv1.IngressList
	Error chan error
}

func GetIngressListChannel(client client.Interface, namespace string) IngressListChannel {
	return GetIngressListChannelWithOptions(client, namespace, ListEverything)
}

func GetIngressListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) IngressListChannel {
	channel := IngressListChannel{
		List:  make(chan *netv1.IngressList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.NetworkingV1().Ingresses(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type EndpointsListChannel struct {
	List  chan *corev1.EndpointsList
	Error chan error
}

func GetEndpointListChannel(client client.Interface, namespace string) EndpointsListChannel {
	return GetEndpointListChannelWithOptions(client, namespace, ListEverything)
}

func GetEndpointListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) EndpointsListChannel {
	channel := EndpointsListChannel{
		List:  make(chan *corev1.EndpointsList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.CoreV1().Endpoints(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err

	}()

	return channel

}

type EventListChannel struct {
	List  chan *corev1.EventList
	Error chan error
}

func GetEventListChannel(client client.Interface, namespace string) EventListChannel {
	return GetEventListChannelWithOptions(client, namespace, ListEverything)
}

func GetEventListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) EventListChannel {
	channel := EventListChannel{
		List:  make(chan *corev1.EventList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.CoreV1().Events(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type ConfigMapListChannel struct {
	List  chan *corev1.ConfigMapList
	Error chan error
}

func GetConfigMapListChannel(client client.Interface, namespace string) ConfigMapListChannel {
	return GetConfigMapListChannelWithOptions(client, namespace, ListEverything)
}

func GetConfigMapListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) ConfigMapListChannel {
	channel := ConfigMapListChannel{
		List:  make(chan *corev1.ConfigMapList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.CoreV1().ConfigMaps(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type SecretListChannel struct {
	List  chan *corev1.SecretList
	Error chan error
}

func GetSecretListChannel(client client.Interface, namespace string) SecretListChannel {
	return GetSecretListChannelWithOptions(client, namespace, ListEverything)
}

func GetSecretListChannelWithOptions(client client.Interface, namespace string, options metav1.ListOptions) SecretListChannel {
	channel := SecretListChannel{
		List:  make(chan *corev1.SecretList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.CoreV1().Secrets(namespace).List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}

type NamespaceListChannel struct {
	List  chan *corev1.NamespaceList
	Error chan error
}

func GetNamespaceListChannel(client client.Interface) NamespaceListChannel {
	return GetNamespaceListChannelWithOptions(client, ListEverything)
}

func GetNamespaceListChannelWithOptions(client client.Interface, options metav1.ListOptions) NamespaceListChannel {
	channel := NamespaceListChannel{
		List:  make(chan *corev1.NamespaceList, 1),
		Error: make(chan error),
	}

	go func() {
		list, err := client.CoreV1().Namespaces().List(context.TODO(), options)

		channel.List <- list
		channel.Error <- err
	}()

	return channel
}
