package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后2位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 10001
const REUQEST_PARAM_ERROR uint32 = 10002
const TOKEN_EXPIRE_ERROR uint32 = 10003
