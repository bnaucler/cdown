package main

import (
	"fmt"
	"flag"
	"os"
	"os/signal"
	"strings"
	"strconv"
	"os/exec"
	"time"
	"github.com/lukesampson/figlet/figletlib"
)

func cherr(e error) {
	if e != nil { panic(e) }
}

func usage() {
	fmt.Printf("Usage: %s [num]\n", os.Args[0])
}

func resterm() {
	fmt.Print("\033[?25h")
	fmt.Print("\033[0m\n")
}

func main() {

	mptr := flag.Int("m", 0, "minutes to count down")
	sptr := flag.Int("s", 0, "seconds to count down")
	fptr := flag.String("f", "univers", "font to use")
	msgptr := flag.String("msg", "Time up!",
		"message to display when done")
	flag.Parse()

	min := *mptr
	sec := *sptr
	font := *fptr

	if min == 0 && sec == 0 { min = 5 }

	// Get term size for placement
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	rtsz, err := cmd.Output()
	ssp := strings.Split(string(rtsz), " ")
	trow, err := strconv.Atoi(ssp[0])
	ssp[1] = strings.TrimSuffix(ssp[1], "\n")
	twid, err := strconv.Atoi(ssp[1])
	cherr(err)
	mid := trow / 2 - 5

	cwd, err := os.Getwd()
	cherr(err)

	f, err := figletlib.GetFontByName(cwd, font)
	cherr(err)

	// Clean up at interrupt
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)

	go func () {
		for range sigc {
			resterm()
			os.Exit(1)
		}
	}()

	// Invisible cursor
	fmt.Print("\033[?25l")

	for tot > -1 {
		cmd = exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		for a := 0; a < mid; a++ {
			fmt.Println()
		}
		pstr := fmt.Sprintf("%02d:%02d", min, sec)
		figletlib.PrintMsg(pstr, f, twid, f.Settings(), "center")
		time.Sleep(1 * time.Second)

		// Red text last minute
		if min == 1 && sec == 0 {
			fmt.Printf("\033[31m")
		}

		sec--
		if sec < 0 {
			sec = 59
			min--
		}
		tot = min * 60 + sec
	}

	// Time up!
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	figletlib.PrintMsg(*msgptr, f, twid, f.Settings(), "center")
	resterm()

}
