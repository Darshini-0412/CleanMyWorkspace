package main

import (
	"fmt"

	"github.com/Mentor-Paris/CleanMyWorkspace"
)

func main() {
	workspace := CleanMyWorkspace.GenererateWorkSpace()

	for _, row := range *workspace {
		for _, item := range row {
			if item == nil {
				fmt.Print("|nil")
			} else {
				fmt.Printf("|%s", *item)
			}
		}
		fmt.Println("|")
	}
}
