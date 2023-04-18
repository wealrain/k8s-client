import storage from "."

const CURRENT_CLUSTER = 'currentCluster'

export const currentCluster = () => storage.getItem(CURRENT_CLUSTER)
export const setCurrentCluster = (cluster) => {
    console.log('set cluster', cluster)
    storage.setItem(CURRENT_CLUSTER, cluster)
}