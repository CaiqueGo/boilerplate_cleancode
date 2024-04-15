package config

import usercreate "github.com/boilerplate_cleancode/user_create"

func (ref Config) UserCreateService() *usercreate.ServiceImpl {
	return usercreate.NewService(ref.UserCreateRepository())
}

func (ref Config) UserCreateRepository() *usercreate.RepositoryImpl {
	return usercreate.NewRepository(ref.Database)
}

func (ref Config) UserCreateHandlerHttp() *usercreate.HandlerHttp {
	return usercreate.NewHandlerHttp(ref.UserCreateService())
}
