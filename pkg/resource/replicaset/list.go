package replicaset

import (
	"k8s-client/pkg/resource/common"

	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type ReplicaSet struct {
	Name    string `json:"name"`
	Desired int32  `json:"desired"`
	Current int32  `json:"current"`
	Ready   int32  `json:"ready"`
	Age     int64  `json:"age"`
}

func GetReplicaSetList(client k8sClient.Interface, namespace string) ([]ReplicaSet, error) {

	channel := common.GetReplicaSetListChannelWithOptions(client, namespace, metav1.ListOptions{})

	replicaSets := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var replicaSetList []ReplicaSet

	for _, replicaSet := range replicaSets.Items {
		replicaSetInfo := toReplicaSet(&replicaSet)
		replicaSetList = append(replicaSetList, replicaSetInfo)
	}

	return replicaSetList, nil
}

func toReplicaSet(replicaSet *appv1.ReplicaSet) ReplicaSet {
	return ReplicaSet{
		Name:    replicaSet.Name,
		Desired: *replicaSet.Spec.Replicas,
		Current: replicaSet.Status.Replicas,
		Ready:   replicaSet.Status.ReadyReplicas,
		Age:     replicaSet.CreationTimestamp.Unix(),
	}
}

func ToInterface(replicaSetList []ReplicaSet) []interface{} {
	var result []interface{}
	for _, replicaSet := range replicaSetList {
		result = append(result, replicaSet)
	}
	return result
}
