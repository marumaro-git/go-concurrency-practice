package service

import (
	"context"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime/trace"
	"sync"
)

func ConvertAllStep02(ctx context.Context, files []string) error {
	ctx, task := trace.NewTask(ctx, "convertAllStep02")
	defer task.End()

	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	var rerr error

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			if err := convertStep02(ctx, f); err != nil {
				mu.Lock()
				rerr = errors.Join(rerr, err)
				mu.Unlock()
			}
		}(file)
	}
	wg.Wait()

	return rerr
}

func convertStep02(ctx context.Context, file string) (rerr error) {
	defer trace.StartRegion(ctx, "convertStep02"+file).End()
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
