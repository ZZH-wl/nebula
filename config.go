package nebula

import (
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"os"
)

func loadConfig() (err error) {
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
