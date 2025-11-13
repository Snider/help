package help

import (
	"testing"

	"github.com/Snider/Core"
	"github.com/stretchr/testify/assert"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// MockDisplay is a mock implementation of the core.Display interface.
type MockDisplay struct {
	ShowCalled bool
}

func (m *MockDisplay) Show() error {
	m.ShowCalled = true
	return nil
}

func (m *MockDisplay) ShowAt(anchor string) error {
	m.ShowCalled = true
	return nil
}

func (m *MockDisplay) Hide() error                                { return nil }
func (m *MockDisplay) HideAt(anchor string) error                 { return nil }
func (m *MockDisplay) OpenWindow(opts ...core.WindowOption) error { return nil }

// MockCore is a mock implementation of the *core.Core type.
type MockCore struct {
	Core         *core.Core
	ActionCalled bool
	ActionMsg    core.Message
}

// ACTION matches the signature required by RegisterAction.
func (m *MockCore) ACTION(c *core.Core, msg core.Message) error {
	m.ActionCalled = true
	m.ActionMsg = msg
	return nil
}

func setupService(t *testing.T) (*Service, *MockCore, *MockDisplay) {
	s, err := New()
	assert.NoError(t, err)

	app := application.New(application.Options{})
	c, err := core.New(core.WithWails(app))
	assert.NoError(t, err)
	mockCore := &MockCore{Core: c}
	mockDisplay := &MockDisplay{}

	s.core = c
	s.display = mockDisplay
	// Register our mock handler. When the real s.Core().ACTION is called,
	// our mock handler will be executed.
	c.RegisterAction(mockCore.ACTION)

	return s, mockCore, mockDisplay
}

func TestNew(t *testing.T) {
	s, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestRegister(t *testing.T) {
	app := application.New(application.Options{})
	c, err := core.New(core.WithWails(app))
	assert.NoError(t, err)

	s, err := Register(c)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	service, ok := s.(*Service)
	assert.True(t, ok)
	assert.NotNil(t, service.core)
	assert.Equal(t, c, service.core)
}

func TestShow(t *testing.T) {
	s, mockCore, _ := setupService(t)

	err := s.Show()
	assert.NoError(t, err)
	assert.True(t, mockCore.ActionCalled)

	msg, ok := mockCore.ActionMsg.(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])
}

func TestShowAt(t *testing.T) {
	s, mockCore, _ := setupService(t)

	err := s.ShowAt("test-anchor")
	assert.NoError(t, err)
	assert.True(t, mockCore.ActionCalled)

	msg, ok := mockCore.ActionMsg.(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])

	opts, ok := msg["options"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "/#test-anchor", opts["URL"])
}

func TestHandleIPCEvents_ServiceStartup(t *testing.T) {
	s, _, _ := setupService(t)
	err := s.HandleIPCEvents(s.core, core.ActionServiceStartup{})
	assert.NoError(t, err)
}

func TestHandleIPCEvents_Default(t *testing.T) {
	s, _, _ := setupService(t)

	// Define a custom message type that is not handled by HandleIPCEvents.
	type unhandledMessage struct{}
	err := s.HandleIPCEvents(s.core, unhandledMessage{})
	assert.NoError(t, err)
}

func TestShow_Errors(t *testing.T) {
	t.Run("NoDisplay", func(t *testing.T) {
		s, err := New()
		assert.NoError(t, err)
		app := application.New(application.Options{})
		c, err := core.New(core.WithWails(app))
		assert.NoError(t, err)
		s.core = c

		err = s.Show()
		assert.Error(t, err)
		assert.Equal(t, "display service not initialized", err.Error())
	})

	t.Run("NoCore", func(t *testing.T) {
		s, err := New()
		assert.NoError(t, err)
		s.display = &MockDisplay{}

		err = s.Show()
		assert.Error(t, err)
		assert.Equal(t, "core runtime not initialized", err.Error())
	})
}

func TestShowAt_Errors(t *testing.T) {
	t.Run("NoDisplay", func(t *testing.T) {
		s, err := New()
		assert.NoError(t, err)
		app := application.New(application.Options{})
		c, err := core.New(core.WithWails(app))
		assert.NoError(t, err)
		s.core = c

		err = s.ShowAt("some-anchor")
		assert.Error(t, err)
		assert.Equal(t, "display service not initialized", err.Error())
	})

	t.Run("NoCore", func(t *testing.T) {
		s, err := New()
		assert.NoError(t, err)
		s.display = &MockDisplay{}

		err = s.ShowAt("some-anchor")
		assert.Error(t, err)
		assert.Equal(t, "core runtime not initialized", err.Error())
	})
}
