package profiling

import profiling "."

func ExampleStart() {
	finishProfiling := profiling.Start("cpu.prof", "mem.prof")
	defer finishProfiling()
}
