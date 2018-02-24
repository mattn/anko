package packages

import (
	"io"
)

func init() {
	Packages["io"] = map[string]interface{}{
		"Copy":             io.Copy,
		"CopyN":            io.CopyN,
		"EOF":              io.EOF,
		"ErrClosedPipe":    io.ErrClosedPipe,
		"ErrNoProgress":    io.ErrNoProgress,
		"ErrShortBuffer":   io.ErrShortBuffer,
		"ErrShortWrite":    io.ErrShortWrite,
		"ErrUnexpectedEOF": io.ErrUnexpectedEOF,
		"LimitReader":      io.LimitReader,
		"MultiReader":      io.MultiReader,
		"MultiWriter":      io.MultiWriter,
		"NewSectionReader": io.NewSectionReader,
		"Pipe":             io.Pipe,
		"ReadAtLeast":      io.ReadAtLeast,
		"ReadFull":         io.ReadFull,
		"TeeReader":        io.TeeReader,
		"WriteString":      io.WriteString,
	}
}
