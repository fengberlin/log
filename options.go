package log

type options struct {
}

type Option interface {
	apply(*options)
}

type optionFunc struct {
	f func(*options)
}

func (fo *optionFunc) apply(opts *options) {
	fo.f(opts)
}

func newOptionFunc(f func(*options)) *optionFunc {
	return &optionFunc{f: f}
}
