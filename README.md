# nebula

#### Install
> go get -u github.com/Wall-js/nebula

#### Demo
```
func main() {
    nebula.Run(func(service micro.Service) {
        service.Init(
            micro.Name("nebula.core.srv.hello"),
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