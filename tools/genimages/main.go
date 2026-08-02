// Command genimages downscales the source images in assets/ and writes
// web-sized JPEGs into ui/static/img/, which is what //go:embed ships.
//
// Run it from the repository root after adding or replacing anything in assets/:
//
//	go run ./tools/genimages
//
// The output is committed, so building the server needs no tooling at all.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
)

const (
	srcDir      = "assets"
	dstDir      = "ui/static/img"
	maxWidth    = 1200
	jpegQuality = 82
)

func main() {
	err := filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch strings.ToLower(filepath.Ext(path)) {
		case ".jpg", ".jpeg", ".png":
			return convert(path)
		default:
			return nil
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}

func convert(srcPath string) error {
	f, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("decode %s: %w", srcPath, err)
	}

	// Preserve the aspect ratio, and never upscale.
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	if w > maxWidth {
		h = h * maxWidth / w
		w = maxWidth
	}

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, bounds, draw.Src, nil)

	outPath, err := destFor(srcPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}

	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if err := jpeg.Encode(out, dst, &jpeg.Options{Quality: jpegQuality}); err != nil {
		return err
	}

	fmt.Printf("%s -> %s (%dx%d)\n", srcPath, outPath, w, h)
	return out.Close()
}

// destFor mirrors a path under assets/ into ui/static/img/, always as .jpg
// since that is what we encode.
func destFor(srcPath string) (string, error) {
	rel, err := filepath.Rel(srcDir, srcPath)
	if err != nil {
		return "", err
	}
	out := filepath.Join(dstDir, rel)
	return strings.TrimSuffix(out, filepath.Ext(out)) + ".jpg", nil
}
