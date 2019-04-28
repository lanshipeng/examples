package main

import (
	cnf "github.com/micro/examples/config/grpc/config"
	"github.com/micro/go-config"
	grpcConfig "github.com/micro/go-config/source/grpc"
	"github.com/micro/go-log"
)

func main() {
	// create new source
	source := grpcConfig.NewSource(
		grpcConfig.WithAddress("127.0.0.1:8600"),
		grpcConfig.WithPath("/micro"),
	)

	// create new config
	conf := config.NewConfig()

	// load the source into config
	if err := conf.Load(source); err != nil {
		log.Fatal(err)
	}

	// get the config
	configs := &cnf.Micro{}
	if err := conf.Scan(configs); err != nil {
		log.Fatal(err)
	}

	log.Logf("Read config: %s", string(conf.Bytes()))

	// watch the config for changes
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Watching for changes ...")

	for {
		v, err := watcher.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Logf("Watching for changes: %v", string(v.Bytes()))
	}
}