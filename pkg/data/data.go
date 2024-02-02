// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package data generated by go-bindata.// sources:
// chart/crds/network.harvesterhci.io_ippools.yaml
// chart/crds/network.harvesterhci.io_virtualmachinenetworkconfigs.yaml
package data

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _chartCrdsNetworkHarvesterhciIo_ippoolsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x19\x5d\x6f\xdb\xba\xf5\xdd\xbf\xe2\x0c\x7b\x48\x0b\x44\x0e\x8a\x0e\xc3\x60\x20\xd8\x3c\xc7\x6b\x8d\x66\x59\x60\xbb\x1d\x86\xe1\x3e\x1c\x4b\xc7\x36\x1b\x8a\xd4\x25\x8f\xdc\xe4\xb6\xfd\xef\x17\xa4\x24\x5b\x92\x25\x59\x71\xdb\x8b\xf2\xc9\x3e\x24\xcf\xf7\x17\x8f\x82\x20\x18\x60\x22\x3e\x90\xb1\x42\xab\x11\x60\x22\xe8\x91\x49\xb9\x7f\x76\xf8\xf0\x37\x3b\x14\xfa\x6a\xf7\x6a\xf0\x20\x54\x34\x82\x49\x6a\x59\xc7\x73\xb2\x3a\x35\x21\xdd\xd0\x5a\x28\xc1\x42\xab\x41\x4c\x8c\x11\x32\x8e\x06\x00\xa8\x94\x66\x74\x60\xeb\xfe\x02\x7c\xfe\x3a\x00\x50\x18\xd3\x08\x44\x92\x68\x2d\xed\x50\x11\x7f\xd2\xe6\x61\xb8\x45\xb3\x23\xcb\x64\xb6\xa1\x18\x0a\x3d\xb0\x09\x85\xee\xd2\xc6\xe8\x34\x19\x41\xdb\xb1\x0c\x5d\x8e\x3e\x63\x6d\x76\x7f\xaf\xb5\xf4\x00\x29\x2c\xbf\x2b\x01\x6f\x85\x65\xbf\x91\xc8\xd4\xa0\xdc\x73\xe1\x61\x76\xab\x0d\xdf\x1d\xb0\x05\x6e\x57\x96\x7e\xe6\xc7\x84\xda\xa4\x12\x4d\x71\x79\x00\x60\x43\x9d\xd0\x08\xfc\xdd\x04\x43\x8a\x06\x00\xbb\x4c\x8f\x1e\x57\x00\x18\x45\x5e\x3d\x28\xef\x8d\x50\x4c\x66\xa2\x65\x1a\xab\x3d\xa5\x8f\x56\xab\x7b\xe4\xed\x08\x86\x4e\xf0\x42\x2b\x0e\xa3\x3f\x51\x68\xed\x6e\xba\xfc\xef\x7f\xe6\xef\x72\x18\x3f\x39\xb2\x96\x8d\x50\x9b\x06\x44\x8c\x9c\xda\xa1\x48\x76\x7f\x19\xe2\x0e\x85\xc4\x95\xac\x62\x1b\x7f\x18\xcf\x6e\xc7\xff\xbc\x9d\x56\xf0\x39\xfe\x36\x64\xba\x11\xa6\xd6\x4b\x79\xc0\xf5\x7e\x31\xbd\x79\x16\x9a\x50\xab\x4c\x27\xf6\xff\x7f\x7f\xf1\x8f\xa1\xbb\x74\x7d\x7d\x31\xa7\x8d\x70\xe6\xa5\xe8\xe2\xe5\x2f\xf9\xd1\x0a\x9d\xf9\xf4\xcd\x6c\xb1\x9c\xce\x6b\xd4\x4e\x28\xa1\x99\xd8\x04\xc3\x2d\xcd\x09\xa3\xa7\x16\x62\x93\xf1\xe4\xed\x74\x3e\x1d\xdf\xfc\xef\xdb\x89\x8d\x37\xa4\xb8\x8b\xd8\xf8\xcd\xf4\x6e\xd9\x9f\x58\x11\x68\xc3\xd0\x90\x8f\xb1\xa5\x88\xc9\x32\xc6\x49\x1d\x6b\x05\x5d\x84\x9c\x39\x41\xb6\xbd\x7b\x85\x32\xd9\xe2\xab\xcc\xb5\xc3\x2d\xc5\x38\xca\xcf\xeb\x84\xd4\xf8\x7e\xf6\xe1\xf5\xa2\x02\x06\x48\x8c\x4e\xc8\xb0\x28\x02\x25\x5b\xa5\xdc\x51\x82\x02\x44\x64\x43\x23\x12\xf6\x49\xe5\x4b\x50\xd9\x03\x70\x04\xb2\x5b\x10\xb9\x24\x42\x16\x78\x4b\x45\xf4\x50\x94\xf3\x04\x7a\x0d\xbc\x15\x16\x0c\x25\x86\x2c\xa9\x2c\xad\x38\x30\x2a\xd0\xab\x8f\x14\xf2\xb0\x86\x7a\x41\xc6\xa1\x71\x71\x9d\xca\x08\x42\xad\x76\x64\x18\x0c\x85\x7a\xa3\xc4\x6f\x7b\xdc\x16\x58\x7b\xa2\x12\x99\x2c\x7b\xc7\x35\x0a\x25\xec\x50\xa6\x74\x09\xa8\xa2\x1a\xe6\x18\x9f\xc0\x90\xa3\x09\xa9\x2a\xe1\xf3\x17\x6c\x9d\x8f\x7f\x6b\x43\x20\xd4\x5a\x8f\x60\xcb\x9c\xd8\xd1\xd5\xd5\x46\x70\x91\x51\x43\x1d\xc7\xa9\x12\xfc\x74\x15\x6a\xc5\x46\xac\x52\xd6\xc6\x5e\x45\xb4\x23\x79\x65\xc5\x26\x40\x13\x6e\x05\x53\xc8\xa9\xa1\x2b\x4c\x44\xe0\x05\x51\xde\xbf\x86\x71\xf4\x67\x93\xe7\x60\x5b\x21\x7b\xe4\x3b\xd9\xf2\x19\xf2\x19\xe6\x71\xc9\x13\x84\x05\xcc\x51\x65\x22\x1e\xac\xe0\x40\x4e\x75\xf3\xe9\x62\x09\x05\x27\x99\xa5\x32\xa3\x1c\x8e\x1e\xe9\xa5\xb0\x8f\xd3\xa6\x50\x6b\x32\xd9\xbd\xb5\xd1\xb1\xc7\x49\x2a\x4a\xb4\x50\xec\xff\x84\x52\x90\x62\xb0\xe9\x2a\x16\xec\xdc\xe0\xd7\x94\x2c\x3b\xd3\xd5\xd1\x4e\x7c\xd5\x81\x15\x41\x9a\x38\x67\x8f\xea\x07\x66\x0a\x26\x18\x93\x9c\xa0\xa5\x3f\xd8\x56\xce\x2a\x36\x70\x46\xe8\x65\xad\x72\x2d\xad\x1f\xce\xd4\x5b\xda\x28\x0a\xe6\x61\x35\xc7\xa9\x5b\x2e\x85\x4f\xb4\x5a\x8b\x4d\x7d\xa7\xeb\x96\x5b\xa1\x88\x4c\x13\xbc\x55\x86\xc3\x7a\x0c\x1e\xd2\x15\x19\x45\x4c\x36\xd8\xa1\x14\x51\xb9\x35\xa8\xaf\x00\x62\xb2\x16\x37\x2e\x0b\xcf\x6e\xe6\xce\x09\x45\x1c\xa7\x5c\x2a\x62\xf5\x65\x52\xe9\x38\x20\xb9\x86\xeb\x6b\xd0\x32\x5a\x90\x5c\x37\x9c\x8d\xda\x68\xae\xb5\x89\x91\x47\x5e\x3d\x8d\x07\x04\x53\xdc\x72\xb7\x87\x02\x62\x7c\x9c\x79\x04\xf0\xba\x43\x83\x68\x0c\x3e\x35\x71\xad\x63\x14\xca\x75\x04\x67\xe9\x3f\xbb\xbe\x20\xe7\xa2\xcd\x08\xbe\x4d\xb8\x6e\xe6\x25\xa1\x25\x57\xa0\xba\x78\x2f\x77\x0c\xd5\xa5\x38\xf9\x11\x3c\x1f\x0c\xd2\x6c\xee\x6e\x99\x5c\xf3\xd7\x4c\xba\x3b\x84\xdc\xa2\x7a\x1a\x2e\xaf\x93\x6e\xd8\x4b\x38\x78\x76\xc8\x41\x35\xec\xa6\x59\xea\x3f\x15\x75\xf0\x8c\xc8\xf3\xa2\x3f\x86\x32\x8d\x5a\x1c\xa1\xb7\xf8\x9d\x86\x87\xbe\xfa\xe9\x36\x70\xb6\xbe\x4d\x87\x99\xb0\x3f\x42\x8f\x96\xd1\xf0\x4f\xef\x44\x0b\xc7\xe5\xf7\x17\xdf\xd5\x7f\x61\xa8\x25\x88\x02\x17\x5f\x2d\x3b\x5e\x6d\x1d\xd1\x7e\x54\x57\x0f\xeb\xec\x02\x96\x47\x52\xc1\x34\x68\x15\x12\x58\x6a\xa6\x52\xa8\xe1\xe2\x4f\x5b\xb4\x2f\x72\x25\x0c\xf3\xa8\x79\x09\x5f\xbe\x80\x83\xdb\x32\xf0\xa2\x01\x91\xd1\x29\x53\x4b\xa9\x3e\xe9\x1b\x3f\xae\x96\xcf\x3d\x5b\xdf\xb3\x9a\x5b\xdf\x46\xce\xee\x7f\x3a\x51\x17\x39\x63\xdf\x4f\xd8\x76\xaf\x0f\x7c\x63\xd6\x00\x4e\x8a\x51\x48\x15\x5c\x28\xed\x68\xab\x33\x08\xfa\xab\xa2\xd1\xe2\x7d\xfc\xbf\xc9\xf7\x33\x57\xae\xba\x7e\x0e\xab\x7b\x7e\x69\x6e\x72\xcc\x54\x8c\x8f\xb7\xa4\x36\xee\xf5\xfc\xd7\x63\x67\xe8\x74\x84\xb3\x24\xbf\x3b\x30\x73\xca\x07\xfa\xd8\x3f\xc1\xd4\x36\xd9\x3e\x63\x7c\xa5\xb5\x24\x54\x95\xdd\x66\x7f\x09\xe0\x78\xba\x54\xc6\x74\xfc\xb2\xf0\xb3\x8a\xbe\x6f\x0b\xdc\x90\xe2\x7b\x1d\xcd\x69\xfd\xdc\xc7\x85\x3a\xb7\xb9\x55\xc5\xe4\xed\x8c\xdb\x1d\x1e\x7f\x18\xe4\x1c\xe3\x6d\x6d\x40\x4e\x35\x7f\x12\x2d\x2f\x0d\x2a\x2b\x8a\xa1\x4d\x5b\x21\xad\x3c\xcd\x6f\xd1\x32\xb0\x88\x29\x7b\x0e\x17\x9c\x01\xef\x51\x51\x94\xbd\x9d\xb5\x22\xa8\x8c\x97\x1a\x14\xa2\x01\x95\xe6\x2d\x99\xfa\x03\xb8\x97\xca\x0a\x31\xde\xfb\x07\x76\x6f\x11\x96\x7e\xc6\x72\x10\x43\xd8\x92\x1c\x9f\xd0\xb6\x3d\xd8\x7b\xf3\x54\x44\x5e\x1f\x66\xde\xa6\x31\xaa\xc0\x10\x46\x2e\x24\x8b\xab\x20\x54\x24\x42\xf4\x73\x8d\x88\x18\x85\xb4\x80\x2b\x9d\xb6\x95\x6a\xc8\x05\xda\x1b\xe1\x5c\xd6\x0d\xa1\xad\x4f\xce\x5a\x38\x77\x6a\xcc\x8e\xbb\xea\x56\x75\x87\x0b\x5b\x67\xe8\x6c\x65\x36\x45\x7d\x0b\x47\x0b\x7f\x34\x9b\xd1\x95\x98\xb9\xf4\xae\xa8\xd7\xb0\x34\x29\x5d\xc2\xbf\x50\x5a\xba\x84\xf7\xea\x41\xe9\x4f\xe7\xf3\xe5\x0f\xf4\xd2\xd3\x53\xe2\xa9\x87\x32\xb5\xae\x02\xed\xf9\x3a\x93\x74\x57\xcf\x19\xb4\x47\x5c\xe0\xf1\x36\x6c\x74\x96\xda\xf6\x97\x89\xeb\x64\x9e\x9b\x58\x51\x4a\x1d\xba\xd0\x6a\x56\x5c\xf9\xeb\xc4\xa9\x97\x6b\xcf\x41\x40\x6b\x1b\xbd\xff\x12\x71\xde\x24\xa0\xb9\x04\x9e\xbe\xd9\xd5\x3a\xd5\x3f\x8e\x94\xf7\x4a\xdf\x39\x7a\x89\x78\x48\x8b\xc7\x94\x8a\x5e\xd4\xed\x06\x2e\x07\xf6\xef\x41\x1a\x29\x1e\x01\x7d\x57\x17\x8d\x80\x4d\x9a\xe1\xb6\xac\x8d\x6f\x45\x0e\x90\x74\xb5\x1f\x18\x17\x1c\xe6\x91\x0e\x9f\xbf\x0e\x7e\x0f\x00\x00\xff\xff\xfb\x31\x8f\xae\x05\x1c\x00\x00")

func chartCrdsNetworkHarvesterhciIo_ippoolsYamlBytes() ([]byte, error) {
	return bindataRead(
		_chartCrdsNetworkHarvesterhciIo_ippoolsYaml,
		"chart/crds/network.harvesterhci.io_ippools.yaml",
	)
}

func chartCrdsNetworkHarvesterhciIo_ippoolsYaml() (*asset, error) {
	bytes, err := chartCrdsNetworkHarvesterhciIo_ippoolsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "chart/crds/network.harvesterhci.io_ippools.yaml", size: 7173, mode: os.FileMode(420), modTime: time.Unix(1706866345, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartCrdsNetworkHarvesterhciIo_virtualmachinenetworkconfigsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\x38\x60\x0f\x6d\x81\xc8\x46\xb0\x61\x1b\x0c\x04\x5b\xe0\x76\x5b\xb0\xa4\x0b\x9a\x34\x2f\xc3\x1e\xce\xe2\xd9\x66\x43\x91\x1a\xef\xe8\x24\xeb\xfa\xdd\x07\x92\x52\x2c\xcb\x76\xe2\x1a\xc1\xf8\x26\xf2\xfe\xfc\xee\xff\xa9\x28\x8a\x01\xd6\xfa\x86\x3c\x6b\x67\xc7\x80\xb5\xa6\x7b\x21\x1b\xbf\x78\x78\xfb\x23\x0f\xb5\x1b\x2d\x8f\x07\xb7\xda\xaa\x31\x4c\x02\x8b\xab\x3e\x10\xbb\xe0\x4b\x7a\x4b\x33\x6d\xb5\x68\x67\x07\x15\x09\x2a\x14\x1c\x0f\x00\xd0\x5a\x27\x18\xaf\x39\x7e\x02\x7c\xfe\x32\x00\xb0\x58\xd1\x18\x96\xda\x4b\x40\x53\x61\xb9\xd0\x96\x2c\xc9\x9d\xf3\xb7\xa5\xb3\x33\x3d\xe7\x61\xf3\x39\x5c\xa0\x5f\x12\x0b\xf9\x45\xa9\x87\xda\x0d\xb8\xa6\x32\x4a\x9a\x7b\x17\xea\x31\xec\x22\xcb\x3a\x1a\x9d\x19\xef\x4d\x56\x77\x91\xd5\xbd\xcf\x8c\x93\xa4\x2e\x51\x19\xcd\xf2\xfb\x73\x94\xe7\x9a\x25\x51\xd7\x26\x78\x34\x4f\x1b\x91\x08\x79\xe1\xbc\xbc\x5f\x81\x29\x60\x59\x59\x92\x72\x36\xef\x7d\x36\xe4\xda\xce\x83\x41\xff\xa4\xe4\x01\x00\x97\xae\xa6\x31\x24\xc1\x35\x96\xa4\x06\x00\xcb\x1c\xb8\xa4\xa8\x00\x54\x2a\xc5\x03\xcd\xa5\xd7\x56\xc8\x4f\x9c\x09\x95\x7d\x84\xf1\x89\x9d\xbd\x44\x59\x8c\x61\x18\x9d\x3a\x5c\x56\x51\x58\x7a\x6c\x23\x74\x73\xf1\xfe\xf4\xe2\x5d\x73\x25\x0f\x51\x21\x8b\xd7\x76\xbe\x45\x84\xa0\x04\x1e\x96\xce\x66\xad\xfc\xe7\x4f\xaf\x7f\x1e\x46\x9e\x93\x93\x57\xa7\xc6\xb8\x12\x85\xd4\xab\x37\x7f\x35\x94\x6b\x7a\x4e\xcf\xcf\xff\x98\x9c\x5e\xbf\x7b\xbb\x97\xaa\x36\xbf\x86\xa5\xa7\x94\x5a\xd7\xba\x22\x16\xac\xea\x75\xa1\xbf\xae\x23\x57\x28\xd9\xba\x26\xfb\x8e\xd1\xd4\x0b\x3c\xce\x5e\x2f\x17\x54\xe1\xb8\xa1\x77\x35\xd9\xd3\xcb\xb3\x9b\x6f\xaf\xd6\xae\x01\x6a\xef\x6a\xf2\xa2\xdb\x58\xe6\xd3\x29\x99\xce\x2d\x80\x22\x2e\xbd\xae\x25\xd5\xd2\xbf\xc5\xda\x1b\x40\x54\x90\xb9\x40\xc5\xda\x21\x06\x59\x50\x1b\x43\x52\x0d\x26\x70\x33\x90\x85\x66\xf0\x54\x7b\x62\xb2\xb9\x9a\xe2\x35\x5a\x70\xd3\x4f\x54\xca\xb0\x27\xfa\x8a\x7c\x14\x13\x53\x2f\x18\x05\xa5\xb3\x4b\xf2\x02\x9e\x4a\x37\xb7\xfa\x9f\x47\xd9\x0c\xe2\x92\x52\x83\x42\x2c\x90\xb2\xc4\xa2\x81\x25\x9a\x40\x47\x80\x56\xf5\x24\x57\xf8\x00\x9e\xa2\x4e\x08\xb6\x23\x2f\x31\x70\x1f\xc7\x85\xf3\x04\xda\xce\xdc\x18\x16\x22\x35\x8f\x47\xa3\xb9\x96\xb6\x91\x94\xae\xaa\x82\xd5\xf2\x30\x2a\x9d\x15\xaf\xa7\x41\x9c\xe7\x91\xa2\x25\x99\x11\xeb\x79\x81\xbe\x5c\x68\xa1\x52\x82\xa7\x11\xd6\xba\x48\x86\xd8\x94\x5d\xc3\x4a\x7d\xe3\x9b\xd6\xc3\x6b\x6a\x37\x72\x27\x9f\xd4\x03\xbe\x22\x3c\xb1\x13\x80\x66\xc0\x46\x54\x36\x71\x15\x85\x78\x15\x5d\xf7\xe1\xdd\xd5\x35\xb4\x48\x72\xa4\x72\x50\x56\xa4\x1b\x7e\x69\xe3\x13\xbd\xa9\xed\x8c\x7c\xe6\x9b\x79\x57\x25\x99\x64\x55\xed\xb4\x95\xf4\x51\x1a\x4d\x56\x80\xc3\xb4\xd2\x12\xd3\xe0\xef\x40\x2c\x31\x74\x7d\xb1\x93\xd4\x6c\x61\x4a\x10\xea\x98\xec\xaa\x4f\x70\x66\x61\x82\x15\x99\x09\x32\xfd\xcf\xb1\x8a\x51\xe1\x22\x06\x61\xaf\x68\x75\x47\x48\x9f\x38\xbb\xb7\xf3\xd0\x8e\x84\xd5\xd9\x5e\xa7\xf1\xd8\x6e\x33\xdf\x78\x05\xd0\x42\xd5\x96\xeb\xa7\x44\x36\x8c\xf5\xa9\x52\x9e\x78\xc7\x33\xc0\xcc\xf9\x0a\x65\x0c\xba\x5e\x7e\xb7\x83\x64\x87\x33\x56\xa7\xc2\xf2\x19\x2d\x15\xde\x9f\x93\x9d\xc7\x3e\x79\xfc\xc3\xa1\x6a\x1a\x27\xc5\x71\xb0\x87\x9e\xef\x0f\x34\x27\x66\xb2\xf6\xa4\xb6\xa9\x28\x3a\xa6\x6e\x7d\xee\x40\xdc\xf2\xbe\x23\x51\x1e\xa1\x9f\xa5\x28\xc3\x26\xf0\xcc\x88\xde\xe3\xc3\xc6\xdb\x7d\x71\x1b\xa6\xe4\x2d\x09\x71\xb1\x44\xa3\x55\x77\xaf\xe9\xa1\x27\x66\x9c\xc7\xf9\xdc\xcd\xb7\x54\xf0\xce\x9a\x87\x58\xa3\xa8\x14\xf5\xdb\x6b\x72\x4b\x30\x11\xbb\x51\x57\x64\x66\x43\x34\xe6\xf5\xfd\x11\xdc\x83\xb6\xc0\x64\x66\x6f\x7a\x1c\x35\x06\xde\xe6\xc2\x6c\xc8\xd4\x39\x43\x68\x7b\xaf\x79\xd0\x6f\xf2\x3c\x1d\xd3\x27\xa3\x79\x90\x6f\x6e\x2e\x22\x8e\xd8\x64\x75\x55\x05\xc1\xa9\xd9\x16\xcb\xec\x8f\x68\x3b\x9c\x9c\xb4\x7e\x59\xa3\xdb\x9e\x48\x05\xac\x2d\x34\x5d\x23\x36\x3b\x48\x5a\x49\xf6\xed\x21\xab\x1d\xe7\x05\xfb\x87\x41\x96\x6b\x8f\x96\x75\xbb\xd1\xec\x2a\xbd\xb5\xb9\x75\x8e\x2c\x20\xba\xa2\x3c\x2b\x5a\x64\x20\x8f\xa2\x48\xe5\xc1\xe2\x2c\xc1\xda\xea\xb5\x79\xc4\x01\x5a\x27\x0b\xf2\xfd\xe9\xf0\x48\xf1\x5c\xeb\x88\x66\x7c\x4c\xd3\x67\x6f\x13\xae\xd3\x02\xb2\x32\x43\x73\xc7\x8e\x3b\xe4\x5d\xd3\x6c\x6f\x4c\x6d\xc2\xed\x03\xe6\xb7\x50\xa1\x2d\x3c\xa1\x8a\xe9\xd8\xb2\x82\xb6\x4a\x97\x98\x86\xbe\x22\x41\x6d\x18\x70\xea\xc2\x66\x73\xe9\xfa\xa1\x13\x84\x43\xa1\x7b\x42\xee\xaf\x95\x3b\x90\x47\x37\x66\xf2\x38\x6a\xd6\xd3\xe1\x15\xf7\x01\x1d\xec\xcc\x6d\xa5\xb2\x03\xd1\x55\x22\xcd\x0b\x6c\x07\xcc\x51\x4a\x45\x37\x83\x6b\x1f\x97\xcc\x5f\xd0\x30\x1d\xc1\x47\x7b\x6b\xdd\xdd\xe1\xb8\x12\xc1\x5e\x7e\x7a\xa8\x93\xf6\xd2\x84\xf8\xbf\xb8\xc2\x75\xa0\xea\xa7\xc7\xd8\xce\x8a\x2b\x92\xdc\xaf\x9d\x5d\xbb\xe7\xd3\xda\x62\xf3\x82\x7d\x09\xdb\x5f\xb7\xb3\xcb\x67\x56\x8f\x17\xd8\x5e\x5e\x62\x33\xd9\x2b\x83\x0f\xe5\x3e\x28\x38\x5b\x99\x36\x2e\x39\xfe\x14\xa8\x31\x88\x0f\x39\x2d\x58\x9c\x4f\x73\x72\x75\x13\xa6\x8f\xff\x3c\xad\x01\x4d\x3d\xc2\xe7\x2f\x83\xff\x02\x00\x00\xff\xff\xe9\x50\xbe\xbe\xbf\x11\x00\x00")

func chartCrdsNetworkHarvesterhciIo_virtualmachinenetworkconfigsYamlBytes() ([]byte, error) {
	return bindataRead(
		_chartCrdsNetworkHarvesterhciIo_virtualmachinenetworkconfigsYaml,
		"chart/crds/network.harvesterhci.io_virtualmachinenetworkconfigs.yaml",
	)
}

func chartCrdsNetworkHarvesterhciIo_virtualmachinenetworkconfigsYaml() (*asset, error) {
	bytes, err := chartCrdsNetworkHarvesterhciIo_virtualmachinenetworkconfigsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "chart/crds/network.harvesterhci.io_virtualmachinenetworkconfigs.yaml", size: 4543, mode: os.FileMode(420), modTime: time.Unix(1706866345, 0)}
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
	"chart/crds/network.harvesterhci.io_ippools.yaml":                      chartCrdsNetworkHarvesterhciIo_ippoolsYaml,
	"chart/crds/network.harvesterhci.io_virtualmachinenetworkconfigs.yaml": chartCrdsNetworkHarvesterhciIo_virtualmachinenetworkconfigsYaml,
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
	"chart": &bintree{nil, map[string]*bintree{
		"crds": &bintree{nil, map[string]*bintree{
			"network.harvesterhci.io_ippools.yaml":                      &bintree{chartCrdsNetworkHarvesterhciIo_ippoolsYaml, map[string]*bintree{}},
			"network.harvesterhci.io_virtualmachinenetworkconfigs.yaml": &bintree{chartCrdsNetworkHarvesterhciIo_virtualmachinenetworkconfigsYaml, map[string]*bintree{}},
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
