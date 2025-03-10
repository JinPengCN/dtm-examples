package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/examples"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmcli/logger"
)

func hintExit(msg string) {
	if msg != "" {
		fmt.Print(msg, "\n")
	}
	fmt.Printf("Usage: %s <command>\n\nCommand can be one of the following:\n\n", filepath.Base(os.Args[0]))
	fmt.Printf("%4s%-28srun a quick start example\n", "", "qs")
	for _, cmd := range examples.Commands {
		fmt.Printf("%4s%-28srun an example includes %s\n", "", cmd.Arg, strings.ReplaceAll(cmd.Arg, "_", " "))
	}
	os.Exit(0)
}
func main() {
	if len(os.Args) == 1 {
		hintExit("")
	}
	logger.InitLog("debug")
	busi.BusiConf = dtmimp.DBConf{
		Driver: "mysql",
		Host:   "localhost",
		Port:   3306,
		User:   "root",
	}
	busi.Startup()
	time.Sleep(2000 * time.Millisecond)
	cmd := os.Args[1]
	if cmd == "qs" {
		examples.QsStartSvr()
		examples.QsFireRequest()
	} else if examples.IsExists(cmd) {
		examples.Call(cmd)
	} else {
		hintExit("unknown command: " + cmd)
	}
	time.Sleep(3 * time.Second)
	select {}
}
