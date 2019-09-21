package jcli_test

import (
	"fmt"

	"github.com/rinor/jorcli/jcli"
)

func ExampleUtilsBech32Convert() {
	var (
		bech32    = "ta1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ulu8sap"
		newPrefix = "ca"
	)

	bech32WithNewPrefix, err := jcli.UtilsBech32Convert(bech32, newPrefix)

	if err != nil {
		fmt.Printf("UtilsBech32Convert: %s", err)
	} else {
		fmt.Printf("%s", string(bech32WithNewPrefix))
	}
	// Output:
	//
	// ca1s4uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4ugqz89h
}
