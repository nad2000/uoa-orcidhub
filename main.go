//go:generate swagger generate client -f assets/uoa-employment-v1.json
//go:generate swagger generate client -f assets/orcidhub-api-v0.1-spec.json  -A orcidhub

package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
