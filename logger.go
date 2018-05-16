package potion

import (
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/golang/glog"
)

type loggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
	startTime  time.Time
	length     int
}

func (lrw *loggerResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggerResponseWriter) Write(b []byte) (n int, err error) {
	n, err = lrw.ResponseWriter.Write(b)
	lrw.length += n
	return
}

func buildLogLine(method string, uri string, startTime time.Time, code int, length int) []byte {
	buf := make([]byte, 0, len(method)+len(uri)+50)
	buf = append(buf, method...)
	buf = append(buf, " - "...)
	buf = append(buf, uri...)
	buf = append(buf, " "...)
	buf = append(buf, strconv.Itoa(code)...)
	buf = append(buf, " [dur:"...)
	buf = append(buf, strconv.FormatInt(time.Since(startTime).Nanoseconds(), 10)...)
	buf = append(buf, "] [len:"...)
	buf = append(buf, strconv.Itoa(length)...)
	buf = append(buf, "]"...)
	return buf
}

func loggerHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggerResponseWriter{w, http.StatusOK, time.Now(), 0}

		defer loggerPanicHandler(lrw, r)

		next.ServeHTTP(lrw, r)

		logLine := buildLogLine(r.Method, r.RequestURI, lrw.startTime, lrw.statusCode, lrw.length)
		if lrw.statusCode < http.StatusBadRequest {
			glog.Infof("%s", logLine)
		} else {
			glog.Errorf("%s", logLine)
		}
	}
}

func loggerPanicHandler(lrw *loggerResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		glog.Errorf("%s", err)
		glog.Errorf("%s", debug.Stack())
		Error(lrw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		logLine := buildLogLine(r.Method, r.RequestURI, lrw.startTime, lrw.statusCode, lrw.length)
		glog.Errorf("%s", logLine)
	}
}
