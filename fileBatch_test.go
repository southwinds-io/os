package os

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFileBatch(t *testing.T) {
	buf := bytes.Buffer{}
	for i := 0; i < 5; i++ {
		tt := &Test{Message: fmt.Sprintf("HELLO %d", i)}
		content, _ := json.Marshal(tt)
		_, err := buf.Write(AddLenHeader(content))
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	items, err := ReadFileBatchFromBytes(buf.Bytes())
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, item := range items {
		fmt.Printf("%s\n", string(item[:]))
	}
}

type Test struct {
	Message string
}

func TestAppendFileBatch(t *testing.T) {
	for i := 0; i < 5; i++ {
		tt := &Test{Message: fmt.Sprintf("HELLO %d", i)}
		content, _ := json.Marshal(tt)
		err := AppendFileBatch(content, "test.json", 0664)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	items, err := ReadFileBatch("test.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, item := range items {
		fmt.Printf("%s\n", string(item[:]))
	}
}
