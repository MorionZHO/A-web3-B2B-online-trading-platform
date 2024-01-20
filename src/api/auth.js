import axiosInstance from "../utils/request"
import { message } from "antd";


export const handleLogin = ((payload)=>{
    return axiosInstance.post('/user/login',payload).then(res=>{
        return res;
    })
    .catch(error=>{
        message.error(error)
    })
})
    
export const handleRegister = ((payload)=>{
    return axiosInstance.post('/user/register',payload).then(res=>{
        return res;
    })
    .catch(error=>{
        message.error(error)
    })
})