package main

import "C"
import "log"

// We assume that inter-thread communication by using channels is
// more efficient than connecting to unix sockets.
//
// Try to get channel pointer from server thread and return it to C function
// When GetChanPointer exits the channel remains alive in server.

//export GetChanPointer
func GetChanPointer(cChan *chan string) {
	signal := make(chan bool)
	log.Printf("cChan=%p", cChan)

	go createPointer(cChan, signal)

	log.Println("signal<-true")
	signal <- true
	log.Println("entering select")
	select {
	case rc := <-(*cChan):
		log.Printf("rc = %s\n", rc)
	default:
		log.Println("default case")
	}
	log.Printf("cChan=%p", cChan)
	log.Printf("*cChan=%p", *cChan)
}

func createPointer(s *chan string, signal <-chan bool) {
	for {
		log.Println("enter select")
		select {
		case <-signal:
			log.Println("received signal")
			if *s == nil {
				log.Println("*s==nil")
				*s = make(chan string)
				log.Printf("*s==%p", *s)
				*s <- "*s==nil"
			} else {
				log.Println("*s!=not nil")
				*s <- "*s!=not nil"
			}
		}
	}
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
