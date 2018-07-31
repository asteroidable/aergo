/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */
package p2p

import actor "github.com/AsynkronIT/protoactor-go/actor"
import mock "github.com/stretchr/testify/mock"

// MockActorService is an autogenerated mock type for the MockActorService type
type MockActorService struct {
	mock.Mock
}

// CallRequest provides a mock function with given fields: _a0, msg
func (_m *MockActorService) CallRequest(_a0 string, msg interface{}) (interface{}, error) {
	ret := _m.Called(_a0, msg)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}) interface{}); ok {
		r0 = rf(_a0, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(_a0, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FutureRequest provides a mock function with given fields: _a0, msg
func (_m *MockActorService) FutureRequest(_a0 string, msg interface{}) *actor.Future {
	ret := _m.Called(_a0, msg)

	var r0 *actor.Future
	if rf, ok := ret.Get(0).(func(string, interface{}) *actor.Future); ok {
		r0 = rf(_a0, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actor.Future)
		}
	}

	return r0
}

// SendRequest provides a mock function with given fields: _a0, msg
func (_m *MockActorService) SendRequest(_a0 string, msg interface{}) {
	_m.Called(_a0, msg)
}
