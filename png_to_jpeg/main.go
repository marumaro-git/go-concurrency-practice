package main

import (
	"context"
	"fmt"
	"os"
	"png_to_jpeg/service"
	"runtime/trace"
)

func main() {
	ctx := context.Background()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: main <step>")
		os.Exit(1)
	}

	step := os.Args[1]
	files := os.Args[2:]

	if err := run(ctx, files, step); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, files []string, traceName string) error {
	f, err := os.Create(traceName)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		return err
	}
	switch traceName {
	case "step01":
		if err := service.ConvertAllStep01(ctx, files); err != nil {
			return err
		}
	case "step02":
		if err := service.ConvertAllStep02(ctx, files); err != nil {
			return err
		}
	case "step03":
		if err := service.ConvertAllStep03(ctx, files); err != nil {
			return err
		}
	case "step04":
		if err := service.ConvertAllStep04(ctx, files); err != nil {
			return err
		}
	case "step05":
		if err := service.ConvertAllStep05(ctx, files); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown step: %s", traceName)
	}

	trace.Stop()

	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}
