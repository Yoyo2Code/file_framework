package system

import (
	"os"
	"os/exec"
)

// creates file and return the PID
func CreateFile(args ProgramArgs) int {
	f, err := os.Create(args.FullFilePath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	return os.Getpid()
}

// execute and return the PID
func ExecuteFile(args ProgramArgs) int {
	// execute the command in spawned process
	cmd := exec.Command(args.Data[0],args.Data...)
	err := cmd.Start()

	if err != nil {
		panic(err)
	}

	return cmd.Process.Pid
}

// update file and return the PID
func UpdateFile(args ProgramArgs) int {

	// if overwriting file
	if(args.OverwriteFileFlag) {
		OverWriteFile(args.FullFilePath, args.WriteData)
	}
	
	// if add to file
	if(args.AddToFileFlag) {
		AddToFile(args.FullFilePath, args.WriteData)
	}

	return os.Getpid() // process id not spawned
}

// delete file and return PID
func DeleteFile(args ProgramArgs) int {
	
	err := os.Remove(args.FullFilePath)

	if err != nil {
		panic(err)
	}

	return os.Getpid()
}

// overwrite the file and return PID
func OverWriteFile(file string, data string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        panic(err)
    }

	defer f.Close()
	
	if _, err = f.WriteString(data); err != nil {
		panic(err)
	}
}

// append text to a file
func AddToFile(file string, data string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	
	defer f.Close()
	
	if _, err = f.WriteString(data); err != nil {
		panic(err)
	}
}