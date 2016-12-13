package main

import (
	"fmt"
	"flag"
	"os"
	"os/user"
	"os/exec"
	"os/signal"
	"strings"
	"strconv"
	"time"
	"path/filepath"
	"github.com/lukesampson/figlet/figletlib"
)

func cherr(e error) {
	if e != nil { panic(e) }
}

func resterm() {
	fmt.Print("\033[?25h")
	fmt.Print("\033[0m\n")
}

func chred(tot int) {
	if tot < 60{
		fmt.Printf("\033[31m")
	}
}

func main() {

	mptr := flag.Int("m", 0, "minutes to count down")
	sptr := flag.Int("s", 0, "seconds to count down")
	fptr := flag.String("f", "univers", "font to use")
	msgptr := flag.String("msg", "Time up!",
		"message to display when done")
	flag.Parse()

	if *mptr == 0 && *sptr == 0 { *mptr = 5 }
	tot := *mptr * 60 + *sptr

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

	// Put univers.flf in $HOME/.fonts
	usr, err := user.Current()
	cherr(err)
	fdir := filepath.Join(usr.HomeDir, ".fonts")

	f, err := figletlib.GetFontByName(fdir, *fptr)
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

	// Check for red at start
	chred(tot)

	for tot > 0 {
		cmd = exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		for a := 0; a < mid; a++ {
			fmt.Println()
		}

		// Red text last minute
		chred(tot)

		pstr := fmt.Sprintf("%02d:%02d", *mptr, *sptr)
		figletlib.PrintMsg(pstr, f, twid, f.Settings(), "center")
		time.Sleep(1 * time.Second)

		*sptr--
		tot--
		if *sptr < 0 {
			*sptr = 59
			*mptr--
		}
	}

	// Time up!
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	figletlib.PrintMsg(*msgptr, f, twid, f.Settings(), "center")
	resterm()
}
