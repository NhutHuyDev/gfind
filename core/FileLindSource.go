package core

type FileLindSource struct {
}

func (fileLindSource FileLindSource) ReadLine() (Line, error) {
	return Line{}, nil
}

func (fileLindSource FileLindSource) Open() {
}

func (fileLindSource FileLindSource) Close() {
}
