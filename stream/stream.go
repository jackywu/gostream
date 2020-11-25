package stream

type StringStream struct {
	data []string
}

func NewStringStream(streamData []string) *StringStream {
	s := new(StringStream)
	s.data = streamData
	return s
}

func (s *StringStream) Filter(f func(string)bool) *StringStream {
	result := make([]string, 0, len(s.data))
	for _, value := range s.data {
		if f(value) {
			result = append(result, value)
		}
	}
	s.data = result
	return s
}

func (s *StringStream) Map(f func(string)string ) *StringStream {
	result := make([]string, 0, len(s.data))
	for _, value := range s.data {
		result = append(result, f(value))
	}
	s.data = result
	return s
}

func (s *StringStream) Collect() []string {
	return s.data
}

