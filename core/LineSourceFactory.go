package core

type LineSourceFactory struct {
}

func CreateInstance(LineSourceFactory LineSourceFactory) []ILineSource {
	return make([]FileLindSource, 0)
}
