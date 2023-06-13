package configurator

import (
	"fmt"
	"testing"
)

func TestConfig_ParseFile(t *testing.T) {
	c := New("base_config/default.yaml")
	fmt.Println(c.Buffer)
}
