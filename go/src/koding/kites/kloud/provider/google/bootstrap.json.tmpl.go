// Code generated by go-bindata.
// sources:
// bootstrap.json.tmpl
// DO NOT EDIT!

package google

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

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

var _bootstrapJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x53\xbb\xae\xdb\x30\x0c\xdd\xfd\x15\x02\xd1\xf1\x3a\xbd\x1d\xfa\xca\x2f\xdc\xa5\x4b\xa7\xa2\x30\x18\x99\x71\xdd\x28\xa6\x41\x49\x09\xd2\x40\xff\x5e\xc8\x4f\xe6\xd1\xc2\x8b\x48\x1e\x1e\x1d\x1e\x5a\xd7\xc2\x18\xe8\x85\x4f\x6d\x4d\x02\x5b\x93\x63\x63\xa0\x61\x6e\x1c\x2d\xb1\x31\x60\x85\x6a\xea\x42\x8b\xce\xc3\xd6\xc0\xbb\xeb\x09\x65\x33\xc2\x2a\x55\x4b\xf0\x32\x77\xf4\xc2\xbf\xc9\x86\x07\xf4\x94\x57\x48\xa1\xa6\xe5\xee\x01\x38\xa6\x13\x0c\xb0\x54\x18\x93\x72\x07\x70\x0c\x7d\x0c\x93\xb6\x31\x25\xe4\x39\x8a\xa5\xfb\x01\x2a\xcb\xc7\x3e\x06\xaa\xf6\xad\xd0\x19\x9d\xd3\x13\x1d\xb8\x6e\xbb\x46\x65\x8c\x81\x0e\x8f\x99\x04\x1a\x4b\xe5\x58\x2f\x0f\x8e\x63\x5d\x2e\x04\x2f\x0a\x4c\xe1\xcc\x72\xc8\xf8\x9a\xf6\x18\x5d\xd0\x55\x74\x8e\xcf\x37\xec\xa3\x27\x81\x2d\x67\x1d\x10\x6c\xaf\xf0\xb9\xc8\x12\xb2\xb9\x3f\x54\xd2\x18\xf8\xf8\xe9\xf3\x97\xaf\xa0\x72\x3f\x97\x73\x52\xf7\x8d\x0e\x54\x82\x5d\x43\xf7\x2c\xf0\xba\x19\xbe\xf7\xaf\x2b\xcf\xcc\x92\x0a\xc5\x74\xef\xdb\xb4\xab\xea\x48\x01\x6b\x0c\xa8\xfd\x9b\x67\xbe\x31\xf0\x09\x70\x54\xe7\x7f\xbd\xd1\x45\xfd\x3a\xd1\x93\x54\xd9\xee\xb4\x1d\x13\x7d\xdc\xb9\xd6\x56\x07\xba\xa4\x55\x64\xba\x15\xb9\x2c\xfc\x84\xd2\xe2\x4e\xfd\xa1\xf0\x87\xbb\x61\xec\xff\xcb\x43\xdf\x62\x49\xe8\xc3\x87\xac\x64\x8d\x4a\xd4\xab\xa3\x28\xdc\x53\x79\xa6\x09\xa7\xe3\x72\xa7\x91\xd1\x97\x96\xba\x20\xe8\x06\xa0\x0a\x6f\x19\xa3\x5f\x6f\x9d\xcf\x0f\x4c\xcb\x7d\xf3\xb9\x44\x78\xba\xa4\xc5\xbb\xe7\xf3\xc2\xf5\xba\xf9\xee\x49\x06\x77\xe7\xe7\x33\xb5\xae\x2e\xff\xbb\xf7\xdb\x80\x79\xa3\x4b\xd2\x6f\xaf\x48\xc5\xdf\x00\x00\x00\xff\xff\x03\x06\x71\xf7\x2d\x04\x00\x00")

func bootstrapJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapJsonTmpl,
		"bootstrap.json.tmpl",
	)
}

func bootstrapJsonTmpl() (*asset, error) {
	bytes, err := bootstrapJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.json.tmpl", size: 1069, mode: os.FileMode(420), modTime: time.Unix(1475345133, 0)}
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
	"bootstrap.json.tmpl": bootstrapJsonTmpl,
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
	"bootstrap.json.tmpl": &bintree{bootstrapJsonTmpl, map[string]*bintree{}},
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
