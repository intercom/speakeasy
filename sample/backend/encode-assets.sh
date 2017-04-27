#!/bin/sh

command -v identify >/dev/null 2>&1 || { echo >&2 "imagemagick is required but it's not installed."; exit 1; }

go-bindata -nometadata -ignore=\\.gitignore -nocompress -o bindata.go -pkg backend assets/...

OUTPUT_PATH="encoded_assets.go"

printf "package backend\nfunc GetAsset(path string) *EncodedAsset {\nswitch path {" > $OUTPUT_PATH 

## For 'find' to work, we need to have extended POSIX regular expressions.
## Different UNIXes do this differently:
if [ "$(uname)" = "Linux" ]; then
    find_option=""
    find_execution_option="-regextype posix-extended "
else ## This is for BSD-like, such as Darwin OS:
    find_option="-E "
    find_execution_option=""
fi

asset_paths=$(find $find_option assets $find_execution_option -regex '.*\.(jpg|jpeg|png|gif)')
for asset_path in $asset_paths
do :
	asset=$(basename $asset_path)
	mime=$(file -b --mime-type $asset_path)
	width=$(identify -ping -format '%w,' $asset_path | cut -d "," -f 1)
	height=$(identify -ping -format '%h,' $asset_path | cut -d "," -f 1)
	printf "\ncase \"$asset_path\":\nreturn &EncodedAsset{ \nPath: \"$asset_path\",\nFilename: \"$asset\",\nMimeType: \"$mime\",\nWidth: $width,\nHeight: $height,\nData: func() []byte { asset, _ := _bindata[\"assets/$asset_path\"]()\nif asset == nil {\nreturn nil\n}\nreturn asset.bytes },}" >> $OUTPUT_PATH
done

printf "\n}\nreturn nil\n}" >> $OUTPUT_PATH

go fmt $OUTPUT_PATH > /dev/null

