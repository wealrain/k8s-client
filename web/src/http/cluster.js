import http from ".";

function listClusters() {
  return http.get("/cluster/list");
}

function createCluster(data) {
  return http.post("/cluster/add", data,null,{
    needShowSuccessMsg:true,
    msg: "add cluster success"
  });
}

function updateCluster(data) {
  return http.put(`/cluster/update`, data);
}

function deleteCluster(id) {
   return http.delete(`/cluster/${id}`);
}

export default {
    listClusters,
    createCluster,
    updateCluster,
    deleteCluster
}
