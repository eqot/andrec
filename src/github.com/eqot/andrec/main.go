package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "andrec"
	app.Usage = "screen recorder for Android"
	app.UsageText = app.Name + " [global options] filename"
	app.Version = "v0.2.0"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "landscape", Aliases: []string{"l"}},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			log.Fatal("output file not found")
		}

		filename := c.Args().Get(0)
		isLandscape := c.Bool("landscape")
		record(filename, isLandscape)

		return nil
	}

	app.Run(os.Args)
}

func record(filename string, isLandscape bool) {
	adbOptions := []string{"shell", "screenrecord", "--size 360x640", "--output-format=h264"}
	if isLandscape {
		adbOptions = append(adbOptions, "--rotate")
	}
	adbOptions = append(adbOptions, "-")

	ffmpegOptions := []string{"-i", "-", "-y"}
	if isLandscape {
		ffmpegOptions = append(ffmpegOptions, "-vf", "transpose=2")
	}
	ffmpegOptions = append(ffmpegOptions, filename)

	c1 := exec.Command("adb", adbOptions...)
	c2 := exec.Command("ffmpeg", ffmpegOptions...)

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	if err := c1.Start(); err != nil {
		log.Fatal(err)
	}
	if err := c2.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Recording screen to", filename)
	fmt.Println("Press enter key to stop recording")

	bufio.NewReader(os.Stdin).ReadBytes('\n')

	c1.Process.Kill()
	w.Close()
	c2.Wait()
}
