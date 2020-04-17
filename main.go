package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
	"strings"
	"sort"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {

	highestSum := int32(0)
	sort.Slice(ar, func(i, j int) bool { return ar[i] > ar[j] })
	for ii:=0; ii<len(ar); ii++ {
		if ii+1 < len(ar) {
			highestSum++
			if ar[ii+1] < ar[ii] {
				ii=len(ar)
			}
		} else {
			if ar[ii] == ar[ii-1] {
				highestSum++
			}
		}	
	}
	return highestSum
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(arCount); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := birthdayCakeCandles(ar)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
