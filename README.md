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
    
*使用swagger自动更新API文档

    go get -u github.com/swaggo/swag/cmd/swag
    
更新API文档

    swag init
    
## handlerHelper

`handlerHelper`是一个自动把`handlerFunc`添加到`Gin`并且设置路由的工具包

1. 在handler文件夹中添加`main.go`文件
```
type helper struct {
}

func Build(r gin.IRoutes) {
    h := new(helper)
    valueOfh := reflect.ValueOf(h)
    numMethod := valueOfh.NumMethod()
    for i := 0; i < numMethod; i++ {
        rt := valueOfh.Method(i).Call(nil)[0].Interface().(*handlerHelper.Router)
        rt.AddHandler(r)
    }
}
```  
2. 之后每次写`handlerFunc`的时候都类似下方的`helloHandler`前面加上一个
`*helper`的一个方法`HelloHandler()`中设置路由。

```
func (h *helper) HelloHandler() (r *handlerHelper.Router) {
    return &handlerHelper.Router{
        Path:   "/HelloHandler",
        Method: "GET",
        Handlers: []gin.HandlerFunc{
            helloHandler,
        }}
}
func helloHandler(c *gin.Context) {

    c.String(http.StatusOK, "Hello world!")
}
```