/*
   Created by jinhan on 17-8-1.
   Tip:
   Update:
*/
package lib

import (
	"fmt"
	"os"
	"testing"
)

func TestMd5FS(t *testing.T) {
	file, err := os.Open("/home/jinhan/code/src/github.com/hunterhug/GoWeb/favicon.ico")
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Println(Md5FS(file))
}
