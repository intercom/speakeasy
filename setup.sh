#!/bin/sh
type go > /dev/null || { echo "Couldn't find go command'" ; exit 1; }

# Print a dot for every second the setup takes
while :;do printf .;sleep 1;done &
trap "kill $!" EXIT

echo "\nGetting packages"
go get -u github.com/gorilla/mux github.com/nu7hatch/gouuid github.com/kevinburke/go-bindata/... golang.org/x/mobile/cmd/gomobile
echo "\ngomobile init 📱"
type gomobile > /dev/null || { echo "Couldn't find gomobile, is \$GOPATH/bin on your \$PATH?" ; exit 1; }
if [ -n "$ANDROID_NDK" ]; then
	gomobile init -ndk $ANDROID_NDK;
else
	gomobile init;
fi
echo "\nInstalling speakeasy tool"
go install github.com/intercom/speakeasy/cmd/speakeasy

echo "\nDone!"

# Kill the dot printing process
kill $! && trap " " EXIT
