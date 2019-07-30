package nebula

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/etcd"
	"github.com/micro/go-plugins/registry/etcdv3"
	"time"
)

var (
	Conf    = config.NewConfig()
	Service micro.Service
)

func loadConfig(cancel func()) (err error) {
	if err := Conf.Load(file.NewSource(
		file.WithPath("./nebula.json"),
	)); err != nil {
		log.Fatalf("[loadConfig] load error，%s", err)
		return err
	}
	configAddr := Conf.Get("configAddr").String("unknown")
	cluster := Conf.Get("cluster").String("unknown")
	namespace := Conf.Get("namespace").String("unknown")
	system := Conf.Get("system").String("unknown")
	version := Conf.Get("version").String("unknown")
	appId := Conf.Get("appId").String("")
	prefix := "/" + cluster + "/" + namespace + "/" + system + "/" + version
	if appId != "" {
		prefix = prefix + "/" + appId
	}
	log.Logf("[loadConfig] configAddr: %s", configAddr)
	log.Logf("[loadConfig] configPrefix: %s", prefix)
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

	if err := Conf.Load(etcdSource); err != nil {
		log.Fatalf("[loadConfig] load error，%s", err)
		return err
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
	return
}

func Run(f func(service micro.Service)) micro.Service {
	for {
		log.Log("[service] starting service")

		var ctx, cancel = context.WithCancel(context.Background())

		if err := loadConfig(cancel); err != nil {
			log.Fatal(err)
		}
		version := Conf.Get("version").String("unknown")

		reg := etcdv3.NewRegistry(func(op *registry.Options) {
			//op.Addrs = []string{"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",}
			op.Addrs = Conf.Get("registryAddr").StringSlice([]string{"localhost:2379"})
		})

		Service = micro.NewService()

		f(Service)

		Service.Init(
			micro.Context(ctx),
			micro.Registry(reg),
			micro.Version(version),
			micro.RegisterTTL(time.Second*30),
			micro.RegisterInterval(time.Second*15),
		)
		//graceful shutdown
		if err := Service.Server().Init(
			server.Wait(nil),
		); err != nil {
			log.Fatal(err)
		}

		log.Log("[service] service options: %s", Service.Options())
		log.Log("[service] server options: %s", Service.Server().Options())

		//service start
		if err := Service.Run(); err != nil {
			log.Fatal(err)
		}
		log.Log("[service] ending service: %s", Service.String())

	}
}
