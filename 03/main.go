package main

import (
	"runtime"
	"sync"
	"time"

	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()

	var wg sync.WaitGroup

	fooModule := python.PyImport_ImportModule("foo")
	rnd := fooModule.GetAttrString("rand")

	state := python.PyEval_SaveThread()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			runtime.LockOSThread()
			defer wg.Done()

			s := python.PyGILState_Ensure()
			rnd.Call(python.PyTuple_New(0), python.PyDict_New())
			time.Sleep(100 * time.Millisecond)
			rnd.Call(python.PyTuple_New(0), python.PyDict_New())
			python.PyGILState_Release(s)
		}()
	}

	wg.Wait()
	python.PyEval_RestoreThread(state)
	python.Finalize()
}
