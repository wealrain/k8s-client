package client

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

type RESTClient interface {
	Delete() *restclient.Request
	Put() *restclient.Request
	Get() *restclient.Request
}

type Verber struct {
	client            RESTClient // core client (pods, services, etc.)
	appsClient        RESTClient // apps client (deployments, statefulsets, etc.)
	batchClient       RESTClient // batch client (jobs, cronjobs, etc.)
	autoscalingClient RESTClient // autoscaling client (horizontalpodautoscalers, etc.)
	storageClient     RESTClient // storage client (storageclasses, etc.)
	rbacClient        RESTClient // rbac client (roles, rolebindings, etc.)
	networkingClient  RESTClient // networking client (ingresses, etc.)
	//apiExtensionsClient RESTClient // apiextensions client (customresourcedefinitions, etc.)
	pluginsClient RESTClient         // plugins client (podsecuritypolicies, etc.)
	config        *restclient.Config // client config
}

func NewVerber(k8sClient kubernetes.Interface, config *restclient.Config) *Verber {
	return &Verber{
		client:            k8sClient.CoreV1().RESTClient(),
		appsClient:        k8sClient.AppsV1().RESTClient(),
		batchClient:       k8sClient.BatchV1().RESTClient(),
		autoscalingClient: k8sClient.AutoscalingV1().RESTClient(),
		storageClient:     k8sClient.StorageV1().RESTClient(),
		rbacClient:        k8sClient.RbacV1().RESTClient(),
		networkingClient:  k8sClient.NetworkingV1().RESTClient(),
		//	apiExtensionsClient: k8sClient.ApiextensionsV1().RESTClient(), todo
		pluginsClient: k8sClient.PolicyV1beta1().RESTClient(),
		config:        config,
	}
}

func (v *Verber) getResourceClient(kind string) (RESTClient, APIMapping, error) {
	resource, ok := KindToAPIMapping[kind]
	if !ok {
		// todo crd resource
	}

	return v.getRESTClientByType(resource.ClientType), resource, nil
}

func (v *Verber) getRESTClientByType(clientType ClientType) RESTClient {
	switch clientType {
	case ClientTypeAppsClient:
		return v.appsClient
	case ClientTypeBatchClient:
		return v.batchClient
	case ClientTypeAutoscalingClient:
		return v.autoscalingClient
	case ClientTypeStorageClient:
		return v.storageClient
	case ClientTypeRbacClient:
		return v.rbacClient
	case ClientTypeNetworkingClient:
		return v.networkingClient
	// case ClientTypeAPIExtensionsClient:
	// 	return v.apiExtensionsClient
	case ClientTypePluginsClient:
		return v.pluginsClient
	default:
		return v.client
	}
}

func (v *Verber) Delete(kind string, namespace string, name string) error {
	client, resource, err := v.getResourceClient(kind)
	if err != nil {
		return err
	}

	// 采用前端删除策略，级联删除子资源
	defaultPropagationPolicy := metav1.DeletePropagationForeground
	defaultDeleteOptions := &metav1.DeleteOptions{
		PropagationPolicy: &defaultPropagationPolicy,
	}

	req := client.Delete().Resource(resource.Resource).Name(name).Body(defaultDeleteOptions)

	if resource.Namespaced {
		req = req.Namespace(namespace)
	}

	return req.Do(context.TODO()).Error()
}

func (v *Verber) Put(kind string, namespace string, name string, data []byte) error {
	client, resource, err := v.getResourceClient(kind)
	if err != nil {
		return err
	}

	req := client.Put().Resource(resource.Resource).Name(name).Body(data).SetHeader("Content-Type", "application/json")

	if resource.Namespaced {
		req = req.Namespace(namespace)
	}

	return req.Do(context.TODO()).Error()
}

func (v *Verber) Get(kind string, namespace string, name string) (runtime.Object, error) {
	client, resource, err := v.getResourceClient(kind)
	if err != nil {
		return nil, err
	}

	req := client.Get().Resource(resource.Resource).Name(name).SetHeader("Accept", "application/json")

	if resource.Namespaced {
		req = req.Namespace(namespace)
	}
	result := &runtime.Unknown{}
	err = req.Do(context.TODO()).Into(result)

	return result, err
}
