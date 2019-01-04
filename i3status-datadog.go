package main

import (
	"log"
	"github.com/zorkian/go-datadog-api"
	"fmt"
	"flag"
)

const (
	Green = "#008000"
	Yellow = "#FFFF00"
	Grey = "#808080"
	Red = "#FF0000"
)

type monitorCounter struct {
	Ok     int
	Warn   int
	Alert  int
	NoData int
}

func main() {

	apiKey := flag.String("apiKey", "", "Datadog API Key")
	appKey := flag.String("appKey", "", "Datadog APP Key")
	flag.Parse()

	client := datadog.NewClient(*apiKey, *appKey)

	monitors, err := client.GetMonitors()
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	monitorCounter := monitorCounter{
		Ok: 0,
		Warn: 0,
		Alert: 0,
		NoData: 0,
	}

	for _, monitor := range monitors {
		switch *monitor.OverallState {
		case "OK":
			monitorCounter.Ok++
		case "No Data":
			monitorCounter.NoData++
		case "Warn":
			monitorCounter.Warn++
		case "Alert":
			monitorCounter.Alert++
		}
	}

	fmt.Printf("%vN %vW %vA %vT", monitorCounter.NoData, monitorCounter.Warn, monitorCounter.Alert, len(monitors))

	switch {
	case monitorCounter.Alert >  0:
		fmt.Printf("\n\n%v", Red)
		break
	case monitorCounter.Warn >  0:
		fmt.Printf("\n\n%v", Yellow)
		break
	case monitorCounter.NoData >  0:
		fmt.Printf("\n\n%v", Grey)
		break
	default:
		fmt.Printf("\n\n%v", Green)
	}
}
