import http from ".";
import { currentCluster } from "../store/cluster";
const getLog = (namespace, name,container) => {
    return http.get(`/log/${currentCluster()}/download/${namespace}/${name}/${container}`,null,null,{
        raw: true
    });
}

const getLogSession = (namespace, name,container) => {
    return http.get(`/log/${currentCluster()}/session/${namespace}/${name}/${container}`);
}

export default {
    getLog,
    getLogSession,
}