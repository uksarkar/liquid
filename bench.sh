#!/bin/bash

# Install benchstat
go install golang.org/x/perf/cmd/benchstat@latest

# Create benchmark test file
cat > expressions/benchmark_test.go << 'EOF'
package expressions

import (
	"testing"
)

var (
	asciiExpr    = "cart.total >= 18 and cart.status == 'paid' or cart.is_admin?"
	unicodeExpr  = "用户.年龄 >= 18 and 用户.国家 == '日本' or 用户.활동?"
)

func BenchmarkScanASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scanExpression(asciiExpr)
	}
}

func BenchmarkScanUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scanExpression(unicodeExpr)
	}
}
EOF

echo "=== Benchmarking OLD commit ==="
git checkout a71aedf199fcdbe30a92a87891e7bd11665ccc34 > /dev/null 2>&1
go test -bench=. -benchmem -count=5 ./expressions > old_bench.txt 2>&1

echo "=== Benchmarking NEW commit ==="  
git checkout - > /dev/null 2>&1
go test -bench=. -benchmem -count=5 ./expressions > new_bench.txt 2>&1

echo "=== RESULTS ==="
echo "Old commit:"
grep -A 5 "Benchmark" old_bench.txt || echo "No benchmarks found in old version"
echo ""
echo "New commit:"
grep -A 5 "Benchmark" new_bench.txt || echo "No benchmarks found in new version"

# If we have proper benchmark output, compare with benchstat
if grep -q "Benchmark" old_bench.txt && grep -q "Benchmark" new_bench.txt; then
    echo ""
    echo "=== BENCHSTAT COMPARISON ==="
    benchstat old_bench.txt new_bench.txt
else
    echo ""
    echo "=== MANUAL COMPARISON ==="
    echo "Old test duration: $(grep '^ok' old_bench.txt)"
    echo "New test duration: $(grep '^ok' new_bench.txt)"
fi

# Cleanup
echo ""
echo "=== CLEANING UP ==="
rm -f old_bench.txt new_bench.txt expressions/benchmark_test.go
echo "Temporary files removed"