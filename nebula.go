package nebula

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
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
	nebulaVersion = "v1.17.4"
	Conf          = config.NewConfig()
	Service       = micro.NewService()
	Web           = web.NewService()
	appId         = ""
	serviceName   = ""
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
	initProcess()
}

func initProcess() {
	var ctx, cancel = context.WithCancel(context.Background())
	if err := loadConfig(cancel); err != nil {
		log.Fatal(err)
	}

	log.Logf("Nebula process start %s", nebulaVersion)
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

func loadConfig(cancel func()) (err error) {
	confPath := "nebula.json"

	// try to use config in runtime
	if _, e := os.Stat("runtime/nebula.json"); e == nil {
		confPath = "runtime/nebula.json"
	}

	if err := Conf.Load(file.NewSource(
		file.WithPath(confPath),
	)); err != nil {
		log.Fatalf("[loadConfig] load error，%s", err)
		return err
	}
	configAddr := Conf.Get("configAddr").String("unknown")
	cluster := Conf.Get("cluster").String("unknown")
	namespace := Conf.Get("namespace").String("unknown")
	_type := Conf.Get("type").String("unknown")
	system := Conf.Get("system").String("unknown")
	version := Conf.Get("version").String("unknown")
	appId := Conf.Get("appId").String("")
	prefix := "/" + cluster + "/" + namespace + "/" + _type + "/" + system + "/" + version
	if appId != "" {
		prefix = prefix + "/" + appId
	}
	//log.Logf("[loadConfig] configAddr: %s", configAddr)
	//log.Logf("[loadConfig] configPrefix: %s", prefix)
	etcdSource := etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress(configAddr),
		// optionally specify prefix; defaults to /micro/conf
		//etcd.WithPrefix("/nebula/nebula-core"),
		//etcd.WithPrefix("/micro/conf/"),
		etcd.WithPrefix(prefix),
		// optionally strip the provided prefix from the keys, defaults to false
		etcd.StripPrefix(true),
	)
	go func() {
		// watch changes
		watcher, err := Conf.Watch()
		if err != nil {
			log.Fatalf("[loadConfig] start watching files error，%s", err)
			return
		}
		v, err := watcher.Next()
		if err != nil {
			log.Fatalf("[loadConfig] watch files error，%s", err)
			return
		}

		log.Logf("[loadConfig] file change， %s", string(v.Bytes()))
		cancel()
	}()

	if err := Conf.Load(etcdSource); err != nil {
		log.Logf("[loadConfig] load error，%s", err.Error())
		//return err
		return nil
	}
	go func() {
		// watch etcd changes
		watcher, err := etcdSource.Watch()
		if err != nil {
			log.Fatalf("[loadConfig] start watching etcd error，%s", err)
			return
		}
		v, err := watcher.Next()
		if err != nil {
			log.Fatalf("[loadConfig] watch etcd error，%s", err)
			return
		}
		log.Logf("[loadConfig] etcd change， %s", string(v.Data))
		cancel()
	}()
	return
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
			initProcess()
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
			initProcess()
		}
	}
}
