package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"go-chassis-demo/schema"
	"log"
)

func main() {
	// 初始化 GoChassis
	chassis.RegisterSchema("rest", &schema.Hello{})
	err := chassis.Init()
	if err != nil {
		log.Fatalf("Failed to initialize GoChassis: %v", err)
	}
	chassis.Run()
}
