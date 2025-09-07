package service

import (
	"context"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime/trace"
)

func ConvertAllStep01(ctx context.Context, files []string) error {
	ctx, task := trace.NewTask(ctx, "convertAll")
	defer task.End()

	for _, file := range files {
		if err := convert(ctx, file); err != nil {
			return err
		}
	}
	return nil
}

func convert(ctx context.Context, file string) (rerr error) {
	defer trace.StartRegion(ctx, "convert"+file).End()
	src, err := os.Open(file)
	if err != nil {
		return err
	}
	defer src.Close()
	pngimg, err := png.Decode(src)
	if err != nil {
		return err
	}

	ext := filepath.Ext(file)
	jpegfile := file[:len(file)-len(ext)] + ".jpg"

	dst, err := os.Create(jpegfile)
	if err != nil {
		return err
	}
	defer func() {
		dst.Close()
		if rerr != nil {
			rerr = errors.Join(rerr, os.Remove(jpegfile))
		}
	}()

	if err := jpeg.Encode(dst, pngimg, nil); err != nil {
		return err
	}

	if err := dst.Sync(); err != nil {
		return err
	}

	return nil
}
