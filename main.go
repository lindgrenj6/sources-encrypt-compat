package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RedHatInsights/sources-api-go/util"
)

func main() {
	encrypt := flag.String("encrypt", "", "encrypt string")
	decrypt := flag.String("decrypt", "", "decrypt string")
	flag.Parse()

	if *encrypt == "" && *decrypt == "" {
		flag.Usage()
		os.Exit(0)
	}

	var out string
	var err error

	switch {
	case *encrypt != "":
		out, err = util.Encrypt(*encrypt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to encrypt: %v\n", err.Error())
			os.Exit(1)
		}
	case *decrypt != "":
		out, err = util.Decrypt(*decrypt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to decrypt: %v\n", err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "no invalid operation found\n")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(out)
}
