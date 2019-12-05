package monitor

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	//logger "cn/monitor/log"
	"strconv"
)

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
	//readStat(process)
	readStatus(process)
}

func readStatus(process *Process) (err error) {
	statPath := fmt.Sprintf("/proc/%s/%s", process.Pid, "status")
	f, err := os.Open(statPath)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, _, err := bfRd.ReadLine()
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			return err
		}
		data := strings.Split(string(line), ":")
		//		logger.Logger(string(line) +" "+data[1])
		switch data[0] {
		case "PPid":
			process.Ppid = strings.Trim(data[1], "\t")
			break
		case "State":
			{
				stata := strings.Replace(strings.Replace(data[1], "\t", "|", -1), " ", "|", -1)
				process.State = strings.Split(stata, "|")[1]
				break
			}
		case "Uid":
			{
				uid, _ := strconv.Atoi(strings.Split(strings.Replace(data[1], "\t", "|", -1), "|")[1])
				process.Uid = uid
				break
			}
		case "Gid":
			{
				gid, _ := strconv.Atoi(strings.Split(strings.Replace(data[1], "\t", "|", -1), "|")[1])
				process.Gid = gid
				break
			}
		}
	}

	return nil
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
