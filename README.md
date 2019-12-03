# nebula

## Install
> go get github.com/Wall-js/nebula
> export GOPROXY=https://goproxy.io
## Demo

## 服务配置规范
####节点命名
```
/集群名/命名空间/类型(srv,api,web,evt)/系统/版本(/id)(括号部分可不使用)

/default/nebula/srv/nebula-core/v0.1(/10086)

/cluster/namespace/type/system/version(/appId)
```

#### Type
```
web,api,srv,evt
```

#### Starting
```
	nebula.Service.Init(
		micro.Name("nebula.core.srv.hello"),
	)
	hello.RegisterHelloHandler(nebula.Service.Server(), new(handler.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), new(subscriber.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), subscriber.Handler)

	nebula.Run()
```

#### Config

```
version := nebula.Conf.Get("version").String("unknown")
```

#### Web
```
func main() {
	nebula.Web.Init(
		web.Name("nebula.core.greeter"),
	)

	router := gin.New()
	say := new(Say)
	router.GET("/greeter", say.Anything)

	// Register Handler
	nebula.Web.Handle("/", router)
	nebula.RunWeb()
}
```
> micro api --handler=http