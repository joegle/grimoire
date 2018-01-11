package main

import "time"
import "fmt"
import "os"
import "bufio"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readLoop() {
	pn := "/dev/ttyp5"

	f, err := os.Open(pn)
	check(err)

	scanner := bufio.NewScanner(f)

	for {
		hasInput := scanner.Scan()
				
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		} else {
			if hasInput {
				fmt.Println(scanner.Text())
			}
		}
	}
}

func main(){
	fmt.Println("hello")
	
	go readLoop()
	
	time.Sleep(50 * time.Second)
}

