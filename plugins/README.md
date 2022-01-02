# golauncher plugins

This set of plugins are the "core" ones and can be used as examples to write plugins in other languages.

A plugin is a standalone binary installed in your system. golauncher will scan your $PATH for binaries starting with `golauncher-` and call them while receiving the user input.

Currently in this folder you will find plugins written in bash/go.

Each plugin receives the event and the user input via args:

Events name are passed by as first argument (`submit`, or `search`) and a json payload containing the user input is passed as second argument ( e.g. `{ "data": "{ "Term": "foo" }" }` , if we enter `foo` in the golauncher input box).

This means that every plugin can be also tested individually, for example by running:

```bash
$ ./plugins/golauncher-process search '{ "data": { "Term": "firefox" } }'    
{
  "state": "",
  "data": "{ \"response\": [ \"Kill process /usr/lib64/firefox/firefox pid:10633\" ] }",
  "error": ""
}
```