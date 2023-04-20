package handler

import (
	"k8s-client/pkg/resource/common"
	"k8s-client/pkg/resource/pod"
	"strconv"
	"time"

	"github.com/emicklei/go-restful/v3/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/runtime"
)

var upgrader = websocket.Upgrader{}

func HandleDelete(c *gin.Context) {
	ns := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")

	err := managerK8SVerber(c).Delete(kind, ns, name)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// todo 如果他把本资源删除了 ， 即当前k8s-client相关的资源

	c.JSON(200, gin.H{"message": "success"})
}

func HandleGet(c *gin.Context) {
	ns := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")

	result, err := managerK8SVerber(c).Get(kind, ns, name)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, result)
}

func HandleUpdate(c *gin.Context) {
	ns := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")

	data := runtime.Unknown{}
	err := c.BindJSON(&data)
	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	err = managerK8SVerber(c).Put(kind, ns, name, data.Raw)

	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func HandleGetPodLog(c *gin.Context) {
	ns := c.Param("namespace")
	name := c.Param("name")
	container := c.Param("container")
	client := managerK8SClient(c)

	sessionId, err := genSessionId()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	SockSessions.Set(sessionId, SockSession{
		SessionId: sessionId,
		Namespace: ns,
		Name:      name,
		Container: container,
		Client:    client,
		SockType:  LOG,
		bound:     make(chan error),
	})

	go WaitFor(sessionId)

	c.JSON(200, gin.H{"sessionId": sessionId})

}

func HandleDownloadPodLog(c *gin.Context) {
	ns := c.Param("namespace")
	name := c.Param("name")
	container := c.Param("container")
	client := managerK8SClient(c)

	logs, err := pod.GetLogsFile(client, ns, name, container)

	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// 生成下载文件名 容器 + pod名 + 时间
	name = container + "-" + name + "-" + time.Now().Format("2006-01-02-15-04-05")
	log.Printf("download file name: %s", name)
	c.Header("Content-Disposition", "attachment; filename="+name+".log")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.String(200, logs)

}

func HandleScale(c *gin.Context) {
	ns := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")
	replicas := c.Param("replicas")

	count, err := strconv.Atoi(replicas)
	if err != nil {
		c.JSON(400, gin.H{"error": "replicas must be a number"})
		return
	}

	err = common.ScaleResource(managerK8SConig(c), kind, ns, name, int32(count))

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
