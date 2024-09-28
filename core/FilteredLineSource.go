package core

type FilteredLineSource struct {
	wrappee ILineSource
	filter  func(Line) bool
}

func (filteredLineSource *FilteredLineSource) GetSourceName() string {
	return filteredLineSource.wrappee.GetSourceName()
}

func (filteredLineSource *FilteredLineSource) ReadLine() (Line, error) {
	line, err := filteredLineSource.wrappee.ReadLine()
	if err != nil {
		return Line{}, err
	}

	for {
		if filteredLineSource.filter(line) {
			return line, nil
		} else {
			line, err = filteredLineSource.wrappee.ReadLine()
			if err != nil {
				return Line{}, err
			}
		}
	}
}

func (filteredLineSource *FilteredLineSource) Open() error {
	return filteredLineSource.wrappee.Open()
}

func (filteredLineSource *FilteredLineSource) Close() error {
	return filteredLineSource.wrappee.Close()
}

func NewFilteredLineSource(source ILineSource, f func(Line) bool) *FilteredLineSource {
	return &FilteredLineSource{wrappee: source, filter: f}
}
