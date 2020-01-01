package nebula

import (
	"context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	Conf           = config.NewConfig()
	Service        micro.Service
	Web            = web.NewService()
	prefix         string
	registryAddr   []string
	AppId          string
	dataCenter     string
	confAddr       string
	serviceVersion string
	serviceName    = "unknown"
	ctx            context.Context
	cancel         func()
)

func SetName(s string) {
	serviceName = s
	Service.Init(
		micro.Name(s),
	)
	if err := Web.Init(
		web.Name(s),
	); err != nil {
		log.Fatal(err)
	}
}

func SetPrefix(s string) {
	if prefix == "/nebula/config" {
		prefix = s
	}
}

func SetVersion(version string) {
	serviceVersion = version
}

func init() {
	Service = micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "dataCenter",
				Usage: "dataCenter dc1",
				Value: "",
			},
			cli.StringFlag{
				Name:  "prefix",
				Usage: "prefix /nebula/config",
				Value: "/nebula/config",
			},
			cli.StringFlag{
				Name:  "confAddr",
				Usage: "confAddr localhost:8500",
				Value: "localhost:8500",
			},
			cli.StringFlag{
				Name:  "prefix",
				Usage: "prefix /test",
				Value: "",
			},
			cli.StringFlag{
				Name:  "appId",
				Usage: "appId default",
				Value: "default",
			},
		),
		micro.Action(func(c *cli.Context) {
			dataCenter = c.String("dataCenter")
			prefix = c.String("prefix")
			confAddr = c.String("confAddr")
			AppId = c.String("appId")
		}),
	)
}

func CommonProcess() {
	log.Logf("-----Nebula Process Start!-----")
	ctx, cancel = context.WithCancel(context.Background())
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}

	registryAddr = Conf.Get("default", "registryAddr").StringSlice([]string{"localhost:8500"})

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = registryAddr
	})

	//graceful shutdown
	wg := new(sync.WaitGroup)

	if err := Service.Server().Init(
		server.Wait(wg),
	); err != nil {
		log.Fatal(err)
	}

	if err := Web.Options().Service.Server().Init(
		server.Wait(wg),
	); err != nil {
		log.Fatal(err)
	}

	// config service
	Service.Init(
		micro.Name(serviceName),
		micro.Context(ctx),
		micro.Registry(reg),
		micro.Version(serviceVersion),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	if err := Web.Init(
		web.Name(serviceName),
		// Alternative Options
		web.Registry(reg),
		web.MicroService(Service),
		web.Context(ctx),
		web.Version(serviceVersion),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	); err != nil {
		log.Fatal(err)
	}
}

func Run() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	for {
		CommonProcess()

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
		}
	}
}

func RunWeb() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	for {
		CommonProcess()
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
		}
	}
}

func RunProcess(process func(context.Context, context.CancelFunc) error) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	for {
		CommonProcess()

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
		}
	}
}
