// Code generated by mockery v1.0.0
package mocks

import dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"

import mock "github.com/stretchr/testify/mock"

// DynamoInterface is an autogenerated mock type for the DynamoInterface type
type DynamoInterface struct {
	mock.Mock
}

// PutItem provides a mock function with given fields: _a0
func (_m *DynamoInterface) PutItem(_a0 *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	ret := _m.Called(_a0)

	var r0 *dynamodb.PutItemOutput
	if rf, ok := ret.Get(0).(func(*dynamodb.PutItemInput) *dynamodb.PutItemOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.PutItemOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dynamodb.PutItemInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: _a0
func (_m *DynamoInterface) Query(_a0 *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *dynamodb.QueryOutput
	if rf, ok := ret.Get(0).(func(*dynamodb.QueryInput) *dynamodb.QueryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.QueryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dynamodb.QueryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}