package serviceboot

type Option func(s *Service)

type Service struct {
	service       interface{}
	initFn        func() (interface{}, error)
	beforeStartFn func(service interface{}) error                // service可能为nil
	startFn       func(service interface{}) (interface{}, error) // service可能为nil
	stopFn        func(service interface{}) error
	typ           string // 服务类型
	running       bool
	startPriority StartPriority
}

func NewService(typ string, priority StartPriority, opts ...Option) *Service {
	s := new(Service)
	s.typ = typ
	s.startPriority = priority
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Service) WithInitFn(fn func() (interface{}, error)) {
	s.initFn = fn
}

func (s *Service) WithBeforeStartFn(fn func(service interface{}) error) {
	s.beforeStartFn = fn
}

func (s *Service) WithStartFn(fn func(service interface{}) (interface{}, error)) {
	s.startFn = fn
}

func (s *Service) WithStopFn(fn func(service interface{}) error) {
	s.stopFn = fn
}

func (s *Service) init() error {
	if s.initFn == nil {
		return nil
	}
	service, err := s.initFn()
	if err != nil {
		return err
	}
	s.service = service
	return nil
}

func (s *Service) beforeStart() error {
	if s.beforeStartFn == nil {
		return nil
	}
	if err := s.beforeStartFn(s.service); err != nil {
		return err
	}
	return nil
}

func (s *Service) Start() error {
	if s.running {
		return nil
	}

	if err := s.init(); err != nil {
		return err
	}

	if err := s.beforeStart(); err != nil {
		return err
	}

	if s.startFn != nil {
		// w.service可能为空，如果服务不需要初始化
		service, err := s.startFn(s.service)
		if err != nil {
			return err
		}
		s.service = service
	}

	s.running = true
	return nil
}

func (s *Service) Stop() error {
	if !s.running || s.stopFn == nil {
		s.running = false
		return nil
	}
	if err := s.stopFn(s.service); err != nil {
		return err
	}

	s.running = false
	return nil
}

func (s *Service) Type() string {
	return s.typ
}
