# Serial IO

Simple go code to read from and write to a serial TTY.

## Usage

```bash
./srialio -h
Usage of srialio:
      --tty string         Serial TTY device (e.g., /dev/ttyUSB0)
  -o, --operation string   Operation: read or write
  -d, --data string        Input data to write to the serial TTY
  -f, --file string        File containing data to write to the serial TTY
  -b, --baud int           Baud rate for serial communication (default 9600)
  -s, --sleep duration     Sleep duration between write operations
      --debug              Enable debug messages

Examples:
    srialio --tty /dev/cu.usbmodem14301 -o write -f 'serial-input.txt'  --debug
    srialio --tty /dev/cu.usbmodem14301 -o read -b 115200
```
