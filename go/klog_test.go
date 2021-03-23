package log

import (
	// "flag"
	"fmt"
	"k8s.io/klog/v2"
	"testing"

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
