package json

import "github.com/MAZEN-Kenjrawi/pwd/internal/model"

type ProfileRepository struct {
	filePath string
}

func (r *ProfileRepository) Save(p *model.Profile) error {
	return nil
}

func (r *ProfileRepository) GetProfileByUsername(username string) (*model.Profile, error) {
	return nil, nil
}

func NewProfileRepository(filePath string) model.ProfileRepository {
	return &ProfileRepository{filePath}
}
