package main

import (
	app "github.com/dynamicgo/mesh-app"
	_ "github.com/dynamicgo/mesh-libp2p-network"
	_ "github.com/dynamicgo/mesh-servicehub/load"
)

func main() {
	app.Run()
}
