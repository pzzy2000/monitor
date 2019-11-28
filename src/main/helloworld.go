package main

import (
	scan "cn/monitor"
	"fmt"
)

func main() {
	//	scan.Scan();
	group,_,groupErr,userErr:=scan.ScanGroupAndUser();
	if(groupErr!=nil ||userErr !=nil){
		scan.ListJson(group);
	}else{
		scan.ListJson(group);
//		scan.ListJson(user);
		s := scan.JsonTostring(group);
		fmt.Printf(s);
//		scan.JsonTostring(user);
	}
	fmt.Println("Hello world!");
}
