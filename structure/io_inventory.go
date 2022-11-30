package structure

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// file should be located in the folder called "inventory_objects"
func GetListOfAllItems(filename string) []string {
	file, err := os.Open(filename)
	buffer := make([]string, 0)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		sc := bufio.NewScanner(file)
		for sc.Scan() {
			ii := sc.Text()
			buffer = append(buffer, ii)
		}
		for i, e := range buffer {
			fmt.Println(i, "-> ", strings.Split(e,":"))
		}
	}
	return buffer
}
