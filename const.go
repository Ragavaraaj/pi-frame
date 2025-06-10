package main

import (
	"os"
	"strconv"
)

// === CONFIGURATION ===
const (
	DEFAULT_INTERVAL = 5          // default interval in seconds
	DEFAULT_DIR      = "./images" // directory containing images
)

var (
	intervalSeconds int    = DEFAULT_INTERVAL // interval in seconds for slideshow
	imageDir        string = DEFAULT_DIR      // directory containing images
)

func initEnv() {
	if value, ok := os.LookupEnv("SLIDESHOW_INTERVAL"); ok {
		if v, err := strconv.Atoi(value); err == nil {
			intervalSeconds = v
		}
	}

	if value, ok := os.LookupEnv("SLIDESHOW_DIR"); ok {
		imageDir = value
	}
}

// ====================

// === SYSTEM CALLS ===

const (
	FBIOGET_FSCREENINFO = 0x4602
	FBIOGET_VSCREENINFO = 0x4600
)

// =====================
