package utils

import (
	"encoding/xml"
	"fmt"
	"encoding/json"
)

func PrintDebugXML(v interface{}) {
	var RQ = make([]byte, 1024)
	RQ, _ = xml.MarshalIndent(v, " ", "	 ")
	fmt.Println(string(RQ))
}

func PrintDebugJson(v interface{}) {
	var RQ = make([]byte, 1024)
	RQ, _ = json.MarshalIndent(v, " ", "  ")
	fmt.Println(string(RQ))
}
