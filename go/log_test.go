package log

import (
	// "flag"
	"bytes"
	"errors"
	"fmt"
	"log"
	"testing"

	"k8s.io/klog/v2"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func TestLogging(t *testing.T) {
	klog.InitFlags(nil)
	klog.Info("Prepare to repel boarders")

	klog.InfoS("i earn respect", "point", 100)
}

func TestZapr(t *testing.T) {
	var log logr.Logger

	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}
	log = zapr.NewLogger(zapLog)

	log.Info("Logr in action!", "the answer", 42)
}

func TestSet(t *testing.T) {
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}
	klog.SetLogger(zapr.NewLogger(zapLog))
	klog.InfoS("i earn respect", "point", 100)
}

func TestWith(t *testing.T) {
	var mainLogger logr.Logger

	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}
	mainLogger = zapr.NewLogger(zapLog)

	log := mainLogger.WithName("reconcilers").WithValues("target-type", "foo")
	log.Info("setting foo on object", "value", 10, "object-name", "little-rain")

	err = errors.New("something is wrong")
	log.Error(err, "unable to reconcile object", "object-name", "little rain")
}

func TestLogOutput(t *testing.T) {
	log.SetFlags(log.Llongfile)
	One()
}

func One() {
	Two()
}

func Two() {
	Three()
}

func Three() {
	log.Output(0, "message")
	log.Output(1, "message")
	log.Output(2, "message")
	log.Output(3, "message")
}

func TestLoggerOutput(t *testing.T) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)

		infof = func(info string) {
			logger.Output(1, info)
		}
	)

	infof("one")

	logger.Output(1, "two")

	fmt.Print(&buf)
}
