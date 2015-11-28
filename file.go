package dropbox

import (
	"bytes"
	"errors"
	"io"
	"os"
	"time"
)

// FileInfo wraps Dropbox file MetaData to implement os.FileInfo.
type FileInfo struct {
	meta *Metadata
}

// Name of the file.
func (f *FileInfo) Name() string {
	return f.meta.Name
}

// Size of the file.
func (f *FileInfo) Size() int64 {
	return int64(f.meta.Size)
}

// IsDir returns true if the file is a directory.
func (f *FileInfo) IsDir() bool {
	return f.meta.Tag == "folder"
}

// Sys is not implemented.
func (f *FileInfo) Sys() interface{} {
	return nil
}

// ModTime returns the modification time.
func (f *FileInfo) ModTime() time.Time {
	return f.meta.ServerModified
}

// Mode returns the file mode flags.
func (f *FileInfo) Mode() os.FileMode {
	var m os.FileMode

	if f.IsDir() {
		m |= os.ModeDir
	}

	return m
}

type File struct {
	files *Files
	path  string
	buf   bytes.Buffer

	io.ReadCloser
}

// Readdir implementation, see os package for details.
func (f *File) Readdir(n int) (ents []os.FileInfo, err error) {
	var cursor string

	if n <= 0 {
		n = -1
	}

	for {
		var out *ListFolderOutput

		if cursor == "" {
			out, err = f.files.ListFolder(&ListFolderInput{Path: f.path})
			cursor = out.Cursor
		} else {
			out, err = f.files.ListFolderContinue(&ListFolderContinueInput{cursor})
			cursor = out.Cursor
		}

		if err != nil {
			return
		}

		for _, ent := range out.Entries {
			ents = append(ents, &FileInfo{ent})
		}

		if n >= 0 && len(ents) >= n {
			ents = ents[:n]
			break
		}

		if !out.HasMore {
			break
		}
	}

	if n >= 0 && len(ents) == 0 {
		err = io.EOF
		return
	}

	return
}

// Stat implementation, see os package for details.
func (f *File) Stat() (os.FileInfo, error) {
	out, err := f.files.GetMetadata(&GetMetadataInput{
		Path: f.path,
	})

	if err != nil {
		return nil, err
	}

	return &FileInfo{&out.Metadata}, nil
}

// Seek is not supported.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return 0, errors.New("seek is not supported")
}

// Write the given blog.
func (f *File) Write(b []byte) (int, error) {
	return f.buf.Write(b)
}

// WriteString writes the given string.
func (f *File) WriteString(s string) (int, error) {
	return f.buf.WriteString(s)
}

// Close uploads the buffered contents.
func (f *File) Close() error {
	_, err := f.files.Upload(&UploadInput{
		Path:   f.path,
		Reader: &f.buf,
		Mode:   WriteModeOverwrite,
		Mute:   true,
	})

	return err
}
