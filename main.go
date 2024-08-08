package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	// Read the logs from the file
	data, err := ioutil.ReadFile("karmada-scheduler-master-metrics.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	logs := string(data)

	// Define the regex pattern to find the specific metric
	re := regexp.MustCompile(`karmada_scheduler_schedule_attempts_total{.*}\s(\d+)`)

	matches := re.FindAllStringSubmatch(logs, -1)

	if len(matches) > 0 {
		fmt.Println("karmada_scheduler_schedule_attempts_total value:", matches[0][1])
	} else {
		fmt.Println("No matches found")
	}
}
