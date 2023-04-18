package handler

import (
	"fmt"
	"k8s-client/api/client"
	"k8s-client/pkg/cluster"
	"log"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func manager(c *gin.Context) *client.ClientManager {
	// 通过id查看集群是否存在
	id := c.Param("cluster")
	log.Printf("cluster-%s", id)
	clientManager := client.ClientManagerMapInstance.GetClientManager(id)
	if clientManager == nil {
		clus, err := cluster.GetClusterById(id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return nil
		}

		if clus == nil {
			c.JSON(404, gin.H{"error": "Cluster not found"})
			c.Abort()
			return nil
		}
		log.Printf("cluster-%d", clus.Id)
		clientManager, err = client.NewClientManager(fmt.Sprintf("cluster-%d", clus.Id), clus.Config)

		if err != nil {
			c.JSON(400, gin.H{"error": "集群配置错误"})
			c.Abort()
			return nil
		}

	}

	return clientManager
}

func managerK8SConig(c *gin.Context) *rest.Config {
	return manager(c).Config
}

func managerK8SClient(c *gin.Context) kubernetes.Interface {
	return manager(c).Client
}

func managerK8SVerber(c *gin.Context) *client.Verber {
	return manager(c).VerberClient
}
