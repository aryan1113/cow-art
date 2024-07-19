package main

import (
	"bufio" // buffered input
	"os"
	"io"
	"fmt"
	"strings"
	"unicode/utf8"
)

// add margni/boundary to ballon
func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

func calculateMaxLen(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return w
}

// make simpler lines
func normalizeStringsLen(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func main(){
	
	fmt.Println("enter text blob to throw :")
	reader := bufio.NewReader(os.Stdin)  // catch user input into a

	var lines[]string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		// if line == ""{
		// 	break
		// }
		if len(strings.TrimSpace(line)) == 0{
			break
		}
		lines = append(lines, string(line))
	} 
	// line, _ := reader.ReadString('\n')
	// // if err == Nil:

	// fmt.Println("i saw this", line)


	var cow = `        \  ^__^
         \ (oo)\_______
	   (__)\       )\/\
	       ||----w |
	       ||     ||
	`
	motapa := calculateMaxLen(lines)
	text_blob := normalizeStringsLen(lines, motapa)
	boundary := buildBalloon(text_blob, motapa)
	fmt.Println(boundary)
	fmt.Println(cow)

}