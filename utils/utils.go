package utils

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func GetLocalIPs() []string {
	result := []string{"localhost"}
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		log.Println("Get local IP addr failed!!!")
		return result
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	return result
}

func listDir(dirPth string) (a []string) {
	dir, err := ioutil.ReadDir(dirPth)
	var files = []string{}
	if err != nil {
		return nil
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		files = append(files, PthSep+fi.Name())
	}
	return files
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
