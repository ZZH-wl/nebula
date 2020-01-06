package nebula

import (
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/consul"
	"os"
)

func loadConfig() (err error) {
	var consulSource source.Source
	if dataCenter != "" {
		consulSource = consul.NewSource(
			consul.WithAddress(confAddr),
			consul.WithPrefix(prefix),
			consul.StripPrefix(true),
			consul.WithDatacenter(dataCenter),
		)
	} else {
		consulSource = consul.NewSource(
			consul.WithAddress(confAddr),
			consul.WithPrefix(prefix),
			consul.StripPrefix(true),
		)
	}

	if _, e := os.Stat("runtime/nebula.json"); e == nil {
		if err := Conf.Load(file.NewSource(file.WithPath("runtime/nebula.json"))); err != nil {
			log.Logf("[loadConfig] load error，%s", err.Error())
			return err
		}
	}

	if err := Conf.Load(consulSource); err != nil {
		log.Logf("[loadConfig] load error，%s", err.Error())
	}
	log.Logf("Config %s", string(Conf.Bytes()))

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
