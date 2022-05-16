package main

import (
	"fmt"
	"time"

	"github.com/Yoyo2Code/file_framework/logger"
	"github.com/Yoyo2Code/file_framework/system"
)

var (
	startTime = time.Now()
)

func main() {
	programArguments := system.GetArguments()
	handleFlags(programArguments)
}

func handleFlags(args system.ProgramArgs) {
	switch args.Action {
	case "create":
		pid := system.CreateFile(args)
		logger.LogFileModification(args, pid)
	case "execute":
		pid := system.ExecuteFile(args)
		logger.LogProcessStart(args, pid)
	case "update":
		pid := system.UpdateFile(args)
		logger.LogFileModification(args, pid)
	case "delete":
		pid := system.DeleteFile(args)
		logger.LogFileModification(args, pid)
	case "send":
		networkData, pid := system.SendMessage(args.ConnectionAddress, args.WriteData)
		logger.LogNetworkTransmit(args, networkData, pid)
	default:
		fmt.Println("no arg")
	}
}