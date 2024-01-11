//在index.js中引入axios
import axios from 'axios';
import QS from 'qs';
//ant-design的message提示组件，大家可根据自己的ui组件更改。
import { message } from 'antd'
import config from './../config'


//设置axios基础路径
const service = axios.create({
  baseURL: config.baseUrl,
  timeout: 10000, // 设置超时时间10s
})

service.interceptors.request.use(config => {
  // 每次发送请求之前本地存储中是否存在token，也可以通过Redux这里只演示通过本地拿到token
  // 如果存在，则统一在http请求的header都加上token，这样后台根据token判断你的登录情况
  // 即使本地存在token，也有可能token是过期的，所以在响应拦截器中要对返回状态进行判断 
  const token = window.localStorage.getItem('userToken') || window.sessionStorage.getItem('userToken');
  //在每次的请求中添加token
  config.data = Object.assign({}, config.data, {
    token: token,
  })
  //设置请求头
  config.headers = {
    'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
  }
  //序列化请求参数，不然post请求参数后台接收不正常
  config.data = QS.stringify(config.data)
  return config
}, error => {
  return error;
})



// 响应拦截器
service.interceptors.response.use(response => {
  //根据返回不同的状态码做不同的事情
  // 这里一定要和后台开发人员协商好统一的错误状态码
  if (response.code) {
    switch (response.code) {
      case 200:
        return response.data;
      case 401:
        //未登录处理方法
        break;
      case 403:
        //token过期处理方法
        break;
      default:
        message.error(response.data.msg)
    }
  } else {
    return response;
  }
})
function request(options) {
  const whiteList = ['/api/user/login', '/api/user/register', '/api/task/get_tasks']
  options.method = options.method || 'get';
  if (options.method.toLowerCase() === 'get') {
    options.params = options.data;
  }

  // 添加token
  console.log(options.url)
  if (!whiteList.includes(options.url.split("?")[0])) {
    if (storage.getItem("token")) {
      if (diffTokenTime()) {
        storage.clearAll();
        message.error("Token已过期,请重新登录!");
        setTimeout(() => {
          this.$router.replace("/login").then(r => r);
        }, 2000);
      }
      options["headers"] = {
        token: storage.getItem("token").toString()
      }
    } else {
      message.error("账号异常，请重新登录!");
      return Promise.reject(new Error("账号异常"));
    }
  }

  // 判断使用mock接口还是真实后端接口
  if (config.env === 'prod') {
    service.defaults.baseURL = config.baseApi;
  } else {
    service.defaults.baseURL = config.mock ? config.mockApi : config.baseApi;
  }
  // console.log(options)
  // debugger
  return service(options).catch(err => {
    message.error(err.toString());
  });
}

// 为request函数添加get/post等调用方法
['get', 'post', 'put', 'delete', 'patch'].forEach(item => {
  request[item] = (url, data, options) => {
    return request({
      url,
      data,
      method: item,
      ...options
    })
  }
})

export default request;
