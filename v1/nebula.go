package nebula

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Version     = "v1.17.5"
	Conf        = config.NewConfig()
	Service     = micro.NewService()
	Web         = web.NewService()
	appId       = ""
	serviceName = ""
	ctx         context.Context
	cancel      func()
)

func SetName(name string) {
	serviceName = name
	Service.Init(
		micro.Name(name),
	)
	if err := Web.Init(
		web.Name(name),
	); err != nil {
		log.Fatal(err)
	}
}

func init() {
	InitProcess()
}

func InitProcess() {
	ctx, cancel = context.WithCancel(context.Background())
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}

	log.Logf("Nebula process start %s", Version)
	//if nebulaVersion != Conf.Get("nebulaVersion").String("unknown") {
	//	log.Fatalf("Nebula config version error:%s", Conf.Get("nebulaVersion").String("unknown"))
	//}

	//serviceType = Conf.Get("type").String("unknown")
	appId = Conf.Get("appId").String("unknown")
	version := Conf.Get("version").String("unknown")

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		//op.Addrs = []string{"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",}
		op.Addrs = Conf.Get("registryAddr").StringSlice([]string{"localhost:2379"})
	})
	name := appId
	if serviceName != "" {
		name = serviceName
	}
	//graceful shutdown
	if err := Service.Server().Init(
		server.Wait(nil),
	); err != nil {
		log.Fatal(err)
	}
	if err := Web.Options().Service.Server().Init(
		server.Wait(nil),
	); err != nil {
		log.Fatal(err)
	}
	// config service
	Service.Init(
		micro.Name(name),
		micro.Context(ctx),
		micro.Registry(reg),
		micro.Version(version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	if err := Web.Init(
		//web.MicroService(Service),
		web.Name(name),
		web.Context(ctx),
		web.Registry(reg),
		web.Version(version),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	); err != nil {
		log.Fatal(err)
	}

}

func NewService() micro.Service {
	return micro.NewService()
}

func Run() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	for {
		log.Log("[service] init service")
		log.Log("[service] service options: ", Service.Options())
		log.Log("[service] server options: ", Service.Server().Options())
		//service start
		if err := Service.Run(); err != nil {
			log.Fatal(err)
		}
		select {
		// wait on kill signal
		case <-ch:
			log.Log("[service] ending service: ", Service.Server().String())
			return
		// wait on context cancel
		default:
			log.Log("[service] restart service: ", Service.Server().String())
			InitProcess()
		}
	}
}

func RunWeb() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	for {
		log.Log("[service] init service")
		log.Log("[service] service options: ", Service.Options())
		log.Log("[service] server options: ", Service.Server().Options())
		//service start
		if err := Web.Run(); err != nil {
			log.Fatal(err)
		}
		select {
		// wait on kill signal
		case <-ch:
			log.Log("[service] ending service: ", Service.Server().String())
			return
		// wait on context cancel
		default:
			log.Log("[service] restart service: ", Service.Server().String())
			InitProcess()
		}
	}
}

func RunProcess(process func(context.Context, context.CancelFunc) error) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	for {
		log.Log("[service] init service")
		log.Log("[service] service options: ", Service.Options())
		log.Log("[service] server options: ", Service.Server().Options())
		//service start
		if err := process(ctx, cancel); err != nil {
			log.Fatal(err)
		}
		select {
		// wait on kill signal
		case <-ch:
			log.Log("[service] ending service: ", Service.Server().String())
			return
		// wait on context cancel
		default:
			log.Log("[service] restart service: ", Service.Server().String())
			InitProcess()
		}
	}
}
