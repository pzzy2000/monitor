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

//func LogJson(sm sync.Map) {
//	sm.Range(func(k, v interface{}) bool {
//		
//		 jsonBytes, err := json.Marshal(v)
//        if err != nil {
//               logger.Error(err.Error())
//        }else{
//        
//        	logger.Logger(" Process -> "+string(jsonBytes));
//        }
//		return true
//	})
//
//}
//
//func listProcess(sm *sync.Map) {
//	
//	sm.Range(func(k, v interface{}) bool {
//	
//
//		 jsonBytes, err := json.Marshal(v)
//        if err != nil {
//               logger.Error(err.Error())
//        }else{
//        
//        	logger.Logger(" Process -> "+string(jsonBytes));
//        }
//
//		return true
//	})
//
//}


