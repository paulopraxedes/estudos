package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/paulopraxedes/estudos/go/gotrim"
)

func main() {

	if len(os.Args) == 2 {
		if isMenu(os.Args[1]) {
			fmt.Println("Args Error! Only menu options inserted. Check the values! ")
			return
		}
		fmt.Println(gotrim.RightTrim(gotrim.LeftTrim(os.Args[1])))

	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "-t", "--top":
			fmt.Println(gotrim.TopTrim(os.Args[2]))
		case "-b", "--bottom":
			fmt.Println(gotrim.BottomTrim(os.Args[2]))
		case "-l", "--left":
			fmt.Println(gotrim.LeftTrim(os.Args[2]))
		case "-r", "--right":
			fmt.Println(gotrim.RightTrim(os.Args[2]))
		case "-V", "--vertical":
			fmt.Println(gotrim.BottomTrim(gotrim.RightTrim(os.Args[2])))
		case "-H", "--horizontal":
			fmt.Println(gotrim.RightTrim(gotrim.LeftTrim(os.Args[2])))
		}
	}

}

func isMenu(argumento string) bool {
	if strings.Contains(argumento, "-t") || strings.Contains(argumento, "--top") ||
		strings.Contains(argumento, "-b") || strings.Contains(argumento, "--bottom") ||
		strings.Contains(argumento, "-l") || strings.Contains(argumento, "--left") ||
		strings.Contains(argumento, "-r") || strings.Contains(argumento, "--right") ||
		strings.Contains(argumento, "-V") || strings.Contains(argumento, "--vertical") ||
		strings.Contains(argumento, "-H") || strings.Contains(argumento, "--horizontal") {
		return true
	}

	return false
}
