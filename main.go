package main

import (
    "fmt"
    "bufio"
    "os"
    //"os/exec"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var text string
    var compareText string = os.Args[1]
    //c := exec.Command("clear")
    for text != "q" {  // break the loop if text == "q"
	    fmt.Print("$ ")
        scanner.Scan()
        text = scanner.Text()
        if text != "q"{
		//if text == "all right" {
		if text == compareText {
			fmt.Println(": ", text)
		    //c.Stdout = os.Stdout
		    fmt.Println("\033[H\033[2J")
		    //fmt.Println("\033[2J")
		    //c.Run()
	    }else{
		    fmt.Println("\033[H\033[2J")
		    fmt.Println(compareText, ".\n")

	    }
        }
    }
}
