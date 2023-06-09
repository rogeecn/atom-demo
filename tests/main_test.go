package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_main(t *testing.T) {
	fmt.Println("xxx", filepath.Dir("a"), "xxx")

}
