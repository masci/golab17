package main

import (
	"sync"

	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()
	defer python.Finalize()

	var wg sync.WaitGroup
	wg.Add(2)

	fooModule := python.PyImport_ImportModule("foo")
	odds := fooModule.GetAttrString("print_odds")
	even := fooModule.GetAttrString("print_even")

	go func() {
		_gstate := python.PyGILState_Ensure()
		odds.Call(python.PyTuple_New(0), python.PyDict_New())
		python.PyGILState_Release(_gstate)

		wg.Done()
	}()

	go func() {
		_gstate := python.PyGILState_Ensure()
		even.Call(python.PyTuple_New(0), python.PyDict_New())
		python.PyGILState_Release(_gstate)

		wg.Done()
	}()

	wg.Wait()
}
