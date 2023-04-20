import axios from 'axios'
import * as React from 'react'
import Snackbar from '@mui/material/Snackbar';
import MuiAlert from '@mui/material/Alert';
import { userToken,removeUserToken } from '../store/token';

class HttpRequest {
    constructor(config) {
        this.http = axios.create(config);
    }

    get(url,params,headers,options) {
       return new Promise((resolve,reject) => {
            this.http.get(url,{params,headers}).then(res => {
                if(options && options.needShowSuccessMsg) {
                    resolve( {
                        msg: options.msg,
                        type: "success"
                    })
                }

                if(options && options.raw) {
                    resolve(res);
                }

                resolve(res.data);
            }).catch(err => {
                if (err.response.status === 401) {
                    // token失效，跳转到登录页面
                    removeUserToken();
                  }
                if(options && options.needShowSuccessMsg) {
                    reject ({
                        msg: err.response.data.error,
                        type: "error"
                    })
                }
                reject(err);
            });
       });
    }

    post(url,data,headers,options) {
        return new Promise((resolve,reject) => {
            this.http.post(url,data,{headers}).then(res => {
                if(options && options.needShowSuccessMsg) {
                    resolve( {
                        msg: options.msg,
                        type: "success"
                    })
                }
                resolve(res.data);
            }).catch(err => {
                if (err.response.status === 401) {
                    // token失效，跳转到登录页面
                    removeUserToken();
                  }
                if(options && options.needShowSuccessMsg) {
                    reject ({
                        msg: err.response.data.error,
                        type: "error"
                    })
                }
                reject(err);
            });
        });
    }

    put(url,data,headers,options) {
        return new Promise((resolve,reject) => {
            this.http.put(url,data,{headers}).then(res => {
                if(options && options.needShowSuccessMsg) {
                    resolve( {
                        msg: options.msg,
                        type: "success"
                    })
                }
                resolve(res.data);
            }).catch(err => {
                if (err.response.status === 401) {
                    // token失效，跳转到登录页面
                    removeUserToken();
                  }
                if(options && options.needShowSuccessMsg) {
                    reject ({
                        msg: err.response.data.error,
                        type: "error"
                    })
                }
                reject(err);
            });
        });
    }
         

    delete(url,headers,options) {
        return new Promise((resolve,reject) => {
            this.http.delete(url,{headers}).then(res => {
                if(options && options.needShowSuccessMsg) {
                    resolve( {
                        msg: options.msg,
                        type: "success"
                    })
                }
                resolve(res.data);
            }).catch(err => {
                if (err.response.status === 401) {
                    // token失效，跳转到登录页面
                    removeUserToken();
                    
                  }
                if(options && options.needShowErrorMsg) {
                    reject ({
                        msg: err.response.data.error,
                        type: "error"
                    })
                }
                reject(err);
            });
        });
    }

    setRequestInterceptor(onFulfilled,onRejected) {
        this.http.interceptors.request.use(onFulfilled,onRejected);
        return this;
    }

    setResponseInterceptor(onFulfilled,onRejected) {
        this.http.interceptors.response.use(onFulfilled,onRejected);
        return this;
    }
}

export const HttpContext = React.createContext({
    showMsg: (msg,type) => {}
});

const Alert = React.forwardRef(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

export function HttpProvider({children}) {
    const [open, setOpen] = React.useState(false);
    const [msg, setMsg] = React.useState("");
    const [type, setType] = React.useState("error");

    const handleClick = (msg,type) => {
        setOpen(true);
        setMsg(msg);
        setType(type);
    };

    const handleClose = (event, reason) => {
        setOpen(false);
    };

    return (
        <HttpContext.Provider value={{showMsg:handleClick}}>
            {children}
            <Snackbar open={open} autoHideDuration={3000} onClose={handleClose}>
            <Alert onClose={handleClose} severity={type} sx={{ width: '100%' }}>
                    {msg}
            </Alert>
            </Snackbar>
        </HttpContext.Provider>
    );
}

const http = new HttpRequest({
    baseURL: '/api',
    timeout: 1000 * 60 * 5,
});

http.setRequestInterceptor((config) => {
    // 在请求头中加入token
    const token = userToken();
    if(token) {
        config.headers.Authorization = token;
    }
    return config;
},(error) => {
    return Promise.reject(error);
});

export default http;