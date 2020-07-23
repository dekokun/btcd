// Code generated by go-bindata.
// sources:
// sample-bchd.conf
// DO NOT EDIT!

package main

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

var _sampleBchdConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x7b\x7f\x6f\x1b\x37\xd2\xff\xff\x7a\x15\x83\xc3\x1d\xec\x14\xb2\x2c\x39\x76\x92\xca\x51\x00\xc7\x69\xef\xfc\xfd\x26\xb1\x11\x27\xbd\x3b\x14\x45\x41\xed\xce\x6a\x79\xe6\x92\x5b\x92\x2b\x59\x7d\x70\x7d\xed\x0f\x66\x48\xee\x0f\x39\xae\xd3\x3e\x71\xfe\x68\xad\x5d\x72\x38\x9c\x9f\x9f\x19\x72\x7f\x3c\xab\x6b\x25\x33\xe1\xa5\xd1\x70\x59\xd3\xff\xdc\x4f\xa3\xd1\x29\x1c\x7c\xd5\x7f\xa3\x53\x78\x23\xbc\x00\x87\xde\x4b\xbd\x72\x5f\x7f\x81\xd1\x29\x7c\x2c\x11\x72\x69\x31\xf3\xc6\x6e\xc1\x1b\x70\xde\x58\x84\x9c\x17\x6e\xb2\x12\x84\x03\x5f\x22\x2c\x95\xc9\x6e\x20\x2b\x85\xd4\x20\x74\x0e\x35\xa2\x05\x91\xe7\x16\x9d\x43\x37\x01\x22\x34\x3a\x1d\x0c\xf3\xe2\x06\x1d\x38\x5c\xa3\x15\x0a\xfe\xfe\x7a\x0c\xce\x80\x2f\xa5\x03\x65\xa2\xf0\xaa\xc6\x79\x28\xc5\x1a\x41\x80\x32\x1e\x4c\x01\x85\x45\x04\x57\x8b\x0c\x27\x89\x3d\x2c\x44\xa3\x3c\x48\x07\xbf\x1d\x4e\x96\x59\x99\x1f\x32\x7b\x46\xc3\xd5\xe5\xf5\xc5\xbf\xe0\xf2\x1a\xdd\x18\xfe\xfa\xf6\xf2\xfc\xec\xed\xd9\xd5\xd5\x9b\xb3\x8f\x67\x87\xaf\xfb\xc3\xfe\x29\x75\x6e\x36\x6e\x3c\x3a\x85\xdf\x0e\xdf\xca\xa5\x15\x76\x7b\xd8\x57\xe2\x75\x53\xd7\xc6\xfa\xe1\xac\x77\x22\x83\xcb\xeb\x31\x6f\xf7\xaf\xa5\xa9\xf0\xb0\xbf\xf6\xe8\x14\xae\x94\xd0\xdf\x4e\x00\xbe\xd3\x6b\x69\x8d\xae\x50\x7b\x58\x0b\x2b\xc5\x52\xa1\x03\x61\x11\xf0\xb6\x16\x3a\xc7\x3c\xec\x1c\xb7\x50\x89\x2d\x2c\x11\x1a\x87\xf9\x04\xe0\xfd\xe5\xc7\xef\xe6\x89\xbb\xd1\x29\xe0\xbd\x84\xfc\xb6\x96\x99\x50\x6a\x0b\x7f\xfb\xe1\xec\xc3\xc5\xd9\xeb\xb7\xdf\xfd\x6d\x0c\xcb\xc6\x47\xb2\x24\xc7\x25\x82\xc8\x32\xd2\x47\x0e\x1b\xe9\xcb\xd1\x29\xfc\x35\x0d\x86\x12\x2d\x4e\x00\xce\x94\x33\x63\xf8\x8d\x64\xd9\xf2\xe6\xcd\x50\x76\x3d\x89\x91\x0a\x68\xbf\xb9\xb4\x8b\xbe\xec\x47\x8f\x62\xed\xef\xd1\x6f\x8c\xbd\x79\x5c\x83\xff\xe4\x10\x3c\x3a\xaf\xd1\xd3\xee\xe2\x9f\x8b\x59\xfb\xae\x44\xb0\xb8\x22\xbb\x26\xcb\xa0\xf7\xa0\x03\x63\x34\xde\xe2\x8a\x1e\x0d\xc7\x3b\x59\x35\x2a\x58\x52\x7f\xfc\xe8\x94\xde\xb4\xd4\xcf\x8d\xd6\x98\x79\x58\x4b\x01\x02\xae\x2f\xcf\xff\xff\xf5\x09\xd4\xd6\xdc\x6e\x5b\x53\xb8\xae\x31\x93\xc5\x56\xea\x15\x88\xf0\x0a\x36\x52\x29\xc8\xa5\x23\x5b\x00\x25\x9d\x47\x2d\xf5\x6a\x74\x0a\x85\xb1\x20\x75\x66\x2a\x1a\x9d\x05\xda\x14\x89\xa0\xd1\x0a\x9d\x8b\x63\x3b\x1f\x65\x3b\xaa\xad\x59\x4b\x52\x3a\x31\x41\xac\xef\x85\x61\x7b\xa3\x53\x30\x1c\xc9\x68\x93\xbc\xf2\x62\x76\xf4\x7c\x32\x9d\x4c\x27\xb3\xf9\xb7\xd3\x93\x69\x7a\xdc\x38\xb4\x8b\xf4\xa3\x16\xce\x2d\x52\x18\xe9\xef\x08\xc4\xd2\xac\x91\xcc\x4c\x38\xd7\x54\xc1\xca\x96\x08\x1f\x8d\x85\xfd\xd2\xfb\xda\xcd\x0f\x0f\x37\x9b\xcd\xc4\x1b\x5b\x5b\xf3\x1f\xcc\xfc\xc4\xd8\xd5\x13\x5a\xfd\xa2\x60\xce\x02\x1d\xe9\x40\x1b\x0f\xde\x58\x7e\x58\x18\xa5\xcc\x86\x76\xdc\xf3\x24\xa2\x5d\x5b\x5c\x93\xdb\x34\x8e\x5e\x7a\x63\x49\xf8\x2c\x4d\x99\x05\xce\xe0\x97\x06\xad\x44\x47\xa3\x95\x31\x37\x4d\xdd\x93\xcd\x3e\xc7\x25\xa9\x33\x8b\x82\x65\xa5\x8d\xde\x56\xd2\x6f\x61\x53\xa2\x8e\xf4\xa4\x0b\xcb\x2d\xb7\x69\x39\x5a\x6b\x6b\x1a\x0b\x17\x57\xb0\x44\xfa\xa5\x50\xdc\x44\xf1\xbe\x79\x7f\xcd\xfb\xd1\xc6\x68\x69\x74\x67\x32\x42\x83\x50\x1e\xad\x16\x5e\xae\xd3\x46\xbd\x49\x4a\xa4\x3f\x27\x3c\xa5\x1f\x60\x3f\xf6\x45\x12\x85\x4a\xf6\xcb\x62\x15\x2c\x58\x6d\x72\x9c\xc0\x7b\xa3\xef\x4c\x0f\x56\xb4\x44\x5a\xc2\x8b\xcc\xc7\x08\xc1\x22\xad\x28\x56\x33\x65\xb2\x01\xcb\x2f\x4c\xe3\x5b\x03\x94\x05\x68\xa3\x59\x97\x2e\xf8\x4c\xdc\x4e\xdf\x3c\x66\xe9\x71\x32\x0f\xfe\xd1\x9a\xc7\x77\x9a\xcd\x97\x98\x74\xde\xa2\xa8\x40\x3a\x13\x3d\x66\xb9\x05\x2b\x74\x6e\x2a\xf9\x2b\x09\x30\x2c\x4a\x64\x20\xb3\x98\x93\x90\x85\x72\x60\xd1\x35\x8a\xe5\x2d\x29\xf4\x12\x25\xd2\x95\x0f\xae\xa2\x71\x03\x99\xb4\x59\x23\x3d\xfb\x05\x8a\xac\xec\xf9\x04\xa7\x27\xe9\xa0\xe2\x8c\x24\x3d\x54\x9c\xe3\x64\x51\xc8\xac\x51\x3e\x88\x31\x33\xd6\xa2\x12\x1e\xfb\xce\xc4\x21\xc2\xd8\x96\xdb\x4e\x89\x9f\xb4\x5c\xa3\x75\x42\xc1\x95\x6a\x56\x9c\x24\xae\x94\xd8\xc2\xfe\xa7\x2b\x7d\xf5\x84\x08\x8a\xc6\x9b\x4a\xf8\x18\xb7\x4d\x4d\x76\x54\x26\x1f\x06\xca\x37\xa3\x53\x9e\x67\x96\x9e\xf3\x65\x49\x09\x83\xed\x42\x91\x41\x45\xed\x41\x61\x4d\x05\x2e\x64\x28\xcc\x21\xc7\xb5\xcc\x38\xe5\x86\x98\xc1\x3b\x0b\xbe\x3b\x3a\x0d\x8a\xe6\x8c\xaa\x0d\x60\x51\x90\x3d\xc9\x02\xf0\x76\x97\x6c\x0c\x08\xd1\x47\x30\xa7\x9d\x36\xb5\xae\xc3\x0e\x63\x20\xba\x8f\x27\x74\xc1\xec\xc9\xe2\xc8\x30\xda\xb8\x04\x2c\xec\x4b\x8d\x2d\xf7\x35\x92\xeb\x28\xa9\x29\xfd\x50\x02\x09\x2c\x92\x53\x47\x63\x84\xa7\x36\x3f\xa8\x85\xf5\x5b\x70\xd2\x07\x07\x8d\x22\x69\x97\x96\x3d\x67\x0d\xb8\x80\xd4\x89\x42\x3b\xda\xdd\xd6\x34\xbc\x99\x25\x96\x52\xe7\xf0\xfe\xec\xe3\xb8\xc7\x5f\xbb\x1e\x39\x0a\x99\x21\xe9\x26\x5f\xa3\xf5\x92\x9c\x91\x83\xba\xc8\x4a\x7e\x95\xb8\x8e\x31\x94\x08\xbb\x28\x0a\xe9\x39\x89\x92\x8b\x60\x30\x67\x8e\x9f\x24\xb3\xbd\x28\x7f\xd8\x17\x9a\x9c\x32\x22\x9a\x5d\x9d\x71\x34\x48\x5b\x92\xf5\x62\x36\x39\x9a\x3c\x9d\x1c\x0f\x1f\x1e\x4d\xa7\x47\xf3\xf9\xec\xe8\xe9\x31\xe9\xe1\x9b\xaf\xfa\x8f\x14\xdb\x54\x95\xb0\x5b\x42\x5a\x7b\x22\xcf\x09\xc6\xed\x01\x19\x72\xe3\x60\x2f\x1a\xfe\xde\x64\x74\x3a\x3a\x85\x4b\x4d\x66\xab\x91\xc6\x0e\x63\xaf\xdf\x98\xb8\x63\x37\xee\x91\x21\x5b\x6e\x69\x8c\x63\x84\xee\xe2\x30\x43\x9e\xd7\xc6\x97\x20\x88\x10\x0b\x97\x00\x67\x94\x2f\xd1\x20\xac\x29\x3c\xbf\xd9\x08\xed\x03\x1e\x15\xdb\xe4\x91\x1c\xf1\x28\x3e\xb5\xd0\x87\x54\x2e\xd6\x92\x52\xa1\x03\xa7\xe4\xaa\xf4\x6a\xcb\x9e\x8d\x16\xb5\xa7\x05\xb7\x09\x40\x8e\x7b\xe6\x47\xb9\x60\x4b\x11\x3a\xf8\x7e\x21\x23\xa4\x75\xc3\x38\xcc\x90\xb5\x67\x0b\x49\xb1\x29\xb1\x50\xb6\x30\x3a\x01\x65\xa2\x55\x1a\x47\x88\xd5\x65\x56\x2e\x29\x55\xa0\x32\x1b\x36\x46\x0a\x6c\x4b\xb1\x54\x5b\xd8\x18\xbd\x47\x18\x21\x24\xae\xca\xe4\xb4\x7b\xa1\xb7\xbe\x24\xd9\x32\x50\x63\xf9\x77\x82\xcd\x0d\x86\x34\x98\xb2\x5c\x3f\xdd\x93\xb7\xf8\x12\x6d\xe4\x3f\x97\x2e\x33\x6b\xb4\x98\x73\xe0\x60\x9e\x30\xbe\x4b\x7e\xd2\x8a\x93\x5d\x41\xe7\x20\x94\x33\xa0\xd0\xbb\x08\x7f\x2a\xe3\xd3\x9c\x1b\x1d\x55\x25\x2c\xe9\x52\xac\x85\x54\x6c\xfd\x09\xd2\x66\x42\x13\x6f\xb4\x89\x3e\x1f\xed\xbb\x61\x62\xdb\x9a\x26\x46\xe3\x16\x71\x40\x45\x6a\x8b\xc9\xbc\x68\x54\xdf\xa3\x49\xb9\x21\x29\x2c\x15\x56\x8e\x15\x15\x43\x3e\xb9\x36\xc5\x7a\x67\x2a\x0c\x3e\x4c\xaa\xd8\xaf\xd1\x96\xa2\x76\x90\x37\xc1\xd1\xa1\x90\x16\x37\x42\xa9\x27\x51\xaa\x9d\x81\x9a\x10\x73\x03\xd7\xa5\xd0\xf9\x38\x18\xc7\xe5\xfb\xb7\xff\xee\xf3\xcc\x10\x2f\xd9\x70\xdc\x5e\x70\x74\x1d\x65\x4f\xd1\xf8\xc2\x07\x31\x46\xac\xd6\x0f\x8a\xfb\x3d\x13\xc2\x5b\x2a\x3b\x24\x99\xa9\x43\x1f\x07\x91\x60\xdb\xc8\xba\x0b\xcd\xa2\x98\x9e\xb0\xa6\xde\xbc\xbf\x06\x87\x48\x42\x60\xe3\x64\x57\xe9\x02\x1c\x13\x8a\xa1\x2d\xa7\xda\x8d\xb0\x46\xab\x32\x2e\xda\xe2\x86\x3a\x8b\xe8\xed\x94\x56\x08\xe6\x49\x95\x44\x4d\x09\xaa\x6f\x6a\x9c\x89\x84\x1e\x28\x7a\x02\x70\x6d\xc6\x81\xe1\x24\xda\xa4\xd8\x90\x7f\xe4\x1a\xd5\x36\xf8\x3c\x29\x3d\xba\xbd\xa1\xe8\xd2\x5b\xfa\x2f\xde\x36\xce\x63\xfe\x97\x48\xf6\xeb\x07\xbf\xd1\x29\x9c\xe5\xa4\x3f\xeb\x58\xb0\xfe\x73\x1e\x4f\x32\xcb\xd1\x49\xcb\xd1\x8a\x12\x19\x0b\xad\x46\x1b\x72\xd8\xe8\x14\xfe\x6d\x1a\x8e\x6d\x29\x70\x31\xd8\xe8\xe5\x6b\x46\x56\x43\x20\x65\xac\x0f\x25\x72\x5b\xcc\xd2\x23\x56\x1c\x15\xcd\x9c\x77\x48\x5f\x03\xc4\x20\x0b\x88\xb8\x8b\x74\xdb\x19\x60\x8c\x10\x10\xc3\xc3\x62\xf6\xed\xd1\x64\xf6\xec\xc5\x64\x36\x99\xf5\x9f\x4e\x19\x9d\x1d\xcd\x5f\x3c\x7d\xfa\xb4\xf7\xbc\xc0\x17\xd3\xf9\xbc\x3f\xf2\xc7\xf0\xe8\xe8\xa7\x30\xf4\x5e\x31\xa5\xc8\xcc\xee\x91\xc2\xf3\x43\x92\xa3\x4a\x21\xc9\x0e\xfe\x4f\xa2\xa3\x52\x74\x57\x78\x7f\x56\x74\x77\xaa\x2d\xdf\x81\x28\x28\x85\x8b\x06\xee\x64\x8e\xd1\x88\x5d\xdc\x5e\x8c\xeb\xb1\xbc\xd1\x31\xbc\xde\x9f\x4a\xc1\xc5\x84\xeb\x22\x14\xed\x5c\x6a\x47\x71\xed\xd3\x1d\xc5\xa5\xe7\x9d\xe2\xd2\x93\xbb\x8a\x7b\x27\x6e\x65\xd5\x54\xa0\x9b\x6a\x89\x96\x12\xb7\xd4\x4b\xd3\x50\x84\x27\x9c\xd9\xf8\xf0\xa3\xf5\xb0\x4a\xdc\xf2\xdf\x8b\xd9\xd1\x49\x9c\xff\x45\x73\x59\xa7\x17\x57\x7d\x12\x35\x5a\x59\x2f\x98\xca\x1b\x4a\x41\xc1\x2c\xdc\x56\x67\x71\x8a\x53\x66\x43\xe1\xa7\xa4\x9c\x40\xe2\xf6\xa5\x45\x57\x1a\x95\x83\xa4\x2a\xc0\xa3\x3b\x74\x98\x31\x4d\xa9\x69\x22\xcd\x8b\xa5\xb4\xab\x11\xf3\xc5\xc9\xec\x68\x3a\xa5\x15\xde\xb7\x3c\xb6\x7c\xed\xa4\x44\xaa\x6a\x08\x42\x32\x80\x17\x76\x85\x3e\x8d\x0c\x1b\x7e\x11\x19\xe5\xb8\xb8\x14\x3a\x80\xd7\x02\x2a\xe9\x02\xa4\xa0\xac\x93\xc4\xa4\x4d\x1c\x11\xc0\x71\x92\x32\x87\x34\x4a\xf2\x42\x83\xcb\x0c\x03\xd0\x22\x54\x16\x49\x00\x9c\xab\x74\xde\xae\xf0\x59\xf2\x4b\xa1\x5b\x59\x2c\x66\x61\x87\xff\x30\x1b\x50\x86\x4b\x59\xa6\x7f\x77\x22\xfc\x20\x94\xcc\xc1\xcb\x0a\xa1\xd1\xd2\x07\x3c\xff\x3f\x6e\x0c\xd5\x18\xca\xff\x12\xe1\x77\x52\x33\xa3\xb3\xb4\x4c\xde\xd8\x50\xc6\x1c\x1d\x97\x3b\x4f\x66\xb3\xf2\xe9\xb4\x9a\x9d\xb8\x14\x00\x36\xa5\xf4\xc8\x29\x28\x27\x17\x4d\x8a\xe0\xed\x5c\x5c\xb9\x49\xea\x67\xb4\x29\x71\xc3\xd8\xe7\xe2\x0a\x2a\xe1\xb3\x92\xea\x0b\x42\x6b\x89\x4a\x97\xa5\x18\x44\xf9\x12\xa5\xed\x49\x2e\x95\xde\x5c\x87\xb4\x93\xba\x22\x73\xf0\x34\xb8\x41\x6f\x54\x74\xa4\xe9\x64\x7a\x78\x74\x3c\x78\x55\xe4\xd3\xe9\x7c\x7e\x38\x7b\xd6\xd7\x77\x2f\x89\x32\x84\x48\x89\xac\x8f\x15\xb9\xee\x67\xc0\xe8\xbc\xb0\xde\x8d\xa9\x00\xe0\x3d\x34\x8e\x62\x11\xd1\xf0\x26\x02\x49\x22\x32\x4c\xb3\x83\xb4\x42\xe1\x2c\xd8\x51\xae\x1d\x2d\x7c\xb7\xc8\x92\xda\xa3\x2d\x44\x16\xfb\x13\xa1\x46\x6c\x8b\xa9\x61\x2f\x67\x90\x8d\x52\x11\xb8\x93\x5a\xa8\x3c\x22\x64\x29\x43\xbf\x82\x40\x52\x82\xec\x6d\x5f\x77\x2f\x36\xbf\xf6\x18\x49\x48\x9a\xc4\x40\x2a\x33\x55\x85\xa9\x35\xd8\x05\xd0\x6d\x0c\xc7\x11\x31\x12\x84\xe7\x1e\x01\x71\x93\xd6\x0e\xed\x88\x8c\x2c\x81\x62\xe3\xc3\xd0\xd9\x1b\xc8\x23\x88\xda\x48\xc7\x3b\x3a\x53\xaa\x2f\x0e\xa3\x87\x3b\x8b\xad\x9a\x00\x55\xe3\x9b\x27\xf3\xd1\x29\x44\x29\x2d\x12\x89\x7a\x7d\xfc\x3b\x74\xfa\x33\x42\xc4\x9d\x76\x13\x9f\x3d\x34\x31\xcd\x9c\xcf\x3f\xcb\x30\x33\x4a\x41\x79\x38\x38\x46\xf4\x7b\xb8\xfb\xfc\xa4\xc8\xdb\xce\xdc\x5d\x06\x3f\x3f\xf7\xc7\xf9\xfc\xa7\x34\x91\xab\x39\x5e\x55\x99\x4c\xa8\xd2\x38\x7f\xff\xc4\xae\xb7\xb3\x33\xfb\xd9\x97\xcc\xfe\x71\x3e\x9f\x3d\xb4\xae\x36\xfa\xc0\x79\xa1\x73\x61\xf3\x96\xcc\xb3\xfb\x99\x78\xf6\x59\x39\x7f\x01\x95\xc1\xe4\xbb\x42\xff\x02\x0a\x3d\x0d\x3c\xbb\x5f\x03\x5f\x40\x28\xa9\x63\x10\x8b\xba\x2a\xe1\xbe\x86\x6e\xea\x5e\x71\xf0\xa1\xd2\xce\x12\x2c\x11\x4a\xc5\xb9\x6d\x92\x4a\x72\xeb\x93\x67\xf4\xba\x54\xc6\x54\x50\x48\xe5\xd1\x4a\xbd\x22\xc8\x8e\x08\xaf\x2f\xae\xa6\xb3\xd9\x2c\xcc\xa5\x71\x3c\x2c\x8c\x72\x81\x0a\xe5\x01\x91\xe7\x92\xf8\x10\x0a\xb2\x12\xb3\x9b\xda\x48\xed\xdd\x04\xbe\x37\xb6\x12\x7e\x0e\x7b\x2f\x4b\xa4\x02\xee\xd5\xfc\x65\x29\x5c\xf9\x6a\x2f\x40\xcb\x6e\xec\x62\x67\xc0\x20\xf3\x36\x52\xf9\x03\xa9\x87\xa4\xe1\x0d\xd7\x20\x79\x3c\x2e\xea\x45\x11\xae\x46\x37\x11\x89\xee\x51\xaa\x35\xb4\x21\xde\x42\x8f\x44\xc7\x7d\x88\x65\xde\xa5\xb2\x87\x1b\x8b\x62\x45\xb0\x96\x01\xad\x74\xfd\x82\xc9\x9b\x2e\x2b\xbf\x6b\x9c\xe7\xe8\x27\x75\xa6\x9a\x9c\xa2\x9a\xb0\x22\x23\xe1\xc0\xde\xe1\xde\x18\xf6\xe6\xf4\x9f\xfd\xd8\xf7\x78\xb2\xc7\x1d\x34\x11\x17\x5c\xf4\x77\x49\xcf\xa4\x4f\x99\xb2\x53\x04\xec\x9f\x7f\x1f\x5b\xc4\x59\x4f\xee\x8f\x71\xb6\xf2\xe1\xea\x1c\x1c\xda\x35\xa1\xa6\x98\x06\x0e\x38\x6b\x74\x5d\x9d\xf4\x3c\x33\xda\x5b\xa3\x42\x7b\x25\xe9\xa7\x9b\x1f\xd2\x6b\x56\xb6\xed\xf0\x90\xe8\x78\x0a\x49\x22\x64\x44\xa9\x0b\xb6\x0f\x02\xd4\xa1\x6c\x04\xdb\x04\x0c\xc4\x49\xb5\xb6\x26\x43\xe7\x42\x4d\xde\xe5\xb0\x1e\x9b\xd2\xa5\x52\x9a\x33\x58\x7b\x0a\x58\x80\xad\x33\x56\xe3\xd9\xfb\x37\xf4\x77\x2d\x9c\x1b\x03\x77\xe8\x6d\x9d\x29\x59\x49\xdf\x7f\xcd\x0f\xc2\x18\x02\x48\x83\x82\x60\xf2\x28\x47\x4c\xd7\x98\x35\x36\xb4\x07\x69\x3f\x67\x57\x17\x9c\x82\xfb\xd5\x46\x30\x44\x2d\x2a\x0c\xa7\xa9\xc2\xb9\x8d\xb1\x79\x2c\x91\x32\x3e\x21\x70\xa6\xed\x9d\x51\xea\xe5\x7d\x60\xfe\xbb\x13\xf9\x1c\xb0\x9d\xe2\x41\xa1\xe0\x70\x4b\x80\xa5\x68\x94\xe2\x96\x83\x29\x06\x9d\xf5\x83\x96\x32\x81\x98\xbc\x92\x1a\x0e\x20\x1e\xb7\xf4\xd4\xd1\xd5\xaa\x49\x2b\x93\x20\x70\xee\xf8\x93\x4b\xe2\x1a\xed\xcf\x4c\xe0\xe7\xc4\xe3\xcf\x5b\xd3\xfc\x4c\xa5\x62\x18\x1a\xce\x03\x86\x6a\xea\xa6\x46\x36\xee\x9b\xdc\xea\x71\xf1\x3b\xd8\xa9\xb8\xcb\xf8\xc3\x58\xaa\xeb\x4f\x7f\x15\x30\x45\x5a\x8b\x70\xea\x2b\x80\x29\xaa\xf7\x18\x4e\xfd\x09\x30\x35\x44\xb4\x9e\x6b\xe8\x1d\x95\x86\x3e\x4c\x2b\xa3\x5e\x92\x26\x51\x5e\x5c\xad\x8f\x23\xe0\x5f\x3f\x7b\x18\x9b\x85\x6c\xc7\xba\xfa\xa3\x48\xac\x37\xeb\x4f\xa0\xb1\x6e\xf2\x03\x80\xec\xf8\xce\x78\x7a\xf8\x30\x26\xbb\x33\xaf\x07\x0a\x8e\x1f\x86\x65\x77\xa6\x27\x28\x70\xfc\x30\x32\xbb\x33\x77\x80\x8b\x8e\x1f\x06\x67\x9f\x5b\x7c\xf6\xd0\xea\x9f\x85\x33\xcf\x7f\x97\x95\xe7\x5f\x0e\xd1\xee\x10\x1a\xcc\xff\x42\x94\x76\x87\x48\x4f\x27\xcf\xff\x20\x50\xbb\x43\x2b\x29\xe8\x39\x05\x9b\xef\xa5\x8a\x47\xa7\x52\xa7\xf8\x9d\x11\x66\x28\x64\x26\x3c\x52\x42\xc7\x10\xa4\xe8\x69\x7b\x53\xc2\xd6\xd9\x84\x1e\x7c\x09\x89\x1b\xdc\x06\x0a\x37\xb8\x1d\x10\xa0\x17\x3b\xf1\xae\xba\xd3\xe5\xc9\x8c\xce\x1a\x6b\x09\xd5\x90\x7f\x67\x4a\x32\xe6\xe1\xf6\x78\xda\xea\xce\xa9\xa6\xad\xb3\x4a\xdc\xc6\x91\x8b\xd9\xf4\x0f\x2f\xb2\xc1\xa5\x33\xd9\x0d\xfa\xb4\x5c\x47\xb5\x7d\xe5\x16\x9f\xeb\x2b\xed\x10\xb2\xf8\x4b\x83\xce\xc7\x0e\x63\x3c\xd7\x8f\xf8\x00\xf3\xde\x68\xb5\xed\x31\xde\x3e\xb5\xf8\x8b\x5b\x1c\x31\xff\xef\xa4\xb5\xf1\x44\x00\xfe\xdf\xf5\xe5\xfb\x03\x22\xff\x4b\x23\xed\x8d\xa3\x75\x5f\x4b\x9f\x19\xa9\xe1\xdc\x58\x84\x83\x83\x18\xed\xb9\x5b\xd5\x58\xb1\xa2\xdc\xca\x21\x76\x74\x1a\x4c\x86\x82\xb1\x58\x4a\x25\xfd\x16\xa4\x73\x0d\xba\xf6\xd4\x66\x89\xb0\x31\xf6\x06\x73\x10\xd6\x34\x3a\xe5\xc2\xb0\xd6\xf0\x72\x48\x07\xb0\xe2\xbd\x16\x0e\xd3\x11\x19\xee\xe4\x2a\x5c\xa3\x26\x84\xc3\xa7\x30\x11\xe5\x84\x93\x86\x98\x3d\x87\x67\xb6\xa1\x0d\x99\xea\x83\xd0\x61\xe7\xc6\x27\x37\x26\x64\x76\xc3\x67\x61\x83\x95\x28\x39\xa5\xe8\x1f\x7a\xac\xb1\xa9\xe4\x0d\x9f\xfb\xac\x71\x00\x0e\x18\xba\xb1\xad\x1a\x5d\xc8\x15\x5b\x7a\x00\xac\xb6\xce\xfe\xc0\x3e\x3f\xbe\xbd\xfe\x4c\x6e\x1e\x9c\x61\x77\xe7\x41\x9c\x92\x42\xb3\x29\xca\x62\x08\x0b\x43\xc3\x8f\xaf\x71\xa4\x88\xd5\x73\xf1\xfd\x84\x4e\x63\x6f\x36\xb5\x08\x02\xdb\x5e\x3d\x1a\xc4\x5e\xfd\x39\x8c\x3d\x3a\x85\x6f\xf0\xb6\x46\x2b\xa9\x6e\x10\xea\x9b\x01\xa1\x87\xa1\x36\x5b\xeb\x9f\x02\xdb\xfd\x75\xf8\xe8\xaa\x71\xe8\x92\xed\x51\x70\xe2\x45\x42\x4c\x4a\x9c\x87\x2e\xbc\xd4\xc8\x5d\xdb\x4e\x37\x1c\x58\xa2\x3d\x3e\x0a\xa8\xe6\x52\x54\x77\x8a\x3e\xe4\xd0\xdd\xf5\xca\xf8\xda\x47\x4f\x8c\x61\x77\xbd\xa0\x37\x3a\x85\xfd\x01\x72\xa0\xb8\x7f\x32\x4e\x37\xc0\xe6\x30\xa3\xdf\x6c\x26\xab\x2e\x0f\xd0\x33\x5e\x5e\x83\x68\x7c\x49\x7e\x11\xef\x0c\x7a\x73\x13\x97\xf5\x49\x96\x84\xf4\xc3\xc5\x8f\x34\x10\xfb\xc1\x91\xc8\xd2\x3b\x9e\xb9\x78\x69\xe8\xef\xa3\x03\xfe\xf5\xea\x71\x4c\xf2\x8d\xf0\x62\x29\x1c\xc2\x75\xbc\x52\xf7\x45\x45\x5f\x9e\x66\x71\x40\x2e\x8d\xca\x93\x2f\xf5\xae\x77\x3e\x8e\x96\x5b\x86\x97\x22\xbb\xc1\xe0\xe5\x8d\xc3\x56\xcc\xaf\x99\x81\xf3\xc4\x40\x68\x81\xe7\x96\xaf\xe2\xcc\xa1\x28\x54\xbe\x24\x0b\x5d\xfa\x6d\x8d\x8b\xf0\x93\xa8\xa2\x42\x8f\x50\x4a\xe7\x8d\x95\x99\x50\x61\x23\xfd\xf0\xc6\x14\xe1\x0c\x96\x4d\x51\x84\x2c\x15\x87\xc4\x63\x26\x6e\x8c\x52\xf2\x0e\x86\x9f\x11\x8f\xa6\x20\x47\x43\x63\x57\xe1\xd2\x5c\xa3\x31\x84\x44\x12\x71\x97\xee\x22\x21\x0e\xb0\x7c\xab\x85\x5b\xcb\xc9\x41\xf9\xf2\x55\x43\x64\x2b\x93\xb3\xef\x9c\x0b\x1d\x6f\xab\x30\xde\xe7\x03\x93\xa3\x17\x2f\xda\x35\x72\xac\x7d\xb9\x38\x7e\x1a\x72\xde\x07\xa4\x1a\x3d\xe7\x5d\x7c\xfa\xf8\xaf\xcb\x4e\x7b\xbc\xb9\x36\x75\x82\xd4\x39\xde\x52\x01\x13\xd8\x21\x1c\x2d\x5d\xbc\xd3\xc8\xef\x58\x06\xce\x0b\x8f\x8b\x69\xda\x45\x42\x01\x4e\xfe\x4a\x75\x16\xbc\x93\xaf\xd3\x75\x90\x76\x9d\x4c\x64\x25\x33\x9e\x2f\xf9\x4f\x1a\xbb\x38\x99\x4e\xef\x4a\xc2\x61\x66\x74\xee\x60\x89\x7e\x83\xa8\x7b\xac\xaa\xc6\x95\xe1\x66\x4f\xbe\xe4\x1f\xec\xe7\x6b\xa1\x16\xb3\x17\x44\xe9\x31\x9c\xe3\x7a\xab\xb3\xd2\x1a\x2d\x7f\x8d\x97\x80\xbf\xd4\x47\x4a\xb3\x61\x09\xb4\xb7\x8b\x28\xa9\xb6\xc4\x10\xa4\xa7\xb1\xf5\x36\x49\xea\xd1\xbd\x86\x76\x12\xea\xfb\x5d\xbb\x56\x54\xff\x77\x8d\xb1\xd4\x05\xf3\xb2\x06\x2b\xf8\x8c\x93\xcd\x8b\xc7\xaf\x50\xa3\x93\xac\x84\x42\x38\x4f\x3b\x7a\xac\x54\xf9\x0e\xab\xda\x18\xf5\xa0\xc8\x1f\xed\xca\xfb\xc0\xae\x83\x7e\xf6\xd3\x29\xe4\x93\xd0\x75\xec\xee\x8e\x85\xbb\x10\xf7\xb9\xe6\xd3\xa3\x29\xff\x0b\x97\xbb\x28\xcf\xca\x35\x06\x3d\x90\x23\xa4\xd7\xa1\x31\x14\x2e\x2d\x55\xf1\x5c\xce\x5b\xa1\x9d\x08\xa9\xb3\x40\x4c\xa7\x27\x46\x3b\x99\xf3\x0d\x1e\xc1\xd5\xca\xaf\x68\x0d\xbd\x1f\x87\xa3\x51\x8b\x4a\x6c\xfd\x6d\x81\x48\x15\xcf\x74\x3a\xe5\x98\xf3\x41\x78\x3c\xe0\x1e\x49\xb8\x42\xdf\xa3\xdd\x36\x3f\xd7\x42\x35\x08\xb3\x13\xf8\x06\x66\xd3\xe9\x34\x6c\x37\x76\x40\x2a\xa9\x1b\xcf\x6e\xcc\x44\x88\x06\x2f\xb4\x98\x9d\x84\x30\x43\xd8\x96\x62\xe8\xaa\x84\xda\x4a\x63\x09\x15\x53\x58\xe6\x51\xdc\xbf\xa6\x65\x8d\x05\x65\x36\x07\xc5\x0e\x07\x11\x33\xd2\xd0\x34\x79\x11\x21\x7b\x14\x85\xac\xb0\x8d\x0b\xc2\x7b\xac\xea\xd0\xb4\x75\x94\x05\x34\x6e\x40\xea\x35\xea\xf4\x95\x82\x18\x5c\x09\xab\x91\xb1\xe4\xc7\xe1\x51\x68\x50\x1b\xe6\x73\xd0\x0e\xf6\xb5\xd0\x26\xc6\x9f\x27\x63\x68\x1c\xec\x57\x32\xb3\xdd\x23\x12\x01\x3f\x54\x4a\x76\xe3\x1c\xec\x77\x3f\x2a\x7a\x4d\x52\xa2\x1f\x25\xec\x97\xa6\xb1\x8e\x01\x83\xb7\x04\xb6\xb1\x0d\x5a\x27\xd3\x8a\xcf\x51\xdf\xb2\x3e\x8c\xad\xf9\xd0\xbb\xa7\x6d\xb6\x7e\x6f\x48\x0d\x77\xe4\x54\x89\xdb\x30\xc3\xdf\xa6\xd3\xe0\x37\xa1\x75\x14\x76\x34\x54\x2d\xbb\x6e\xff\xc6\xd6\x24\x7d\x86\xe1\x08\x47\x87\x84\xf4\x81\x04\x3f\xac\x7c\x07\x44\x2c\xae\x84\xcd\xd9\xa6\x4d\xd1\xf6\x57\xf4\xce\x7d\xfc\x90\x2b\x94\xd8\x6a\xa3\x9d\xcf\x13\xe9\xff\x60\xe6\xbf\x12\x6d\x22\xd5\x27\xfe\x40\x0a\xe2\x7c\xd7\xa6\x9f\xc6\xdf\x1a\xfe\x51\x89\x5b\x76\xbd\xe3\x93\x47\xca\x1c\xe1\x9b\x1f\xa1\xe0\x82\x73\xea\xe3\x84\xa9\xd7\x9c\xd6\x09\x81\xa7\xbb\x0c\x20\x42\x90\x2f\x85\x2b\x0f\x28\x6f\x0e\x24\x1d\x12\x7c\x2c\x18\xc2\x4d\x67\xc1\xa9\x69\xa0\x8d\xee\x94\x3a\x5d\x33\x5b\xa1\xb7\x62\xd3\x27\xf4\xe1\xea\x9c\x8d\xfa\x96\x29\x06\x4d\xdc\xcf\x4d\x6c\xa4\x7e\x11\x43\xa1\xdc\x71\x28\x6c\x56\x0e\x17\x75\x5c\x47\xb4\xdc\xc5\xfb\x4d\xb6\xe5\xe0\x51\x00\x80\x5c\x69\xe1\x1b\x8b\xf0\x03\xda\xd0\x88\x21\xb6\xcf\xc9\x86\x1e\x45\xa5\x21\x1c\x84\x2f\x46\xd2\xd2\x6c\xb1\x21\xa4\x55\xe2\x96\x2c\xfb\x84\x62\x3a\xa0\xf6\x56\x06\x48\xe4\xe4\x6a\x60\xd7\x27\x21\x9f\x3c\x86\x48\xce\x8d\xd4\xf0\x77\xd4\x18\x6e\x85\xc0\xfe\x3b\x6e\x57\x3d\xf9\x43\x05\x04\x95\x3a\x1d\x89\x84\x80\x81\x62\xba\x12\x94\x71\xd2\x57\x14\xed\xf5\xee\x2a\x34\xc5\x44\xf7\x49\x16\x5f\x0f\x32\xcd\xaa\x8c\x4d\x8b\x4c\x28\xc5\x77\xb0\x36\x48\x36\xe7\x86\xa7\x53\xe7\x57\x9f\x88\x06\x5a\xd8\x97\x45\xbc\x9e\x9d\x3f\x79\x1c\xb0\x15\xbf\x65\xd8\x5d\x3b\x1c\x43\xf6\xea\xeb\x78\x49\xa1\xfd\x70\x8b\xbb\x1a\xf1\x86\x2b\x97\x33\xe8\xb8\x4b\x52\x37\xb6\x36\x54\x78\xb7\x9f\xdb\x85\x82\x34\x9c\x68\x85\x0f\x88\xc0\x49\x9d\x85\x36\x4b\xfb\xd5\x02\x51\x64\x67\xa4\xf7\xd2\x41\x21\x2c\x78\x63\x42\x86\xa6\x05\x3a\xc6\xda\x93\x84\x8d\xb1\xbe\x0c\x57\xe3\x37\x65\x6c\xe5\x44\x55\xe1\xa2\x10\xca\x61\xef\xd4\x37\x5e\x47\xf5\x06\x6a\xb1\x65\xf1\xe6\xbd\x5a\x66\x67\x05\x3e\x8c\x32\x3e\xf4\x8d\xb8\xda\x88\x96\xb0\xab\xfb\xb4\x5c\xde\x95\xce\xe8\x79\x50\x1a\xc3\xf1\xe7\xce\x47\x04\xed\x85\x97\xb0\x20\xbd\x59\xcc\x68\x27\xcb\xd0\xd5\x8b\x43\x1f\x1c\x70\xf4\xe0\x88\xa7\x77\x5a\xa1\x11\xa3\x84\xad\xa4\x3c\x14\xc0\x93\x37\xe1\x3b\x94\xdd\xc3\x16\xd2\xf6\x6e\xec\x0d\x91\x90\xcf\x6e\x50\xb3\x6d\x17\x88\x8e\x8f\xdd\x82\xd6\xe2\xd3\x16\x5e\xa5\x2b\x8e\xf1\x04\x3a\xe7\x56\x4f\x4f\x82\x3b\xb2\x9d\xc0\xf0\x03\xb2\xcf\xf0\xcd\x14\xa9\xca\x43\x41\xd4\x22\xd0\x0a\x57\xe0\x95\xba\x9f\x34\x34\x75\xfc\x62\xa8\xbf\xa1\x46\x7b\xa9\xe2\x49\x97\xf0\x04\x1d\xf9\xbc\x71\x78\x5f\xba\x57\x08\x93\xc4\x5a\x78\x52\x49\xcd\xd1\xec\xde\xce\xf3\x43\xe2\x66\x17\x0b\x90\x36\x09\x2a\x1e\x21\x86\x4b\xfd\xb1\x6d\xe9\x50\xbb\xc6\x85\x57\x20\x8b\xc8\xae\x12\x76\xd5\xdd\x46\x14\x3e\x10\xea\x98\x8b\xa1\xf6\xf9\x49\x8b\xdd\x7b\x1c\x0e\x79\x8a\x60\x84\x34\x78\x90\x54\x77\x98\x40\xb0\xb0\x28\x42\xc1\xdf\x7d\xae\x94\x16\xa1\x3d\xdc\xb5\x0f\x86\xff\x8e\x8f\x69\x99\x49\x10\x95\x69\xb4\x77\x63\x08\x57\x25\xeb\x86\xfe\x26\x7f\x73\x55\x48\xf0\xc4\x8e\x6b\xaf\xd5\x31\x23\xe8\x7c\x6b\x46\xd1\x97\x58\x14\xbf\x34\x48\x81\x84\x59\x4e\xb7\xa8\x84\xa7\xa8\x13\x19\x0e\xdf\x80\xdd\xad\x20\xc4\x0a\xc7\xe1\x4e\xbf\x15\xd2\x61\x68\x5b\x44\x33\xa5\x3a\x12\x97\x5b\xca\xf1\xf1\x68\xa3\xe2\x2f\xac\xb8\x96\xb8\x41\xb5\x8d\x85\x4d\x32\xe3\x10\x1c\xf8\x7b\xb2\x2c\xa5\x88\xf0\x0d\x57\x2b\x96\x4e\xb5\xd2\xf5\xb4\xba\xdc\xc6\x82\x5a\xb4\x2a\x4a\x4d\x63\xbe\xba\x3e\x34\x03\xda\x96\x46\xcc\xc3\x51\x35\xd3\x4e\x3c\xf7\xf3\xe8\x69\x1b\xac\x85\x8e\xd4\x84\x8a\xab\x53\x5e\xae\xeb\xd8\x88\x0a\x06\x25\x35\xb7\x2c\x58\x11\x41\xb3\x29\xd9\xb1\xd3\xf1\x32\xfc\x61\x69\x70\x9d\x09\x5c\xf8\x14\x16\xd8\x7e\xc3\x57\xc4\x5c\xda\x10\x08\x50\x41\x03\x5d\x9f\x60\x23\x5c\x0c\xb6\xec\x70\x34\x7a\x02\x17\x45\xbc\x39\x9f\x87\x4a\x20\x5c\xbf\x26\xb6\x8b\x46\x67\x81\x65\x52\xc5\x36\xde\x1e\x70\x94\x13\xda\x2b\xfd\xe4\xe3\x5b\x70\xde\xc6\xbb\x32\xd9\xb2\x50\x62\xe5\x16\x81\x95\xc7\xe9\x3c\xe2\xb2\x59\x3d\x4e\x8b\x90\x28\x83\x32\xab\x55\xf8\xde\x72\x8d\xaa\x6b\x05\xf2\xcf\x78\x13\xd6\x5b\x91\xe1\x18\x72\x1a\x3f\xe6\x1e\xf9\x18\x36\xc2\xea\x31\xa0\xb5\xc6\x8e\x21\xb3\x92\x2f\x8c\xff\xb7\x77\xa9\x9f\xfb\xe0\xe9\x88\xfe\xa5\x6b\x96\x6e\xeb\x3c\x56\xaf\x16\x2f\x99\xf4\xab\x71\xf7\xec\xa8\x7b\x38\x99\x4c\x42\xad\xca\x41\xd0\x44\xb6\xe2\x8d\xae\x5c\xae\x65\xde\x08\x05\xed\x4c\x8a\xd4\x9f\x1c\x86\xe6\xfc\xc1\x01\x73\xc8\x33\x16\x8e\x7b\x4b\xa1\xa9\x3d\xfc\xda\xa6\x9b\xcb\x2d\xf8\x76\x06\xed\x2b\x95\x4a\xdc\x12\x4f\x07\x05\xbd\xbe\xf8\x3f\x3e\x7e\xbc\x82\xda\x9a\x42\x2a\x6c\x8f\xdb\xe2\x07\x09\xe9\xf1\xbd\xd7\x3c\xc2\x19\x4c\x77\x3d\x7e\xf7\x52\xfd\x0e\x9d\xfe\x59\x04\x59\x62\xf8\x26\x25\x7d\x97\x2e\x3c\x94\xde\xd7\xf3\xc3\xc3\xf6\xec\x66\xfe\x32\x4e\x25\xee\x5f\x1d\xf2\xd6\x0e\x6b\x7a\x06\x86\x62\x55\x6c\x94\xc6\xef\x92\x69\xe0\xe2\xd9\xf4\x19\x17\x25\xff\xb4\xd2\x23\xa3\x90\xb4\xfa\x9d\xaf\x75\xd2\x81\x55\x56\x37\x69\xf6\xa1\xaf\x6a\xb6\xfc\x09\x3d\x19\xfd\x6f\x00\x00\x00\xff\xff\x99\xe3\x5b\x95\xce\x41\x00\x00")

func sampleBchdConfBytes() ([]byte, error) {
	return bindataRead(
		_sampleBchdConf,
		"sample-bchd.conf",
	)
}

func sampleBchdConf() (*asset, error) {
	bytes, err := sampleBchdConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sample-bchd.conf", size: 16846, mode: os.FileMode(420), modTime: time.Unix(1582740231, 0)}
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
	"sample-bchd.conf": sampleBchdConf,
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
	"sample-bchd.conf": &bintree{sampleBchdConf, map[string]*bintree{}},
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