package main

import (
	"fmt"
	"log"
	"os"
	"plugin"

	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/app"
	"github.com/babakgh/kata/design_patterns/abstract_factory/go/pkg/core"
)

func main() {
	factory := loadFactory()
	app.App.Start(factory)
}

func loadPlugin(version string) plugin.Symbol {
	mod := fmt.Sprintf("./bin/%s.so", version)

	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Printf("mod %v is loaded", version)

	symfactory, err := plug.Lookup("Factory")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return symfactory
}

func loadFactory() core.ShapeFactory {
	version := "v1"
	if len(os.Args) >= 2 {
		version = os.Args[1]
	}

	factorySymbol := loadPlugin(version)

	factory, ok := factorySymbol.(core.ShapeFactory)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	return factory
}
