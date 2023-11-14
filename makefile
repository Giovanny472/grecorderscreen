build-gscreenshot:
	@go build -ldflags -H=windowsgui -o .\bin\gscreenshot.exe  .\cmd\main.go