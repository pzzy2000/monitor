package main

import (
	scan "cn/monitor"
//	scanet "cn/monitor/net"
	"fmt"
)

func main() {
//		scan.Scan();
	s:= scan.ProcessToJsonString()
	
//
	fmt.Println(s); 
}
