APP_NAME ?= Cleanium

# Detect OS for sed compatibility
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	SED_INPLACE = sed -i ''
else
	SED_INPLACE = sed -i
endif

test:
	cd src && go run *.go

clean:
	rm -rf build

compile: 
	cd src && GOOS=darwin GOARCH=amd64 go build -o ../build/macos/amd64/$(APP_NAME).app/Contents/MacOS/$(APP_NAME)
	chmod +x build/macos/amd64/$(APP_NAME).app/Contents/MacOS/$(APP_NAME)
	cd src && GOOS=darwin GOARCH=arm64 go build -o ../build/macos/arm64/$(APP_NAME).app/Contents/MacOS/$(APP_NAME)
	chmod +x build/macos/arm64/$(APP_NAME).app/Contents/MacOS/$(APP_NAME)

define scaffold_arch
	mkdir -p build/macos/$(1)/$(APP_NAME).app/Contents/MacOS
	mkdir -p build/macos/$(1)/$(APP_NAME).app/Contents/Resources
	cp -R Info.plist build/macos/$(1)/$(APP_NAME).app/Contents/
	$(SED_INPLACE) 's/Cleanium/$(APP_NAME)/g' build/macos/$(1)/$(APP_NAME).app/Contents/Info.plist
	$(if $(APP_URL),$(SED_INPLACE) 's|<string></string>|<string>$(APP_URL)</string>|' build/macos/$(1)/$(APP_NAME).app/Contents/Info.plist)
	$(if $(APP_ICON),cp $(APP_ICON) build/macos/$(1)/$(APP_NAME).app/Contents/Resources/Icon.icns,cp -R icon/Icon.icns build/macos/$(1)/$(APP_NAME).app/Contents/Resources/)
endef

scaffold:
	$(call scaffold_arch,amd64)
	$(call scaffold_arch,arm64)

build: clean scaffold compile
	
build-release: build
	cd build/macos/amd64 && zip -r ../../$(APP_NAME)-macos-amd64.zip $(APP_NAME).app
	cd build/macos/arm64 && zip -r ../../$(APP_NAME)-macos-arm64.zip $(APP_NAME).app