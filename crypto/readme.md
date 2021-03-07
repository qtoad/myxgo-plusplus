```
import (
	"fmt"

	"github.com/eaglexpf/crypto"
)

data, err := crypto.NewRsa()
if err != nil {
	fmt.Println("error")
} else {
	fmt.Println(string(data.PrivateKey))
	fmt.Println(string(data.PublicKey))
}
encodeData, _ := crypto.EncodeRSA(data.PublicKey, []byte("Hello World!!!this is a success message!"))
decodeData, _ := crypto.DecodeRSA(data.PrivateKey, encodeData)
fmt.Println(encodeData)
fmt.Println(string(decodeData))

```