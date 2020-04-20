package namegen

import (
	"fmt"
	"github.com/wondenge/namegen"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(namegen.GetRandomName(0))
}
