import http from ".";
import md5 from "md5";
function createToken(username,password) {
  // 使用md5加密密码
    const data = {
        username: md5(username),
        password: md5(password)
    }

  return http.post("/token/create", data);
}

function verifyToken(token) {
  return http.get("/token/verify",null,{
    Authorization: token
  },{
    needShowSuccessMsg:true,
    msg: "login success"
  });
}

export default {
    createToken,
    verifyToken
}