package app

import "context"

func (a *App) Startup(ctx context.Context) {
  a.ctx = ctx
}
