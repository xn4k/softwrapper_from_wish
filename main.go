package main

import (
	"bytes"
	"fmt"
)

func body() {
	var lineChars int
	var inputText string
	//Max. Characters & Getting text input
	fmt.Println("How much characters max. should every line contain? ")
	fmt.Scanln(&lineChars)
	fmt.Println("Enter your text: ")
	fmt.Scanln(&inputText)
	//check length
	lenIn := 0
	for range inputText {
		lenIn++
	}

	switch {
	case lenIn <= lineChars:
		fmt.Println("There are some problems here, this text can't be smaller" +
			" or equal to the max. char. per line, try again!")

	default:
		fmt.Println("The action can proceed since everything is ok!")
		linesShouldBe := lenIn / lineChars
		fmt.Println("In this new formatted text are", linesShouldBe, "lines:")

		/*var s = inputText
		for i := 5; i < len(s); i += 6 {
			s = s[:i] + "\n" + s[i:]
		}
		fmt.Println(s)
		*/
		fmt.Println(insertNth(inputText, lineChars))
	}

}

func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	//Loop
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}

func main() {
	body()

	/*fmt.Println("Enter your text: ")
	var inputText string
	fmt.Scanln(&inputText)
	//check and print out length
	lenIn := 0
	for range inputText {
		lenIn++
	}
	fmt.Println(lenIn)
	*/

	/*var x string = inputText

	scanner := bufio.NewScanner(strings.NewReader(x))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}*/

}
