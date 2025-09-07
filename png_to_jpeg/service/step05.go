package service

import (
	"context"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime/trace"

	"github.com/sourcegraph/conc/panics"
	"github.com/sourcegraph/conc/pool"
)

func ConvertAllStep05(ctx context.Context, files []string) error {
	ctx, task := trace.NewTask(ctx, "convertAllStep05")
	defer task.End()

	pool := pool.New().WithErrors().WithContext(ctx)

	for _, file := range files {
		file := file
		pool.Go(func(ctx context.Context) (rerr error) {

			var c panics.Catcher
			defer func() {
				if r := c.Recovered(); r != nil {
					rerr = r.AsError()
				}
			}()
			c.Try(func() {
				rerr = convertStep05(ctx, file)
			})

			if rerr != nil {
				return rerr
			}
			return nil
		})
	}
	if err := pool.Wait(); err != nil {
		return err
	}
	return nil
}

func convertStep05(ctx context.Context, file string) (rerr error) {
	defer trace.StartRegion(ctx, "convertStep05"+file).End()

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
