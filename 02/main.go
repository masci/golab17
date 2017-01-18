package main

import (
	"sync"

	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()

	var wg sync.WaitGroup
	wg.Add(2)

	fooModule := python.PyImport_ImportModule("foo")
	odds := fooModule.GetAttrString("print_odds")
	even := fooModule.GetAttrString("print_even")

	// Initialize locks the GIL but at this point we don't need it
	// anymore. We release it so that goroutines can acquire it
	state := python.PyEval_SaveThread()

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

	python.PyEval_RestoreThread(state)
	python.Finalize()
}
