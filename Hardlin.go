package main

import (
	"fmt"
	"os"
	"syscall"
)
	func main() {
	fi, err := os.Stat("/home/")
	if err != nil {
		fmt.Println(err)
		return
	}
	nlink := uint64(0)
	if sys := fi.Sys(); sys !=nil {
		if stat, ok := sys.(*syscall.Stat_t); ok{
			nlink = uint64(stat.Nlink)
			}
		}
		fmt.Println("\n Total Number : ",nlink)
	}
