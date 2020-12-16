// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.graphql
// types/comment.graphql
// types/todo.graphql
// types/user.graphql
package schema

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8d\x3d\x0e\xc2\x30\x0c\x85\x77\x9f\xe2\xb1\x95\x2b\x78\x66\x61\x60\x40\xc0\x84\x18\xa2\xc6\x82\x0c\xad\xc1\x71\x86\x0a\xf5\xee\x28\xa5\xad\x3a\xd9\x7a\x3f\xdf\xcb\xed\x4b\xba\x80\x2f\x01\x9f\x22\x36\x30\xce\xf5\x10\xd0\x15\x0f\x9e\xb4\x67\x9c\xe6\x8f\x46\x22\x1f\xde\xf2\x8f\x4c\x1d\xd7\xa8\x4d\x49\x91\x71\x3c\xec\xf6\x8c\xab\x46\x9d\xe5\xbc\xd5\xef\xd5\x78\x10\x50\xb2\xd8\xd6\xb8\x65\xb1\x95\xbb\x0c\x4d\xe8\xd6\x24\xb8\xd4\x5e\x53\x71\x8c\x8b\x5b\xea\x9f\xeb\xca\x48\xbf\x00\x00\x00\xff\xff\x14\x12\x38\xda\xbc\x00\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 188, mode: os.FileMode(420), modTime: time.Unix(1608112962, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesCommentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\xce\xcf\xcd\x4d\xcd\x2b\x51\xa8\xe6\x52\x50\x28\xcd\x4c\xb1\x52\x08\x2e\x29\xca\xcc\x4b\xe7\x52\x50\xc8\x4b\xcc\x4d\x85\x73\x6b\xb9\x00\x01\x00\x00\xff\xff\xe2\xa3\xf6\xd5\x2e\x00\x00\x00")

func typesCommentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesCommentGraphql,
		"types/comment.graphql",
	)
}

func typesCommentGraphql() (*asset, error) {
	bytes, err := typesCommentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/comment.graphql", size: 46, mode: os.FileMode(420), modTime: time.Unix(1608063930, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesTodoGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\xc9\x4f\xc9\x57\xa8\xe6\x52\x50\x28\xcd\x4c\xb1\x52\xf0\x74\xe1\x52\x50\xc8\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\xe7\x52\x50\x48\xce\xcf\xcd\x4d\xcd\x2b\xd1\x00\xc9\x43\x04\x35\xad\x14\xa2\x9d\x21\xa2\xb1\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\xe3\x42\xa3\x08\x48\x00\x00\x00")

func typesTodoGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesTodoGraphql,
		"types/todo.graphql",
	)
}

func typesTodoGraphql() (*asset, error) {
	bytes, err := typesTodoGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/todo.graphql", size: 72, mode: os.FileMode(420), modTime: time.Unix(1608107140, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2d\x4e\x2d\x52\xa8\xe6\x52\x50\x28\xcd\x4c\xb1\x52\xf0\x74\xe1\x52\x50\xc8\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\xe7\xaa\xe5\x02\x04\x00\x00\xff\xff\xf7\xa6\x57\x89\x27\x00\x00\x00")

func typesUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesUserGraphql,
		"types/user.graphql",
	)
}

func typesUserGraphql() (*asset, error) {
	bytes, err := typesUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/user.graphql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1608113362, 0)}
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
	"schema.graphql":        schemaGraphql,
	"types/comment.graphql": typesCommentGraphql,
	"types/todo.graphql":    typesTodoGraphql,
	"types/user.graphql":    typesUserGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
	"types": &bintree{nil, map[string]*bintree{
		"comment.graphql": &bintree{typesCommentGraphql, map[string]*bintree{}},
		"todo.graphql":    &bintree{typesTodoGraphql, map[string]*bintree{}},
		"user.graphql":    &bintree{typesUserGraphql, map[string]*bintree{}},
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
