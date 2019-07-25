package nebula

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/etcd"
)

var Conf = config.NewConfig()

func loadConfig(cancel func()) (err error) {
	if err := Conf.Load(file.NewSource(
		file.WithPath("./nebula.json"),
	)); err != nil {
		log.Fatalf("[loadConfig] load error，%s", err)
		return err
	}
	configAddress := Conf.Get("configAddress").String("unknown")
	cluster := Conf.Get("cluster").String("unknown")
	namespace := Conf.Get("namespace").String("unknown")
	system := Conf.Get("system").String("unknown")
	version := Conf.Get("version").String("unknown")
	appId := Conf.Get("appId").String("")
	prefix := "/" + cluster + "/" + namespace + "/" + system + "/" + version
	if appId != "" {
		prefix = prefix + "/" + appId
	}
	log.Logf("[loadConfig] configAddress: %s", configAddress)
	log.Logf("[loadConfig] configPrefix: %s", prefix)
	etcdSource := etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress(configAddress),
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

func Run(f func(service micro.Service)) {
	for {
		var ctx, cancel = context.WithCancel(context.Background())

		if err := loadConfig(cancel); err != nil {
			log.Fatal(err)
		}
		//go func() {
		//	<-time.After(time.Second * 5)
		//	log.Logf("Shutdown example: shutting down service")
		//	cancel()
		//}()
		service := micro.NewService()

		f(service)
		log.Log("[service] Register service")

		service.Init(
			micro.Context(ctx),
		)

		if err := service.Server().Init(
			server.Wait(nil),
		); err != nil {
			log.Fatal(err)
		}

		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
		log.Log("[service] restart service")
	}
}
