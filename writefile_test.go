/*
  os library - © 2022 SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package os

import (
	"fmt"
	"os"
	"testing"
)

func TestWriteS3(t *testing.T) {
	content, err := os.ReadFile("writefile_test.go")
	if err != nil {
		t.Fatal(err)
	}
	if err = WriteFile(content, "s3://127.0.0.1:9000/test/writefile_test.go", "abcdefgh:12345678"); err != nil {
		t.Fatal(err)
	}
	fmt.Println("file uploaded successfully")
}
