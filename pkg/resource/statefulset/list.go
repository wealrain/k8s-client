package statefulset

import (
	"fmt"
	"k8s-client/pkg/resource/common"

	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type StatefulSet struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Pods      string `json:"pods"`
	Age       int64  `json:"age"`
	Replicas  int32  `json:"replicas"`
}

func GetStatefulSetList(client k8sClient.Interface, namespace string) ([]StatefulSet, error) {

	channel := common.GetStatefulSetListChannelWithOptions(client, namespace, metav1.ListOptions{})

	statefulSets := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var statefulSetList []StatefulSet

	for _, statefulSet := range statefulSets.Items {
		statefulSetInfo := toStatefulSet(&statefulSet)
		statefulSetList = append(statefulSetList, statefulSetInfo)
	}

	return statefulSetList, nil
}

func toStatefulSet(statefulSet *appv1.StatefulSet) StatefulSet {
	// 获取ready的pod数量和总的pod数量之比
	readyPods := statefulSet.Status.ReadyReplicas
	totalPods := statefulSet.Status.Replicas

	return StatefulSet{
		Name:      statefulSet.Name,
		Namespace: statefulSet.Namespace,
		Pods:      fmt.Sprintf("%d/%d", readyPods, totalPods),
		Age:       statefulSet.CreationTimestamp.Unix(),
		Replicas:  statefulSet.Status.Replicas,
	}
}

func ToInterface(StatefulSetList []StatefulSet) []interface{} {
	var result []interface{}
	for _, StatefulSet := range StatefulSetList {
		result = append(result, StatefulSet)
	}
	return result
}
