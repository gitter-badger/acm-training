package main

import (
	"fmt"
	"github.com/uestc-acm/acm-training/config"
	"github.com/uestc-acm/acm-training/router"
	"os"
)

func main() {
	config.InitCommandLineFlags()
	if config.VersionPresent() {
		fmt.Printf("acm-training version %s\n", config.Version)
		os.Exit(0)
	}

	r := router.Create()
	r.Run()
}
