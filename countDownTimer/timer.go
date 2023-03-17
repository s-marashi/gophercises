package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type countDown struct {
	t int
	d int
	h int
	m int
	s int
}

func main() {
	deadline := flag.String(
		"deadline",
		"",
		"The deadline for the countdown timer in DateTime format (e.g. 2019-12-25 15:00:00)",
	)
	flag.Parse()

	if *deadline == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	v, err := time.ParseInLocation(time.DateTime, *deadline, time.Local)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeDifference(v, time.Now())
		if timeRemaining.t <= 0 {
			fmt.Println("Countdown reached!")
			break
		}

		fmt.Printf(
			"Days: %d Hours: %d Minutes: %d Seconds: %d\n",
			timeRemaining.d,
			timeRemaining.h,
			timeRemaining.m,
			timeRemaining.s,
		)
	}
}

func getTimeDifference(t1 time.Time, t2 time.Time) countDown {
	difference := t1.Sub(t2)
	total := int(difference.Seconds())

	return countDown{
		t: total,
		d: total / (60 * 60 * 24),
		h: total / (60 * 60) % 24,
		m: (total / 60) % 60,
		s: total % 60,
	}
}
