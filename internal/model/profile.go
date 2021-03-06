package model

import "errors"

var (
	ErrorProfileDoesNotExist         = errors.New("profile does not exist")
	ErrorProfileAlreadyExists        = errors.New("profile already exists")
	ErrorLoginAlreadyExistsInProfile = errors.New("login already exists in profile")
	ErrorLoginDoesNotExistInProfile  = errors.New("login does not exist in profile")
)

type ProfileRepository interface {
	Save(*Profile) error
	GetProfileByUsername(string) (*Profile, error)
}

type Profile struct {
	Username string
	secret   hash
	logins   loginList
}

func (p *Profile) AddLogin(username, domain, password string) error {
	if p.HasLogin(username, domain) {
		return ErrorLoginAlreadyExistsInProfile
	}

	newLogin, err := newLogin(username, domain, password)
	if err != nil {
		return err
	}

	p.logins = append(p.logins, newLogin)

	return nil
}

func (p *Profile) RemoveLogin(username, domain string) error {
	k := p.findLogin(username, domain)
	if k == -1 {
		return ErrorLoginDoesNotExistInProfile
	}

	p.logins = append(p.logins[:k], p.logins[k+1:]...)

	return nil
}

func (p *Profile) HasLogin(username, domain string) bool {
	return p.findLogin(username, domain) > -1
}

func (p *Profile) GetLogins() loginList {
	return p.logins
}

func (p *Profile) findLogin(username, domain string) int {
	for k, l := range p.logins {
		if l.Username == username && l.Domain == domain {
			return k
		}
	}

	return -1
}

func NewProfile(username, secret string) (*Profile, error) {
	hash, err := hashString(secret)
	if err != nil {
		return nil, err
	}

	return &Profile{
		username,
		hash,
		nil,
	}, nil
}
