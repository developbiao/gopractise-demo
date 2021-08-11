package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	glog.MaxSize = 1024 * 1024 * 1024 // 1GiB split log file
	flag.Parse()
	defer glog.Flush()

	glog.Info("This is info messages")
	glog.Infof("This is info messages: %v", 23232)

	glog.Warning("This is warning message")
	glog.Warningf("This is warning message: %v", 382312)

	glog.Error("This is error message")
	glog.Errorf("This is error message: %v", 382312)

	//glog.Fatal("This is fatal message")
	//glog.Fatalf("This is fatal message", 7832)

	glog.V(3).Info("LEVEL 3 message")
	glog.V(4).Info("LEVEL 4 message")
	glog.V(5).Info("LEVEL 5 message")
	glog.V(8).Info("LEVEL 8 message")

}
