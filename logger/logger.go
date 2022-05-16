package logger

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/Yoyo2Code/file_framework/system"
)

var (
	logfile        = "./framework.csv"
)

// logs when a process is started
func LogProcessStart(pa system.ProgramArgs, processID int) {
	// get the process name and start time
	cmd1 := exec.Command("ps", "-p", fmt.Sprint(processID), "-o", "comm=,lstart=")
	nameAndStart, _ := cmd1.Output()

	// grab the username
	pUser, _ := user.Current()
	uName := pUser.Username

	// separate the process name and time stamp
	nameAndStartFields := strings.Fields(string(nameAndStart))
	processName, timeStamp := nameAndStartFields[0], strings.Join(nameAndStartFields[1:], " ")

	// change to string
	logOutput := fmt.Sprintf(
		"%s,%s,%s,%s,%d",
		timeStamp,
		uName,
		processName,
		strings.Join(pa.Data, " "),
		processID,
	)

	writeToLogFile(logOutput)
}

// logs when a file is modified (created, updated, deleted)
func LogFileModification(pa system.ProgramArgs, processID int) {
	// get the process name and start time
	cmd1 := exec.Command("ps", "-p", fmt.Sprint(processID), "-o", "comm=,lstart=")
	cmd2 := exec.Command("ps", "-p", fmt.Sprint(processID), "-o", "command=")

	// grab the output
	nameAndStart, _ := cmd1.Output()
	processCommandLine, _ := cmd2.Output()

	// get the current user
	pUser, _ := user.Current()
	uName := pUser.Username
	
	// remove new line
	processCommandLine = processCommandLine[:len(processCommandLine)-1]
	nameAndStartFields := strings.Fields(string(nameAndStart))

	// separate the process name and timestamp
	processName, timeStamp := nameAndStartFields[0], strings.Join(nameAndStartFields[1:], " ")

	// format string
	logOutput := fmt.Sprintf(
		"%s,%s,%s,%s,%s,%s,%d",
		timeStamp,
		pa.FullFilePath,
		pa.Action,
		uName,
		processName,
		processCommandLine,
		processID,
	)

	writeToLogFile(logOutput)
}

// logs when data is transmitted across TCP
func LogNetworkTransmit(pa system.ProgramArgs, networkData system.NetworkData, processID int) {
	// get the process name and start time
	cmd1 := exec.Command("ps", "-p", fmt.Sprint(processID), "-o", "comm=,lstart=")
	nameAndStart, _ := cmd1.Output()

	// get the current user
	pUser, _ := user.Current()
	uName := pUser.Username

	// separate the process name and time stamp
	nameAndStartFields := strings.Fields(string(nameAndStart))
	processName, timeStamp := nameAndStartFields[0], strings.Join(nameAndStartFields[1:], " ")

	// format the string
	logOutput := fmt.Sprintf(
		"%s,%s,%s,%s,%s,%s,%s,%s,%d",
		timeStamp,
		uName,
		fmt.Sprintf("%s:%s", networkData.DestinationIP, networkData.DestinationPort),
		fmt.Sprintf("%s:%s", networkData.SourceIP, networkData.SourcePort),
		networkData.BytesSent,
		networkData.Protocol,
		processName,
		strings.Join(pa.Data, " "),
		processID,
	)

	writeToLogFile(logOutput)}

// write to the logfile defined
// NOTE: can be modified to print csv, yaml, etc.
func writeToLogFile(entry string) {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	
	defer f.Close()
	
	if _, err = f.WriteString(entry + "\n"); err != nil {
		panic(err)
	}
}
