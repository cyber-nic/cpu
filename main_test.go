package main

import (
	"testing"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

// MockCPUInfoProvider mocks CPU info retrieval
type MockCPUProvider struct{}

func (m *MockCPUProvider) Info() ([]cpu.InfoStat, error) {
	return []cpu.InfoStat{
		{
			ModelName: "Intel Core i7-9700K",
			CoreID:    "0",
			Mhz:       3600.0,
		},
		{
			ModelName: "Intel Core i7-9700K",
			CoreID:    "1",
			Mhz:       3600.0,
		},
	}, nil
}

func (m *MockCPUProvider) Percent(interval time.Duration, percpu bool) ([]float64, error) {
	return []float64{10.5, 20.3}, nil
}

func TestGetCPUInfo(t *testing.T) {
	// Test successful CPU info retrieval
	t.Run("Successful CPU Info Retrieval", func(t *testing.T) {
		mockProvider := &MockCPUProvider{}

		// Call the function
		result := getCPUInfo(mockProvider)

		// Assertions
		assert.Equal(t, "Intel Core i7-9700K", result.name)
		assert.Equal(t, float64(3600.0), result.speed)
		assert.Equal(t, 2, result.cores)
		assert.Equal(t, []float64{10.5, 20.3}, result.usage)
	})
}

func TestCenterText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		width    int
		expected string
	}{
		{
			name:     "TextFitsExactly",
			text:     "Hello",
			width:    5,
			expected: "Hello",
		},
		{
			name:     "TextNeedsPadding",
			text:     "Hello",
			width:    11,
			expected: "   Hello   ",
		},
		{
			name:     "TextLongerThanWidth",
			text:     "Hello, World!",
			width:    5,
			expected: "Hello, World!",
		},
		{
			name:     "TextWithEvenPadding",
			text:     "Hi",
			width:    6,
			expected: "  Hi  ",
		},
		{
			name:     "TextWithOddPadding",
			text:     "Hi",
			width:    7,
			expected: "  Hi   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := centerText(tt.text, tt.width)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCreateThreadBox(t *testing.T) {
	tests := []struct {
		name      string
		threadNum int
		width     int
		usage     float64
		expected  []string
	}{
		{
			name:      "LowUsage",
			threadNum: 0,
			width:     10,
			usage:     10.0,
			expected: []string{
				"┌────────┐",
				"│Thread 0│",
				"│" + aurora.Green("█").String() + "       │",
				"└────────┘",
			},
		},
		{
			name:      "MediumUsage",
			threadNum: 1,
			width:     10,
			usage:     50.0,
			expected: []string{
				"┌────────┐",
				"│Thread 1│",
				"│" + aurora.Yellow("████").String() + "    │",
				"└────────┘",
			},
		},
		{
			name:      "HighUsage",
			threadNum: 2,
			width:     10,
			usage:     75.0,
			expected: []string{
				"┌────────┐",
				"│Thread 2│",
				"│" + aurora.BrightRed("██████").String() + "  │",
				"└────────┘",
			},
		},
		{
			name:      "MaxUsage",
			threadNum: 3,
			width:     10,
			usage:     100.0,
			expected: []string{
				"┌────────┐",
				"│Thread 3│",
				"│" + aurora.Red("████████").String() + "│",
				"└────────┘",
			},
		},
		{
			name:      "UsageExceedsWidth",
			threadNum: 4,
			width:     5,
			usage:     100.0,
			expected: []string{
				"┌───┐",
				"│Thread 4│",
				"│" + aurora.Red("███").String() + "│",
				"└───┘",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := createThreadBox(tt.threadNum, tt.width, tt.usage)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestCreateCoreBox(t *testing.T) {
	tests := []struct {
		name           string
		coreNum        int
		threadsPerCore int
		width          int
		threadWidth    int
		usage          []float64
		expected       []string
	}{
		{
			name:           "SingleThreadCore",
			coreNum:        0,
			threadsPerCore: 1,
			width:          20,
			threadWidth:    10,
			usage:          []float64{10.0},
			expected: []string{
				"┌──────────────────┐",
				"│      Core 0      │",
				"│    ┌────────┐    │",
				"│    │Thread 0│    │",
				"│    │" + aurora.Green("█").String() + "       │    │",
				"│    └────────┘    │",
				"└──────────────────┘",
			},
		},
		{
			name:           "MultiThreadCore",
			coreNum:        0,
			threadsPerCore: 2,
			width:          20,
			threadWidth:    10,
			usage:          []float64{10.0, 50.0},
			expected: []string{
				"┌──────────────────┐",
				"│      Core 0      │",
				"│    ┌────────┐    │",
				"│    │Thread 0│    │",
				"│    │" + aurora.Green("█").String() + "       │    │",
				"│    └────────┘    │",
				"│    ┌────────┐    │",
				"│    │Thread 1│    │",
				"│    │" + aurora.Yellow("████").String() + "    │    │",
				"│    └────────┘    │",
				"└──────────────────┘",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := createCoreBox(tt.coreNum, tt.threadsPerCore, tt.width, tt.threadWidth, tt.usage)
			assert.Equal(t, tt.expected, result)
		})
	}
}
