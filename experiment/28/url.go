package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	rawurl := "https://a b.com/"
	u, err := url.Parse(rawurl)
	if err != nil {

		fmt.Printf("%#v\n", err)
		if errs, ok := err.(*url.Error); ok {
			// fmt.Printf("%+v\n", errs)
			fmt.Println(errs.Op)
			fmt.Println(errs.URL)
			fmt.Println(errs.Err)
		}
		os.Exit(1)
	}
	fmt.Println(u)
}
