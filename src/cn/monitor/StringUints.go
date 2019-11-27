package monitor

import (
	"regexp"
	 "strings"  
)

func IsDigit(data string) bool{

	pattern := "\\d+"; //反斜杠要转义

    result,_ := regexp.MatchString(pattern,data);
    return result;
}

func  Trim(src string) string{
	 return  strings.Replace(src, "\n", "", -1)  
}

func  TrimUnicode(src string) string{
	 return  strings.Replace(src, "\u0000", " ", -1)  
}