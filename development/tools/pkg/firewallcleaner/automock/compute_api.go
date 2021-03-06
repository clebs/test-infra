// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	compute "google.golang.org/api/compute/v1"
	container "google.golang.org/api/container/v1"

	mock "github.com/stretchr/testify/mock"
)

// ComputeAPI is an autogenerated mock type for the ComputeAPI type
type ComputeAPI struct {
	mock.Mock
}

// DeleteFirewallRule provides a mock function with given fields: project, firewall
func (_m *ComputeAPI) DeleteFirewallRule(project string, firewall string) {
	_m.Called(project, firewall)
}

// LookupClusters provides a mock function with given fields: project
func (_m *ComputeAPI) LookupClusters(project string) ([]*container.Cluster, error) {
	ret := _m.Called(project)

	var r0 []*container.Cluster
	if rf, ok := ret.Get(0).(func(string) []*container.Cluster); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*container.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupFirewallRule provides a mock function with given fields: project
func (_m *ComputeAPI) LookupFirewallRule(project string) ([]*compute.Firewall, error) {
	ret := _m.Called(project)

	var r0 []*compute.Firewall
	if rf, ok := ret.Get(0).(func(string) []*compute.Firewall); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*compute.Firewall)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupInstances provides a mock function with given fields: project
func (_m *ComputeAPI) LookupInstances(project string) ([]*compute.Instance, error) {
	ret := _m.Called(project)

	var r0 []*compute.Instance
	if rf, ok := ret.Get(0).(func(string) []*compute.Instance); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*compute.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupNodePools provides a mock function with given fields: clusters
func (_m *ComputeAPI) LookupNodePools(clusters []*container.Cluster) ([]*container.NodePool, error) {
	ret := _m.Called(clusters)

	var r0 []*container.NodePool
	if rf, ok := ret.Get(0).(func([]*container.Cluster) []*container.NodePool); ok {
		r0 = rf(clusters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*container.NodePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*container.Cluster) error); ok {
		r1 = rf(clusters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
