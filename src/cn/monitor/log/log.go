package log

import (
"fmt"
)

func  Logger(format string, v ...interface{}){
	fmt.Println(format);
}

func  Error(format string, v ...interface{}){
	  fmt.Println(format);
}


