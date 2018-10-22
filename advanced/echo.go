package main

import (
	"flag"
	"os"
)

// nFlag 的类型会自动推断
var nFlag = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	// flag.PrintDefaults() 打印 flag 的使用帮助信息，本例中打印的是：
	// -n=false: print newline
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *nFlag {
				s += Newline
			}
		}
		s += flag.Arg(i)
	}

	os.Stdout.WriteString(s)
}
