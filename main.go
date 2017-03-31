package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	header = flag.Bool("header", true, "use header")
	field  = flag.String("field", "", "collect field")
	raw    = flag.Bool("raw", false, "raw string")
	obj    = flag.Bool("obj", true, "output object")
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if *obj && *raw {
		fatal(errors.New("-obj and -raw both should not be specified"))
	}

	var ff []string
	if *field != "" {
		ff = strings.Split(*field, ",")
	}

	fm := make(map[string]int)
	csvr := csv.NewReader(os.Stdin)

	if *header {
		record, err := csvr.Read()
		if err != nil {
			fatal(err)
		}
		for i, v := range record {
			fm[v] = i
		}
	}

	first := true
	for {
		record, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fatal(err)
		}

		if first {
			first = false
			if !*raw {
				fmt.Print("[")
			}
			if !*header {
				for i := 0; i < len(record); i++ {
					fm[fmt.Sprintf("f%d", i+1)] = i
				}
			}
		} else if !*raw {
			fmt.Print(",")
		}

		if *obj {
			out := make(map[string]string)
			if len(ff) > 0 {
				for _, fi := range ff {
					i := fm[fi]
					out[fi] = record[i]
				}
			} else {
				for k, v := range fm {
					out[k] = record[v]
				}
			}
			b, err := json.Marshal(out)
			if err != nil {
				fatal(err)
			}
			fmt.Print(string(b))
		} else {
			if len(ff) > 0 {
				out := []string{}
				for _, fi := range ff {
					i := fm[fi]
					out = append(out, record[i])
				}
				record = out
			}
			if *raw {
				fmt.Print(strings.Join(record, ""))
			} else {
				b, err := json.Marshal(record)
				if err != nil {
					fatal(err)
				}
				fmt.Print(string(b))
			}
		}
	}
	if !*raw {
		fmt.Print("]")
	}
}
