package common

import (
	"github.com/crispuscrew/pgxray/src/internal/opt"

	"time"
)

type CompleteLoadingMsg struct{}

// Toast messages
// Zero timeout means never expire
// Not set timeout means use default timeout
type AddCriticalToast 	struct { Item string; Timeout opt.Opt[time.Duration] }
type AddErrorToast 		struct { Item string; Timeout opt.Opt[time.Duration] }
type AddWarningToast 	struct { Item string; Timeout opt.Opt[time.Duration] }
type AddInfoToast 		struct { Item string; Timeout opt.Opt[time.Duration] }