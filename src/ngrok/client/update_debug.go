// +build !release,!autoupdate

package client

import (
	"github.com/hasura/hasuractl-go/pkg/ngrok/src/ngrok/client/mvc"
)

// no auto-updating in debug mode
func autoUpdate(state mvc.State, token string) {
}
