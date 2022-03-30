package auth

import "cv_app/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
