import http from ".";
import { currentCluster } from "../store/cluster";

const listPods = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/pods/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listDeployments = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/deployments/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listStatefulSets = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/statefulsets/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listReplicasets = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/replicasets/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listServices = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/services/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listIngresses = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/ingresses/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listEndpoints = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/endpoints/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listConfigmaps = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/configmaps/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listSecrets= (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/secrets/` + namespace,dataFilter,null,{
        needShowErrorMsg:true
    });
};

const listEvents = (namespace,dataFilter) =>  {
    return http.post(`/list/${currentCluster()}/events/${namespace}`,dataFilter,null,{
        needShowErrorMsg:true
    });
};

export default {
    listPods,
    listDeployments,
    listStatefulSets,
    listReplicasets,
    listServices,
    listIngresses,
    listEndpoints,
    listConfigmaps,
    listSecrets,
    listEvents
};
