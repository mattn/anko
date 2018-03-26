// +build go1.6

package packages

import (
	"os"
	"reflect"
)

func init() {
	if _, ok := Packages["os"]; !ok {
		Packages["os"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["os"]; !ok {
		PackageTypes["os"] = make(map[string]interface{})
	}
	var file os.File
	var fileinfo os.FileInfo
	var filemode os.FileMode
	var linkerror os.LinkError
	var patherror os.PathError
	var procattr os.ProcAttr
	var process os.Process
	var processstate os.ProcessState
	var signal os.Signal
	var syscallerror os.SyscallError
	Packages["os"]["Args"] = os.Args
	Packages["os"]["Chdir"] = os.Chdir
	Packages["os"]["Chmod"] = os.Chmod
	Packages["os"]["Chown"] = os.Chown
	Packages["os"]["Chtimes"] = os.Chtimes
	Packages["os"]["Clearenv"] = os.Clearenv
	Packages["os"]["Create"] = os.Create
	Packages["os"]["DevNull"] = os.DevNull
	Packages["os"]["Environ"] = os.Environ
	Packages["os"]["ErrExist"] = os.ErrExist
	Packages["os"]["ErrInvalid"] = os.ErrInvalid
	Packages["os"]["ErrNotExist"] = os.ErrNotExist
	Packages["os"]["ErrPermission"] = os.ErrPermission
	Packages["os"]["Exit"] = os.Exit
	Packages["os"]["Expand"] = os.Expand
	Packages["os"]["ExpandEnv"] = os.ExpandEnv
	Packages["os"]["FindProcess"] = os.FindProcess
	Packages["os"]["Getegid"] = os.Getegid
	Packages["os"]["Getenv"] = os.Getenv
	Packages["os"]["Geteuid"] = os.Geteuid
	Packages["os"]["Getgid"] = os.Getgid
	Packages["os"]["Getgroups"] = os.Getgroups
	Packages["os"]["Getpagesize"] = os.Getpagesize
	Packages["os"]["Getpid"] = os.Getpid
	Packages["os"]["Getppid"] = os.Getppid
	Packages["os"]["Getuid"] = os.Getuid
	Packages["os"]["Getwd"] = os.Getwd
	Packages["os"]["Hostname"] = os.Hostname
	Packages["os"]["Interrupt"] = os.Interrupt
	Packages["os"]["IsExist"] = os.IsExist
	Packages["os"]["IsNotExist"] = os.IsNotExist
	Packages["os"]["IsPathSeparator"] = os.IsPathSeparator
	Packages["os"]["IsPermission"] = os.IsPermission
	Packages["os"]["Kill"] = os.Kill
	Packages["os"]["Lchown"] = os.Lchown
	Packages["os"]["Link"] = os.Link
	Packages["os"]["LookupEnv"] = os.LookupEnv
	Packages["os"]["Lstat"] = os.Lstat
	Packages["os"]["Mkdir"] = os.Mkdir
	Packages["os"]["MkdirAll"] = os.MkdirAll
	Packages["os"]["ModeAppend"] = os.ModeAppend
	Packages["os"]["ModeCharDevice"] = os.ModeCharDevice
	Packages["os"]["ModeDevice"] = os.ModeDevice
	Packages["os"]["ModeDir"] = os.ModeDir
	Packages["os"]["ModeExclusive"] = os.ModeExclusive
	Packages["os"]["ModeNamedPipe"] = os.ModeNamedPipe
	Packages["os"]["ModePerm"] = os.ModePerm
	Packages["os"]["ModeSetgid"] = os.ModeSetgid
	Packages["os"]["ModeSetuid"] = os.ModeSetuid
	Packages["os"]["ModeSocket"] = os.ModeSocket
	Packages["os"]["ModeSticky"] = os.ModeSticky
	Packages["os"]["ModeSymlink"] = os.ModeSymlink
	Packages["os"]["ModeTemporary"] = os.ModeTemporary
	Packages["os"]["ModeType"] = os.ModeType
	Packages["os"]["NewFile"] = os.NewFile
	Packages["os"]["NewSyscallError"] = os.NewSyscallError
	Packages["os"]["O_APPEND"] = os.O_APPEND
	Packages["os"]["O_CREATE"] = os.O_CREATE
	Packages["os"]["O_EXCL"] = os.O_EXCL
	Packages["os"]["O_RDONLY"] = os.O_RDONLY
	Packages["os"]["O_RDWR"] = os.O_RDWR
	Packages["os"]["O_SYNC"] = os.O_SYNC
	Packages["os"]["O_TRUNC"] = os.O_TRUNC
	Packages["os"]["O_WRONLY"] = os.O_WRONLY
	Packages["os"]["Open"] = os.Open
	Packages["os"]["OpenFile"] = os.OpenFile
	Packages["os"]["PathListSeparator"] = os.PathListSeparator
	Packages["os"]["PathSeparator"] = os.PathSeparator
	Packages["os"]["Pipe"] = os.Pipe
	Packages["os"]["Readlink"] = os.Readlink
	Packages["os"]["Remove"] = os.Remove
	Packages["os"]["RemoveAll"] = os.RemoveAll
	Packages["os"]["Rename"] = os.Rename
	Packages["os"]["SEEK_CUR"] = os.SEEK_CUR
	Packages["os"]["SEEK_END"] = os.SEEK_END
	Packages["os"]["SEEK_SET"] = os.SEEK_SET
	Packages["os"]["SameFile"] = os.SameFile
	Packages["os"]["Setenv"] = os.Setenv
	Packages["os"]["StartProcess"] = os.StartProcess
	Packages["os"]["Stat"] = os.Stat
	Packages["os"]["Stderr"] = os.Stderr
	Packages["os"]["Stdin"] = os.Stdin
	Packages["os"]["Stdout"] = os.Stdout
	Packages["os"]["Symlink"] = os.Symlink
	Packages["os"]["TempDir"] = os.TempDir
	Packages["os"]["Truncate"] = os.Truncate
	Packages["os"]["Unsetenv"] = os.Unsetenv
	PackageTypes["os"]["File"] = reflect.TypeOf(&file).Elem()
	PackageTypes["os"]["FileInfo"] = reflect.TypeOf(&fileinfo).Elem()
	PackageTypes["os"]["FileMode"] = reflect.TypeOf(&filemode).Elem()
	PackageTypes["os"]["LinkError"] = reflect.TypeOf(&linkerror).Elem()
	PackageTypes["os"]["PathError"] = reflect.TypeOf(&patherror).Elem()
	PackageTypes["os"]["ProcAttr"] = reflect.TypeOf(&procattr).Elem()
	PackageTypes["os"]["Process"] = reflect.TypeOf(&process).Elem()
	PackageTypes["os"]["ProcessState"] = reflect.TypeOf(&processstate).Elem()
	PackageTypes["os"]["Signal"] = reflect.TypeOf(&signal).Elem()
	PackageTypes["os"]["SyscallError"] = reflect.TypeOf(&syscallerror).Elem()
}
