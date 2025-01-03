// Copyright (c) 2025 The bel2 developers

package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/BeL2Labs/Arbiter_Signer/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
