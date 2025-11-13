---
title: Help
---

# Overview

This module provides an in-app help system for Wails applications, including a simple way to display documentation and handle deep-links to specific sections.

## Quick start
```go
package main

import (
    "github.com/wailsapp/wails/v3/pkg/application"
    "github.com/snider/help"
)

func main() {
    app := application.New(application.Options{})
    helpService, _ := help.Register(app)

    wailsApp := application.New(application.Options{
        Services: []application.Service{
            application.NewService(helpService),
        },
    })
    wailsApp.Run()
}
```

## Usage
```go
package main

import (
    "github.com/wailsapp/wails/v3/pkg/application"
    "github.com/snider/help"
)

type MyService struct {
    help help.Service
}

func (s *MyService) ShowHelp() {
    s.help.Show()
}

func (s *MyService) ShowHelpAt(anchor string) {
    s.help.ShowAt(anchor)
}
```
