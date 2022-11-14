package go_logger_cmp

import (
	"github.com/golang/glog"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"io"
	"log"
	"os"
	"testing"
)

func Benchmark_zap(b *testing.B) {
	withBenchedLogger(b, func(log *zap.Logger) {
		log.Error("Error Test")
		log.Info("Info Test")
		log.Warn("Warning Test")
	})
}

func Benchmark_logrus(b *testing.B) {
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(io.Discard)

	for i := 0; i < b.N; i++ {
		logrusLogger.Error("Error Test")
		logrusLogger.Info("Info Test")
		logrusLogger.Warn("Warning Test")
	}
}

func Benchmark_glog(b *testing.B) {
	defer quietGlog()()

	for i := 0; i < b.N; i++ {
		glog.Error("Error Test")
		glog.Info("Info Test")
		glog.Warning("Warning Test")
	}
}

func Benchmark_zerolog(b *testing.B) {
	zeroLogger := zerolog.New(io.Discard)
	for i := 0; i < b.N; i++ {
		zeroLogger.Printf("Test log")
	}
}

func quietGlog() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

func withBenchedLogger(b *testing.B, f func(*zap.Logger)) {
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionConfig().EncoderConfig),
			&zaptest.Discarder{},
			zap.DebugLevel,
		))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f(logger)
		}
	})
}