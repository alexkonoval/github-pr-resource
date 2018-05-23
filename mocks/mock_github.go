// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/itsdalmo/github-pr-resource (interfaces: Github)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	github_pr_resource "github.com/itsdalmo/github-pr-resource"
	reflect "reflect"
)

// MockGithub is a mock of Github interface
type MockGithub struct {
	ctrl     *gomock.Controller
	recorder *MockGithubMockRecorder
}

// MockGithubMockRecorder is the mock recorder for MockGithub
type MockGithubMockRecorder struct {
	mock *MockGithub
}

// NewMockGithub creates a new mock instance
func NewMockGithub(ctrl *gomock.Controller) *MockGithub {
	mock := &MockGithub{ctrl: ctrl}
	mock.recorder = &MockGithubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGithub) EXPECT() *MockGithubMockRecorder {
	return m.recorder
}

// GetCommitByID mocks base method
func (m *MockGithub) GetCommitByID(arg0 string) (*github_pr_resource.CommitObject, error) {
	ret := m.ctrl.Call(m, "GetCommitByID", arg0)
	ret0, _ := ret[0].(*github_pr_resource.CommitObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitByID indicates an expected call of GetCommitByID
func (mr *MockGithubMockRecorder) GetCommitByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitByID", reflect.TypeOf((*MockGithub)(nil).GetCommitByID), arg0)
}

// GetPullRequestByID mocks base method
func (m *MockGithub) GetPullRequestByID(arg0 string) (*github_pr_resource.PullRequestObject, error) {
	ret := m.ctrl.Call(m, "GetPullRequestByID", arg0)
	ret0, _ := ret[0].(*github_pr_resource.PullRequestObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequestByID indicates an expected call of GetPullRequestByID
func (mr *MockGithubMockRecorder) GetPullRequestByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequestByID", reflect.TypeOf((*MockGithub)(nil).GetPullRequestByID), arg0)
}

// ListModifiedFiles mocks base method
func (m *MockGithub) ListModifiedFiles(arg0 int) ([]string, error) {
	ret := m.ctrl.Call(m, "ListModifiedFiles", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModifiedFiles indicates an expected call of ListModifiedFiles
func (mr *MockGithubMockRecorder) ListModifiedFiles(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModifiedFiles", reflect.TypeOf((*MockGithub)(nil).ListModifiedFiles), arg0)
}

// ListOpenPullRequests mocks base method
func (m *MockGithub) ListOpenPullRequests() ([]*github_pr_resource.PullRequest, error) {
	ret := m.ctrl.Call(m, "ListOpenPullRequests")
	ret0, _ := ret[0].([]*github_pr_resource.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenPullRequests indicates an expected call of ListOpenPullRequests
func (mr *MockGithubMockRecorder) ListOpenPullRequests() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenPullRequests", reflect.TypeOf((*MockGithub)(nil).ListOpenPullRequests))
}

// PostComment mocks base method
func (m *MockGithub) PostComment(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "PostComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostComment indicates an expected call of PostComment
func (mr *MockGithubMockRecorder) PostComment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostComment", reflect.TypeOf((*MockGithub)(nil).PostComment), arg0, arg1)
}

// UpdateCommitStatus mocks base method
func (m *MockGithub) UpdateCommitStatus(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "UpdateCommitStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCommitStatus indicates an expected call of UpdateCommitStatus
func (mr *MockGithubMockRecorder) UpdateCommitStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommitStatus", reflect.TypeOf((*MockGithub)(nil).UpdateCommitStatus), arg0, arg1, arg2)
}
