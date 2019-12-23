# nebula

## Version
go版本：go 1.13 
micro版本：micro 自建镜像1.17.2  
nebula版本：nebula tag版本 v1.17.4

## Install
> go get github.com/Wall-js/nebula  
> export GOPROXY=https://goproxy.cn,direct
golang替换国源
```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=*.hiqio.com,*.gitlab.com,*.gitee.com //跳过私有库
```
## Demo

## 服务配置规范
####节点命名
```
/集群名/命名空间/类型(srv,api,web,evt)/系统/版本(/id)(括号部分可不使用)

/default/nebula/srv/nebula-core/v0.1(/nebula.core.srv)

/cluster/namespace/type/system/version(/appId)
```

#### Type
页面web,接口api,服务srv,事件evt

#### Starting
```
func main() {
	nebula.SetName("nebula.core.srv")
	hello.RegisterHelloHandler(nebula.Service.Server(), new(handler.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), new(subscriber.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), subscriber.Handler)
	nebula.Run()
}
```

#### Config
```
    version := nebula.Conf.Get("version").String("unknown")
    name := nebula.Conf.Get("config","name").String("unknown")
```

#### Web
```
func main() {
	nebula.SetName("go.micro.api.greeter")
	router := gin.New()
	say := new(Say)
	router.GET("/greeter", say.Anything)
	// Register Handler
	nebula.Web.Handle("/", router)
	nebula.RunWeb()
}
```
