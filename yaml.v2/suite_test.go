package yaml_test

import (
	. "github.com/qtoad/xgo-plusplus/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})
