package main

// === CONFIGURATION ===
const (
	intervalSeconds = 5          // slideshow interval (in seconds)
	imageDir        = "./images" // directory containing images
)

// ====================

// === SYSTEM CALLS ===

const (
	FBIOGET_FSCREENINFO = 0x4602
	FBIOGET_VSCREENINFO = 0x4600
)

// =====================
