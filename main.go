package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const MAX_CELLS uint16 = 65535

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

	var content []byte
	buf := make([]byte, 4096)
	for {
		count, err := file.Read(buf)
		if err == io.EOF || count == 0 {
			break
		}
		content = append(content, buf[:count]...)
	}
	interpreter(content)
}

func interpreter(content []byte) {
	var tape [MAX_CELLS]int32
	cell_ptr := 0
	ch_ptr := 0
	loop := 0

	for ch_ptr < len(content) {
		switch content[ch_ptr] {
		case '>':
			cell_ptr++
		case '<':
			cell_ptr--
		case '+':
			tape[cell_ptr]++
		case '-':
			tape[cell_ptr]--
		case '.':
			fmt.Print(string(tape[cell_ptr]))
		case ',':
			fmt.Scanln(tape[cell_ptr])
		case '[':
			if tape[cell_ptr] == 0 {
				loop = 1
				for loop > 0 {
					ch_ptr++
					next_ch := content[ch_ptr]
					if next_ch == '[' {
						loop++
					} else if next_ch == ']' {
						loop--
					}
				}
			}
		case ']':
			loop = 1
			for loop > 0 {
				ch_ptr--
				prev_ch := content[ch_ptr]
				if prev_ch == '[' {
					loop--
				} else if prev_ch == ']' {
					loop++
				}
			}
			ch_ptr--
		}
		ch_ptr++
	}
}
