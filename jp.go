package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	bytes, err := ioutil.ReadAll(r)
	exitOnErr(err, "Error while reading from Stdin")
	var f interface{}
	err = json.Unmarshal(bytes, &f)
	exitOnErr(err, "Error while decoding json")
	pretty, err := json.MarshalIndent(f, "", "  ")
	exitOnErr(err, "Error while encoding back to pretty json")
	_, err = w.Write(pretty)
	exitOnErr(err, "Error while writing to Stdout")
}

func exitOnErr(err error, message string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, message)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
