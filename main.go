package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	maxPercentPtr := flag.Float64("max-percent", 0.0, "maximum CPU percentage")
	periodPtr := flag.Duration("period", 5*time.Second, "period between checks")
	allowedNamesPtr := flag.String("allowed-names", "", "comma-separated list of allowed process names")
	flag.Parse()

	if *maxPercentPtr == 0.0 {
		log.Fatal("Usage: autokill -max-percent <max-percent> -period <period> -allowed-names <allowed-names>")
	}

	var allowedNames []string
	if *allowedNamesPtr != "" {
		allowedNames = strings.Split(*allowedNamesPtr, ",")
	}

	for range time.Tick(*periodPtr) {
		log.Println("Checking...")
		handle(*maxPercentPtr, allowedNames)
	}
}

func handle(maxPercent float64, allowedNames []string) {
	processes, err := process.Processes()
	if err != nil {
		log.Println(err)
		return
	}
	for _, p := range processes {
		percent, err := p.CPUPercent()
		if err != nil {
			log.Println(err)
			continue
		}
		if percent <= maxPercent {
			continue
		}

		name, err := p.Name()
		if err != nil {
			log.Println(err)
			continue
		}

		if isAllowed(name, allowedNames) {
			continue
		}

		log.Printf("Killing %s, CPU: %f\n", name, percent)
		if err := p.Kill(); err != nil {
			log.Println(err)
		}
	}
}

func isAllowed(name string, allowedNames []string) bool {
	for _, allowedName := range allowedNames {
		if name == allowedName {
			return true
		}
	}
	return false
}
