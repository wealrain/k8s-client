package handler

import (
	"k8s-client/pkg/resource/deployment"

	"github.com/gin-gonic/gin"
)

func HandleDeploymentPause(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	err := deployment.PauseDeployment(managerK8SClient(c), namespace, name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func HandleDeploymentResume(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	err := deployment.ResumeDeployment(managerK8SClient(c), namespace, name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func HandleDeploymentRestart(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	err := deployment.RestartDeployment(managerK8SClient(c), namespace, name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
