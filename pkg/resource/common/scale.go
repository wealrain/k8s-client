package common

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/scale"
)

func ScaleResource(cfg *rest.Config, kind, namespace, name string, count int32) error {
	groupResource := schema.ParseGroupResource(kind)
	if groupResource.Group == "" || groupResource.Resource == "" {
		groupResource = appsv1.Resource(kind)
	}

	sc, err := getScalesGetter(cfg)
	if err != nil {
		return err
	}

	resource, err := sc.Scales(namespace).Get(context.TODO(), groupResource, name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	resource.Spec.Replicas = count

	_, err = sc.Scales(namespace).Update(context.TODO(), groupResource, resource, metav1.UpdateOptions{})
	return err
}

// / 在Kubernetes中，许多资源都可以进行扩展，例如Deployment、StatefulSet、ReplicaSet等。
// / ScalesGetter接口提供了一种标准化的方式来获取这些资源的扩展信息
func getScalesGetter(cfg *rest.Config) (scale.ScalesGetter, error) {

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(cfg)

	if err != nil {
		return nil, err
	}

	cfg.GroupVersion = &appsv1.SchemeGroupVersion
	cfg.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(cfg)
	if err != nil {
		return nil, err
	}

	resolver := scale.NewDiscoveryScaleKindResolver(discoveryClient)
	dc := memory.NewMemCacheClient(discoveryClient)
	drm := restmapper.NewDeferredDiscoveryRESTMapper(dc)

	drm.Reset()

	return scale.New(restClient, drm, dynamic.LegacyAPIPathResolverFunc, resolver), nil
}
