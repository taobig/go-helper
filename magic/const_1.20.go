//go:build go1.20
// +build go1.20

package magic

import "time"

// Deprecated: As of Go 1.20, use time.DateOnly.
const SimpleDateLayout = time.DateOnly //2006-01-02

// Deprecated: As of Go 1.20, use time.DateTime.
const SimpleDatetimeLayout = time.DateTime //2006-01-02 15:04:05

// Deprecated
const SimpleDatetimeNanoLayout = "2006-01-02 15:04:05.999999999" //2014-11-23 14:24:03.1305627
