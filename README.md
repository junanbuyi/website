# api

## xxxx api
method: post  
使用前端使用json表单发送，后端用postform接收;  
使用了 regexp 包进行输入验证，使用了 bcrypt 包对密码进行哈希处理，使用了 sql.ErrNoRows 错误处理返回更具体的错误信息。


## zhuce: saasfxvx
### api1 xxxx:vzczxczx
### api2 xxxxx: /regiest (chuli zhuce luoji)
http: post, 
#### request:
params: xx int, xxx string, xxxx string
SHILI: 
{
    form: "xxx":"xxx",
    "yyy":"yyy"
}
curl: --xxxxx

#### response
respo body
xx int, yy int , zz int
{
    "xxx":"xxx"
    "yyy":"yy"
    zzz:"zz"
}

## 登录页面  ：要有对账户，密码的验证登录功能，设置登录按钮，点击按钮之后跳转到/homepage
### /login处理登录逻辑
- http请求方式：post 

| 参数 | 类型 | 注释 |
| :-: | :-: | :-: |
|  account| string | 用户名|  
|  password| string | 密码|  
```
{
    // 用户名不能是非法字符，位数不能小于8位
    "account":"12345678"，
    //密码不能是非法字符，位数不能小于8位
    "password":"12345678"
}
```  
#### 响应
"message": "注册成功",
| 参数 | 类型 | 注释 |
| :-: | :-: | :-: |
|  message| string | 请求|  
| code | int | 0 success, fail non-0 |
|  error| string| 如果code为零，该字段无效，如果不为零，包含错误信息|  

```
{
    "message":"success",
    "code" : 0,
    "error": nil

}
```

## 首页面:要有个人主页，机器对话，论坛，博客描述等板块，当点击个人主页，页面从/homepage跳转到/personalpage;当点击机器对话页面时，页面从/homepage跳转到/personalpage;当点击论坛时，页面从/homepage跳转到/talkpage;
### /firstpage处理首页面逻辑  
- 
  
    

## 个人主页：要有个人头像从用户自己客户端上传的功能模块，要有本人创建完成的文章预览模块，要有创建，更新和删除文章模块，同时要有自己秘密日记模块
  
  -

##  个人信息模块：包含用户头像，用户个性签名以及修改签名，，用户邮箱账号等  /personalinfo处理页面逻辑
- http请求方式为：post  

| 参数 | 类型 | 注释 |
| :-: | :-: | :-: | 
|  sign| string | 签名|  
| email| string | 邮箱|

```
{
    // 签名不能是非法字符
    "sign":"远方有故乡"，
    //邮件不能是非法字符，位数不能小于8位
    "password":"12345678"
}
```  

## 响应 
| 参数 | 类型 | 注释 |
| :-: | :-: | :-: |
|  message| string | 请求|  
| code | int | 0 success, fail non-0 |
|  error| string| 如果code为零，该字段无效，如果不为零，包含错误信息|  

```
{
    "message":"success",
    "code" : 0,
    "error": nil

}
```
