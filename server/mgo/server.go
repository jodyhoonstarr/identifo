package mgo

import (
	"log"

	"github.com/madappgang/identifo/model"
	"github.com/madappgang/identifo/server"
)

// ServerSettings are server settings.
var ServerSettings model.ServerSettings

func init() {
	ServerSettings = server.ServerSettings
	if ServerSettings.DBType != "mongodb" {
		log.Fatalf("Incorrect database type %s for MongoDB-backed server", ServerSettings.DBType)
	}
}

// NewServer creates new backend service with MongoDB support.
func NewServer(settings model.ServerSettings, options ...func(*server.Server) error) (model.Server, error) {
	dbComposer, err := NewComposer(settings)
	if err != nil {
		return nil, err
	}
	return server.NewServer(settings, dbComposer, options...)
}
