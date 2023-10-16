if [ ! -d "go" ]; then
   echo "[NO GO COMPILER FOUND]"
   echo "[DOWNLOADING GO COMPILER]"
   curl -LO "https://go.dev/dl/go1.21.2.windows-amd64.zip"
   echo "[UNZIPPING GO COMPILER]"
   unzip "go1.21.2.windows-amd64.zip" 1>/dev/null
   echo "[REMOVING ZIP FILE]"
   rm "go1.21.2.windows-amd64.zip"
fi

GOBIN="./go/bin/go.exe"

if [ ! -d "w64devkit" ]; then
   echo "[NO GCC FOUND]"
   echo "[DOWNLOADING GCC]"
   curl -LO "https://github.com/skeeto/w64devkit/releases/download/v1.20.0/w64devkit-1.20.0.zip"
   echo "[UNZIPPING GCC]"
   unzip "w64devkit-1.20.0.zip" 1>/dev/null
   echo "[REMOVING ZIP FILE]"
   rm "w64devkit-1.20.0.zip"
fi

echo "[ATTEMPING TO COMPILE]"
COMPILER="/w64devkit/bin/x86_64-w64-mingw32-gcc.exe" 
CC="$PWD$COMPILER" CGO_ENABLED=1 GOOS=windows $GOBIN build
