package nebula

import (
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/consul"
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

	configAddr = Conf.Get("configAddr").String(":8500")

	consulSource := consul.NewSource(
		consul.WithAddress(configAddr),
		consul.WithPrefix(configKey),
		consul.StripPrefix(true),
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

	if err := Conf.Load(consulSource); err != nil {
		log.Logf("[loadConfig] load error，%s", err.Error())
		//return err
		return nil
	}

	go func() {
		// watch etcd changes
		watcher, err := consulSource.Watch()
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
