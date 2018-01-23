# INICompareKit

INICompareKit it's a tool to compare INI files to find which keys are missing on the compared file.

[![Build Status](https://travis-ci.org/marcelosousaalmeida/INICompareKit.svg?branch=master)](https://travis-ci.org/marcelosousaalmeida/INICompareKit)

## Installation

If you are using Go 1.6+ (or 1.5 with the `GO15VENDOREXPERIMENT=1` environment variable), you can install `INICompareKit` with the following command:

```bash
$ go get -u github.com/marcelosousaalmeida/INICompareKit
```

## Usage

```bash
$ ./INICompareKit --help
Usage of ./INICompareKit:
  -compared-filename string
    	Compared filename
  -count-only
    	Only compares the amount of keys (default true)
  -original-filename string
    	Original filename

$ ./INICompareKit -original-filename original.ini -compared-filename compared.ini -count-only false
```
## Contributing

All contributions are welcome, but if you are considering significant changes, please open an issue beforehand and discuss it with us.

## License

MIT. See the `LICENSE` file for more information.
