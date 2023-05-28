package tl

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogHandler(lvl zapcore.Level, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := With(req.Context(), ID())
		req = req.WithContext(ctx)

		if h, ok := w.(http.Hijacker); ok {
			w = &logResponseHijacker{
				logResponseWriter: logResponseWriter{
					ResponseWriter: w,
					Request:        req,
					Level:          lvl,
				},
				Hijacker: h,
			}
		} else {
			w = &logResponseWriter{
				ResponseWriter: w,
				Request:        req,
				Level:          lvl,
			}
		}

		next.ServeHTTP(w, req)
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	Request *http.Request
	Level   zapcore.Level
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	Log(w.Request.Context(), w.Level, "http request",
		zap.String("method", w.Request.Method),
		zap.Stringer("url", w.Request.URL),
		zap.Int("code", statusCode),
		zap.String("status", http.StatusText(statusCode)))

	w.ResponseWriter.WriteHeader(statusCode)
}

type logResponseHijacker struct {
	logResponseWriter
	http.Hijacker
}
