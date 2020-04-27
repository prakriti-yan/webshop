package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler
}

func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the situation that this handler is the last middle ware in the chain
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}
	encodings := r.Header.Get("Accept-Encoding") // look for the accept encoding header
	if !strings.Contains(encodings, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}
	// if the accept encoding does not contain gzip that means we can not do gzip
	w.Header().Add("Content-Encoding", "gzip")
	gzipwriter := gzip.NewWriter(w)
	defer gzipwriter.Close()
	var rw http.ResponseWriter
	if pusher, ok := w.(http.Pusher); ok {
		rw = gzipPusherResponseWriter{
			gzipResponseWriter: gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gzipwriter,
			},
			Pusher: pusher,
		}
	} else {
		rw = gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzipwriter,
		}
	}
	gm.Next.ServeHTTP(rw, r)
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

type gzipPusherResponseWriter struct {
	gzipResponseWriter
	http.Pusher
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}

// this writing function inplementation on gzipResponseWriter will make it belong to ResponseWriter
// interface, and by using write in gzipwriter, we successfully gzippped our file!
