package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println("-------------1------------")

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))
	fmt.Println("-------------2------------")

	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	fmt.Println("-------------3------------")

	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println("-------------4------------")

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("-------------5------------")

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Println("-------------6------------")

	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))
	fmt.Println("-------------7------------")

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	fmt.Println("-------------8------------")
}
