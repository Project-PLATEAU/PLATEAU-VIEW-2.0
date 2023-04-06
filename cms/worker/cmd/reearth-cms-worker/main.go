package main

import "github.com/reearth/reearth-cms/worker/internal/app"

var version = ""

func main() {
	app.Start(debug, version)
}
