package subject

type Service struct{}

func NewService() *Subject {
	return &Subject{}
}

func (s *Service) List() []Subject {
	// todo

	return []Subject{}
}

func (s *Service) Get(id int) (*Subject, error) {
	// todo

	return &Subject{}, nil
}

func (s *Service) Add(id int) (*Subject, error) {
	// todo

	return &Subject{}, nil
}

func (s *Service) Remove(id int) (*Subject, error) {
	// todo

	return &Subject{}, nil
}

func (s *Service) Update(id int) (*Subject, error) {
	// todo

	return &Subject{}, nil
}
