VERSION=$(git describe --tags --abbrev=0 | sed 's/\./_/g')
rm -r bin/*
echo CurrentVersion = $VERSION
echo "Building Windows x86... [1/8]"
GOOS=windows GOARCH=386 go build -o bin/PCLHomepageChecker_windows_x86_${VERSION}.exe main.go
echo "Building Windows amd64... [2/8]"
GOOS=windows GOARCH=amd64 go build -o bin/PCLHomepageChecker_windows_amd64_${VERSION}.exe main.go
echo "Building Windows arm... [3/8]"
GOOS=windows GOARCH=arm go build -o bin/PCLHomepageChecker_windows_arm_${VERSION}.exe main.go
echo "Building Linux x86... [4/8]"
GOOS=linux GOARCH=386 go build -o bin/PCLHomepageChecker_linux_x86_${VERSION} main.go
echo "Building Linux amd64... [5/8]"
GOOS=linux GOARCH=amd64 go build -o bin/PCLHomepageChecker_linux_amd64_${VERSION} main.go
echo "Building Linux arm... [6/8]"
GOOS=linux GOARCH=arm go build -o bin/PCLHomepageChecker_linux_arm_${VERSION} main.go
echo "Building Darwin(MacOS) x86... [7/8]"
GOOS=darwin GOARCH=386 go build -o bin/PCLHomepageChecker_macos_x86_${VERSION} main.go
echo "Building Darwin(MacOS) arm... [8/8]"
GOOS=darwin GOARCH=arm go build -o bin/PCLHomepageChecker_macos_arm_${VERSION} main.go
echo "Building Finished!"
