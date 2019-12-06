# GoSpree
开发中
## 相关库
* gorm https://gorm.io/
* gin https://gin-gonic.com/
* gin-swagger https://github.com/swaggo/gin-swagger
## 目录结构
>GoSpree
>>core *//核心库*
>>
>>service *//执行相关业务*
>>
>>dao *//与数据库相关的操作*
>>
>>docs *//swagger自动生成的API文档*
>>
>>model *//各种类型model*
>>
>>restAPI *//REST API*
>>>handler *//HandlerFunc*
>>>
>>>middleware *//中间件*
>>>
>>util *//工具包*
>>
>main.go *//程序入口*
>
>config.json *//配置文件*

## 安装注意
    
* 使用swagger自动更新API文档
```
go get -u github.com/swaggo/swag/cmd/swag
``` 
* 更新API文档

```
swag init
```