package utils

import (
	mathRand "math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

// Id generate a ULID
func Id() string {
	for {
		seed := time.Now().UnixNano()
		entropy := mathRand.New(mathRand.NewSource(seed))
		id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
		if err != nil {
			//log.Warn("generate ulid failed", zap.Error(err))
			continue
		}

		return strings.ToLower(id.String())
	}
}
