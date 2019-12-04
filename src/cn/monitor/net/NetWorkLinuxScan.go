package net

import (
	"bufio"
	toolsUnits "cn/monitor"
	//	logger "cn/monitor/log"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

func Sacn(sm *sync.Map) {
	Scan()

}

func Scan() {
	socketPid, err := scanSocket()
	if err != nil {
	} else {
		// key : pid
		var pidSocket map[int]NetWorks = make(map[int]NetWorks)
		scanTcp(pidSocket, socketPid)
		scanTcp6(pidSocket, socketPid)
		mjson, _ := json.Marshal(pidSocket)
		mString := string(mjson)
		fmt.Printf("print mString:%s", mString)
	}

}

func scanSocket() (map[string]string, error) {
	socketPid := make(map[string]string)
	files, err := ioutil.ReadDir("/proc")
	if err == nil {
		for _, f := range files {
			if f.IsDir() && toolsUnits.IsDigit(f.Name()) {
				statPath := fmt.Sprintf("/proc/%s/%s", f.Name(), "fd")
				fdChilds, err_ := ioutil.ReadDir(statPath)
				if err_ == nil {
					for _, fdChild := range fdChilds {
						path := fmt.Sprintf("/proc/%s/%s/%s", f.Name(), "fd", fdChild.Name())
						linkc, err := os.Readlink(path)
						if err == nil && strings.HasPrefix(linkc, "socket:") {
							linkc = linkc[8:strings.LastIndex(linkc, "]")]
							socketPid[linkc] = f.Name()
						}
					}
				}
			}

		}
		return socketPid, nil
	} else {
		return socketPid, err
	}

}

func scanTcp6(pidSocket map[int]NetWorks, socketPid map[string]string) (err error) {
	f, err := os.Open("/proc/net/tcp6")
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	sum := 0
	for {
		lineByte, _, err := bfRd.ReadLine()
		if sum == 0 {
			sum++
			continue
		}
		//		logger.Logger(string(lineByte))
		line := string(lineByte)
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			return err
		}
		lines := strings.Fields(line)
		//		fmt.Println(" %q ", lines)
		st, _ := strconv.ParseInt(lines[3], 16, 8)
		switch st {
		case 10: //var tcp_listen = 10
			{
				ints1 := lines[1][0:4]
				ints2 := lines[1][4:8]
				ints3 := lines[1][8:12]
				ints4 := lines[1][12:16]
				ints5 := lines[1][16:20]
				ints6 := lines[1][20:24]
				ints7 := lines[1][24:28]
				ints8 := lines[1][28:32]
				port, _ := strconv.ParseInt(lines[1][33:37], 16, 16)
				localAddress := fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s:%d", ints1, ints2, ints3, ints4, ints5, ints6, ints7, ints8, port)
				//				logger.Logger(localAddress + "    >   " + lines[9])
				pidString := socketPid[lines[9]]

				if pidString == "" {
					pidString = socket_pid_defult_nil
				}
				pid, _ := strconv.Atoi(pidString)
				netWorks, ok := pidSocket[pid]
				if ok {
					netWorks.NetWork = append(netWorks.NetWork, NetWork{Proto: 6, LocalAddress: localAddress})
					pidSocket[pid] = netWorks
				} else {
					netWorks = NetWorks{Pid: pid}
					netWorks.NetWork = append(netWorks.NetWork, NetWork{Proto: 6, LocalAddress: localAddress})
					pidSocket[pid] = netWorks
				}
			}
			break
		}
	}
	return nil
}

func scanTcp(pidSocket map[int]NetWorks, socketPid map[string]string) (err error) {
	f, err := os.Open("/proc/net/tcp")
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	sum := 0
	for {
		lineByte, _, err := bfRd.ReadLine()
		if sum == 0 {
			sum++
			continue
		}
		//		logger.Logger(string(lineByte))
		line := string(lineByte)
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			return err
		}
		lines := strings.Fields(line)
		//		fmt.Println(" %q ", lines)
		st, _ := strconv.ParseInt(lines[3], 16, 8)

		switch st {
		case 10: //var tcp_listen = 10
			{
				int1, _ := strconv.ParseInt(lines[1][0:2], 16, 8)
				int2, _ := strconv.ParseInt(lines[1][2:4], 16, 8)
				int3, _ := strconv.ParseInt(lines[1][4:6], 16, 8)
				int4, _ := strconv.ParseInt(lines[1][6:8], 16, 8)
				port, _ := strconv.ParseInt(lines[1][9:13], 16, 16)
				localAddress := fmt.Sprintf("%d.%d.%d.%d:%d", int4, int3, int2, int1, port)
				//				logger.Logger(localAddress + "    >   " + lines[9])
				pidString := socketPid[lines[9]]
				if pidString == "" {
					pidString = socket_pid_defult_nil
				}
				pid, _ := strconv.Atoi(pidString)
				netWorks, ok := pidSocket[pid]
				if ok {
					netWorks.NetWork = append(netWorks.NetWork, NetWork{Proto: 4, LocalAddress: localAddress})
					pidSocket[pid] = netWorks
				} else {
					netWorks := NetWorks{Pid: pid}
					netWorks.NetWork = append(netWorks.NetWork, NetWork{Proto: 4, LocalAddress: localAddress})
					pidSocket[pid] = netWorks
				}
			}
			break
		}
	}
	return nil
}
