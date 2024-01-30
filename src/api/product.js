import axios from "axios";
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

export const getProductByType = ((type)=>{
    return axiosInstance.get(`/shop?type=${type}`).then(res=>{
        return res;
    })
})

export const shelfProduct = ((payload)=>{
    return axiosInstance.post(`/seller/shelfProduct`,payload).then(res=>{
        return res;
    })
})

export const getMyProducts = ((companyName)=>{
    return axiosInstance.get(`/shop/manageProduct?companyName=${companyName}`).then(res=>{
        return res;
    })
})

export const editProduct = ((payload)=>{
    return axiosInstance.post(`/seller/editProduct`,payload).then(res=>{
        return res;
    })
})

