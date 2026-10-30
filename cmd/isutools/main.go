package main

import (
	"github.com/grease-the-benchmark/isutools/v2/analysers/db"
	"github.com/grease-the-benchmark/isutools/v2/analysers/echo"
	"github.com/grease-the-benchmark/isutools/v2/analysers/embed"
	"github.com/grease-the-benchmark/isutools/v2/analysers/fasthttp"
	"github.com/grease-the-benchmark/isutools/v2/analysers/fiber"
	"github.com/grease-the-benchmark/isutools/v2/analysers/gin"
	"github.com/grease-the-benchmark/isutools/v2/analysers/http"
	"github.com/grease-the-benchmark/isutools/v2/analysers/importer"
	"github.com/grease-the-benchmark/isutools/v2/analysers/initialize"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		embed.Analyzer,
		echo.Analyzer,
		gin.Analyzer,
		http.Analyzer,
		fiber.Analyzer,
		fasthttp.Analyzer,
		db.Analyzer,
		initialize.Analyzer,
		importer.Analyzer,
	)
}
