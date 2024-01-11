/**
 * 环境配置
 */
const env = process.env.NODE_ENV || 'prod';
const EnvConfig = {
    dev: {
        baseApi: "http://114.212.101.141:8091",
        mockApi: "https://www.fastmock.site/mock/6180fc8b033f967754808c2b71bde554"
    },
    test: {
        baseApi: "http://114.212.101.141:8091",
        mockApi: "https://www.fastmock.site/mock/94566a192434d6fcf3b3ef6869fc99b0"
    },
    
}
export default {
    env,
    mock: true,
    namespace: 'web3trade',
    ...EnvConfig[env]
}