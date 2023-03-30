

test:
	cd src && go run *.go

clean:
	rm -rf build

compile: 
	cd src && GOOS=darwin GOARCH=amd64 go build -o ../build/macos/amd64/Cleanium.app/Contents/MacOS/Cleanium
	chmod +x build/macos/amd64/Cleanium.app/Contents/MacOS/Cleanium
	cd src && GOOS=darwin GOARCH=arm64 go build -o ../build/macos/arm64/Cleanium.app/Contents/MacOS/Cleanium
	chmod +x build/macos/arm64/Cleanium.app/Contents/MacOS/Cleanium

scaffold:
	mkdir -p build/macos/amd64/Cleanium.app/Contents/MacOS
	mkdir -p build/macos/amd64/Cleanium.app/Contents/Resources
	cp -R Info.plist build/macos/amd64/Cleanium.app/Contents/
	cp -R icon/* build/macos/amd64/Cleanium.app/Contents/Resources/
	mkdir -p build/macos/arm64/Cleanium.app/Contents/MacOS
	mkdir -p build/macos/arm64/Cleanium.app/Contents/Resources
	cp -R Info.plist build/macos/arm64/Cleanium.app/Contents/
	cp -R icon/* build/macos/arm64/Cleanium.app/Contents/Resources/

build: clean scaffold compile
	
build-release: build
	cd build/macos/amd64 && zip -r ../../Cleanium-macos-amd64.zip Cleanium.app
	cd build/macos/arm64 && zip -r ../../Cleanium-macos-arm64.zip Cleanium.app