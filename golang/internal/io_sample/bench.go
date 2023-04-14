package io_sample

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Write file that will emulate io stream
// it not bytes precise, but close enough
func WriteFile(size int64, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriterSize(file, 5)
	defer writer.Flush()
	for i := int64(0); i < size; i++ {
		str := fmt.Sprintf("%d ", i)
		// if len(str) >= writer.Available() {
		// 	break
		// }
		if _, err := fmt.Fprint(writer, str); err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(filename string) (io.ReadCloser, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func MyFileReader(reader io.ReadCloser) error {
	defer reader.Close()
	rdr := bufio.NewReader(reader)
	var item int
	for {
		_, err := fmt.Fscanf(rdr, "%d ", &item)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println(item)
	}
}
