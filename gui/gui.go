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
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2/app"
	"gopkg.in/yaml.v2"
)

//go:generate fyne bundle -package gui -o data.go ../Icon.png

func Run(theme, pluginDir string) {
	app := app.New()
	t := NewTheme()

	if theme != "" {
		b, err := ioutil.ReadFile(theme)
		if err == nil {
			fmt.Println("Loading theme", theme)
			yaml.Unmarshal(b, t)
		}
	}
	app.Settings().SetTheme(t)
	app.SetIcon(resourceIconPng)

	c := newLauncher(pluginDir)
	c.loadUI(app)
	app.Run()
}
