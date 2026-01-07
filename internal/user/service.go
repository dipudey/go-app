package user

type Service interface {
	GetAllUsers() ([]Response, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllUsers() ([]Response, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []Response
	for _, u := range users {
		response = append(response, Response{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return response, nil
}
