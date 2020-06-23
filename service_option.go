package serviceboot

func WithInitFn(fn func() (interface{}, error)) Option {
	return func(s *Service) {
		s.initFn = fn
	}
}

func WithBeforeStartFn(fn func(service interface{}) error) Option {
	return func(s *Service) {
		s.beforeStartFn = fn
	}
}

func WithStartFn(fn func(service interface{}) (interface{}, error)) Option {
	return func(s *Service) {
		s.startFn = fn
	}
}

func WithStopFn(fn func(service interface{}) error) Option {
	return func(s *Service) {
		s.stopFn = fn
	}
}
