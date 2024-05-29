package student

type Service struct{}

func NewService() *Student {
	return &Student{}
}

func (s *Service) List() []Student {
	// todo

	return []Student{}
}

func (s *Service) Get(id int) (*Student, error) {
	// todo

	return &Student{}, nil
}

func (s *Service) Add(id int) (*Student, error) {
	// todo

	return &Student{}, nil
}

func (s *Service) Remove(id int) (*Student, error) {
	// todo

	return &Student{}, nil
}

func (s *Service) Update(id int) (*Student, error) {
	// todo

	return &Student{}, nil
}
