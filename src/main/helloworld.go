package main

import (
	scan "cn/monitor"
	scanet "cn/monitor/net"
	"fmt"
)

func main() {
//		scan.Scan();
	 data := scan.Scan();
	 scanet.Sacn(&data)
//
	fmt.Println(data); 
}
