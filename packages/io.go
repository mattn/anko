// +build go1.6

package packages

import (
	"io"
	"reflect"
)

func init() {
	if _, ok := Packages["io"]; !ok {
		Packages["io"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["io"]; !ok {
		PackageTypes["io"] = make(map[string]interface{})
	}
	var bytereader io.ByteReader
	var bytescanner io.ByteScanner
	var bytewriter io.ByteWriter
	var closer io.Closer
	var limitedreader io.LimitedReader
	var pipereader io.PipeReader
	var pipewriter io.PipeWriter
	var readcloser io.ReadCloser
	var readseeker io.ReadSeeker
	var readwritecloser io.ReadWriteCloser
	var readwriteseeker io.ReadWriteSeeker
	var readwriter io.ReadWriter
	var reader io.Reader
	var readerat io.ReaderAt
	var readerfrom io.ReaderFrom
	var runereader io.RuneReader
	var runescanner io.RuneScanner
	var sectionreader io.SectionReader
	var seeker io.Seeker
	var writecloser io.WriteCloser
	var writeseeker io.WriteSeeker
	var writer io.Writer
	var writerat io.WriterAt
	var writerto io.WriterTo
	Packages["io"]["Copy"] = io.Copy
	Packages["io"]["CopyBuffer"] = io.CopyBuffer
	Packages["io"]["CopyN"] = io.CopyN
	Packages["io"]["EOF"] = io.EOF
	Packages["io"]["ErrClosedPipe"] = io.ErrClosedPipe
	Packages["io"]["ErrNoProgress"] = io.ErrNoProgress
	Packages["io"]["ErrShortBuffer"] = io.ErrShortBuffer
	Packages["io"]["ErrShortWrite"] = io.ErrShortWrite
	Packages["io"]["ErrUnexpectedEOF"] = io.ErrUnexpectedEOF
	Packages["io"]["LimitReader"] = io.LimitReader
	Packages["io"]["MultiReader"] = io.MultiReader
	Packages["io"]["MultiWriter"] = io.MultiWriter
	Packages["io"]["NewSectionReader"] = io.NewSectionReader
	Packages["io"]["Pipe"] = io.Pipe
	Packages["io"]["ReadAtLeast"] = io.ReadAtLeast
	Packages["io"]["ReadFull"] = io.ReadFull
	Packages["io"]["TeeReader"] = io.TeeReader
	Packages["io"]["WriteString"] = io.WriteString
	PackageTypes["io"]["ByteReader"] = reflect.TypeOf(&bytereader).Elem()
	PackageTypes["io"]["ByteScanner"] = reflect.TypeOf(&bytescanner).Elem()
	PackageTypes["io"]["ByteWriter"] = reflect.TypeOf(&bytewriter).Elem()
	PackageTypes["io"]["Closer"] = reflect.TypeOf(&closer).Elem()
	PackageTypes["io"]["LimitedReader"] = reflect.TypeOf(&limitedreader).Elem()
	PackageTypes["io"]["PipeReader"] = reflect.TypeOf(&pipereader).Elem()
	PackageTypes["io"]["PipeWriter"] = reflect.TypeOf(&pipewriter).Elem()
	PackageTypes["io"]["ReadCloser"] = reflect.TypeOf(&readcloser).Elem()
	PackageTypes["io"]["ReadSeeker"] = reflect.TypeOf(&readseeker).Elem()
	PackageTypes["io"]["ReadWriteCloser"] = reflect.TypeOf(&readwritecloser).Elem()
	PackageTypes["io"]["ReadWriteSeeker"] = reflect.TypeOf(&readwriteseeker).Elem()
	PackageTypes["io"]["ReadWriter"] = reflect.TypeOf(&readwriter).Elem()
	PackageTypes["io"]["Reader"] = reflect.TypeOf(&reader).Elem()
	PackageTypes["io"]["ReaderAt"] = reflect.TypeOf(&readerat).Elem()
	PackageTypes["io"]["ReaderFrom"] = reflect.TypeOf(&readerfrom).Elem()
	PackageTypes["io"]["RuneReader"] = reflect.TypeOf(&runereader).Elem()
	PackageTypes["io"]["RuneScanner"] = reflect.TypeOf(&runescanner).Elem()
	PackageTypes["io"]["SectionReader"] = reflect.TypeOf(&sectionreader).Elem()
	PackageTypes["io"]["Seeker"] = reflect.TypeOf(&seeker).Elem()
	PackageTypes["io"]["WriteCloser"] = reflect.TypeOf(&writecloser).Elem()
	PackageTypes["io"]["WriteSeeker"] = reflect.TypeOf(&writeseeker).Elem()
	PackageTypes["io"]["Writer"] = reflect.TypeOf(&writer).Elem()
	PackageTypes["io"]["WriterAt"] = reflect.TypeOf(&writerat).Elem()
	PackageTypes["io"]["WriterTo"] = reflect.TypeOf(&writerto).Elem()
}
