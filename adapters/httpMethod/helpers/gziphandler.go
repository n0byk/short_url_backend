package helpers

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var gzPool = sync.Pool{
	New: func() interface{} {
		w := gzip.NewWriter(ioutil.Discard)
		return w
	},
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// func (w *gzipResponseWriter) WriteHeader(status int) {
// 	w.Header().Del("Content-Length")
// 	w.ResponseWriter.WriteHeader(status)
// }

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return

		}

		gz := gzPool.Get().(*gzip.Writer)
		defer gzPool.Put(gz)
		gz.Reset(w)
		defer gz.Close()
		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(&gzipResponseWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

func ReadBodyBytes(r *http.Request) ([]byte, error) {

	var reader io.Reader

	if r.Header.Get("Content-Encoding") == "gzip" || r.Header.Get("Content-Encoding") == "application/gzip" {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
		reader = gz
		defer gz.Close()
	} else {
		reader = r.Body
	}

	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil

}

func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("failed init compress writer: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed write data to compress temporary buffer: %v", err)
	}
	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("failed compress data: %v", err)
	}
	return b.Bytes(), nil
}
