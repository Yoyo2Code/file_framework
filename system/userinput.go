package system

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

// holds the arguments
type ProgramArgs struct {
	FullFilePath string
	Data     []string // stores the data to transmit if TCP
	Action	 string
	CommandLine []string
	OverwriteFileFlag bool
	AddToFileFlag bool
	WriteData string 
	ConnectionAddress string
}

// return arguments from cli
func GetArguments() (pa ProgramArgs) {
	argsWithProgram := os.Args  // get all args for program
	actionPtr := flag.String("action", "", "create, execute, update, delete, send")
	overwritePtr := flag.Bool("o", false, "flag to overwrite")
	addPtr := flag.Bool("a", false, "flag to add to file")
	connectionPtr := flag.String("address", "", "address:port to send udp")

	flag.Parse()

	allFlags := flag.Args()
	fileName := ""

	// updating file requires flags
	if (*actionPtr == "update" && (*overwritePtr == *addPtr)) {
		panic("update action must have -a OR -o flag")
	}

	// sending data via udp requires address:port
	if(*actionPtr == "send" && *connectionPtr == "") {
		panic("sending requires an address:port")
	}

	// if not sending, set the full path
	if(*actionPtr == "create" || *actionPtr == "update" || *actionPtr == "delete") {
		fileName = AbsolutePath(allFlags[0])
	}

	// sending then requires flag
	if(len(allFlags) == 0) {
		panic("requires data to send")
	}
	
	userData := ""

	// if not sending message, then grab just the data
	if(*actionPtr != "send") {
		userData = strings.Join(allFlags[1:], " ")
	} else {
		userData = strings.Join(allFlags, " ")
	}

	return ProgramArgs{
		FullFilePath: fileName, // 1 exe, 2 action, file
		Data: flag.Args(),
		Action: *actionPtr,
		CommandLine: argsWithProgram,
		OverwriteFileFlag: *overwritePtr,
		AddToFileFlag: *addPtr,
		WriteData: userData,
		ConnectionAddress: *connectionPtr,
	}
}

// calculate the absolute path
func AbsolutePath(path string) string {
	f, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}
	return f
}