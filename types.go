package main

// fbFixScreeninfo reflects the Linux fb_fix_screeninfo structure.
// Note: Several fields are 64-bit (unsigned long) on many systems.
type fbFixScreeninfo struct {
	Id           [16]byte
	SmemStart    uint64 // unsigned long (64-bit)
	SmemLen      uint32
	Type         uint32
	TypeAux      uint32
	Visual       uint32
	XPanStep     uint16
	YPanStep     uint16
	YWrapStep    uint16
	LineLength   uint32
	MmioStart    uint64 // unsigned long (64-bit)
	MmioLen      uint32
	Accel        uint32
	Capabilities uint16
	Reserved     [2]uint16
}

// fbVarScreeninfo reflects a minimal version of the Linux fb_var_screeninfo structure.
type fbVarScreeninfo struct {
	Xres         uint32
	Yres         uint32
	XresVirtual  uint32
	YresVirtual  uint32
	Xoffset      uint32
	Yoffset      uint32
	BitsPerPixel uint32
	Grayscale    uint32
	Red          fbBitfield
	Green        fbBitfield
	Blue         fbBitfield
	Transp       fbBitfield
	Nonstd       uint32
	Activate     uint32
	Height       uint32
	Width        uint32
	AccelFlags   uint32
	Pixclock     uint32
	LeftMargin   uint32
	RightMargin  uint32
	UpperMargin  uint32
	LowerMargin  uint32
	HsyncLen     uint32
	VsyncLen     uint32
	Sync         uint32
	Vmode        uint32
	Rotate       uint32
	Colorspace   uint32
	Reserved     [4]uint32
}

type fbBitfield struct {
	Offset   uint32
	Length   uint32
	MsbRight uint32
}
