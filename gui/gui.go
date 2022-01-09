// Copyright © 2021 Ettore Di Giacinto <mudler@mocaccino.org>
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
	"io/ioutil"

	"fyne.io/fyne/v2"
	theme "fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/app"
	"github.com/ghodss/yaml"
)

//go:generate fyne bundle -package gui -o data.go ../Icon.png

// Load theme. Best effort
func loadTheme(t string, app fyne.App) {
	if t != "" {
		b, err := ioutil.ReadFile(t)
		if err == nil {
			jb, err := yaml.YAMLToJSON(b)
			if err == nil {
				t, err := theme.FromJSON(string(jb))
				if err == nil {
					app.Settings().SetTheme(t)
				}
			}
		}
	}
}

func Run(tn, pluginDir string) {
	app := app.New()

	loadTheme(tn, app)
	app.SetIcon(resourceIconPng)

	c := newLauncher(pluginDir)
	c.loadUI(app)
	app.Run()
}
