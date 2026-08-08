package output

import (
	"fmt"
	"strings"
)

type acmeFormatter struct{}

func (f acmeFormatter) Output(buf []byte) string {
	builder := strings.Builder{}

	nbBlocks := len(buf) >> 3
	for i := range nbBlocks {
		fmt.Fprintf(&builder, "; CHAR %d\n", i)

		for j := range 8 {
			fmt.Fprintf(&builder, "!byte %%%08b\n", buf[(i<<3)+j])
		}

		fmt.Fprintln(&builder)
	}

	return builder.String()
}
