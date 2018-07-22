package load

import (
	"github.com/dynamicgo/mesh"
	app "github.com/dynamicgo/mesh-app"
	servicehub "github.com/dynamicgo/mesh-servicehub"
)

func init() {
	app.ImportService(app.ServiceDescriptor{
		Name: mesh.ProtocolServiceHub,
		Main: servicehub.Main,
		ServiceOptions: []mesh.ServiceOption{
			mesh.NoRemoteConfig(),
		},
	})
}
