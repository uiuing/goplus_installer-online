package main

import (
  "embed"

  "github.com/wailsapp/wails/v2"
  "github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
  app := NewApp()

  err := wails.Run(&options.App{
    Title:         "GoPlus Online Installation",
    DisableResize: true,
    Width:         800,
    Height:        500,
    Assets:        assets,
    OnStartup:     app.startup,
    Bind: []interface{}{
      app,
    },
  })

  if err != nil {
    println("Error:", err)
  }
}
