package handler

import (
	"k8s-client/pkg/resource/configmap"
	"k8s-client/pkg/resource/deployment"
	"k8s-client/pkg/resource/endpoint"
	"k8s-client/pkg/resource/event"
	"k8s-client/pkg/resource/ingress"
	"k8s-client/pkg/resource/pod"
	"k8s-client/pkg/resource/replicaset"
	"k8s-client/pkg/resource/secret"
	"k8s-client/pkg/resource/service"
	"k8s-client/pkg/resource/statefulset"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGetPods handles the GET /pods endpoint
func HandleGetPods(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	pods, err := pod.GetPodList(client, ns)
	if err != nil {
		log.Printf("get pod list error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(pod.ToInterface(pods))
	c.JSON(200, result)
}

func HandleGetDeployments(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	deployments, err := deployment.GetDeploymentList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(deployment.ToInterface(deployments))
	c.JSON(200, result)
}

func HandleGetStatefulSets(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	statfulSets, err := statefulset.GetStatefulSetList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(statefulset.ToInterface(statfulSets))
	c.JSON(200, result)
}

func HandleGetServices(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	services, err := service.GetServiceList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(service.ToInterface(services))
	c.JSON(200, result)
}

func HandleGetReplicaSets(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	replicaSets, err := replicaset.GetReplicaSetList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(replicaset.ToInterface(replicaSets))
	c.JSON(200, result)
}

func HandleGetIngresses(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	ingresses, err := ingress.GetIngressList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(ingress.ToInterface(ingresses))
	c.JSON(200, result)
}

func HandleGetEndpoints(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	endpoints, err := endpoint.GetEndpointList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(endpoint.ToInterface(endpoints))
	c.JSON(200, result)
}

func HandleGetConfigMaps(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	configmaps, err := configmap.GetConfigmapList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(configmap.ToInterface(configmaps))
	c.JSON(200, result)
}

func HandleGetSecrets(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	secrets, err := secret.GetSecretList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(secret.ToInterface(secrets))
	c.JSON(200, result)
}

func HandleGetEvents(c *gin.Context) {
	client := managerK8SClient(c)
	ns := c.Param("namespace")
	events, err := event.GetEventList(client, ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 过滤数据
	dataFilter := DataFilter{}
	if err := dataFilter.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := dataFilter.Filter(event.ToInterface(events))
	c.JSON(200, result)
}
