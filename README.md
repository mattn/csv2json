# csv2json

CSV/JSON converter

## Usage

Read CSV from stdin, Writer JSON to stdout.

```
Usage of csv2json:
  -field string
        collect field
  -header
        use header (default true)
  -obj
        output object
  -raw
        raw string
```

list tweet texts
```
$ cat mytweets.csv | csv2json -raw -f text
```

convert CSV to JSON, and format with jq
```
$ cat mytweets.csv | csv2json | jq .
```

## Installation

```
$ go get github.com/mattn/csv2json
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
