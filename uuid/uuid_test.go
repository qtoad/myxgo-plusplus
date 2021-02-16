package uuid_test

import (
	"fmt"
	"github.com/qtoad/xgo-plusplus/uuid"
	"testing"
)

func TestKeyEqual(t *testing.T) {
	var UUID4 = uuid.UUID4()
	fmt.Println("UUID4 is", UUID4)
	fmt.Println("UUID4 is", uuid.UUID4())
	fmt.Println("UUID4 is", uuid.UUID4())
	fmt.Println("UUID4 is", uuid.UUID4())
	fmt.Println("UUID4 is", uuid.UUID4())

	UUID5 := uuid.UUID5(uuid.NamespaceDNS, "SomeName")
	fmt.Println("UUID5", UUID5)
}
