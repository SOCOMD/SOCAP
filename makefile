.PHONY build-win64:
build-win64:
	@GOARCH="amd64"; \
	CGO_ENABLED=1; \
	go build -o ./bin/socap_x64.dll -buildmode=c-shared .

.PHONY build-linux64:
build-linux64:
	@GOARCH="amd64"; \
	CGO_ENABLED=1; \
	go build -o ./bin/socap_x64.so -buildmode=c-shared .

.PHONY build-addon:
build-addon:
	@MakePbo.exe -N ./addon ./bin/socap.pbo