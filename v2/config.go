package nebula

import (
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/consul"
)

func loadConfig() (err error) {
	var consulSource source.Source
	if dataCenter != "" {
		consulSource = consul.NewSource(
			consul.WithAddress(configAddr),
			consul.WithPrefix(configKey),
			consul.StripPrefix(true),
			consul.WithDatacenter(dataCenter),
		)
	} else {
		consulSource = consul.NewSource(
			consul.WithAddress(configAddr),
			consul.WithPrefix(configKey),
			consul.StripPrefix(true),
		)
	}

	if err := Conf.Load(consulSource); err != nil {
		log.Logf("[loadConfig] load error，%s", err.Error())
		return err
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
	//go func() {
	//	// watch etcd changes
	//	watcher, err := consulSource.Watch()
	//	if err != nil {
	//		log.Fatalf("[loadConfig] start watching consul error，%s", err)
	//		return
	//	}
	//	v, err := watcher.Next()
	//	if err != nil {
	//		log.Fatalf("[loadConfig] watch consul error，%s", err)
	//		return
	//	}
	//	log.Logf("[loadConfig] consul change， %s", string(v.Data))
	//	cancel()
	//}()
	return
}
