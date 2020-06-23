package serviceboot

type StartPriority uint8

const (
	StartPriorityMin StartPriority = 0 // 立即运行
	// 下面的交由Start()运行

	StartPriorityMax StartPriority = 9
)

var manager struct {
	services [10][]*Service
}

func AddService(s *Service) error {
	for _, service := range manager.services[s.startPriority] {
		if s.Type() == service.Type() {
			// 服务已启动
			return nil
		}
	}

	manager.services[s.startPriority] = append(manager.services[s.startPriority], s)
	if s.startPriority == StartPriorityMin {
		if err := s.Start(); err != nil {
			//log.Printf("[%s] start failed. error: %v", s.Type(), err)
			return err
		}
		//log.Printf("[%s] start.", s.Type())
	}
	return nil
}

func Start() error {
	for i := StartPriorityMin + 1; i <= StartPriorityMax; i++ {
		for _, s := range manager.services[i] {
			if err := s.Start(); err != nil {
				//log.Printf("[%s] start failed. error: %v", s.Type(), err)
				return err
			}
			//log.Printf("[%s] start.", s.Type())
		}
	}
	return nil
}

func Stop() {
	for i := StartPriorityMax; i <= StartPriorityMax; i-- {
		for _, s := range manager.services[i] {
			if err := s.Stop(); err != nil {
				//log.Printf("[%s] stop failed. error: %v", s.Type(), err)
			}
			//log.Printf("[%s] stop.", s.Type())
		}
	}
}

func GetService(typ string) *Service {
	for _, services := range manager.services {
		for _, service := range services {
			if service.typ == typ {
				return service
			}
		}
	}
	return nil
}
