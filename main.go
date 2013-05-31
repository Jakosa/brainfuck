package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const MAX_CELLS = 65536

func main() {
	// TODO: add -h --help cmd
	var sw bool
	flag.BoolVar(&sw, "s", false, "-s, -s=true, default is false")
	flag.BoolVar(&sw, "stopwatch", false, "--stopwatch, --stopwatch=true, default is false")
	var file string
	flag.StringVar(&file, "f", "", "-f=file.bf")
	flag.StringVar(&file, "file", "", "--file=file.bf")
	flag.Parse()

	if file == "" {
		// TODO: print full usage
		fmt.Println("Please specify a file: brainfuck -f file.bf")
		return
	}
	scanner(file)
}

func scanner(fl string) {
	file, err := os.Open(fl)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var content string
	buf := make([]byte, 4096)
	for {
		count, err := file.Read(buf)
		if err == io.EOF || count == 0 {
			break
		}
		content += string(buf[:count])
	}
	interpreter(content)
}

func interpreter(content string) {
	var tape [MAX_CELLS]uint
	cell_ptr := 0
	ch_ptr := 0

	for ch_ptr < len(content) {
		return
		switch string(content[ch_ptr]) {
		case ">":
			cell_ptr++
		case "<":
			cell_ptr--
		case "+":
			tape[cell_ptr]++
		case "-":
			tape[cell_ptr]--
		case ".":
			fmt.Println(tape[cell_ptr])
		case ",":
			fmt.Scanln(tape[cell_ptr])
		case "[":
			if tape[cell_ptr] == 0 {
				for {
					ch_ptr++
					if string(content[ch_ptr]) == "]" {
						ch_ptr++
						break
					}
				}
			}
		case "]":
			if tape[cell_ptr] != 0 {
				for {
					ch_ptr--
					if string(content[ch_ptr]) == "[" {
						break
					} else {
						ch_ptr--
					}
				}
			}
		}
		ch_ptr++
	}
}
