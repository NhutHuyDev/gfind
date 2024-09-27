package core

type ILineSource interface {
	ReadLine() (Line, error)
	Open()
	Close()
}
