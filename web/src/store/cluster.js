import storage from "."

const CURRENT_CLUSTER = 'currentCluster'
const CURRENT_CLUSTER_NAME = 'currentClusterName'

export const currentCluster = () => storage.getItem(CURRENT_CLUSTER)
export const currentClusterName = () => storage.getItem(CURRENT_CLUSTER_NAME)

export const setCurrentCluster = (cluster) => {
    storage.setItem(CURRENT_CLUSTER, cluster)
}
export const setCurrentClusterName = (clusterName) => {
    storage.setItem(CURRENT_CLUSTER_NAME, clusterName)
}
