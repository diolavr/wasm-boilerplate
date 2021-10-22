SHELL := bash
# .ONESHELL - гарантирует, что каждый сценарий make запускается как один сеанс оболочки
.ONESHELL:
# .SHELLFLAGS - использовать строгий режим
.SHELLFLAGS := -eu -o pipefail -c
# .DELETE_ON_ERROR - удаляет целевой файл, если что-то пошло не так
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(origin .RECIPEPREFIX), undefined)
	$(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

-include .env

SRCDIR = $(shell pwd)

.PHONY: clean
clean:
> @ go clean --cache

.PHONY: install
install: copy-wasm-exec

.PHONY: build
build: build-wasm build-backend

.PHONY: build-backend
build-backend:
> @ echo "> Build wasm-server ..."
> @ go build -ldflags "-w" -o ./build/wasm-server ./backend/main.go

.PHONY: copy-wasm-exec
> @ cp $(GOROOT)/misc/wasm/wasm_exec.js ./assets/wasm_exec.js

.PHONY: build-wasm
build-wasm:
> @ echo "> Build code.wasm ..."
> @ GOOS=js GOARCH=wasm go build -ldflags "-w" -o ./assets/code.wasm ./wasm/main.go

.PHONY: run
run:
> @ ./build/wasm-server
