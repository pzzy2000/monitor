package main

import (
	scan "cn/monitor"
	"fmt"
)

func main() {
//		scan.Scan();
	 data := scan.ProcessToJsonString();
//
	fmt.Println(data); 
}
