package resp

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/utils"
)

func Serialize(rawCmd string) string {
	segments := utils.GetCmdSegments(rawCmd)
	serializedCmd := fmt.Sprintf("*%d\r\n", len(segments))
	for i := 0; i < len(segments); i++ {
		serializedCmd += fmt.Sprintf("$%d\r\n%s\r\n", len(segments[i]), segments[i])
	}
	return serializedCmd
}
