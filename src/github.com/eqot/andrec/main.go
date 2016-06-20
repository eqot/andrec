package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "andrec"
	app.Usage = "screen recorder for Android"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	app.Action = func(c *cli.Context) error {
		filename := "Nefertiti"
		if c.NArg() > 0 {
			filename = c.Args().Get(0)
		}

		record(filename)

		return nil
	}

	app.Run(os.Args)
}

func record(filename string) {
	c1 := exec.Command("adb", "shell", "screenrecord", "--size 360x640", "--output-format=h264", "-")
	c2 := exec.Command("ffmpeg", "-i", "-", "-y", filename)

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

	time.Sleep(time.Second * 3)

	c1.Process.Kill()
	w.Close()
	c2.Wait()
}
