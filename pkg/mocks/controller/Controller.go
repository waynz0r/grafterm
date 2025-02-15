// Code generated by mockery v1.0.0. DO NOT EDIT.

package controller

import context "context"

import mock "github.com/stretchr/testify/mock"
import model "github.com/waynz0r/grafterm/pkg/model"
import time "time"

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// GetRangeMetrics provides a mock function with given fields: ctx, query, start, end, step
func (_m *Controller) GetRangeMetrics(ctx context.Context, query model.Query, start time.Time, end time.Time, step time.Duration) ([]model.MetricSeries, error) {
	ret := _m.Called(ctx, query, start, end, step)

	var r0 []model.MetricSeries
	if rf, ok := ret.Get(0).(func(context.Context, model.Query, time.Time, time.Time, time.Duration) []model.MetricSeries); ok {
		r0 = rf(ctx, query, start, end, step)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.MetricSeries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Query, time.Time, time.Time, time.Duration) error); ok {
		r1 = rf(ctx, query, start, end, step)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingleInstantMetric provides a mock function with given fields: ctx, query
func (_m *Controller) GetSingleInstantMetric(ctx context.Context, query model.Query) (*model.Metric, error) {
	ret := _m.Called(ctx, query)

	var r0 *model.Metric
	if rf, ok := ret.Get(0).(func(context.Context, model.Query) *model.Metric); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Metric)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingleMetric provides a mock function with given fields: ctx, query, t
func (_m *Controller) GetSingleMetric(ctx context.Context, query model.Query, t time.Time) (*model.Metric, error) {
	ret := _m.Called(ctx, query, t)

	var r0 *model.Metric
	if rf, ok := ret.Get(0).(func(context.Context, model.Query, time.Time) *model.Metric); ok {
		r0 = rf(ctx, query, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Metric)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Query, time.Time) error); ok {
		r1 = rf(ctx, query, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
