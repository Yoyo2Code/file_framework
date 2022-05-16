# Key Features

- Start a process, given a path to an executable file and the desired (optional) command-line arguments
- Create a file of a specified type at a specified location
- Modify a file
- Delete a file
- Establish a network connection and transmit data

### Log file (csv headers in progress...)

#### Process start
- Timestamp of start time
- Username that started the process
- Process name
- Process command line
- Process ID
#### File creation, modification, deletion
- Timestamp of activity
- Full path to the file
- Activity descriptor - e.g. create, modified, delete
- Username that started the process that created/modified/deleted the file
- Process name that created/modified/deleted the file
- Process command line
- Process ID
#### Network connection and data transmission
- Timestamp of activity
- Username that started the process that initiated the network activity
- Destination address and port
- Source address and port
- Amount of data sent
- Protocol of data sent
- Process name
- Process command line
- Process ID

# Background

This program functions as a framework to modify files, run executables, or send UDP data. When the commands are run it will log the actions in the order shown above. The process timestamp is grabbed from the cli since there are millisecond differences between code execution and process creation. The process ID is taken from a sub process when running an executable, but grabs the parent process ID during network sending or file modification since it does not spawn a new process.

The application is made up of the `main` and `system` packages with multiple imports. The application **must** be run using `go build` and then running the executable to run as expected. This is due to the location of the build when running `go run`. The flow of the application is:

`userinput.go` (grab the arguments/inputs) -> `main.go` (handle the arguments read in) -> `files.go` or `network.go` (perform logic using arguments) -> `logger.go` (log the process information)

# Requirements

This program was tested with `go1.17.2` on `darwin/amd64` for MacOS. The `ps` commands are available on Linux but not tested. Will not work on Windows 10 nor has been tested.


# Examples

Options shown by running `go build` then the executable with `-h` flag.

Before running examples the project must be built...

## Run executable (sh file creates a new file to verify it works)

`./file_framework -action=execute sh`

## Create file

`./file_framework -action=create ./new.txt`

## Update file

Add to the file

`./file_framework -action=update -a ./new.txt hello person`

Overwrite entire contents of file to new text

`./file_framework -action=update -o ./new.txt goodbye person`

## Delete File

`./file_framework -action=delete ./new.txt`

## Send via UDP

`./file_framework -action=send -address=58:8080 hello ip 58`