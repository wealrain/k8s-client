import http from ".";
import { currentCluster } from "../store/cluster";

const podDetail = (namespace, name) => {
    return http.get(`/detail/${currentCluster()}/pod/${namespace}/${name}`);
};

const deploymentDetail = (namespace, name) => {
    return http.get(`/detail/${currentCluster()}/deployment/${namespace}/${name}`);
};

const configmapDetail = (namespace, name) => {
    return http.get(`/detail/${currentCluster()}/configmap/${namespace}/${name}`);
};

export default {
    podDetail,
    deploymentDetail,
    configmapDetail,

}