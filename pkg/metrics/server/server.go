/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	ctrlruntimemetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

// We have to unregister the ProcessCollector and GoCollector
// from the ctrltuntimemetrics Registry, otherwise Collecting errors
// out because they are both there and in the default prometheus registry.
// This is not extremely nice but as pretty as "collect metrics from the
// two registries" will ever got, unless the ctrltuntimemetrics.Registry
// becomes configurable
func init() {
	ctrlruntimemetrics.Registry.Unregister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	ctrlruntimemetrics.Registry.Unregister(collectors.NewGoCollector())
}

// New returns a brand new *MetricsServer that gathers the metrics
// from both the prometheus default registry and the ctrlruntimemetrics registry
func New(listenAddress string) *MetricsServer {
	return &MetricsServer{
		gatherers:     []prometheus.Gatherer{prometheus.DefaultGatherer, ctrlruntimemetrics.Registry},
		listenAddress: listenAddress,
	}
}

// MetricsServer is our own metrics server implementation that gathers the metrics from
// both the default prometheus registry and the ctrltuntimemetrics registry.
// The background is that the latter is not configurable at all and we don't
// want to force developers into using it, because that is counterintuitive
// and prone to be forgotten
type MetricsServer struct {
	gatherers     prometheus.Gatherers
	listenAddress string
}

// Start implements sigs.k8s.io/controller-runtime/pkg/manager.Runnable
func (m *MetricsServer) Start(ctx context.Context) error {
	if len(m.gatherers) < 1 {
		return errors.New("no gatherers defined")
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer,
		promhttp.HandlerFor(m.gatherers, promhttp.HandlerOpts{}),
	))
	s := http.Server{
		Addr:         m.listenAddress,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return fmt.Errorf("metrics server stopped: %v", s.ListenAndServe())
}

// MetricsServer implements LeaderElectionRunnable to indicate that it does not require to run
// within an elected leader
var _ manager.LeaderElectionRunnable = &MetricsServer{}

func (m *MetricsServer) NeedLeaderElection() bool {
	return false
}
