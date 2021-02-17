package yaml_test

import (
	. "github.com/qtoad/mygo-plusplus/check"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})
