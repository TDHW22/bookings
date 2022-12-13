package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// Appconfig holds the application config
type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
