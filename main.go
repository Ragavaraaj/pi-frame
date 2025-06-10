package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"

	imagedraw "golang.org/x/image/draw"
)

func main() {

	// initialize environment variables and defaults
	initEnv()

	// Open the framebuffer device.
	fb, err := os.OpenFile("/dev/fb0", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf("failed to open framebuffer: %v", err)
	}
	defer fb.Close()

	// Query fixed screen info.
	mappingSize, width, height, err := sysCalls(fb)

	if err != nil {
		log.Fatalf("Failed to get framebuffer info: %v", err)
	}

	fmt.Printf("Framebuffer memory length: %d bytes\n", mappingSize)

	// Memory-map the framebuffer.
	fbData, err := syscall.Mmap(int(fb.Fd()), 0, mappingSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("failed to mmap: %v", err)
	}
	defer syscall.Munmap(fbData)

	// Discover image files in the directory.
	images, err := readImageFiles(imageDir)
	if err != nil {
		log.Fatalf("error reading image directory: %v", err)
	}
	if len(images) == 0 {
		log.Fatalf("no images found in directory %s", imageDir)
	}

	index := 0
	for {
		imgPath := images[index]
		fmt.Printf("Displaying image: %s\n", imgPath)

		img, err := loadImage(imgPath)
		if err != nil {
			log.Printf("failed to load image %s: %v", imgPath, err)
			index = (index + 1) % len(images)
			continue
		}

		log.Println("screenwidth:", width, "screenheight:", height)

		// Resize to the configured screen resolution.
		resized := image.NewRGBA(image.Rect(0, 0, width, height))
		imagedraw.ApproxBiLinear.Scale(resized, resized.Bounds(), img, img.Bounds(), imagedraw.Over, nil)

		// Convert the image to RGB565 and write it.
		if err := drawImageToFramebuffer(resized, fbData, width, height); err != nil {
			log.Printf("failed to draw image to framebuffer: %v", err)
		}

		time.Sleep(time.Duration(intervalSeconds) * time.Second)
		index = (index + 1) % len(images)
	}
}

// sysCalls makes all the system calls needed to read the framebuffer info.
func sysCalls(fb *os.File) (int, int, int, error) {
	var fix fbFixScreeninfo
	var vinfo fbVarScreeninfo
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fb.Fd(), FBIOGET_FSCREENINFO, uintptr(unsafe.Pointer(&fix))); errno != 0 {
		return 0, 0, 0, fmt.Errorf("ioctl FBIOGET_FSCREENINFO failed: %v", errno)
	}

	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fb.Fd(), FBIOGET_VSCREENINFO, uintptr(unsafe.Pointer(&vinfo))); errno != 0 {
		return 0, 0, 0, fmt.Errorf("ioctl FBIOGET_VSCREENINFO failed: %v", errno)
	}

	return int(fix.SmemLen), int(vinfo.Xres), int(vinfo.Yres), nil
}

// readImageFiles searches the directory for image files.
func readImageFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		lower := strings.ToLower(path)
		if strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".jpeg") || strings.HasSuffix(lower, ".png") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// loadImage opens and decodes an image file.
func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if strings.HasSuffix(strings.ToLower(path), ".png") {
		return png.Decode(f)
	}
	return jpeg.Decode(f)
}

// drawImageToFramebuffer writes an RGBA image to the framebuffer in RGB565 format.
func drawImageToFramebuffer(img *image.RGBA, fbData []byte, width, height int) error {
	// Verify enough framebuffer memory.
	if len(fbData) < width*height*2 {
		return fmt.Errorf("framebuffer size is smaller than expected")
	}

	index := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rr, gg, bb, aa := img.At(x, y).RGBA()
			// Convert from 16-bit [0,65535] down to 8-bit.
			r8 := uint8(rr >> 8)
			g8 := uint8(gg >> 8)
			b8 := uint8(bb >> 8)
			a8 := uint8(aa >> 8)

			// Blend pixel over black if not fully opaque.
			if a8 != 255 {
				r8 = uint8((uint16(r8) * uint16(a8)) / 255)
				g8 = uint8((uint16(g8) * uint16(a8)) / 255)
				b8 = uint8((uint16(b8) * uint16(a8)) / 255)
			}

			// Pack the color into RGB565.
			r5 := (r8 >> 3) & 0x1F
			g6 := (g8 >> 2) & 0x3F
			b5 := (b8 >> 3) & 0x1F
			rgb565 := uint16(r5)<<11 | uint16(g6)<<5 | uint16(b5)

			// Write in little-endian.
			fbData[index] = byte(rgb565 & 0xff)
			fbData[index+1] = byte(rgb565 >> 8)
			index += 2
		}
	}
	return nil
}
