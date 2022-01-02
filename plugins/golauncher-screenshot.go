package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"image/png"

	"github.com/kbinani/screenshot"
	"github.com/mudler/go-pluggable"
)

func fork() {
	exec.Command("/proc/self/exe",
		"_screenshot",
	).Start()
}

func takeScreenshot() {
	time.Sleep(1 * time.Second)
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		os.MkdirAll("/tmp/screenshots", os.ModePerm)
		fileName := fmt.Sprintf("/tmp/screenshots/%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
	}
	// Open screenshot folder
	exec.Command("xdg-open", "/tmp/screenshots").Run()
}

func main() {
	var (
		submit pluggable.EventType = "submit"
		search pluggable.EventType = "search"
	)

	if os.Args[1] == "_screenshot" {
		takeScreenshot()
	}

	factory := pluggable.NewPluginFactory()

	factory.Add(submit, func(e *pluggable.Event) pluggable.EventResponse {
		d := &struct {
			Term string
		}{}
		json.Unmarshal([]byte(e.Data), d)

		s := ""

		if strings.Contains(d.Term, "Take screenshot") {
			s = "Taking screenshots in /tmp/screenshots"
			fork()
		}
		return pluggable.EventResponse{State: s}
	})

	factory.Add(search, func(e *pluggable.Event) pluggable.EventResponse {
		d := &struct {
			Term string
		}{}
		json.Unmarshal([]byte(e.Data), d)
		if strings.Contains(strings.ToLower(d.Term), "screen") || strings.Contains(strings.ToLower(d.Term), "take") {
			b, _ := json.Marshal(struct{ Response []string }{Response: []string{"Take screenshot"}})
			return pluggable.EventResponse{Data: string(b)}
		}
		return pluggable.EventResponse{}
	})

	factory.Run(pluggable.EventType(os.Args[1]), os.Args[2], os.Stdout)
}
