package main

import (
	"iboxctl/pkg/root"
	_ "iboxctl/child/start/cmd"
	_ "iboxctl/child/stop/cmd"

)
var(
	stop= make(chan byte)
)
func main(){
	root.Run()
	//<-stop
}