import http from ".";
import { currentCluster } from "../store/cluster";
const deleteResource = (kind,namespace, name) => {
    return http.delete(`/op/${currentCluster()}/${kind}/${namespace}/${name}`);
}

const getResource = (kind,namespace, name) => {
    return http.get(`/op/${currentCluster()}/${kind}/${namespace}/${name}`);
}

const putResource = (kind,namespace, name, body) => {
    return http.put(`/op/${currentCluster()}/${kind}/${namespace}/${name}`, body,null,{
        needShowSuccessMsg: true,
        msg:`${kind} ${name} is updated`
    });
}

const scaleResource = (kind,namespace, name, replicas) => {
    return http.put(`/op/${currentCluster()}/scale/${kind}/${namespace}/${name}/${replicas}`);
}

const restartDeployment = (namespace, name) => {
    return http.put(`/op/${currentCluster()}/restart/${namespace}/${name}`,null,null,{
        needShowSuccessMsg: true,
        msg:`Deployment ${name} is restarted`
    });
}

const pauseDeployment = (namespace, name) => {
    return http.put(`/op/${currentCluster()}/pause/${namespace}/${name}`,null,null,{
        needShowSuccessMsg: true,
        msg:`Deployment ${name} is paused`
    });
}

const resumeDeployment = (namespace, name) => {
    return http.put(`/op/${currentCluster()}/resume/${namespace}/${name}`,null,null,{
        needShowSuccessMsg: true,
        msg:`Deployment ${name} is resumed`
    });
}

export default {
    deleteResource,
    getResource,
    putResource,
    scaleResource,
    restartDeployment,
    pauseDeployment,
    resumeDeployment,
}