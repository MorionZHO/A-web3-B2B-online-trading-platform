import axiosInstance from "../utils/request"
import { message } from "antd";

export const fetchProductDetail = ((id)=>{
    return axiosInstance.get(`/shop?id=${id}`).then(res=>{
        return res;
    })
    .catch(error=>{
        message.error(error)
    })
})

export const getAllProducts = (()=>{
    return axiosInstance.get('/shop').then(res=>{
        return res;
    })
})

export const searchProduct = ((keyword)=>{
    return axiosInstance.get(`/search?KeyWord=${keyword}`).then(res=>{
        return res;
    })
})
