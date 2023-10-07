if [ ! -d "go" ]; then
   curl -LO "https://go.dev/dl/go1.21.2.windows-amd64.zip"
   tar -xf "go1.21.2.windows-amd64.zip"
   rm "go1.21.2.windows-amd64.zip"
fi

GOBIN="go/bin/go.exe"

if [ ! -d "w64devkit" ]; then
   curl -LO "https://github.com/skeeto/w64devkit/releases/download/v1.20.0/w64devkit-1.20.0.zip"
   tar -xf "w64devkit-1.20.0.zip"
   rm "w64devkit-1.20.0.zip"
fi

CC="w64devkit/bin/x86_64-w64-mingw32-gcc.exe"

CGO_ENABLED=1 GOOS=windows $GOBIN build
