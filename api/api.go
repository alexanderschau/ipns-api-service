package api

import (
	"context"

	shell "github.com/ipfs/go-ipfs-api"
)

var s = shell.NewShell(":5001")
var ctx = context.TODO()
