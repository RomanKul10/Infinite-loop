package mylogger

import "log"

func LictenForLog(ch chan string) {

	for {
		msg := <-ch
		log.Println(msg)
	}
}
