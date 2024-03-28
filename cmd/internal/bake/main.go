package main

import (
	"context"
	"fmt"
	"github.com/apple/pkl-go/pkl"
	"os"
	"path/filepath"
	"runtime"
)

func getProjectRootDir() string {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Errorf("failed to get current dir"))
	}
	return filepath.Join(path, "../../../../")
}

func main() {
	ctx := context.Background()
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	projectRoot := getProjectRootDir()
	binary, err := evaluator.EvaluateExpressionRaw(ctx, pkl.FileSource(projectRoot, "pkl/dev/config.pkl"), "module")
	if err != nil {
		panic(err)
	}
	outDir := filepath.Join(projectRoot, "baked/dev/")
	if err = os.MkdirAll(outDir, 0o777); err != nil {
		panic(err)
	}
	if err = os.WriteFile(filepath.Join(outDir, "appConfig"), binary, 0o777); err != nil {
		panic(err)
	}
}
