package deployment

import (
	"fmt"
	"k8s-client/pkg/resource/common"

	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type Deployment struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"` // 命名空间
	Pods      string `json:"pods"`
	Age       int64  `json:"age"`
	Replicas  int32  `json:"replicas"`
}

func GetDeploymentList(client k8sClient.Interface, namespace string) ([]Deployment, error) {

	channel := common.GetDeploymentListChannelWithOptions(client, namespace, metav1.ListOptions{})

	deployments := <-channel.List
	err := <-channel.Error

	if err != nil {
		return nil, err
	}

	var deploymentList []Deployment

	for _, deployment := range deployments.Items {
		deploymentInfo := toDeployment(&deployment)
		deploymentList = append(deploymentList, deploymentInfo)
	}

	return deploymentList, nil
}

func toDeployment(deployment *appv1.Deployment) Deployment {
	// 获取ready的pod数量和总的pod数量之比
	readyPods := deployment.Status.ReadyReplicas
	totalPods := deployment.Status.Replicas

	return Deployment{
		Name:      deployment.Name,
		Namespace: deployment.Namespace,
		Pods:      fmt.Sprintf("%d/%d", readyPods, totalPods),
		Age:       deployment.CreationTimestamp.Unix(),
		Replicas:  deployment.Status.Replicas,
	}
}

func ToInterface(deploymentList []Deployment) []interface{} {
	var result []interface{}
	for _, deployment := range deploymentList {
		result = append(result, deployment)
	}
	return result
}
