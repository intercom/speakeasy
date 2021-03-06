// Code generated by go-bindata.
// sources:
// assets/avatars/justice.png
// DO NOT EDIT!

package backend

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsAvatarsJusticePng = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00H\x00\x00\x00H\b\x03\x00\x00\x00b3Cu\x00\x00\x00cPLTE\x00\x00\x00;\x88\xc3;\x88\xc3\xff\xe8\xb6;\x88\xc3\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe8\xb6\xff\xe4\xae\xff\xc3d\xff\u0655\xff\xb0;\xff\xbf\\\xff\xb4C\xff\xc6l;\x88\xc3\xff\xb7L\xff\xce}\xff\xbbT\xff\u0485\xff\xcau\xff\u054d\xff\xac3\xff\xe1\xa6it\x984\x00\x00\x00\x10tRNS\x00P\xdf\x00\x00\xdf\u03ff\xaf\x80p@\x10 \xef0\xdfS\x89\xd7\x00\x00\x01\xe4IDATx^\xed\xd8\u02ce\xdb0\f\x05\xd0\u012a\x9dX\x96=\xf43\xef\xb4\xff\xff\x95\xa5PLqGnIj\x8c\xecrw\xd9\x1cD\xa2(B\xde}\u027e8\x9bS\xec\u007f`v\xf8c\u007f\xce\nH\tT\xe4A\xc5\u007f\xa1sf^\r\xe5/\xedE\x9b\x8d\xd9Z~\fq~\u0299\x88\xe3\x92\xec\\\x92\x92\xa1\x9b\f\xdd\u0629T\xe8`\x83\x0e&\xe8$C'\x86\x8e*T3t\x95\xa1+C\xb5\ny\x86F\x19\x1a\x19\xf2*\xd40\xd4\xcbP\xcfP\xa3B\x818\x8b\xe4,\xc4\t\x02\x04\xf5\x1f$h`\xa7t:td\xe8)AO(\x9a\x04y\x86\xfaI9\xd7\xde\x00\xb5\xc4\x19\xe4\x95Qk\x80\\%\xfe\xa5\xa9\x87\x06\x91!O\u0499\xbc\xc2\xca\x14(tQ\x1a\x84\x85uA\x86\xb0K\xa8\xbf\xfc\u02f9\xf4\x04\xfd\xa1A\xa1Bi\xedT\xce\b\xb9\x8f\x0e\xa5\xd4\xe9>L\x10J\xf3Wg\x06G\x87P\xa2!\xddgtt\b%\xc5AH\x96\xfe\xd6\xee\x02\x8e\x19\xc2\xdaMp\xa0\xa9\nN\x80d\xe9\xfe\a\xba\xa3c\x87PZ\xa2\xb3\x80\x93\x0f\xb96\xd6\xee\x11\xa1\a;]\xeb\xb2!l\x961Bc\xd2\x18\xb9P\x83P\xb3\x11\xa2\b\xfd\xda\b\u055f\x10m[Z\xdb\x01\x84\x9b\x9d_~\x80\xb4\xf2\x8b\x0e@ \xe5\xb7\bB(\xe57-BZ\u04ea\x0eB\xf9\xd7H\x13\x1d\x84\x04I\xbfj\x11R\xa5\x9d\xe8 \x84\x92}\x1cI\x90}\x1c\xd5$A\xf6\x01\x19:\x11\xb2\x8flO\x1aD\xde\x04U:TY\xa0\x96\x92\xdc\xe2\xa3!Ik\x80\xbc\x05\xf2\x06\xe8h\x81\x8e\x06\xa8\xb4@\xa5\x0e\x05Js\x8a\u03d84A\x85\x9a\x15\xf4\x88c-M\xa3B~\x05\xdd\xe3\xc0N\xe3U\xa8^Ac\x1ckij\xcb\xc3/\rt\b\x96-\x1f\x9a\xe75tP\xa1\x92L\xa9T\x88\x8c\x91\xa1\xf7\xf7\xa3\xf7\xf7\xa3\xef\x95\xff7\x8d\xda\xf7\u060ayi\xd8\x00\x00\x00\x00IEND\xaeB`\x82")

func assetsAvatarsJusticePngBytes() ([]byte, error) {
	return _assetsAvatarsJusticePng, nil
}

func assetsAvatarsJusticePng() (*asset, error) {
	bytes, err := assetsAvatarsJusticePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/avatars/justice.png", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/avatars/justice.png": assetsAvatarsJusticePng,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"avatars": &bintree{nil, map[string]*bintree{
			"justice.png": &bintree{assetsAvatarsJusticePng, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

