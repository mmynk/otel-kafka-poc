// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					Active:          MetricConfig{Enabled: true},
					BlocksReceived:  MetricConfig{Enabled: true},
					BlocksSent:      MetricConfig{Enabled: true},
					Buffered:        MetricConfig{Enabled: true},
					Cached:          MetricConfig{Enabled: true},
					ContextSwitches: MetricConfig{Enabled: true},
					Free:            MetricConfig{Enabled: true},
					IdleTime:        MetricConfig{Enabled: true},
					Inactive:        MetricConfig{Enabled: true},
					Interrupts:      MetricConfig{Enabled: true},
					IoWaitTime:      MetricConfig{Enabled: true},
					RunnableProcs:   MetricConfig{Enabled: true},
					StolenTime:      MetricConfig{Enabled: true},
					SwapIn:          MetricConfig{Enabled: true},
					SwapOut:         MetricConfig{Enabled: true},
					Swapped:         MetricConfig{Enabled: true},
					SystemTime:      MetricConfig{Enabled: true},
					TotalProcs:      MetricConfig{Enabled: true},
					UserTime:        MetricConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					Active:          MetricConfig{Enabled: false},
					BlocksReceived:  MetricConfig{Enabled: false},
					BlocksSent:      MetricConfig{Enabled: false},
					Buffered:        MetricConfig{Enabled: false},
					Cached:          MetricConfig{Enabled: false},
					ContextSwitches: MetricConfig{Enabled: false},
					Free:            MetricConfig{Enabled: false},
					IdleTime:        MetricConfig{Enabled: false},
					Inactive:        MetricConfig{Enabled: false},
					Interrupts:      MetricConfig{Enabled: false},
					IoWaitTime:      MetricConfig{Enabled: false},
					RunnableProcs:   MetricConfig{Enabled: false},
					StolenTime:      MetricConfig{Enabled: false},
					SwapIn:          MetricConfig{Enabled: false},
					SwapOut:         MetricConfig{Enabled: false},
					Swapped:         MetricConfig{Enabled: false},
					SystemTime:      MetricConfig{Enabled: false},
					TotalProcs:      MetricConfig{Enabled: false},
					UserTime:        MetricConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
