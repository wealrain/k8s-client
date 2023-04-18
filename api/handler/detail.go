package handler

import (
	"k8s-client/pkg/resource/configmap"
	"k8s-client/pkg/resource/deployment"
	"k8s-client/pkg/resource/pod"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetPodDetail(c *gin.Context) {
	client := managerK8SClient(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	pod, err := pod.GetPodDetail(client, namespace, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pod)
}

func HandleGetDeploymentDetail(c *gin.Context) {
	client := managerK8SClient(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	deployment, err := deployment.GetDeploymentDetail(client, namespace, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deployment)
}

// func HandleGetStatefulSetDetail(c *gin.Context) {
// 	namespace := c.Param("namespace")
// 	name := c.Param("name")
// 	statefulset, err := client.GetStatefulSet(namespace, name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, statefulset)
// }

// func HandleGetReplicaSetDetail(c *gin.Context) {
// 	namespace := c.Param("namespace")
// 	name := c.Param("name")
// 	replicaset, err := client.GetReplicaSet(namespace, name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, replicaset)
// }

// func HandleGetServiceDetail(c *gin.Context) {
// 	namespace := c.Param("namespace")
// 	name := c.Param("name")
// 	service, err := client.GetService(namespace, name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, service)
// }

// func HandleGetIngressDetail(c *gin.Context) {
// 	namespace := c.Param("namespace")
// 	name := c.Param("name")
// 	ingress, err := client.GetIngress(namespace, name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, ingress)
// }

// func HandleGetEndpointDetail(c *gin.Context) {
// 	namespace := c.Param("namespace")
// 	name := c.Param("name")
// 	endpoint, err := client.GetEndpoint(namespace, name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, endpoint)
// }

func HandleGetConfigMapDetail(c *gin.Context) {
	client := managerK8SClient(c)
	namespace := c.Param("namespace")
	name := c.Param("name")
	configmap, err := configmap.GetConfigMapDetail(client, namespace, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, configmap)
}

// func HandleGetSecretDetail(c *gin.Context) {
// 	namespace := c.Param("namespace")
// 	name := c.Param("name")
// 	secret, err := client.GetSecret(namespace, name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, secret)
// }

// func HandleGetEventDetail(c *gin.Context) {
// 	namespace := c.Param("namespace")
// 	name := c.Param("name")
// 	event, err := client.GetEvent(namespace, name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, event)
// }
