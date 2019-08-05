# nebula

## Install
> go get github.com/Wall-js/nebula

## Demo

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
	nebula.Service.Init(
		micro.Name("nebula.core.srv.hello"),
	)
	hello.RegisterHelloHandler(nebula.Service.Server(), new(handler.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), new(subscriber.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), subscriber.Handler)

	nebula.Run()
```
> micro api --handler=http