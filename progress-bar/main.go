package main

import (
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"time"
)

func main() {
	ProgressBarOption()
	fmt.Println()
	//ProgressBar()
}

// ProgressBarOption
// go get -u github.com/schollz/progressbar/v3
func ProgressBarOption() {
	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		//progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		//progressbar.OptionShowCount(),
		progressbar.OptionSetDescription("[_yellow_][1/5][reset] [_magenta_]Reading book[reset]"),
		//progressbar.OptionSetElapsedTime(false),
		progressbar.OptionShowElapsedTimeOnFinish(),
		//progressbar.OptionSpinnerType(1),
		//progressbar.OptionSetPredictTime(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]█[reset]",
			SaucerHead:    "[green][reset]",
			SaucerPadding: " ",
			BarStart:      "|",
			BarEnd:        "|",
		}))

	bar.Add(200)
	time.Sleep(time.Second * 2)
	bar.Add(200)
	bar.Describe("[_yellow_][2/5][reset] [_magenta_]Listen music[reset]")
	time.Sleep(time.Second * 2)
	bar.Add(200)
	bar.Describe("[_yellow_][3/5][reset] [_magenta_]Watch movie[reset]")
	time.Sleep(time.Second * 2)
	bar.Add(200)
	bar.Describe("[_yellow_][4/5][reset] [_magenta_]Play game[reset]")
	time.Sleep(time.Second * 2)
	//bar.Add(200)
	bar.Describe("[_yellow_][5/5][reset] [_magenta_]ready sleep[reset]")
	bar.Finish()
	//bar.Close()
}

func ProgressBar() {
	bar := progressbar.New(5000)
	for i := 0; i < 10000; i++ {
		bar.Add(i)
		time.Sleep(time.Millisecond)
	}
}
