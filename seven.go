package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var workoutTime = 60
var breakTime = 15

func main() {
	workoutName := flag.String("w", "seven", "workout mode")
	flag.Parse()
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = fmt.Sprintf("%v/.config", os.Getenv("HOME"))
	}
	configFile, err := os.Open(fmt.Sprintf("%v/seven/%v.conf", configHome, *workoutName))
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	scanner := bufio.NewScanner(configFile)

	for scanner.Scan() {
		exercise := scanner.Text()
		if strings.HasPrefix(exercise, "#") {
			continue
		}
		for i := workoutTime - 1; i >= 0; i-- {
			fmt.Printf("\r%v -- %v seconds remaining\033[K", exercise, i)
			time.Sleep(time.Second)
		}
		for i := breakTime - 1; i >= 0; i-- {
			fmt.Printf("\rBreak time -- %v seconds remaining\033[K", i)
			time.Sleep(time.Second)
		}
	}
	fmt.Println()
}
