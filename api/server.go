package api

import (
	"k8s-client/api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

func list(r *gin.Engine) {
	api := r.Group("/list")
	api.POST("/:cluster/pods/:namespace", handler.HandleGetPods)
	api.POST("/:cluster/deployments/:namespace", handler.HandleGetDeployments)
	api.POST("/:cluster/statefulsets/:namespace", handler.HandleGetStatefulSets)
	api.POST("/:cluster/replicasets/:namespace", handler.HandleGetReplicaSets)
	api.POST("/:cluster/services/:namespace", handler.HandleGetServices)
	api.POST("/:cluster/ingresses/:namespace", handler.HandleGetIngresses)
	api.POST("/:cluster/endpoints/:namespace", handler.HandleGetEndpoints)
	api.POST("/:cluster/configmaps/:namespace", handler.HandleGetConfigMaps)
	api.POST("/:cluster/secrets/:namespace", handler.HandleGetSecrets)
	api.POST("/:cluster/events/:namespace", handler.HandleGetEvents)
	api.GET("/:cluster/namespace", handler.HandleGetNamespaces)
}

func op(r *gin.Engine) {
	api := r.Group("/op")
	api.DELETE("/:cluster/:kind/:namespace/:name", handler.HandleDelete)
	api.GET("/:cluster/:kind/:namespace/:name", handler.HandleGet)
	api.PUT("/:cluster/:kind/:namespace/:name", handler.HandleUpdate)
	api.PUT("/:cluster/scale/:kind/:namespace/:name/:replicas", handler.HandleScale)
	api.PUT("/:cluster/resume/:namespace/:name", handler.HandleDeploymentResume)
	api.PUT("/:cluster/pause/:namespace/:name", handler.HandleDeploymentPause)
	api.PUT("/:cluster/restart/:namespace/:name", handler.HandleDeploymentRestart)
}

func log(r *gin.Engine) {
	r.GET("/logs/*any", gin.WrapH(handler.CreateAttachHandler("/logs")))
	r.GET("/log/:cluster/session/:namespace/:name/:container", handler.HandleGetPodLog)
	r.GET("/log/:cluster/session/:namespace/:name/", handler.HandleGetPodLog)
	r.GET("/log/:cluster/download/:namespace/:name/:container", handler.HandleDownloadPodLog)
	r.GET("/log/:cluster/download/:namespace/:name/", handler.HandleDownloadPodLog)
}

func terminal(r *gin.Engine) {
	sockjs.NewHandler("/api/terminal", sockjs.DefaultOptions, handler.HandleTerminalSession)
	api := r.Group("/terminal")
	api.GET("/shell/:namespace/:name/:container", handler.HandleExecShell)
}

func detail(r *gin.Engine) {
	api := r.Group("/detail")
	api.GET("/:cluster/pod/:namespace/:name", handler.HandleGetPodDetail)
	api.GET("/:cluster/deployment/:namespace/:name", handler.HandleGetDeploymentDetail)
	// api.GET("/statefulset/:namespace/:name", handler.HandleGetStatefulSetDetail)
	// api.GET("/replicaset/:namespace/:name", handler.HandleGetReplicaSetDetail)
	// api.GET("/service/:namespace/:name", handler.HandleGetServiceDetail)
	// api.GET("/ingress/:namespace/:name", handler.HandleGetIngressDetail)
	// api.GET("/endpoint/:namespace/:name", handler.HandleGetEndpointDetail)
	api.GET("/:cluster/configmap/:namespace/:name", handler.HandleGetConfigMapDetail)
	// api.GET("/secret/:namespace/:name", handler.HandleGetSecretDetail)
	// api.GET("/event/:namespace/:name", handler.HandleGetEventDetail)

}

func cluster(r *gin.Engine) {
	api := r.Group("/cluster")
	api.POST("/add", handler.HandleAddCluster)
	api.GET("/list", handler.HandleListCluster)
	api.DELETE("/delete/:id", handler.HandleDeleteCluster)
	api.PUT("/update", handler.HandleUpdateCluster)
	api.GET("/info/:id", handler.HandleGetCluster)

}

func user(r *gin.Engine) {
	api := r.Group("/user")
	api.POST("/add", handler.HandleAddUser)
	api.POST("/list", handler.HandleListUser)
	api.POST("/delete", handler.HandleDeleteUser)
	api.POST("/update", handler.HandleUpdateUser)
	// api.GET("/:id", handler.HandleGetUser)
	api.POST("/login", handler.HandleLogin)
}

func token(r *gin.Engine) {
	api := r.Group("/token")
	api.POST("/create", handler.HandleCreateToken)
}

func Start(addr string) {
	// Start the server
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Next()
	})

	r.Use(handler.Recover)
	// 拦截所有请求，验证token，除了登录接口
	r.Use(handler.AuthMiddleware)
	cluster(r)
	user(r)
	list(r)
	detail(r)
	op(r)
	log(r)
	terminal(r)
	token(r)
	r.Run(addr)
}
