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

// MockRuntime is a mock implementation of the *core.Core type.
type MockRuntime struct {
	Runtime      *core.Core
	ActionCalled bool
	ActionMsg    core.Message
}

// ACTION matches the signature required by RegisterAction.
func (m *MockRuntime) ACTION(r *core.Core, msg core.Message) error {
	m.ActionCalled = true
	m.ActionMsg = msg
	return nil
}

func setupService(t *testing.T) (*Service, *MockRuntime, *MockDisplay) {
	s, err := New()
	assert.NoError(t, err)

	app := application.New(application.Options{})
	r, err := core.New(core.WithWails(app))
	assert.NoError(t, err)
	mockRuntime := &MockRuntime{Runtime: r}
	mockDisplay := &MockDisplay{}

	s.runtime = r
	s.display = mockDisplay
	// Register our mock handler. When the real s.runtime.ACTION is called,
	// our mock handler will be executed.
	r.RegisterAction(mockRuntime.ACTION)

	return s, mockRuntime, mockDisplay
}

func TestNew(t *testing.T) {
	s, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestRegister(t *testing.T) {
	app := application.New(application.Options{})
	r, err := core.New(core.WithWails(app))
	assert.NoError(t, err)

	s, err := Register(r)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	service, ok := s.(*Service)
	assert.True(t, ok)
	assert.NotNil(t, service.runtime)
	assert.Equal(t, r, service.runtime)
}

func TestShow(t *testing.T) {
	s, mockRuntime, _ := setupService(t)

	err := s.Show()
	assert.NoError(t, err)
	assert.True(t, mockRuntime.ActionCalled)

	msg, ok := mockRuntime.ActionMsg.(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])
}

func TestShowAt(t *testing.T) {
	s, mockRuntime, _ := setupService(t)

	err := s.ShowAt("test-anchor")
	assert.NoError(t, err)
	assert.True(t, mockRuntime.ActionCalled)

	msg, ok := mockRuntime.ActionMsg.(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "display.open_window", msg["action"])
	assert.Equal(t, "help", msg["name"])

	opts, ok := msg["options"].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "/#test-anchor", opts["URL"])
}

func TestHandleIPCEvents_ServiceStartup(t *testing.T) {
	s, _, _ := setupService(t)
	err := s.HandleIPCEvents(s.runtime, core.ActionServiceStartup{})
	assert.NoError(t, err)
}

func TestHandleIPCEvents_Default(t *testing.T) {
	s, _, _ := setupService(t)

	// Define a custom message type that is not handled by HandleIPCEvents.
	type unhandledMessage struct{}
	err := s.HandleIPCEvents(s.runtime, unhandledMessage{})
	assert.NoError(t, err)
}

func TestShow_Errors(t *testing.T) {
	t.Run("NoDisplay", func(t *testing.T) {
		s, err := New()
		assert.NoError(t, err)
		app := application.New(application.Options{})
		r, err := core.New(core.WithWails(app))
		assert.NoError(t, err)
		s.runtime = r

		err = s.Show()
		assert.Error(t, err)
		assert.Equal(t, "display service not initialized", err.Error())
	})

	t.Run("NoRuntime", func(t *testing.T) {
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
		r, err := core.New(core.WithWails(app))
		assert.NoError(t, err)
		s.runtime = r

		err = s.ShowAt("some-anchor")
		assert.Error(t, err)
		assert.Equal(t, "display service not initialized", err.Error())
	})

	t.Run("NoRuntime", func(t *testing.T) {
		s, err := New()
		assert.NoError(t, err)
		s.display = &MockDisplay{}

		err = s.ShowAt("some-anchor")
		assert.Error(t, err)
		assert.Equal(t, "core runtime not initialized", err.Error())
	})
}
