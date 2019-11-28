package monitor

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var error_message_read_file_error = "-1000"

func readfile(file string, process *Process) (data string, err error) {
	statPath := fmt.Sprintf("/proc/%s/%s", process.Pid, file)
	dataBytes, err := ioutil.ReadFile(statPath)
	if err != nil {
		return
	} else {
		data := string(dataBytes)
		return data, err
	}

}

func readAll(process *Process) {
	readComm(process)
	readCmdline(process)
	readStat(process)
}

func readCmdline(process *Process) (err error) {
	cmdline, err := readfile("cmdline", process)
	if err == nil {
		process.Cmdline = TrimUnicode(cmdline)
	} else {
		process.Comm = error_message_read_file_error
	}
	return
}

func readStat(process *Process) (err error) {
	stat, err := readfile("stat", process)
	if err == nil {
		binStart := strings.IndexRune(stat, '(') + 1
		binEnd := strings.IndexRune(stat[binStart:], ')')

		stat = stat[binStart+binEnd+2:]
		//		fmt.Printf(stat)
		_, err = fmt.Sscanf(stat,
			"%s %s %s %s",
			&process.State,
			&process.Ppid,
			&process.Pgid, &process.Sid)
		if err != nil {
			process.Pid = error_message_read_file_error
			process.Pgid = error_message_read_file_error
			process.State = error_message_read_file_error
			process.Sid = error_message_read_file_error
		} else {

		}
	} else {
		process.Pid = error_message_read_file_error
			process.Pgid = error_message_read_file_error
			process.State = error_message_read_file_error
			process.Sid = error_message_read_file_error
	}
	return
}

func readComm(process *Process) (err error) {

	comm, err := readfile("comm", process)
	if err == nil {
		process.Comm = Trim(comm)
	} else {
		process.Comm = error_message_read_file_error
	}
	return
}
