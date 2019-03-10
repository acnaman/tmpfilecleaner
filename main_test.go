package main 

import (
	"os"
	"testing"
)

func TestFileDelete(t *testing.T){
	os.Mkdir("test/deletedir/1/foo", 0777)
	
	DeleteFile("test/test1.yaml")

	if f, _ := os.Stat("test/deletedir/1/foo"); f.IsDir() {
		t.Fatal()
	}
}


