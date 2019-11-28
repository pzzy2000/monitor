package main

import (
	scan "cn/monitor"
	"fmt"
)

func main() {
	//	scan.Scan();
	group,user,groupErr,userErr:=scan.ScanGroupAndUser();
	if(groupErr!=nil ||userErr !=nil){
		scan.ListJson(group);
	}else{
		scan.ListJson(group);
		scan.ListJson(user);
	}
	fmt.Println("Hello world!");
}
