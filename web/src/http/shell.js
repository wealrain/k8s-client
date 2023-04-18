import http from "."
import { currentCluster } from "../store/cluster";
function execShell(namespace,pod,container) {
    return http.get(`/shell/${namespace}/${pod}/${container}`)
}

export default {
    execShell
}