package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/tarm/serial"
)

func readFromSerial(tty string, baudRate int, debug bool) {
	c := &serial.Config{Name: tty, Baud: baudRate}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Printf("[!] Error opening serial port: %v\n", err)
		os.Exit(1)
	}
	defer s.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\nReceived interrupt signal. Exiting...")
		os.Exit(0)
	}()

	fmt.Printf("[-] Reading from %s... Press Ctrl+C to exit.\n", tty)

	buf := make([]byte, 128)
	for {
		n, err := s.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("[!] Error reading from serial port: %v\n", err)
			os.Exit(1)
		}
		if n > 0 {
			data := buf[:n]
			if debug {
				fmt.Printf("[-] Read: %s", data)
			} else {
				fmt.Print(string(data))
			}
		}
	}
}

func writeToSerial(tty string, baudRate int, data, file string, sleep time.Duration, debug bool) {
	c := &serial.Config{Name: tty, Baud: baudRate}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Printf("[!] Error opening serial port: %v\n", err)
		os.Exit(1)
	}
	defer s.Close()

	fmt.Printf("[-] Writing to %s...\n", tty)

	if data != "" {
		writeStringToSerial(s, data, sleep, debug)
	} else if file != "" {
		writeFileToSerial(s, file, sleep, debug)
	}
}

func writeStringToSerial(s io.Writer, data string, sleep time.Duration, debug bool) {
	for _, char := range data {
		byteData := []byte{byte(char)}
		if debug {
			fmt.Printf("[-] Write: %s", byteData)
		}

		_, err := s.Write(byteData)
		if err != nil {
			fmt.Printf("[!] Error writing to serial port: %v\n", err)
			os.Exit(1)
		}

		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
}

func writeFileToSerial(s io.Writer, filePath string, sleep time.Duration, debug bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("[!] Error opening data file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if debug {
			fmt.Printf("[-] Write: %s\n", line)
		}

		_, err := s.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Printf("[!] Error writing to serial port: %v\n", err)
			os.Exit(1)
		}

		if sleep > 0 {
			time.Sleep(sleep)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("[!] Error reading data: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	tty := flag.String("tty", "", "Serial TTY device (e.g., /dev/ttyUSB0)")
	operation := flag.StringP("operation", "o", "", "Operation: read or write")
	data := flag.StringP("data", "d", "", "Input data to write to the serial TTY")
	file := flag.StringP("file", "f", "", "File containing data to write to the serial TTY")
	baudRate := flag.IntP("baud", "b", 9600, "Baud rate for serial communication")
	sleep := flag.DurationP("sleep", "s", 0, "Sleep duration between write operations")
	debug := flag.Bool("debug", false, "Enable debug messages")

	flag.CommandLine.SortFlags = false
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *tty == "" || *operation == "" {
		fmt.Println("Please provide both TTY and operation parameters.")
		os.Exit(1)
	}

	switch *operation {
	case "read":
		readFromSerial(*tty, *baudRate, *debug)
	case "write":
		if *data == "" && *file == "" {
			fmt.Println("Please provide either data or file for write operation.")
			os.Exit(1)
		}
		writeToSerial(*tty, *baudRate, *data, *file, *sleep, *debug)
	default:
		fmt.Println("[!] ERROR: Invalid operation. Supported operations: read, write.")
		os.Exit(1)
	}
}
