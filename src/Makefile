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

.PHONY: doc manual manual_builtin # handled in other Makefiles

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
	rm -f ./cmd/gui/gui ./cmd/profiler/profiler ./cmd/gui/bindata.go
	$(MAKE) -C docs clean
	$(MAKE) -C manual clean