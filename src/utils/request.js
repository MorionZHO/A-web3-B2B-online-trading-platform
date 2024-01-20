//在index.js中引入axios
import axios from 'axios';
import { message } from 'antd'
import { diffTokenTime } from "./auth";
import storage from './storage';
// apiService.js

// 创建axios实例
const axiosInstance = axios.create({
  baseURL: 'https://www.fastmock.site/mock/602626775b8e8e6f688cc0cbc8f37dbe/api', // 你的API基础URL
  timeout: 8000,
});

// 添加请求拦截器
axiosInstance.interceptors.request.use(config => {
  // 在发送请求之前做些什么，例如添加认证信息
  const token = localStorage.getItem('authToken');
  if (token) {
    if (diffTokenTime()) {
      storage.clearAll();
      message.error('Token has expired, please log in again!')
    }
    config.headers['Authorization'] = `Bearer ${token}`;
  }
  return config;
}, error => {
  // 对请求错误做些什么
  message.error(error)
  return Promise.reject(error);
});

// 添加响应拦截器
axiosInstance.interceptors.response.use(response => {
  // 对响应数据做点什么
  return response.data; // 只返回数据部分
}, error => {
  // 对响应错误做点什么
  if (error.response && error.response.status === 401) {
    window.location.href = '/login'
    // 处理401错误（未授权访问）
    // 例如：跳转到登录页面
  }
  return Promise.reject(error);
});

// API服务模块
const apiService = {
  getData: async () => {
    try {
      return await axiosInstance.get('/data');
    } catch (error) {
      throw error;
    }
  },
  // ...更多API函数
};

export default axiosInstance;
