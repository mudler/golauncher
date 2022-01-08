// Copyright © 2022 Ettore Di Giacinto <mudler@mocaccino.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package gui

import (
	"encoding/json"
	"fmt"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fynexWidget "fyne.io/x/fyne/widget"
	"github.com/0xAX/notificator"
	pluggable "github.com/mudler/go-pluggable"
	"go.deanishe.net/fuzzy"
)

type Launcher struct {
	sync.Mutex
	pluginDir   string
	window      fyne.Window
	manager     *pluggable.Manager
	notifier    *notificator.Notificator
	defaultSize fyne.Size
	responses   map[pluggable.EventType]map[*pluggable.Plugin]*pluggable.EventResponse
}

const (
	searchEvent pluggable.EventType = "search"
	submitEvent pluggable.EventType = "submit"
	appName     string              = "golauncher"
)

func (c *Launcher) Reload(app fyne.App) {

	entry := fynexWidget.NewCompletionEntry([]string{})
	entry.Wrapping = fyne.TextWrapBreak

	var selection interface{}

	entry.OnMenuNavigation = func(w widget.ListItemID) {
		selection = true
	}

	entry.OnSubmitted = func(s string) {
		if len(s) < 3 {
			return
		}

		if selection == nil {
			// Nothing was selected and user hit enter
			// act as user selected the first item
			if len(entry.Options) > 0 {
				s = entry.Options[0]
			}
		}

		res := c.pluginSubmit(submitEvent, s)
		for _, r := range res {
			r.notify(c.notifier)
		}

		app.Quit()
	}

	noResults := func() {
		c.window.Resize(c.defaultSize)
		entry.HideCompletion()
	}

	// When the use typed text, complete the list.
	entry.OnChanged = func(s string) {
		// completion start for text length >= 3
		if len(s) < 3 {
			noResults()
			return
		}

		var results []string
		res := c.pluginSubmit(searchEvent, s)
		for _, r := range res {
			for _, r := range r.Responses() {
				if r != "" {
					results = append(results, r)
				}
			}
		}

		// no results
		if len(results) == 0 {
			noResults()
			return
		}

		fuzzy.SortStrings(results, s)

		// then show them
		entry.SetOptions(results)

		c.window.Resize(c.defaultSize.Add(fyne.NewSize(0, 60*fyne.Min(float32(len(results)), 2))))

		entry.ShowCompletion()
	}
	s := container.NewVSplit(entry, layout.NewSpacer())

	s.SetOffset(0.2)
	c.window.SetContent(
		container.NewAdaptiveGrid(1, s),
	)
	c.window.Canvas().Focus(entry)
}

func (c *Launcher) responseHandler(ev pluggable.EventType) func(p *pluggable.Plugin, r *pluggable.EventResponse) {
	return func(p *pluggable.Plugin, r *pluggable.EventResponse) {
		c.Lock()
		if _, exists := c.responses[ev]; !exists {
			c.responses[ev] = make(map[*pluggable.Plugin]*pluggable.EventResponse)
		}
		c.responses[ev][p] = r
		c.Unlock()
	}
}

type pluginResult struct {
	plugin   *pluggable.Plugin
	response *pluggable.EventResponse
}

func (p pluginResult) Responses() []string {
	pr := &struct {
		Response []string
	}{}
	json.Unmarshal([]byte(p.response.Data), pr)
	return pr.Response
}

func (p pluginResult) notify(n *notificator.Notificator) {
	title := fmt.Sprintf("%s %s", appName, p.plugin.Name)
	if p.response.Errored() {
		n.Push(title, p.response.Error, "", notificator.UR_CRITICAL)
	}
	if p.response.State != "" {
		n.Push(title, p.response.State, "", notificator.UR_NORMAL)
	}
}

func (c *Launcher) pluginSubmit(ev pluggable.EventType, input string) (res []*pluginResult) {
	c.manager.Publish(ev, struct{ Term string }{Term: input})
	for p, r := range c.responses[ev] {
		res = append(res, &pluginResult{plugin: p, response: r})
	}
	return
}

func (c *Launcher) loadUI(app fyne.App) {
	pluginsDir := c.pluginDir
	if pluginsDir == "" {
		pluginsDir = "plugins"
	}

	c.manager.Autoload(appName, pluginsDir)
	c.manager.Register()
	c.manager.Response(submitEvent, c.responseHandler(submitEvent))
	c.manager.Response(searchEvent, c.responseHandler(searchEvent))

	c.window = app.NewWindow(appName)
	c.Reload(app)

	c.window.SetFixedSize(true)
	c.window.Resize(c.defaultSize)
	c.window.SetPadded(true)
	c.window.CenterOnScreen()

	c.window.Show()
}

func newLauncher(pluginDir string) *Launcher {
	return &Launcher{
		pluginDir:   pluginDir,
		defaultSize: fyne.NewSize(640, 30),
		notifier: notificator.New(notificator.Options{
			AppName: appName,
		}),
		manager: pluggable.NewManager(
			[]pluggable.EventType{
				submitEvent,
				searchEvent,
			},
		),
		responses: make(map[pluggable.EventType]map[*pluggable.Plugin]*pluggable.EventResponse),
	}
}
