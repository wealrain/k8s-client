package client

type ClientType string

const (
	ClientTypeDefault             = "restclient"
	ClientTypeAppsClient          = "appsclient"
	ClientTypeBatchClient         = "batchclient"
	ClientTypeAutoscalingClient   = "autoscalingclient"
	ClientTypeStorageClient       = "storageclient"
	ClientTypeRbacClient          = "rbacclient"
	ClientTypeAPIExtensionsClient = "apiextensionsclient"
	ClientTypeNetworkingClient    = "networkingclient"
	ClientTypePluginsClient       = "plugin"
)

type ResourceKind string

const (
	ResourceKindConfigMap                = "configmap"
	ResourceKindDaemonSet                = "daemonset"
	ResourceKindDeployment               = "deployment"
	ResourceKindEvent                    = "event"
	ResourceKindHorizontalPodAutoscaler  = "horizontalpodautoscaler"
	ResourceKindIngress                  = "ingress"
	ResourceKindServiceAccount           = "serviceaccount"
	ResourceKindJob                      = "job"
	ResourceKindCronJob                  = "cronjob"
	ResourceKindLimitRange               = "limitrange"
	ResourceKindNamespace                = "namespace"
	ResourceKindNode                     = "node"
	ResourceKindPersistentVolumeClaim    = "persistentvolumeclaim"
	ResourceKindPersistentVolume         = "persistentvolume"
	ResourceKindCustomResourceDefinition = "customresourcedefinition"
	ResourceKindPod                      = "pod"
	ResourceKindReplicaSet               = "replicaset"
	ResourceKindReplicationController    = "replicationcontroller"
	ResourceKindResourceQuota            = "resourcequota"
	ResourceKindSecret                   = "secret"
	ResourceKindService                  = "service"
	ResourceKindStatefulSet              = "statefulset"
	ResourceKindStorageClass             = "storageclass"
	ResourceKindClusterRole              = "clusterrole"
	ResourceKindClusterRoleBinding       = "clusterrolebinding"
	ResourceKindRole                     = "role"
	ResourceKindRoleBinding              = "rolebinding"
	ResourceKindPlugin                   = "plugin"
	ResourceKindEndpoint                 = "endpoint"
	ResourceKindNetworkPolicy            = "networkpolicy"
	ResourceKindIngressClass             = "ingressclass"
)

type APIMapping struct {
	Resource   string
	ClientType ClientType
	Namespaced bool
}

var KindToAPIMapping = map[string]APIMapping{
	ResourceKindConfigMap:                {"configmaps", ClientTypeDefault, true},
	ResourceKindDaemonSet:                {"daemonsets", ClientTypeAppsClient, true},
	ResourceKindDeployment:               {"deployments", ClientTypeAppsClient, true},
	ResourceKindEvent:                    {"events", ClientTypeDefault, true},
	ResourceKindHorizontalPodAutoscaler:  {"horizontalpodautoscalers", ClientTypeAutoscalingClient, true},
	ResourceKindIngress:                  {"ingresses", ClientTypeNetworkingClient, true},
	ResourceKindIngressClass:             {"ingressclasses", ClientTypeNetworkingClient, false},
	ResourceKindJob:                      {"jobs", ClientTypeBatchClient, true},
	ResourceKindCronJob:                  {"cronjobs", ClientTypeBatchClient, true},
	ResourceKindLimitRange:               {"limitrange", ClientTypeDefault, true},
	ResourceKindNamespace:                {"namespaces", ClientTypeDefault, false},
	ResourceKindNode:                     {"nodes", ClientTypeDefault, false},
	ResourceKindPersistentVolumeClaim:    {"persistentvolumeclaims", ClientTypeDefault, true},
	ResourceKindPersistentVolume:         {"persistentvolumes", ClientTypeDefault, false},
	ResourceKindCustomResourceDefinition: {"customresourcedefinitions", ClientTypeAPIExtensionsClient, false},
	ResourceKindPod:                      {"pods", ClientTypeDefault, true},
	ResourceKindReplicaSet:               {"replicasets", ClientTypeAppsClient, true},
	ResourceKindReplicationController:    {"replicationcontrollers", ClientTypeDefault, true},
	ResourceKindResourceQuota:            {"resourcequotas", ClientTypeDefault, true},
	ResourceKindSecret:                   {"secrets", ClientTypeDefault, true},
	ResourceKindService:                  {"services", ClientTypeDefault, true},
	ResourceKindServiceAccount:           {"serviceaccounts", ClientTypeDefault, true},
	ResourceKindStatefulSet:              {"statefulsets", ClientTypeAppsClient, true},
	ResourceKindStorageClass:             {"storageclasses", ClientTypeStorageClient, false},
	ResourceKindEndpoint:                 {"endpoints", ClientTypeDefault, true},
	ResourceKindNetworkPolicy:            {"networkpolicies", ClientTypeNetworkingClient, true},
	ResourceKindClusterRole:              {"clusterroles", ClientTypeRbacClient, false},
	ResourceKindClusterRoleBinding:       {"clusterrolebindings", ClientTypeRbacClient, false},
	ResourceKindRole:                     {"roles", ClientTypeRbacClient, true},
	ResourceKindRoleBinding:              {"rolebindings", ClientTypeRbacClient, true},
	ResourceKindPlugin:                   {"plugins", ClientTypePluginsClient, true},
}
