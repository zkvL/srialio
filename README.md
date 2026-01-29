# Serial IO

Simple go code to read from and write to a serial TTY.

> ## No further support for this one due to more advanced alternatives
>
> This tool was used to understand and interact with serial communications but more robust solutions exist to that purpose:
>
> - [TIO](https://github.com/tio/tio) - A serial device I/O tool
>
> [![DEPRECATED](https://img.shields.io/badge/STATUS-DEPRECATED-red.svg)]()
> [![License: GNU General Public License v3.0](https://img.shields.io/badge/License-GNU-yellow.svg)](https://www.gnu.org/licenses/gpl-3.0.html)

## Usage

```bash
❯❯❯ ./srialio -h
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

## Install
```bash
❯❯❯ go install github.com/zkvL/srialio@latest
```

