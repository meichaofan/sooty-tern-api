# 用户模块
- [ChangeLog](#ChangeLog)
- [环境](#环境)
- [用户中心相关接口](#用户中心相关接口)
  - [用户登录](#用户登录)
  - [用户退出](#用户退出)
  - [新建用户](#新建用户)
  - [用户资料详情](#用户资料详情)
  - [用户资料修改](#用户资料修改)
  
---

## ChangeLog

版本 | 变更内容 | 变更时间|变更人
---|---|---|---
V0.0.1 | 初始化版本 | 2019/09/19 | 梅超凡(1783590642@qq.com)

## 环境

* 开发环境

* 测试环境

* 正式环境

## 字段定义

### 返回结果
```
{
    result: Result
    request-id: Request-id
}
```

## 用户中心相关接口

### 用户登录
```
POST /api/user/login
```

#### 请求参数
参数名|类型|是否必须|默认值|描述
---|---|---|---|---
mobile|string|是| |手机号
password|string|是| |密码

#### 响应示例
```
{
    "result":{
        "token":ahasshdqwqqwd
    }
    "request-id":xxx-xxx-xx-xxx-xxx    
}
```

### 用户注销

```
POST /api/user/logout
```

#### 请求参数
参数名|类型|是否必须|默认值|描述
---|---|---|---|---
token|string|是| |token秘钥

#### 响应示例
```
{
    "result":{
        "success":true
    }
    "request-id":xxx-xxx-xx-xxx-xxx    
}
```

### 新建用户
```
POST /api/user
```



### 用户资料信息
```
GET /api/user
```

### 修改用户资料
```
PUT /api/user
```