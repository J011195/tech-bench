package main

import (
	"embed"
	"fmt"
	"log/slog"
	"runtime"

	"github.com/J011195/tech-bench/util"
)

//go:embed config/single/*.yaml
var config string

//go:embed config/multiple/**.yaml
var configs embed.FS

const commandName = "main"

func main() {
	RunTechBench()
}

func RunTechBench() {
	PrintBuildCompilerVersion()
	PrintSingleEmbeddedFile()
	PrintMultipleEmbeddedFile()
}

func PrintSingleEmbeddedFile() {
	slog.Info(commandName, util.GetFuncName(), "Printing single config")
	fmt.Println(config)
}

func PrintMultipleEmbeddedFile() {
	slog.Info(commandName, util.GetFuncName(), "Printing multiple configs")
	configs, err := configs.ReadDir("config/multiple")
	if err != nil {
		slog.Error(commandName, util.GetFuncName(), fmt.Sprintf("configs.ReadDir error, %s", err.Error()))
		panic(err)
	}

	for _, file := range configs {
		fmt.Println(file.Name())
	}
}

func PrintBuildCompilerVersion() {
	slog.Info(commandName, util.GetFuncName(), fmt.Sprintf("Compiled with GO version: %s", runtime.Version()))
}
