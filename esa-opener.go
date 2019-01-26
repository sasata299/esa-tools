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
    var url string
    if context.Bool("star") {
      url = "https://" + os.Getenv("ESA_TEAM_NAME") + ".esa.io/stars"
    } else {
      url = "https://" + os.Getenv("ESA_TEAM_NAME") + ".esa.io/members/" + context.Args().Get(0)
    }

    browser.OpenURL(url)
    return nil
  }

  app.Flags = []cli.Flag{
    cli.BoolFlag {
      Name: "star, s",
      Usage: "starを付けたページを開く",
    },
  }

  app.Run(os.Args)
}
