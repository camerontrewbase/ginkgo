package main

import (
	"github.com/camerontrewbase/ginkgo/internal/app"
	"github.com/camerontrewbase/ginkgo/internal/controller"
)

func main() {
	app.NewApp(controller.NewServer()).Start()
}
