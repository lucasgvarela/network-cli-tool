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
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"
	app.UsageText = "cli [COMMAND] --host [HOSTNAME]"
	app.Version = "0.1"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := range ns {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks Up the IP addresses for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for i := range ip {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cn",
			Usage: "Looks Up the CNAME for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cn, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}
				fmt.Println(cn)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks Up the MX records for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}
				for i := range mx {
					fmt.Println(mx[i].Host, mx[i].Pref)
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
