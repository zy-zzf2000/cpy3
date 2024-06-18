package main

import (
	"fmt"
	python3 "github.com/go-python/cpy3"
	"log"
	"os"
)

func ImportModule(dir, name string) *python3.PyObject {
	sysModule := python3.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	pathStr, _ := pythonRepr(path)
	log.Println("before add path is " + pathStr)
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(""))
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))
	pathStr, _ = pythonRepr(path)
	log.Println("after add path is " + pathStr)
	return python3.PyImport_ImportModule(name)
}

func main() {
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}

	python3.PyRun_SimpleString("import sys")
	python3.PyRun_SimpleString("sys.path.append('/data/home/yatorozhang/sourcecode/go/cpy3/examples')")
	helloPy := python3.PyImport_ImportModule("test_python3")
	if helloPy == nil {
		log.Fatalf("helloPy is nil")
		return
	}
	helloFunc := helloPy.GetAttrString("test_func")
	if helloFunc == nil {
		log.Fatalf("helloFunc is nil")
	}
	var args = python3.PyTuple_New(2)
	python3.PyTuple_SetItem(args, 0, python3.PyUnicode_FromString("1+1+1"))
	helloPy3Str := helloFunc.Call(args, python3.Py_None)
	if helloPy3Str == nil {
		log.Fatalf("helloPy3Str is nil")
	}
	funcResultStr, _ := pythonRepr(helloPy3Str)
	log.Println("func result: " + funcResultStr)
}

func pythonRepr(o *python3.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil")
	}

	s := o.Repr()
	if s == nil {
		python3.PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method")
	}
	defer s.DecRef()

	return python3.PyUnicode_AsUTF8(s), nil
}
