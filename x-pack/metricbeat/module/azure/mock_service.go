// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package azure

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-06-01/insights"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-03-01/resources"
	"github.com/stretchr/testify/mock"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
)

// MockService mock for the azure monitor services
type MockService struct {
	mock.Mock
}

// GetResourceDefinitions is a mock function for the azure service
func (client *MockService) GetResourceDefinitions(ID []string, group []string, rType string, query string) (resources.ListResultPage, error) {
	args := client.Called(ID, group, rType, query)
	return args.Get(0).(resources.ListResultPage), args.Error(1)
}

// GetMetricDefinitions is a mock function for the azure service
func (client *MockService) GetMetricDefinitions(resourceID string, namespace string) (insights.MetricDefinitionCollection, error) {
	args := client.Called(resourceID, namespace)
	return args.Get(0).(insights.MetricDefinitionCollection), args.Error(1)
}

// GetMetricNamespaces is a mock function for the azure service
func (client *MockService) GetMetricNamespaces(resourceID string) (insights.MetricNamespaceCollection, error) {
	args := client.Called(resourceID)
	return args.Get(0).(insights.MetricNamespaceCollection), args.Error(1)
}

// GetMetricValues is a mock function for the azure service
func (client *MockService) GetMetricValues(resourceID string, namespace string, timegrain string, timespan string, metricNames []string, aggregations string, filter string) ([]insights.Metric, error) {
	args := client.Called(resourceID, namespace)
	return args.Get(0).([]insights.Metric), args.Error(1)
}

// MockReporterV2 mock implementation for testing purposes
type MockReporterV2 struct {
	mock.Mock
}

// Event function is mock implementation for testing purposes
func (reporter *MockReporterV2) Event(event mb.Event) bool {
	args := reporter.Called(event)
	return args.Get(0).(bool)
}

// Error is mock implementation for testing purposes
func (reporter *MockReporterV2) Error(err error) bool {
	args := reporter.Called(err)
	return args.Get(0).(bool)
}

// NewMockClient instantiates a new client with the mock azure service
func NewMockClient() *Client {
	azureMockService := new(MockService)
	client := &Client{
		AzureMonitorService: azureMockService,
		Config:              Config{},
		Log:                 logp.NewLogger("test azure monitor"),
	}
	return client
}
