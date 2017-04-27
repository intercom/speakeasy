#!/bin/sh
type go > /dev/null || { echo "Couldn't find go command'" ; exit 1; }

while :;do printf .;sleep 1;done &
#Die with parent if we die prematurely
trap "kill $!" EXIT

echo "\nGetting mux files"
go get -u github.com/gorilla/mux
echo "\nGetting gouuid files"
go get -u github.com/nu7hatch/gouuid
echo "\nGetting go-bindata files"
go get -u github.com/jteeuwen/go-bindata/...
echo "\nGetting gomobile files"
go get -u golang.org/x/mobile/cmd/gomobile
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

#Kill the loop and unset the trap or else the pid might get reassigned and we might end up killing a completely different process
kill $! && trap " " EXIT
