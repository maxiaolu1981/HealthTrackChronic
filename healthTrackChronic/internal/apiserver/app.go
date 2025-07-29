package apiserver

import (
	"github.com/maxiaolu1981/healthTrackChronic/internal/apiserver/options"
	"github.com/maxiaolu1981/healthTrackChronic/pkg/app"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more iam-apiserver information at:
    https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-apiserver.md`

func NewApp(baseName string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("healthTrackChronic API Server",
		baseName,
		app.WithOptions(opts),
	)
}
