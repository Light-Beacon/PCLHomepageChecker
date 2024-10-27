GOOS=windows GOARCH=386 go build -o bin/PCLHomepageChecker_windows_x86.exe main.go
GOOS=windows GOARCH=amd64 go build -o bin/PCLHomepageChecker_windows_amd64.exe main.go
GOOS=windows GOARCH=arm go build -o bin/PCLHomepageChecker_windows_arm.exe main.go
GOOS=linux GOARCH=386 go build -o bin/PCLHomepageChecker_linux_x86 main.go
GOOS=linux GOARCH=amd64 go build -o bin/PCLHomepageChecker_linux_amd64 main.go
GOOS=linux GOARCH=arm go build -o bin/PCLHomepageChecker_linux_arm main.go
GOOS=darwin GOARCH=amd64 go build -o bin/PCLHomepageChecker_macos_amd64 main.go