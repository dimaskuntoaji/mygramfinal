package repositorysocialmedia

import (
	entity "finall/entity"
	mock "github.com/stretchr/testify/mock"
)

type RepositorySocialMediaMock struct {
	mock.Mock
}

func (_m *RepositorySocialMediaMock) Create(data entity.SocialMedia) (entity.SocialMedia, error) {
	ret := _m.Called(data)

	var r0 entity.SocialMedia
	if rf, ok := ret.Get(0).(func(entity.SocialMedia) entity.SocialMedia); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.SocialMedia)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.SocialMedia) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RepositorySocialMediaMock) DeleteByID(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *RepositorySocialMediaMock) GetList() ([]entity.SocialMedia, error) {
	ret := _m.Called()

	var r0 []entity.SocialMedia
	if rf, ok := ret.Get(0).(func() []entity.SocialMedia); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.SocialMedia)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RepositorySocialMediaMock) UpdateByID(data entity.SocialMedia) (entity.SocialMedia, error) {
	ret := _m.Called(data)

	var r0 entity.SocialMedia
	if rf, ok := ret.Get(0).(func(entity.SocialMedia) entity.SocialMedia); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.SocialMedia)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.SocialMedia) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}