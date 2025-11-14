package help

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:public/*
var helpStatic embed.FS

// Logger defines the interface for logging messages.
type Logger interface {
	Info(message string, args ...any)
	Error(message string, args ...any)
}

// App defines the interface for the application-level context.
type App interface {
	Logger() Logger
}

// Core defines the interface for the core runtime functionalities
// that the help service depends on.
type Core interface {
	ACTION(msg map[string]any) error
	App() App
}

// Display defines the interface for a display service.
// The help service requires this to be present for dependency checking.
type Display interface{}

// Help defines the interface for the help service.
type Help interface {
	Show() error
	ShowAt(anchor string) error
	ServiceStartup(ctx context.Context) error
}

// Options holds configuration for the help service.
type Options struct {
	Source string
	Assets fs.FS
}

// Service manages the in-app help system.
type Service struct {
	core    Core
	display Display
	assets  fs.FS
	opts    Options
}

// New is the constructor for the help service.
func New(opts Options) (*Service, error) {
	if opts.Source == "" {
		opts.Source = "mkdocs"
	}

	s := &Service{
		opts: opts,
	}

	var err error
	if opts.Assets != nil {
		s.assets = opts.Assets
	} else if s.opts.Source != "mkdocs" {
		s.assets = os.DirFS(s.opts.Source)
	} else {
		s.assets, err = fs.Sub(helpStatic, "public")
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// Init initializes the service with its core dependencies.
func (s *Service) Init(c Core, d Display) {
	s.core = c
	s.display = d
}

// ServiceStartup is called when the app starts, after dependencies are injected.
func (s *Service) ServiceStartup(context.Context) error {
	if s.core == nil {
		return fmt.Errorf("core runtime not initialized")
	}
	s.core.App().Logger().Info("Help service started")
	return nil
}

// Show displays the help window.
func (s *Service) Show() error {
	// Check core first for consistent behavior in both normal and fallback paths
	if s.core == nil {
		return fmt.Errorf("core runtime not initialized")
	}

	if s.display == nil {
		app := application.Get()
		if app == nil {
			return fmt.Errorf("wails application not running")
		}
		// Note: The Wails application must be configured to serve the help assets
		// from the root path ("/") for this fallback to work correctly.
		app.Window.NewWithOptions(application.WebviewWindowOptions{
			Title:  "Help",
			Width:  800,
			Height: 600,
			URL:    "/",
		})
		return nil
	}
	msg := map[string]any{
		"action": "display.open_window",
		"name":   "help",
		"options": map[string]any{
			"Title":  "Help",
			"Width":  800,
			"Height": 600,
		},
	}

	return s.core.ACTION(msg)
}

// ShowAt displays a specific section of the help documentation.
func (s *Service) ShowAt(anchor string) error {
	// Check core first for consistent behavior in both normal and fallback paths
	if s.core == nil {
		return fmt.Errorf("core runtime not initialized")
	}

	if s.display == nil {
		app := application.Get()
		if app == nil {
			return fmt.Errorf("wails application not running")
		}
		// Note: The Wails application must be configured to serve the help assets
		// from the root path for this fallback to work correctly.
		url := fmt.Sprintf("/#%s", anchor)
		app.Window.NewWithOptions(application.WebviewWindowOptions{
			Title:  "Help",
			Width:  800,
			Height: 600,
			URL:    url,
		})
		return nil
	}

	url := fmt.Sprintf("/#%s", anchor)

	msg := map[string]any{
		"action": "display.open_window",
		"name":   "help",
		"options": map[string]any{
			"Title":  "Help",
			"Width":  800,
			"Height": 600,
			"URL":    url,
		},
	}
	return s.core.ACTION(msg)
}

// Ensure Service implements the Help interface.
var _ Help = (*Service)(nil)
