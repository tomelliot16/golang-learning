/*************************************************************************
	> File Name: sorter.go
	> Author: DavidWang
	> Mail: davidwang2006@qq.com 
	> Created Time: Thu 25 Oct 2012 06:47:29 PM CST
 ************************************************************************/

package main
import (
	"fmt"
	"flag"
	"bufio"
	"io"
	"os"
	"strconv"
	"time"
	"sorter/algorithm/bubblesort"
)

var infile *string = flag.String("i","infile","File contains values for sorting")
var outfile *string = flag.String("o","outfile","File to receive sorted values")
var algorithm *string = flag.String("a","qsort","Sort algorithm")

func main(){

	flag.Parse()
	if infile != nil{
		fmt.Println("infile =",*infile,"outfile =",*outfile,"algorithm =",*algorithm)
	}
	values,err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm{
		case "qsort":
			fallthrough
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm",*algorithm,"is either unknown or unsupported")
		}
		t2 := time.Now()

		fmt.Println("The sorting process costs",t2.Sub(t1),"to complete")

		writeValues(values,*outfile)
	}else{
		fmt.Println(err)
	}
}

/**
read values from file
*/
func readValues(infile string)(values []int,err error){
	file,err := os.Open(infile)
	if err != nil{
		fmt.Println("Failed to open the input file ",infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int,0)
	for{
		line,isPrefix,err1 := br.ReadLine()
		if err1 != nil{
			if err1 != io.EOF{
				err = err1
			}
			break
		}
		if isPrefix{
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		value,err1 := strconv.Atoi(str)
		if err1 != nil{
			err = err1
			return
		}
		values = append(values,value)
	}
	return
}
/**write sorted data to file*/
func writeValues(values []int,outfile string) error{
	file,err := os.Create(outfile)
	if err != nil{
		fmt.Println("Failed to create the output file ",outfile)
		return err
	}
	defer file.Close()
	for _,value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str+"\n")
	}
	return nil
}