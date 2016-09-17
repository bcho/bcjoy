package files

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"join.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\x94VM\x8f\xeb4\x14\xddϯ\xb0\xc2\xe6!\xb5\xf9\x987e:\x9d\xb4\x02Ă\x05\xbc\x05\x885r\x92\x9b\xc4o\x1c;r\x9cv\xcah\x96H\x80\x04\x1b\x04\x12\x12\v6H,Y\x83\x9e\xf85e\xfe\x067\xdfn\x93jx\xaa\xd4$\u05fe\xe7\x1e\x9f{\xec\xc4Ou\xc67\x17\x84\xf8)Ш\xba\xc1[\xcd4\x87\xcd\xe1\xbb\xdf\x0e_\xffN\x1e\x1e\x88\xad\x81f\xf6+\x9a\x01y|\xf4\x9df\xb8\x99\x9a\x81\xa6D\xe0\xc8\xda\xda2\xd8\xe5Ri\x8b\x84Rh\x10zm\xedX\xa4\xd3u\x04[\x16¼~\x981\xc14\xa3|^\x84\x94\xc3ڳ\xddY\x86\xa1\xačHY\x80\xaa\x1fi\x80\x11!\xad\x9a\xa0\xd31\xf4#\xb6%!\xa7E\xb1\xb6\x8a\x1c\xaf\xa9ղ1\x06\xb8Ld\xd1\xc6\xc7#$\x00\xaa\xf6aJ\xb5\xb5\xf1\x1d\x1cl\x01\x8c\xdb|\xe3\a\x9b\xf1\xe2\x03\x9c\x9fwSz\x16\x9a\xear\xa8vx\xf3\xe3\xe1\xdb\xef\x0f\xbf\xfe\xf1\xf4\xd7?\xc4\x0f\xbaY4\xd4l\vV\r*\x05g\x02>\x85,\x00U\xb4\xc0\xc41&k\xa9)o\xe6ַ\xc7S[\xb6\x1d\x11:\xac\x8dFL$\x16I\x15\xc4k\xab\xca~-\x99\xf8\xe2\xb3O0\xd1\xda<\xfd\xf0'\xb2j\x1a\xeb;t\xb4\f\x96\b&\xacͿ\xdf\xfc\\a6\x18\xa9\xd6y\xb1r\x9c^\x8a\xcf\xcb \x92\x19e\x02!\xed^G;\x94\x99E4U\t`\xe3\xbf\xd42\xc7r\xbf\xfc}x\xf3SU\xc8\x1e\xb8\xc6RjP\x9dT\xb9܁\x82\x88\x04\xfbqń\xe9\xb4\f*`'\bS\x89\u007f\xaf\xe5ި\x11p*\xee\xacM\x1d\x1eV\xe3\x98\x05\xfcB\xef;\xaf\x12b7n!\x0f\xed3!\xb5'W\xe4\xa5\xeb\xe6\xf7\xb7}Tý\x9eS\x8ej\xacH\x88F\x065\x8c\xc5\xe8\xedyL3\xc6\xf7+b}\f|\v\x9a\x85\x94\xbc\x82\x12\xac\x19\xe9\x033\xf2\x81B\x9f\x0f\x89\x19\xd2f\bxY\xd5\"\xb4Բ\x1b{\xbc\xe8\xf8U\x0e7\xd85)\xf3@j-\xb3\x15\xb92H\x0e9\xb5э\xa4\\\x16\xb8\xc3$VR\xc0i\xe5\xb8\xdbs\x80\x97g\x01\xc7\x12]-M\x85R`I\xaaO\xa3\x11\xab\x04F]X\xed\xeey\xc0ex7\f\a4\xbcK\x94,E4/\xd8W\x80\xda\xca\xed \xed@\xa0\xf7\x94\xc1\xc2\xc8e\x19M0\xb9T\xfc\x85\xe3\x1c\xf9Ϲ\xf4\"w\xb1\\z\xae\x17_]߸W.u\x17@\xaf\xe3\xd8\r\xa3\xe0&^عH\xde\x1d\x15̍:u{\x1bv\xde\xc2\\[\xd7>\f\x12w\x04A\xedv獻\x87;\xe1Xi,H\xa3jn\rv\xe2\xba\bB\xa9h\xd3?!\x85ѻPr\xa9V\xe4\x9d8\x8eOܸk\x9bq\xed\xbaSbc\xcer\x11z\x8b\xe5I)\xad\xa8(b\xa9\xd0\be\x9e\x83\nia\x96+UQ\xd5\xcb\xf1\xf88\xda\x00\xb2\xd4UsW\xc4(VC\xb5\xae3:USF\x1dݬ \x80\xe8s&fd\"h\xb0\x96*³\xbf\xf5\x9c;a\xad#O\xf5\xfa\a\x86\xee&\x99~\x8d\xd3\x05\x87|\x1b)\xefNA\x1aq\xea\xf7\xd2\v\xcf~9v\x8e\xdd\x1c\xe9F^ץc\xc5\a\xab\xd9\xcd\x19{\xc6qޤK<w\xd2s=\x16}\xb6\xfe3暀\\\xa5\xd5ޜހ\xe7jL8\xb4Gn\x0e\xe5\t\xa2\x1f\xbdW\xfdn\x9f\x97\xe3\xed\x8f\xea\xf1\x89;\xa1bKlJÛ\x0f\xab\xdf\xffޠ\xadw\xbb\xc3\x15ɓBr\x16\x9d\x02\x8dJ\x8f\xb4\x9e\xdc\xea\x13\r8\x83\xfb~\x06\x11\xa3\xb8\xfc\xfbf#\x91E%\xc0\x00?~\r\x1e\x1fVޱď\x17\xe6\xd5w\xfaWj\xfb\xb5\x84\x1ff\xf57\xe4\u007f\x01\x00\x00\xff\xff\xb3SWQK\n\x00\x00",
		hash:  "1a2c734602f4e315d8cd64df78faba0a81f7e4a9c566521279d9306a83bbdc89",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1469027702, 0),
		size:  2635,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	f, ok := staticFiles[strings.TrimPrefix(req.URL.Path, "/")]
	if !ok {
		NotFound(rw, req)
		return
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	}
	return gzip.NewReader(strings.NewReader(f.data))
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}
