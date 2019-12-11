package adapter

import (
	"bytes"
	"strconv"
)

type ArrayToStringAdapter struct {
	IntArray []int
}

func (arr *ArrayToStringAdapter) stringGenerator() string {
	var buffer bytes.Buffer
	if arr.IntArray != nil {
		for i := 0; i < len(arr.IntArray); i++ {
			buffer.WriteString(strconv.Itoa(arr.IntArray[i]))
		}
	}
	return buffer.String()
}
