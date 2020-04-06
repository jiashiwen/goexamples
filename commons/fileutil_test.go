package commons

import "testing"

func TestCreateNewFile(t *testing.T) {
	_,err:=CreateNewFile("/tmp/testcreatefile")

	if err!=nil{
		t.Error(err)
	}
}

func TestRemoveFile(t *testing.T) {
	if FileExists("/tmp/testcreatefile") {
		err := RemoveFile("/tmp/testcreatefile")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestIsDir(t *testing.T) {
	dir:="/tmp"
	is:=IsDir(dir)
	if !is {
		t.Error(dir,"is not dir")
	}
}