all: buildgui buildprofiler

buildgui: manual_builtin cmd/gui/*
	cd ./cmd/gui && \
	go-bindata res/ && \
	go build .

buildprofiler: cmd/profiler/*
	cd ./cmd/profiler && \
	go build .

run: buildgui
	./cmd/gui/gui

profile: buildprofiler
	./cmd/profiler/profiler

test: pkg/*
	go test ./pkg/...

# Delivery
pack:

.PHONY: doc manual manual_builtin debian # handled in other Makefiles

# Debian package
debian: buildgui
	mkdir -p ivs-calculator
	cp -R debian/ ivs-calculator/DEBIAN
# set up filesystem
	mkdir -p ivs-calculator/usr/bin
	mkdir -p ivs-calculator/usr/share/applications
	mkdir -p ivs-calculator/usr/share/icons/hicolor/scalable/apps
	cp cmd/gui/gui ivs-calculator/usr/bin/ivs-calculator
	cp assets/ivs-calculator.desktop ivs-calculator/usr/share/applications/ivs-calculator.desktop
	cp assets/ivs_calculator.svg -p ivs-calculator/usr/share/icons/hicolor/scalable/apps/ivs-calculator.svg
# build .deb
	dpkg-deb --build ivs-calculator
# remove fs
	rm -rf ivs-calculator/

# Documentation
doc:
	$(MAKE) -C docs html

# User manual generation
manual:
	$(MAKE) -C manual all

manual_builtin:
	$(MAKE) -C manual builtin

# Cleanup
clean:
	rm -f ./cmd/gui/gui ./cmd/profiler/profiler ./cmd/gui/bindata.go ./ivs-calculator.deb
	$(MAKE) -C docs clean
	$(MAKE) -C manual clean