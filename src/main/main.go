package main

import (
	"bufio"
	"fmt"
	"gosorter/msort"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

//var infile = "file/unsorted.dat"

func main() {
	infile, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return

	}
	infile = infile + "/sortfile/unsort.dat"
	values, err := readValueFromFile(infile)
	if err == nil {
		fmt.Println("read values:", values)
	} else {
		fmt.Println(err)
	}

	//msort.SortValues(values)
	//fmt.Println("read values:", values)

	msort.QuickSort(values, 0, len(values)-1)
	fmt.Println("read values:", values)

}

//get values from file,return []int values
func readValueFromFile(filepath string) (values []int, err error) {
	file, err1 := os.Open(filepath)
	if err1 != nil {
		fmt.Println("Failed to open file", filepath)
		err = err1
		return
	}

	//defer function
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("a too long line,seems unexpeted")
			return
		}

		//make int to string
		str := string(line)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}

	return
}
