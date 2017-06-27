package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

var (
	ConKey    = os.Getenv("CON_KEY_TW")        //Twitter-Consumer-Key
	ConSecKey = os.Getenv("CON_SECRET_KEY_TW") //Twitter-Consumer-Secret-Key
	AccKey    = os.Getenv("ACC_KEY_TW")        //Twitter-Access-Key
	AccSecKey = os.Getenv("ACC_SECRET_KEY_TW") //Twitter-Access-Secret-Key
  　api       *anaconda.TwitterApi
)

func Set() {
	anaconda.SetConsumerKey(ConKey)
	anaconda.SetConsumerSecret(ConSecKey)
	api = anaconda.NewTwitterApi(AccKey, AccSecKey)
}

func init() {
	Set() //keyの設定
	if len(ConKey) == 0 {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter consmer key")
		os.Exit(1)
	}
	if len(ConSecKey) == 0 {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter secret consumer key")
		os.Exit(1)
	}
	if len(AccKey) == 0 {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter access token")
		os.Exit(1)
	}
	if len(AccSecKey) == 0 {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+"Twitter access secret key")
		os.Exit(1)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Twitter CLI"
	app.Usage = "Command line interface for Twitter"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "tweet",
			Aliases: []string{"tw", "t"},
			Usage:   "Tweet sentence",
			Action:  tweet,
		},
  }
  app.Run(os.Args)
}
