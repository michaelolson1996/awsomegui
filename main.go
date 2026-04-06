package main

import (
  "embed"

  "github.com/wailsapp/wails/v2"
  "github.com/wailsapp/wails/v2/pkg/options"
  "github.com/wailsapp/wails/v2/pkg/options/assetserver"

  "awsomegui/backend/app"
  "awsomegui/backend/services/session"
  awsauth "awsomegui/backend/services/auth/awsauth"
  awsclient "awsomegui/backend/services/aws/client"
  vpcsvc "awsomegui/backend/services/aws/vpc"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

  workspace := session.NewWorkspaceManager()
  factory := awsclient.NewClientFactory(workspace)
  vpcService := vpcsvc.NewVPCService(factory)
  profileService := awsauth.NewProfileService()

  // Create an instance of the app structure
  app := app.NewApp(
    workspace,
    profileService,
    vpcService,
  )

  // Create application with options
  err := wails.Run(&options.App{
    Title:  "awsomegui",
    Width:  1024,
    Height: 768,
    AssetServer: &assetserver.Options{
      Assets: assets,
    },
    BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
    OnStartup:        app.Startup,
    Bind: []interface{}{
      app,
    },
  })

  if err != nil {
    println("Error:", err.Error())
  }
}
