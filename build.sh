VERSION=$(git describe --tags --abbrev=0 | sed 's/\./_/g')
rm -r bin/*
echo CurrentVersion = $VERSION
echo "Building Windows x86... [1/7]"
GOOS=windows GOARCH=386 go build -o bin/PCLHomepageChecker_windows_x86_${VERSION}.exe main.go
echo "Building Windows amd64... [2/7]"
GOOS=windows GOARCH=amd64 go build -o bin/PCLHomepageChecker_windows_amd64_${VERSION}.exe main.go
echo "Building Windows arm... [3/7]"
GOOS=windows GOARCH=arm go build -o bin/PCLHomepageChecker_windows_arm_${VERSION}.exe main.go
echo "Building Linux x86... [4/7]"
GOOS=linux GOARCH=386 go build -o bin/PCLHomepageChecker_linux_x86_${VERSION} main.go
echo "Building Linux amd64... [5/7]"
GOOS=linux GOARCH=amd64 go build -o bin/PCLHomepageChecker_linux_amd64_${VERSION} main.go
echo "Building Linux arm... [6/7]"
GOOS=linux GOARCH=arm go build -o bin/PCLHomepageChecker_linux_arm_${VERSION} main.go
echo "Building Darwin(MacOS) amd64... [7/7]"
GOOS=darwin GOARCH=amd64 go build -o bin/PCLHomepageChecker_macos_amd64_${VERSION} main.go
echo "Building Finished!"