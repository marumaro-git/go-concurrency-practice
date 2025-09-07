package service

import (
	"context"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime/trace"

	"golang.org/x/sync/errgroup"
)

func ConvertAllStep03(ctx context.Context, files []string) error {
	ctx, task := trace.NewTask(ctx, "convertAllStep03")
	defer task.End()

	eg, ctx := errgroup.WithContext(ctx)

	for _, file := range files {
		file := file
		eg.Go(func() error {
			if err := convertStep03(ctx, file); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func convertStep03(ctx context.Context, file string) (rerr error) {
	defer trace.StartRegion(ctx, "convertStep03"+file).End()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

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
