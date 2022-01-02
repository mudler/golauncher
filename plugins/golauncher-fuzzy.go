package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mudler/go-pluggable"
	"github.com/sahilm/fuzzy"
)

func listDir(dir string) []string {
	content := []string{}

	filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				content = append(content, path)
			}

			return nil
		})

	return content
}

func ff(f string) (res []string) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	bins := []string{}
	for _, p := range paths {
		bins = append(bins, listDir(p)...)
	}

	m := fuzzy.Find(f, bins)
	for _, match := range m {
		res = append(res, match.Str)
	}

	// Cut results to first ten
	if len(res) > 10 {
		res = res[0:10]
	}

	return
}

type res struct {
	Term string
}

func mapS(s []string, m string) (res []string) {
	for _, r := range s {
		res = append(res, m+r)
	}
	return res
}

func main() {

	var (
		submit pluggable.EventType = "submit"
		search pluggable.EventType = "search"
	)

	factory := pluggable.NewPluginFactory()

	factory.Add(submit, func(e *pluggable.Event) pluggable.EventResponse {
		d := &res{}
		json.Unmarshal([]byte(e.Data), d)
		s := ""
		if strings.Contains(d.Term, "Execute: ") {
			proc := strings.ReplaceAll(d.Term, "Execute: ", "")
			cmd := exec.Command(proc)
			s = fmt.Sprintf("opening %s", proc)
			cmd.Start()
			go func() {
				cmd.Wait()
			}()
		}
		return pluggable.EventResponse{State: s}
	})

	factory.Add(search, func(e *pluggable.Event) pluggable.EventResponse {
		d := &res{}
		json.Unmarshal([]byte(e.Data), d)
		res := mapS(ff(d.Term), "Execute: ")
		b, _ := json.Marshal(struct{ Response []string }{Response: res})
		return pluggable.EventResponse{Data: string(b)}
	})

	factory.Run(pluggable.EventType(os.Args[1]), os.Args[2], os.Stdout)
}
