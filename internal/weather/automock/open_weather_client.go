// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"
import weather "github.com/aszecowka/QWRhbSBTemVjb3drYSByZWNydWl0bWVudCB0YXNr/internal/weather"

// OpenWeatherClient is an autogenerated mock type for the OpenWeatherClient type
type OpenWeatherClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, city
func (_m *OpenWeatherClient) Get(ctx context.Context, city string) (weather.OpenWeatherResponse, error) {
	ret := _m.Called(ctx, city)

	var r0 weather.OpenWeatherResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) weather.OpenWeatherResponse); ok {
		r0 = rf(ctx, city)
	} else {
		r0 = ret.Get(0).(weather.OpenWeatherResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
