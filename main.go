package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"time"
)

func main() {

	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%msg%\n",
	})

	if os.Args[1] == "--version" {
		logrus.Println("0.0.1")
		os.Exit(0)
	}

	if os.Args[1] == "--help" {
		logrus.Println("Usage: ./sum startTime lunchDuration endTime")
		logrus.Println("Example: ./sum 08:30 0:42 17:24")
		os.Exit(0)
	}

	if len(os.Args) != 4 {
		logrus.Errorf("sum requires exactly three args: HH:mm HH:mm HH:mm (start lunch-duration end)")
		os.Exit(1)
	}

	start := os.Args[1]
	lunch := os.Args[2]
	end := os.Args[3]

	parsedStart, err := time.Parse("15:04", start)
	checkError(err)

	parsedEnd, err := time.Parse("15:04", end)
	checkError(err)

	parsedLunch, err := time.Parse("15:04", lunch)
	checkError(err)

	startMinute := parsedStart.Hour()*60 + parsedStart.Minute()
	endMinute := parsedEnd.Hour()*60 + parsedEnd.Minute()
	lunchMinute := parsedLunch.Hour()*60 + parsedLunch.Minute()

	workedMinutes := endMinute - lunchMinute - startMinute
	fmt.Printf("%s:%s\n", padZero(workedMinutes/60), padZero(workedMinutes%60))
}

func padZero(i int) string {
	if i < 10 {
		return fmt.Sprintf("0%d", i)
	}
	return fmt.Sprintf("%d", i)
}

func checkError(err error) {
	if err != nil {
		logrus.Errorf("error: %v\n", err)
		os.Exit(1)
	}
}
