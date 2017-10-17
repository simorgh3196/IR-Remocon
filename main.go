package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

const (
	signal = "5953ffffffff01001e001e000f0f0f0f000f0f0f8007808707808787078003c0c303c003c0c303e0e1e1e1e1e101e0e1e1f000f0f0f0f0f0f0f0707800780078787800383c3c3c003c003c3c3c1c1e1e1e1e1e1e1e1e000f0f0f000f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000045"
	port   = "/dev/tty.usbserial-00002bf6"
	// port = "COM5"
)

func main() {
	// Set up options.
	config := &serial.Config{
		Name:        port,
		Baud:        115200,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
		ReadTimeout: 10 * time.Second,
	}

	// Open the port.
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	b := []byte{0x72} // = 'r' (Receive)
	if _, err := port.Write(b); err != nil {
		log.Fatal("0x72 is failed")
	}

	r := make([]byte, 240)
	if _, err := port.Read(r); err != nil {
		log.Fatal("Read Error 59h")
	}
	log.Println("59h - Y:", r)

	r = make([]byte, 240)
	if _, err := port.Read(r); err != nil {
		log.Fatal("Read Error 53h")
	}
	log.Println("53h - S:", r)

	r = make([]byte, 240)
	if _, err := port.Read(r); err != nil {
		log.Fatal("Read Error Data")
	}
	log.Println("Data:", r)

	r = make([]byte, 240)
	if _, err := port.Read(r); err != nil {
		log.Fatal("Read Error 45h")
	}
	log.Println("45h - E:", r)

	// Dummy
	//b := []byte{0x00}
	//if _, err := port.Write(b); err != nil {
	//	//log.Fatalf("port.Write Dummy Error: %v", err)
	//	log.Printf("port.Write Dummy Error: %v", err)
	//	b = []byte{0x63}
	//	_, _ = port.Write(b)
	//}

	//// LED-ON
	//b = []byte{0x69}
	//if _, err := port.Write(b); err != nil {
	//	log.Fatalf("port.Write LED-ON Error: %v", err)
	//}

	//r := make([]byte, 240)
	//if n, err := port.Read(r); err != nil {
	//	log.Fatalf("port.Read LED-ON Error: %v", err)
	//} else if n == 0 {
	//	log.Fatalln("Can not received 0x59")
	//} else {
	//	fmt.Println("Recieve data of LED-ON: ", r) // 0x4F(= 'ON') or 0x59(= 'off')
	//}

	//// Submit
	//b = []byte{0x74}
	//if _, err := port.Write(b); err != nil {
	//	log.Fatalf("port.Write: %v", err)
	//}

	//r = make([]byte, 240)
	//if n, err := port.Read(r); err != nil {
	//	log.Fatalf("port.Read Submit Error: %v", err)
	//} else if n == 0 {
	//	log.Fatalln("Can not received 0x59")
	//} else {
	//	fmt.Println("Recieve data of Submit: ", r) // 0x59
	//}

	//// Select channel [ch1: 0x31, ch2: 0x32, ch3: 0x33, ch4: 0x34]
	//b = []byte{0x31}
	//if _, err := port.Write(b); err != nil {
	//	log.Fatalf("port.Write: %v", err)
	//}

	//r = make([]byte, 240)
	//if n, err := port.Read(r); err != nil {
	//	log.Fatalf("port.Read Channel Error: %v", err)
	//} else if n == 0 {
	//	log.Fatalln("Can not received 0x59")
	//} else {
	//	fmt.Println("Recieve data of channel: ", r) // 0x59
	//}

	//// Submit Signal
	//b = []byte(signal)
	//if _, err := port.Write(b); err != nil {
	//	log.Fatalf("port.Write: %v", err)
	//}

	//r = make([]byte, 240)
	//if n, err := port.Read(r); err != nil {
	//	log.Fatalf("port.Read Signal Error: %v", err)
	//} else if n == 0 {
	//	log.Fatalln("Can not received 0x59")
	//} else {
	//	fmt.Println("Recieve data of Signal: ", r) // 0x45
	//}

	fmt.Println("Finish.")
}
