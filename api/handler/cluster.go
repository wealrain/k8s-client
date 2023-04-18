package handler

import (
	"fmt"
	"k8s-client/api/client"
	"k8s-client/pkg/cluster"
	"k8s-client/pkg/errors"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/discovery"
)

func HandleAddCluster(c *gin.Context) {
	// 从json中获取数据
	var clus cluster.Cluster
	err := c.ShouldBindJSON(&clus)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	err = cluster.AddCluster(clus)
	if err != nil {
		if _, ok := err.(*errors.ClusterError); ok {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// 根据config的内容获取集群信息
	clientManager, err := client.NewClientManager(fmt.Sprintf("cluster-%d", clus.Id), clus.Config)
	if err != nil {
		clus.Status = "error"
		cluster.UpdateCluster(clus)
		c.JSON(400, gin.H{"error": "集群配置错误"})
		return
	}
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(clientManager.Config)
	if err != nil {
		clus.Status = "error"
		cluster.UpdateCluster(clus)
		c.JSON(400, gin.H{"error": "集群配置错误"})
		return
	}

	// 获取集群的版本信息
	version, err := discoveryClient.ServerVersion()
	if err != nil {
		clus.Status = "running"
		clus.Version = "unknown"
		cluster.UpdateCluster(clus)
		c.JSON(200, gin.H{
			"message": "success",
		})
		return
	}

	clus.Version = version.Major + "." + version.Minor
	clus.Status = "running"
	cluster.UpdateCluster(clus)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func HandleDeleteCluster(c *gin.Context) {
	id := c.Param("id")
	err := cluster.DeleteCluster(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// 删除集群时，删除集群的client
	client.ClientManagerMapInstance.DeleteClientManager(fmt.Sprintf("cluster-%s", id))

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func HandleUpdateCluster(c *gin.Context) {
	var clus cluster.Cluster
	err := c.ShouldBindJSON(&clus)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	err = cluster.UpdateCluster(clus)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// 如果是更新集群的config，需要重新获取集群信息
	if clus.Config != "" {
		clientManager, err := client.NewClientManager(fmt.Sprintf("cluster-%d", clus.Id), clus.Config)
		if err != nil {
			clus.Status = "error"
			cluster.UpdateCluster(clus)
			c.JSON(400, gin.H{"error": "集群配置错误"})
			return
		}

		version := clientManager.Config.GroupVersion.Version
		clus.Version = version
		clus.Status = "running"
		cluster.UpdateCluster(clus)
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func HandleGetCluster(c *gin.Context) {
	id := c.Param("id")
	clus, err := cluster.GetClusterById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, clus)
}

func HandleListCluster(c *gin.Context) {
	cluss, err := cluster.GetClusterList()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, cluss)
}
