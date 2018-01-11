package main

import "time"
import "fmt"
import "os"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func sendLoop() {

	pn := "/dev/ptyp5"
	f, err := os.OpenFile(pn, os.O_WRONLY, 0777)
	check(err)
	
	q := []int{600, 700, 800, 900, 800, 700}
	
	for i := 0 ; ; i++ {
		fmt.Print("sending...")

		output := fmt.Sprintf("%d\n", q[i%len(q)])

		n3, err := f.WriteString(output)
		check(err)
		fmt.Printf(" sent %d\n", n3)
		
		time.Sleep(time.Second)
		
	}
}

func main(){
	fmt.Println("hello")

	go sendLoop()

	// Keep alive
	time.Sleep(50 * time.Second)
}

