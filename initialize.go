package isutools

import (
	isucache "github.com/grease-the-benchmark/isutools/v2/cache"
	"github.com/grease-the-benchmark/isutools/v2/internal/benchmark"
)

func BeforeInitialize() {
	isucache.AllPurge()
}

func AfterInitialize() {
	benchmark.Start()
}
