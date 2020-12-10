package main

import (
	"context"
	"log"

	"github.com/shirou/gopsutil/host"
)

func main() {
	i, err := host.InfoWithContext(context.Background())
	if err != nil {
		log.Fatal("error retrieving host info")
	}

	log.Printf("Hostname: %s\n", i.Hostname)
	log.Printf("KernelArch: %s\n", i.KernelArch)
	log.Printf("KernelVersion: %s\n", i.KernelVersion)
	log.Printf("OS: %s\n", i.OS)
	log.Printf("Platform: %s\n", i.Platform)
	log.Printf("PlatformFamily: %s\n", i.PlatformFamily)
	log.Printf("PlatformVersion: %s\n", i.PlatformVersion)
}
