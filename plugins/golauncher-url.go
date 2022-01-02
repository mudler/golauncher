package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"net/url"

	"github.com/mudler/go-pluggable"
	"mvdan.cc/xurls/v2"
)

func mapS(s []string, m string) (res []string) {
	for _, r := range s {

		URL, _ := url.Parse(r)

		if URL.Scheme == "" {
			URL.Scheme = "https"
		}
		res = append(res, m+URL.String())
	}
	return res
}

func main() {
	xurlsRelaxed := xurls.Relaxed()
	var (
		submit pluggable.EventType = "submit"
		search pluggable.EventType = "search"
	)

	factory := pluggable.NewPluginFactory()

	factory.Add(submit, func(e *pluggable.Event) pluggable.EventResponse {
		d := &struct{ Term string }{}
		json.Unmarshal([]byte(e.Data), d)
		s := ""
		if strings.Contains(d.Term, "Open url: ") {
			u := strings.ReplaceAll(d.Term, "Open url: ", "")
			cmd := exec.Command("xdg-open", u)
			s = fmt.Sprintf("opening %s", u)
			cmd.Start()
			go func() {
				cmd.Wait()
			}()
		}
		return pluggable.EventResponse{State: s}
	})

	factory.Add(search, func(e *pluggable.Event) pluggable.EventResponse {
		d := &struct{ Term string }{}
		json.Unmarshal([]byte(e.Data), d)
		res := mapS(xurlsRelaxed.FindAllString(d.Term, -1), "Open url: ")

		b, _ := json.Marshal(struct{ Response []string }{Response: res})
		return pluggable.EventResponse{Data: string(b)}
	})

	factory.Run(pluggable.EventType(os.Args[1]), os.Args[2], os.Stdout)
}
