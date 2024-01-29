// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	models "pokt_gateway_server/pkg/pokt/pokt_v0/models"

	mock "github.com/stretchr/testify/mock"
)

// PocketService is an autogenerated mock type for the PocketService type
type PocketService struct {
	mock.Mock
}

type PocketService_Expecter struct {
	mock *mock.Mock
}

func (_m *PocketService) EXPECT() *PocketService_Expecter {
	return &PocketService_Expecter{mock: &_m.Mock}
}

// GetLatestBlockHeight provides a mock function with given fields:
func (_m *PocketService) GetLatestBlockHeight() (*models.GetLatestBlockHeightResponse, error) {
	ret := _m.Called()

	var r0 *models.GetLatestBlockHeightResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*models.GetLatestBlockHeightResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *models.GetLatestBlockHeightResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GetLatestBlockHeightResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PocketService_GetLatestBlockHeight_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestBlockHeight'
type PocketService_GetLatestBlockHeight_Call struct {
	*mock.Call
}

// GetLatestBlockHeight is a helper method to define mock.On call
func (_e *PocketService_Expecter) GetLatestBlockHeight() *PocketService_GetLatestBlockHeight_Call {
	return &PocketService_GetLatestBlockHeight_Call{Call: _e.mock.On("GetLatestBlockHeight")}
}

func (_c *PocketService_GetLatestBlockHeight_Call) Run(run func()) *PocketService_GetLatestBlockHeight_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PocketService_GetLatestBlockHeight_Call) Return(_a0 *models.GetLatestBlockHeightResponse, _a1 error) *PocketService_GetLatestBlockHeight_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PocketService_GetLatestBlockHeight_Call) RunAndReturn(run func() (*models.GetLatestBlockHeightResponse, error)) *PocketService_GetLatestBlockHeight_Call {
	_c.Call.Return(run)
	return _c
}

// GetSession provides a mock function with given fields: req
func (_m *PocketService) GetSession(req *models.GetSessionRequest) (*models.GetSessionResponse, error) {
	ret := _m.Called(req)

	var r0 *models.GetSessionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.GetSessionRequest) (*models.GetSessionResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*models.GetSessionRequest) *models.GetSessionResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GetSessionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.GetSessionRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PocketService_GetSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSession'
type PocketService_GetSession_Call struct {
	*mock.Call
}

// GetSession is a helper method to define mock.On call
//   - req *models.GetSessionRequest
func (_e *PocketService_Expecter) GetSession(req interface{}) *PocketService_GetSession_Call {
	return &PocketService_GetSession_Call{Call: _e.mock.On("GetSession", req)}
}

func (_c *PocketService_GetSession_Call) Run(run func(req *models.GetSessionRequest)) *PocketService_GetSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.GetSessionRequest))
	})
	return _c
}

func (_c *PocketService_GetSession_Call) Return(_a0 *models.GetSessionResponse, _a1 error) *PocketService_GetSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PocketService_GetSession_Call) RunAndReturn(run func(*models.GetSessionRequest) (*models.GetSessionResponse, error)) *PocketService_GetSession_Call {
	_c.Call.Return(run)
	return _c
}

// SendRelay provides a mock function with given fields: req
func (_m *PocketService) SendRelay(req *models.SendRelayRequest) (*models.SendRelayResponse, error) {
	ret := _m.Called(req)

	var r0 *models.SendRelayResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.SendRelayRequest) (*models.SendRelayResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*models.SendRelayRequest) *models.SendRelayResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SendRelayResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.SendRelayRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PocketService_SendRelay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendRelay'
type PocketService_SendRelay_Call struct {
	*mock.Call
}

// SendRelay is a helper method to define mock.On call
//   - req *models.SendRelayRequest
func (_e *PocketService_Expecter) SendRelay(req interface{}) *PocketService_SendRelay_Call {
	return &PocketService_SendRelay_Call{Call: _e.mock.On("SendRelay", req)}
}

func (_c *PocketService_SendRelay_Call) Run(run func(req *models.SendRelayRequest)) *PocketService_SendRelay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.SendRelayRequest))
	})
	return _c
}

func (_c *PocketService_SendRelay_Call) Return(_a0 *models.SendRelayResponse, _a1 error) *PocketService_SendRelay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PocketService_SendRelay_Call) RunAndReturn(run func(*models.SendRelayRequest) (*models.SendRelayResponse, error)) *PocketService_SendRelay_Call {
	_c.Call.Return(run)
	return _c
}

// NewPocketService creates a new instance of PocketService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPocketService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PocketService {
	mock := &PocketService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
