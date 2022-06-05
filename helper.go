package main

import "fmt"

func appName() {
	fmt.Println("calogoflavorings")
}

func version() {
	fmt.Println("1.0.0")
}

func mylog(entries ...interface{}) {
	for _, entry := range entries {
		fmt.Println("Status:", entry)
	}
}
