package core

type ConsoleFindSource struct {
}

func (consoleFileSource ConsoleFindSource) ReadLine() (Line, error) {
	return Line{}, nil
}

func (consoleFileSource ConsoleFindSource) Open() {
}

func (consoleFileSource ConsoleFindSource) Close() {

}
