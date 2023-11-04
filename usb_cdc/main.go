package main

import (
	"bufio"
	"fmt"
	"machine"
	"os"
	"time"
)

func waitInitializedSerial() {
	for !machine.Serial.DTR() {
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	waitInitializedSerial()

	fmt.Printf("hello tinygo\r\n")

	msg := ""
	fmt.Scanf("%s\r\n", &msg)
	fmt.Printf("msg: %q\r\n", msg)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("You typed: %s\r\n", scanner.Text())
	}
}
