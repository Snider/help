package help

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
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

	return s, mockRuntime, mockDisplay, mockLogHandler
}

func TestNew(t *testing.T) {
	s, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestServiceStartup(t *testing.T) {
	s, _, _ := setupService(t)
	err := s.ServiceStartup(context.Background())
	assert.NoError(t, err)
}

func TestShow(t *testing.T) {
	s, mockRuntime, _, _ := setupService(t)

	err := s.Show()
	assert.NoError(t, err)
	assert.True(t, mockRuntime.ActionCalled)

	msg := mockCore.ActionMsg
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])
}

func TestShowAt(t *testing.T) {
	s, mockRuntime, _, _ := setupService(t)

	err := s.ShowAt("test-anchor")
	assert.NoError(t, err)
	assert.True(t, mockRuntime.ActionCalled)

	msg := mockCore.ActionMsg
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])

	opts, ok := msg["options"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "/#test-anchor", opts["URL"])
}

func TestServiceStartup_CoreNotInitialized(t *testing.T) {
	s, _, _ := setupService(t)
	s.core = nil
	err := s.ServiceStartup(context.Background())
	assert.Error(t, err)
	assert.Equal(t, "core runtime not initialized", err.Error())
}

func TestShow_DisplayNotInitialized(t *testing.T) {
	s, _, _ := setupService(t)
	s.display = nil
	err := s.Show()
	assert.Error(t, err)
	assert.Equal(t, "display service not initialized", err.Error())
}

func TestShow_CoreNotInitialized(t *testing.T) {
	s, _, _ := setupService(t)
	s.core = nil
	err := s.Show()
	assert.Error(t, err)
	assert.Equal(t, "core runtime not initialized", err.Error())
}

func TestShowAt_DisplayNotInitialized(t *testing.T) {
	s, _, _ := setupService(t)
	s.display = nil
	err := s.ShowAt("some-anchor")
	assert.Error(t, err)
	assert.Equal(t, "display service not initialized", err.Error())
}

func TestShowAt_CoreNotInitialized(t *testing.T) {
	s, _, _ := setupService(t)
	s.core = nil
	err := s.ShowAt("some-anchor")
	assert.Error(t, err)
	assert.Equal(t, "core runtime not initialized", err.Error())
}
