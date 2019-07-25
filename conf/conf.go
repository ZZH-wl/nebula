package conf

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-plugins/config/source/etcd"
)

func run() {
	//e := toml.NewEncoder()
	//e := json.NewEncoder()
	//conf := conf.NewConfig()
	etcdSource := etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress("localhost:2379"),
		// optionally specify prefix; defaults to /micro/conf
		//etcd.WithPrefix("/nebula/nebula-core"),
		//etcd.WithPrefix("/micro/conf/"),
		etcd.WithPrefix("/test/conf/database/v0.1"),
		// optionally strip the provided prefix from the keys, defaults to false
		etcd.StripPrefix(true),
		//source.WithEncoder(e),
	)
	if err := config.Load(etcdSource); err != nil {
		fmt.Print(err)
	}
	//database := conf.Get("conf","database","address").String("none")
	//database := conf.Get("conf","database","address").String("none")
	//database := conf.Get("test","conf","database","port").String("unknown")
	conf := config.Map()
	fmt.Print(conf)
	//database := conf.Get("/micro/conf/database/address").String("none")
	//fmt.Print(conf.Get("database","address").String("unknown"))
	w, err := config.Watch("address")
	if err != nil {
	}
	_, err = w.Next()
	if err != nil {
	}

	//fileSource := file.NewSource(file.WithPath("./conf/conf.json"))
	//if err := conf.Load(fileSource); err != nil {
	//	fmt.Print(err)
	//}
	//conf.LoadFile("./conf/conf.json")
	//address := conf.Get("hosts", "database", "address").String("err\n")

}
