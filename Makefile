# Use bash as shell (Note: Ubuntu now uses dash which doesn't support PIPESTATUS).
SHELL := /bin/bash

.PHONY: benchmark
benchmark:
	@mkdir -p target/test
	@mkdir -p target/report
	go test -covermode=atomic \
		-run=NONE \
		-bench=. \
		-race \
		-cpuprofile=target/report/cpu.prof \
		-memprofile=target/report/mem.prof \
		-mutexprofile=target/report/mutex.prof \
		-coverprofile=target/report/coverage.prof \
		-v . | \
		tee >(go-junit-report > target/test/report.xml); \
		test $${PIPESTATUS[0]} -eq 0

.PHONY: pprof_cpu
pprof_cpu: benchmark
	go tool pprof target/report/cpu.prof

.PHONY: pprof_mem
pprof_mem: benchmark
	go tool pprof target/report/mem.prof

