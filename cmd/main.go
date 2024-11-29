package main

import (
	"flag"
	"rwg/pkg"
	"time"
)

func main() {
	serivename := flag.String("service", "wg-quick@wg0", "Service name to restart")
	interfaceName := flag.String("interface", "wg0", "Interface name to restart")
	SecondInterval := flag.Int64("interval", 5, "Interval to check the service")
	service := flag.Bool("service", false, "Service to restart")
	ipCheck := flag.String("ip", "", "IP to check the connection")
	flag.Parse()
	interval := time.Duration(*SecondInterval) * time.Second

	if *ipCheck == "" {
		panic("IP to check is required")
	}

	pkg.CheckOut(interval, *ipCheck, *serivename, *interfaceName, *service)

}
