package mark

type Service struct{}

func NewService() *Mark {
	return &Mark{}
}

func (s *Service) List() []Mark {
	// todo

	return []Mark{}
}

func (s *Service) Get(id int) (*Mark, error) {
	// todo

	return &Mark{}, nil
}

func (s *Service) Add(id int) (*Mark, error) {
	// todo

	return &Mark{}, nil
}

func (s *Service) Remove(id int) (*Mark, error) {
	// todo

	return &Mark{}, nil
}

func (s *Service) Update(id int) (*Mark, error) {
	// todo

	return &Mark{}, nil
}
