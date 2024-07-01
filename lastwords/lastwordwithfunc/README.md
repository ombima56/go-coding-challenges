# lastword
## Instructions

Write a function LastWord that takes a stringand returns its last word followed by a\n`.

-    A word is a section of string delimited by spaces or by the start/end of the string.

Expected function
```
func LastWord(s string) string{

}
```
## Usage

Here is a possible program to test your function :
```bash
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(piscine.LastWord(" lorem,ipsum "))
	fmt.Print(piscine.LastWord(" "))
}
```
And its output :
```
$ go run . | cat -e
not$
lorem,ipsum$
$
$
```