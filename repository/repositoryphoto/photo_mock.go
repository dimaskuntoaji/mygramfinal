package repositoryphoto

import (
	entity "finall/entity"
	mock "github.com/stretchr/testify/mock"
)

type RepositoryPhotoMock struct {
	mock.Mock
}

func (_m *RepositoryPhotoMock) Create(data entity.Photo) (entity.Photo, error) {
	ret := _m.Called(data)

	var r0 entity.Photo
	if rf, ok := ret.Get(0).(func(entity.Photo) entity.Photo); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.Photo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Photo) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RepositoryPhotoMock) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *RepositoryPhotoMock) GetPhotoByUserID(id uint) (entity.Photo, error) {
	ret := _m.Called(id)

	var r0 entity.Photo
	if rf, ok := ret.Get(0).(func(uint) entity.Photo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Photo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RepositoryPhotoMock) GetPhotos() ([]entity.Photo, error) {
	ret := _m.Called()

	var r0 []entity.Photo
	if rf, ok := ret.Get(0).(func() []entity.Photo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Photo)
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

func (_m *RepositoryPhotoMock) Update(data entity.Photo) (entity.Photo, error) {
	ret := _m.Called(data)

	var r0 entity.Photo
	if rf, ok := ret.Get(0).(func(entity.Photo) entity.Photo); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.Photo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Photo) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}