// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/CodFrm/cxmooc-tools/server/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// TopicRepository is an autogenerated mock type for the TopicRepository type
type TopicRepository struct {
	mock.Mock
}

func (_m *TopicRepository) FindByHash(hash string) (*entity.TopicEntity, error) {
	ret := _m.Called(hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*entity.TopicEntity) *entity.TopicEntity); ok {
		r0 = rf(topicEntity)
	} else {
		r0 = ret.Get(0).(*entity.TopicEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.TopicEntity) error); ok {
		r1 = rf(topicEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exist provides a mock function with given fields: topicEntity
func (_m *TopicRepository) Exist(topicEntity *entity.TopicEntity) (bool, error) {
	ret := _m.Called(topicEntity.GetTopic())

	var r0 bool
	if rf, ok := ret.Get(0).(func(*entity.TopicEntity) bool); ok {
		r0 = rf(topicEntity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.TopicEntity) error); ok {
		r1 = rf(topicEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: topicEntity
func (_m *TopicRepository) Save(topicEntity *entity.TopicEntity) error {
	ret := _m.Called(topicEntity.GetTopic())

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.TopicEntity) error); ok {
		r0 = rf(topicEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchTopic provides a mock function with given fields: topic
func (_m *TopicRepository) SearchTopic(topic *entity.TopicEntity) ([]*entity.TopicEntity, error) {
	ret := _m.Called(topic.GetTopic())

	var r0 []*entity.TopicEntity
	if rf, ok := ret.Get(0).(func(*entity.TopicEntity) []*entity.TopicEntity); ok {
		r0 = rf(topic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.TopicEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.TopicEntity) error); ok {
		r1 = rf(topic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
