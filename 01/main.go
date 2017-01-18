package main

import "github.com/sbinet/go-python"

func main() {
	python.Initialize()
	defer python.Finalize()

	fooModule := python.PyImport_ImportModule("foo")
	if fooModule == nil {
		panic("Error importing module")
	}

	helloFunc := fooModule.GetAttrString("hello")
	if helloFunc == nil {
		panic("Error importing function")
	}

	helloFunc.Call(python.Py_None, python.Py_None)
}
