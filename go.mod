module github.com/taobig/go-helper

go 1.15

require golang.org/x/text v0.3.7

//replace golang.org/x/text => github.com/golang/text v0.3.3
// or export GOPROXY=https://goproxy.io

// go get -u golang.org/x/text
// go mod tidy
