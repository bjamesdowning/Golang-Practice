package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Test CLI Applet"
	app.Usage = "Lookup IP for particular host"
	app.Version = "0.0.1"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "networktechstudy.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Looks up IP of host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for _, i := range ip {
					fmt.Println(i)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
