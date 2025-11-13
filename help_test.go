package help

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// MockLogger is a mock implementation of the Logger interface.
type MockLogger struct {
	InfoCalled  bool
	ErrorCalled bool
}

func (m *MockLogger) Info(message string, args ...any)  { m.InfoCalled = true }
func (m *MockLogger) Error(message string, args ...any) { m.ErrorCalled = true }

// MockApp is a mock implementation of the App interface.
type MockApp struct {
	logger Logger
}

func (m *MockApp) Logger() Logger { return m.logger }

// MockCore is a mock implementation of the Core interface.
type MockCore struct {
	ActionCalled bool
	ActionMsg    map[string]any
	app          App
}

func (m *MockCore) ACTION(msg map[string]any) error {
	m.ActionCalled = true
	m.ActionMsg = msg
	return nil
}

func (m *MockCore) App() App { return m.app }

// MockDisplay is a mock implementation of the Display interface.
type MockDisplay struct{}

func setupService(t *testing.T) (*Service, *MockCore, *MockDisplay) {
	s, err := New()
	assert.NoError(t, err)

	mockLogger := &MockLogger{}
	mockApp := &MockApp{logger: mockLogger}
	mockCore := &MockCore{app: mockApp}
	mockDisplay := &MockDisplay{}

	s.Init(mockCore, mockDisplay)

	return s, mockCore, mockDisplay
}

func TestNew(t *testing.T) {
	s, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestServiceStartup(t *testing.T) {
	s, _, _ := setupService(t)
	err := s.ServiceStartup(context.Background(), application.ServiceOptions{})
	assert.NoError(t, err)
}

func TestShow(t *testing.T) {
	s, mockCore, _ := setupService(t)

	err := s.Show()
	assert.NoError(t, err)
	assert.True(t, mockCore.ActionCalled)

	msg := mockCore.ActionMsg
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])
}

func TestShowAt(t *testing.T) {
	s, mockCore, _ := setupService(t)

	err := s.ShowAt("test-anchor")
	assert.NoError(t, err)
	assert.True(t, mockCore.ActionCalled)

	msg := mockCore.ActionMsg
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])

	opts, ok := msg["options"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "/#test-anchor", opts["URL"])
}
