package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)
var(
	close chan bool = make(chan bool)
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "当前 IP：",strings.Split(r.RemoteAddr,":")[0],"  来自于：中国 四川 成都  电信")
}

func main() {
	port:="8000"
	if len(os.Args) >1 {
		port=os.Args[1]
	}
	http.HandleFunc("/", IndexHandler)
	err:=http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-close
}
