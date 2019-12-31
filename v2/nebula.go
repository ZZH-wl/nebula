package nebula

import (
	"context"
	"flag"
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
	Service        = micro.NewService()
	Web            = web.NewService()
	registryAddr   []string
	AppId          string
	dataCenter     string
	configAddr     string
	configKey      string
	serviceVersion string
	serviceName    = ""
	ctx            context.Context
	cancel         func()
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

func SetConfigKey(key string) {
	if configKey == "config/" {
		configKey = key
	}
}

func SetVersion(version string) {
	serviceVersion = version
}

func NewService() micro.Service {
	return micro.NewService()
}

func init() {
	//flag.StringVar(&configAddr, "configAddr", "", "consul Addr")
	flag.StringVar(&dataCenter, "dataCenter", "", "dc1")
	flag.StringVar(&configKey, "configKey", "/nebula/config", "/nebula/config")
	flag.StringVar(&AppId, "appId", "default", "default")
	flag.StringVar(&configAddr, "configAddr", "localhost:8500", "localhost:8500")
	flag.Parse()
}

func CommonProcess() {
	log.Logf("-----Nebula Process Start!-----")

	ctx, cancel = context.WithCancel(context.Background())
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}

	registryAddr = Conf.Get(AppId, "registryAddr").StringSlice([]string{"localhost:8500"})

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
		web.Context(ctx),
		web.Registry(reg),
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
