PREFIX ?= /usr/local
UBINDIR ?= $(PREFIX)/bin
BUNDLE_NAME ?= golauncher

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

build:
	$(MAKE) -C plugins/ build
	go build -o golauncher

release: $(current_dir)/out
	mkdir -p $(current_dir)/out/$(PREFIX)/share/applications
	mkdir -p $(current_dir)/out/$(PREFIX)/share/pixmaps
	cp -rf $(current_dir)/resources/assets/logo.png $(current_dir)/out/$(PREFIX)/share/pixmaps/golauncher.png
	cp -rf $(current_dir)/resources/assets/golauncher.desktop $(current_dir)/out/$(PREFIX)/share/applications
	cp -rf $(current_dir)/resources/assets/Makefile $(current_dir)/out/
	cd out && tar -cJf $(current_dir)/$(BUNDLE_NAME).tar.xz .

install: build
	$(MAKE) -C plugins/ install
	install -d $(DESTDIR)/$(UBINDIR)
	install -m 0755 golauncher $(DESTDIR)/$(UBINDIR)/

run: build
	./golauncher

$(current_dir)/out:
	mkdir out
	$(MAKE) DESTDIR=$(current_dir)/out install

test: $(current_dir)/out
	THEME=themes/sand.yaml PLUGIN_DIR=$(current_dir)/out/usr/local/bin $(current_dir)/out/usr/local/bin/golauncher

clean:
	$(MAKE) -C plugins/ clean
	rm -rf $(current_dir)/out
	rm -rf golauncher
