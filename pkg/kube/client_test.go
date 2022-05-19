//go:build integration
// +build integration

package kube

import (
	"context"
	"os"
	"testing"

	"github.com/odpf/entropy/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var envClusterCACert = os.Getenv("TEST_K8S_CLUSTER_CA_CERT")
var envClientKey = os.Getenv("TEST_K8S_CLIENT_KEY")
var envClientCert = os.Getenv("TEST_K8S_CLIENT_CERT")
var envHost = os.Getenv("TEST_K8S_HOST")
var envNamespace = os.Getenv("TEST_K8S_NAMESPACE")
var envPod = os.Getenv("TEST_K8S_POD")
var envContainer = os.Getenv("TEST_K8S_CONTAINER")

func TestGetStreamingLogs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name          string
		ClusterCACert string
		ClientKey     string
		ClientCert    string
		Host          string
		Namespace     string
		Pod           string
		Container     string
		wantErr       error
	}{
		{
			Name:          "InvalidCredentials",
			ClusterCACert: "invalid cluster certificate",
			ClientKey:     "invalid client key",
			ClientCert:    "invalid client certificate",
			Host:          "invalid host",
			Namespace:     "invalid namespace",
			Pod:           "invalid pod",
			Container:     "invalid container",
			wantErr:       errors.New("invalid credentials"),
		},
		{
			Name:          "StreamFromOnePod",
			ClusterCACert: envClusterCACert,
			ClientKey:     envClientKey,
			ClientCert:    envClientCert,
			Host:          envHost,
			Namespace:     envNamespace,
			Pod:           envPod,
			Container:     envContainer,
			wantErr:       nil,
		},
		{
			Name:          "StreamFromAllPods",
			ClusterCACert: envClusterCACert,
			ClientKey:     envClientKey,
			ClientCert:    envClientCert,
			Host:          envHost,
			Namespace:     envNamespace,
			Pod:           "",
			Container:     "",
			wantErr:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			filter := make(map[string]string)
			filter["pod"] = tt.Pod
			filter["container"] = tt.Container

			config := DefaultClientConfig()
			config.Host = tt.Host
			config.ClusterCACertificate = tt.ClusterCACert
			config.ClientKey = tt.ClientKey
			config.ClientCertificate = tt.ClientCert

			client := NewClient(config)

			ctx := new(context.Context)
			_, err := client.StreamLogs(*ctx, tt.Namespace, filter)
			if tt.wantErr != nil {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
