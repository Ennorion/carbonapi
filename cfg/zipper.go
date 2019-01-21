package cfg

import (
	"io"

	"github.com/bookingcom/carbonapi/pathcache"
)

// Zipper is the zipper config
type Zipper struct {
	Common    `yaml:",inline"`
	PathCache pathcache.PathCache
}

// ParseZipperConfig reads the zipper config from a supplied reader
func ParseZipperConfig(r io.Reader) (Zipper, error) {
	cfg, err := ParseCommon(r)

	if err != nil {
		return Zipper{}, err
	}

	return fromCommon(cfg), nil
}

func fromCommon(c Common) Zipper {
	return Zipper{
		Common:    c,
		PathCache: pathcache.NewPathCache(c.ExpireDelaySec),
	}
}
