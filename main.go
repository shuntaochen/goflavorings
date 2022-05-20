package main

import (
	"fmt"

	"flavorings/domain"
)

func main() {
	testentry()
	testmux()
	fmt.Println("go flavorings...")
	x := domain.MyData{}
	fmt.Println(x)
	appName()
	version()
	// testmysql()
	mymigrate()
	testjson()
	testtemplate()
	testjwtinit()
	testjwt()
	testjwtdecode()
	testsession()
	testcache()

}
