# If PREFIX isn't provided, we check for $(DESTDIR)/usr/local and use that if it exists.
# Otherwice we fall back to using /usr.

LOCAL != test -d $(DESTDIR)/usr/local && echo -n "/local" || echo -n ""
LOCAL ?= $(shell test -d $(DESTDIR)/usr/local && echo "/local" || echo "")
PREFIX ?= /usr$(LOCAL)

default:
	# Run "sudo make install" to install the application.
	# Run "sudo make uninstall" to uninstall the application.

install:
	install -Dm00644 usr/local/share/applications/golauncher.desktop $(DESTDIR)$(PREFIX)/share/applications/golauncher.desktop
	install -Dm00755 usr/local/bin/golauncher $(DESTDIR)$(PREFIX)/bin/golauncher
	install -Dm00755 usr/local/bin/golauncher-url $(DESTDIR)$(PREFIX)/bin/golauncher-url
	install -Dm00755 usr/local/bin/golauncher-screenshot $(DESTDIR)$(PREFIX)/bin/golauncher-screenshot
	install -Dm00755 usr/local/bin/golauncher-fuzzy $(DESTDIR)$(PREFIX)/bin/golauncher-fuzzy
	install -Dm00755 usr/local/bin/golauncher-process $(DESTDIR)$(PREFIX)/bin/golauncher-process
	install -Dm00755 usr/local/bin/golauncher-window-switch $(DESTDIR)$(PREFIX)/bin/golauncher-window-switch
	install -Dm00644 usr/local/share/pixmaps/golauncher.png $(DESTDIR)$(PREFIX)/share/pixmaps/golauncher.png

uninstall:
	-rm $(DESTDIR)$(PREFIX)/share/applications/golauncher.desktop
	-rm $(DESTDIR)$(PREFIX)/bin/golauncher
	-rm $(DESTDIR)$(PREFIX)/share/pixmaps/golauncher.png