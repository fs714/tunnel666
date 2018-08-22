package main

import (
	"flag"
	"fmt"
	"github.com/fs714/tunnel666/utils/log"
	"os"
)

var BuildTime string
var AppVersion = "0.0.1 build on " + BuildTime

var (
	version        bool
	logLevel       string
	localTunAddr   string
	remoteTunAddr  string
	remoteHostAddr string
	remoteHostPort int
)

func init() {
	flag.BoolVar(&version, "v", false, "Version")
	flag.StringVar(&logLevel, "-loglevel", "info", "Log Level")
	flag.StringVar(&localTunAddr, "l", "10.6.66.1/30", "Local Tunnel Address like 10.6.66.1/30")
	flag.StringVar(&remoteTunAddr, "r", "10.6.66.2/30", "Remote Tunnel Address like 10.6.66.2/30")
	flag.StringVar(&remoteHostAddr, "rh", "", "Remote Host IP Address")
	flag.IntVar(&remoteHostPort, "rp", 0, "Remote Host Port")
	flag.Parse()

	if version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	log.SetLevel(logLevel)
	log.SetFormat("text")
	log.SetOutput(os.Stdout)
}

func main() {
	log.Info("Start Tunnel666...")
}
