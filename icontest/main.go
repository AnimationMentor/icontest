package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/getlantern/systray"

	ico "github.com/Kodeworks/golang-image-ico"
)

var imageFilesNames []string

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %v image.png image.ico...\n", os.Args[0])
		return
	}

	imageFilesNames = os.Args[1:]
	fmt.Printf("images: %v\n", strings.Join(imageFilesNames, ", "))

	systray.Run(onReady, onExit)
}

func onReady() {
	for _, name := range imageFilesNames {
		addMenuAction(name, "load me", loadImageHandler(name))
	}
	systray.AddSeparator()
	addMenuAction("Quit", "", func(_ *systray.MenuItem) { onExit() })

	loadImage(imageFilesNames[0])
}

func onExit() {
	fmt.Println("cheers")
	os.Exit(0)
}

func addMenuAction(label, tooltip string, f func(*systray.MenuItem)) *systray.MenuItem {
	m := systray.AddMenuItem(label, tooltip)
	go func() {
		for {
			<-m.ClickedCh
			f(m)
		}
	}()
	return m
}

func loadImageHandler(name string) func(*systray.MenuItem) {
	return func(m *systray.MenuItem) {
		loadImage(name)
	}
}

func loadImage(name string) {
	fmt.Printf(">> loading %v\n", name)
	buf, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		// convert to ico if it's not already
		ext := strings.ToLower(filepath.Ext(name))
		if ext == ".png" {
			fmt.Println("converting from png to ico")
			buf = toIco(png.Decode, buf)
		} else if ext == ".jpg" || ext == ".jpeg" {
			fmt.Println("converting from jpg to ico")
			buf = toIco(jpeg.Decode, buf)
		}
	}

	systray.SetIcon(buf)
}

func toIco(f func(r io.Reader) (image.Image, error), inBits []byte) []byte {
	img, err := f(bytes.NewReader(inBits))
	if err != nil {
		log.Fatal(err)
	}

	buf := &bytes.Buffer{}
	if err := ico.Encode(buf, img); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
