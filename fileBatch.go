/*
  os library - © 2022 SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package os

import (
	"bytes"
	"io"
	"os"
	"unsafe"
)

// AppendFileBatch appends a []byte to an existing file or to a new file if not there
func AppendFileBatch(content []byte, path string, perm os.FileMode) error {
	contentWithLen := AddLenHeader(content)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = file.Write(contentWithLen)
	return err
}

func AddLenHeader(content []byte) []byte {
	buf := new(bytes.Buffer)
	l := len(content)
	buf.Write(intToByteSlice(int64(l))) // first 8 bytes contains length of request
	buf.Write(content)
	return buf.Bytes()
}

func ReadFileBatch(path string) ([][]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ReadFileBatchFromBytes(content)
}

func ReadFileBatchFromBytes(content []byte) ([][]byte, error) {
	files := make([][]byte, 0)
	lenOfFile := make([]byte, 8)
	reader := bytes.NewReader(content)
	var ix int64 = 0
	for {
		lenOfFile = make([]byte, 8)
		_, err := io.ReadFull(reader, lenOfFile)
		if err != nil {
			return nil, err
		}
		ix += 8
		fileLen := byteSliceToInt(lenOfFile)
		fileBuffer := make([]byte, fileLen)
		_, err = reader.Seek(ix, 0)
		if err != nil {
			return nil, err
		}
		_, err = io.ReadFull(reader, fileBuffer)
		if err != nil {
			return nil, err
		}
		ix += int64(len(fileBuffer))
		files = append(files, fileBuffer)
		if reader.Len() == 0 {
			break
		}
		_, err = reader.Seek(ix, 0)
	}
	return files, nil
}

func intToByteSlice(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func byteSliceToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
