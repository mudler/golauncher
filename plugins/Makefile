DESTDIR ?=
PREFIX ?= /usr/local
UBINDIR ?= $(PREFIX)/bin

build:
	CGO_ENABLED=0 go build -o golauncher-fuzzy ./golauncher-fuzzy.go
	CGO_ENABLED=0 go build -o golauncher-url ./golauncher-url.go
	CGO_ENABLED=0 go build -o golauncher-screenshot ./golauncher-screenshot.go

clean:
	rm -rf golauncher-fuzzy golauncher-url golauncher-screenshot

install: build
	install -d $(DESTDIR)/$(UBINDIR)
	install -m 0755 golauncher-screenshot $(DESTDIR)/$(UBINDIR)/
	install -m 0755 golauncher-url $(DESTDIR)/$(UBINDIR)/
	install -m 0755 golauncher-fuzzy $(DESTDIR)/$(UBINDIR)/
	install -m 0755 golauncher-process $(DESTDIR)/$(UBINDIR)/
	install -m 0755 golauncher-window-switch $(DESTDIR)/$(UBINDIR)/