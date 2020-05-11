package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func TestDemo(t *testing.T) {
	f := fibonacci()
	results := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i := 0; i < 10; i++ {
		result := f()
		logrus.WithFields(logrus.Fields{
			"step":   i,
			"result": result,
		}).Info()
		if result != results[i] {
			t.Logf("expected: %d -> got: %d\n", results[i], result)
			t.Fail()
		}
	}
}
