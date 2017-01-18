package main

import (
	"fmt"
	"sync"

	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()
	// `Initialize` acquires the GIL but we need to leave it
	// so that our gouroutines can acquire it when needed
	state := python.PyEval_SaveThread()

	var wg sync.WaitGroup

	_gstate := python.PyGILState_Ensure()
	fooModule := python.PyImport_ImportModule("foo")
	rnd := fooModule.GetAttrString("print_rand")
	python.PyGILState_Release(_gstate)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			s := python.PyGILState_Ensure()
			val := rnd.Call(python.PyTuple_New(0), python.PyDict_New())
			python.PyGILState_Release(s)

			fmt.Printf("%v\n", python.PyFloat_AsDouble(val))
		}()
	}

	wg.Wait()
	python.PyEval_RestoreThread(state)
	python.Finalize()
}
