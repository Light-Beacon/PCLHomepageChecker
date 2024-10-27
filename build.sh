VERSION=$(git describe --tags --abbrev=0 | sed 's/\./_/g')
rm -r bin/*
echo CurrentVersion = $VERSION
GOOS=windows GOARCH=386 go build -o bin/PCLHomepageChecker_windows_x86_${VERSION}.exe main.go
GOOS=windows GOARCH=amd64 go build -o bin/PCLHomepageChecker_windows_amd64_${VERSION}.exe main.go
GOOS=windows GOARCH=arm go build -o bin/PCLHomepageChecker_windows_arm_${VERSION}.exe main.go
GOOS=linux GOARCH=386 go build -o bin/PCLHomepageChecker_linux_x86_${VERSION} main.go
GOOS=linux GOARCH=amd64 go build -o bin/PCLHomepageChecker_linux_amd64_${VERSION} main.go
GOOS=linux GOARCH=arm go build -o bin/PCLHomepageChecker_linux_arm_${VERSION} main.go
GOOS=darwin GOARCH=amd64 go build -o bin/PCLHomepageChecker_macos_amd64_${VERSION} main.go