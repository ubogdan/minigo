package os

import "syscall"

const O_RDONLY = 0

var Stdout *File = &File{
	id: 1,
}

var Stderr *File = &File{
	id: 2,
}

// File represents an open file descriptor.
type File struct {
	id int
}

func openFileNolog(name string, flag int, perm int) (*File, error) {
	fd, err := syscall.Open(name, flag, perm)
	f := &File{
		id:fd,
	}

	return f, err
}

func OpenFile(name string, flag int, perm int) (*File, error) {
	return openFileNolog(name, flag, perm)
}

func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}

func (f *File) write(b []byte) (int, error) {
	var fid int = f.id
	var n int
	var addr *byte = &b[0]
	n = write(fid, addr, len(b))
	return n,nil
}

// Write writes len(b) bytes to the File.
func (f *File) Write(b []byte) (int, error) {
	n, err := f.write(b)
	return n, err
}

func Exit(i int) {
}

var Args []string

func init() {
	Args = runtime_args()
}

