DBG_MAKEFILE=1
PROJECT_NAME=cosmos-checkers-sdk

# BuildVersion is uses the target release date from year to day seperated by DOTs (.)
# 2022.12.31
#
# following this is the character `f` which defines the number of Features & fixes seperated by DOT (.)
# f3.14
#
# this states that this release has 3 features and/or breaking changes plus 14 fixes/changes
#
# full example is
# 2022.12.31f3.14
#
# this states that the this version was released on 31st December 2022 and it contains 3 features and 14 fixes/changes
# remember to update the BUILDVERION value below this line before official releases and to set the Commit Hash tag to the same value
BUILDVERSION=2022.12.20f1
export VERSION=$(BUILDVERSION)
BUILDDATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
export LASTBUILDDATE=$(BUILDDATE)

build:
	$(info ***** building $(PROJECT_NAME) *****)
	@GOOS=js GOARCH=wasm go build -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildDate=$(LASTBUILDDATE)" -o ./bin/$(PROJECT_NAME).wasm
	cp ./bin/$(PROJECT_NAME).wasm ../frontend/static/$(PROJECT_NAME).wasm
	
clean:
	$(info ***** cleaning $(PROJECT_NAME) *****)
	@rm -f ./bin/$(PROJECT_NAME).wasm
