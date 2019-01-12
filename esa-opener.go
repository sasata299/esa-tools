package main

import (
    "os"
    "github.com/urfave/cli"
    "github.com/pkg/browser"
)

func main() {
  app := cli.NewApp()

  app.Name = "esa-opener"
  app.Usage = "screen_nameを引数で指定する"
  app.Version = "0.0.1"

  app.Action = func (context *cli.Context) error {
    var url = "https://" + os.Getenv("ESA_TEAM_NAME") + ".esa.io/members/" + context.Args().Get(0)
    browser.OpenURL(url)
    return nil
  }

  app.Run(os.Args)
}
