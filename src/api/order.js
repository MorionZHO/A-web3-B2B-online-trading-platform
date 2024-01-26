import axios from "axios";
import axiosInstance from "../utils/request"
import { message } from "antd";

export const submitOrder = ((payload)=>{
    return axiosInstance.post(`/buyer/submitOrder`,payload).then(res=>{
        return res;
    })
})