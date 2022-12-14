package biz

import (
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(bizsteam.NewSteamUseCase)
