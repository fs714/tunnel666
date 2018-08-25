package main

import (
	"flag"
	"fmt"
	"github.com/fs714/tunnel666/utils/exec"
	"github.com/fs714/tunnel666/utils/log"
	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
	"net"
	"os"
	"strconv"
	"strings"
)

var BuildTime string
var AppVersion = "0.0.1 build on " + BuildTime

var (
	version        bool
	logLevel       string
	localTunAddr   string
	localHostAddr  string
	remoteHostAddr string
	ifaceName      string
)

func init() {
	flag.BoolVar(&version, "v", false, "Version")
	flag.StringVar(&logLevel, "loglevel", "info", "Log Level")
	flag.StringVar(&localTunAddr, "lt", "10.6.66.1/30", "Local Tunnel Address like 10.6.66.1/30")
	flag.StringVar(&localHostAddr, "l", "0.0.0.0:6666", "Local Host Listening IP and Port like 0.0.0.0:6666")
	flag.StringVar(&remoteHostAddr, "r", "", "Remote Host Address like 66.6.6.66:6666")
	flag.StringVar(&ifaceName, "i", "tunnel666", "Local Tunnel Interface Name")
	flag.Parse()

	if version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	log.SetLevel(logLevel)
	log.SetFormat("text")
	log.SetOutput(os.Stdout)
}

func argsValidate(localTunAddr string, localHostAddr string, remoteHostAddr string) (err error) {
	_, _, err = net.ParseCIDR(localTunAddr)
	if err != nil {
		log.Error("Invalid Local Tunnel IP Address: ", localTunAddr)
		return
	}

	la := strings.Split(localHostAddr, ":")
	localHostIp := net.ParseIP(la[0])
	if localHostIp == nil {
		log.Error("Invalid Local Host IP Address: ", la[0])
		return
	}
	_, err = strconv.Atoi(la[1])
	if err != nil {
		log.Error("Invalid Local Host Port: ", la[1])
		return
	}

	ra := strings.Split(remoteHostAddr, ":")
	remoteHostIp := net.ParseIP(ra[0])
	if remoteHostIp == nil {
		log.Error("Invalid Remote Host IP Address: ", ra[0])
		return
	}
	_, err = strconv.Atoi(ra[1])
	if err != nil {
		log.Error("Invalid Remote Host Port: ", ra[1])
		return
	}

	return
}

func createIface(localTunAddr string) (iface *water.Interface, err error) {
	ifaceConf := water.Config{
		DeviceType: water.TAP,
	}
	ifaceConf.Name = ifaceName
	iface, err = water.New(ifaceConf)
	if err != nil {
		return
	}

	_, err = exec.ExecCommand("ip addr add " + localTunAddr + " dev " + ifaceConf.Name)
	if err != nil {
		return
	}

	_, err = exec.ExecCommand("ip link set dev " + ifaceConf.Name + " up")
	if err != nil {
		return
	}

	return
}

func main() {
	log.Info("Start Tunnel666...")

	err := argsValidate(localTunAddr, localHostAddr, remoteHostAddr)
	if err != nil {
		return
	}

	iface, err := createIface(localTunAddr)
	if err != nil {
		log.Error(err.Error())
		return
	}

	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := iface.Read([]byte(frame))
		if err != nil {
			log.Error(err)
		}
		frame = frame[:n]
		log.Info("SRC: ", frame.Source())
		log.Info("DST: ", frame.Destination())
		log.Info(fmt.Sprintf("Type: % x", frame.Ethertype()))
		log.Info(fmt.Sprintf("Payload: % x", frame.Payload()))
	}
}
