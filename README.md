# nebula

## Install
> go get github.com/Wall-js/nebula

## Demo

#### Starting
```
func main() {
    nebula.Run(func(service micro.Service) {
        service.Init(
            micro.Name("nebula.core.srv.hello"),
            micro.Version("latest"),
        )
    })
    
    // Register Handler
    hello.RegisterHelloHandler(service.Server(), new(handler.Hello))
    // Register Struct as Subscriber
    micro.RegisterSubscriber("nebula.core.srv.hello", service.Server(), new(subscriber.Hello))
    // Register Function as Subscriber
    micro.RegisterSubscriber("nebula.core.srv.hello", service.Server(), subscriber.Handler)
}
```

#### Config

```
	version := Conf.Get("version").String("unknown")
```

