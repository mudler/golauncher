<h1 align="center">
  <br>
	<img src="https://user-images.githubusercontent.com/2420543/147978379-b9097fd4-89d9-4119-bef6-459fa4554d7d.png" width=128
         alt="logo"><br>
    golauncher

<br>
</h1>

<h3 align="center">A go application launcher </h3>
<p align="center">
  <a href="https://opensource.org/licenses/">
    <img src="https://img.shields.io/badge/licence-MIT-brightgreen"
         alt="license">
  </a>
  <a href="https://github.com/mudler/golauncher/issues"><img src="https://img.shields.io/github/issues/mudler/golauncher"></a>
  <img src="https://img.shields.io/badge/made%20with-Go-blue">
  <img src="https://goreportcard.com/badge/github.com/mudler/golauncher" alt="go report card" />
</p>

<p align="center">
	 <br>
      A simple, highly extensible, customizable application launcher and window switcher written in less than 300 lines of Golang and fyne <br>
</p>

<h1 align="center">
  <br>
	<img src="https://user-images.githubusercontent.com/2420543/147978694-6a5797f7-d2f7-49ea-9c61-48ef027cb5e5.png" 
         alt="screenshot"><br>
	<img src="https://user-images.githubusercontent.com/2420543/147978695-ff18c4df-33ac-4b65-9f4b-3676c6edd5de.png" 
         alt="screenshot"><br>
</h1>

golauncher is a simple, highly extensible application launcher written in Golang. Its written using [fyne](https://github.com/fyne-io/fyne) in less than 300 lines of Go (actually, most of them are just layouts!). 

Works on i3, Xfce, GNOME, Plasma, fynedesk, ...

# :computer: Installation

Download the [release](https://github.com/mudler/golauncher/releases), extract the tarball in your system and run `make install`.

As it does use `fyne`, does not depend on GTK or either QT installed on your system.

## :construction_worker: Build from source

You can also build golauncher locally with `go build`. 

Note: plugins are standalone binaries and not part of golauncher, you need to install them separately, or if you are developing, use `--plugin-dir` to point to a specific plugin directory.

# :runner: Run

Once you have `golauncher` installed you can either run it from the terminal with `golauncher`, or either start it from the application menu.

```
GLOBAL OPTIONS:
   --theme value        [$THEME]
   --plugin-dir value   [$PLUGIN_DIR]
```

Golauncher takes optionally a theme with `--theme` and an additional directory to scan for plugins (`--plugin-dir`).

The plugin directory must contains binary prefixed with `golauncher-` in order to be loaded.

# :pencil2: Extensible

golauncher is extensible from the bottom up, and indeed the core does provide no functionalities besides the GUI displaying.

## :gear: Writing extensions

Extensions can be written in any language. `golauncher` reads binaries starting with `golauncher-` prefix inside  `$PATH` and automatically invokes them while the user is submitting inputs to retrieve results to be displayed to the user. Optionally, golauncher takes a `PLUGIN_DIR` environment variable (or `--plugin-dir` as args) to specify an additional plugin directory to use.

All the current functionalities of golauncher are split into separate plugins. Plugins can be written in any language. For examples on how to extend golauncher, see the [plugin](https://github.com/mudler/golauncher/tree/master/plugins) folder in this repository.


# :art: Themes

Golauncher supports custom themes, you can find examples inside the [themes](https://github.com/mudler/golauncher/tree/master/themes) folder.

To run golauncher with a theme, run:

```
$ THEME=/home/mudler/.golauncher/monokai.yaml golauncher
or
$ golauncher --theme /home/mudler/.golauncher/monokai.yaml
```

## `sand`

![Screenshot from 2022-01-07 23-05-33](https://user-images.githubusercontent.com/2420543/148613691-552ab66b-f8e8-45b0-98a0-9cf57807375f.png)

## `blue_light`

![Screenshot from 2022-01-07 23-06-34](https://user-images.githubusercontent.com/2420543/148613697-4ca0c300-b646-4866-8ce3-4308d639d06e.png)

## `blue_night`

![Screenshot from 2022-01-07 23-07-04](https://user-images.githubusercontent.com/2420543/148613703-f9faf52a-4d30-47ec-a70f-1ca9fca9cd7a.png)


# License

MIT License

Copyright (c) 2022 Ettore Di Giacinto
