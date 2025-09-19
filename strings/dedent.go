package strings

import (
	"strings"

	. "github.com/tobiashort/utils-go/must"
)

func Dedent(s string) string {
	builder := strings.Builder{}
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if i == 0 {
			Must2(builder.WriteString(line))
		} else if index := strings.IndexByte(line, '|'); index > -1 {
			Must2(builder.WriteString(line[index+1:]))
		} else {
			Must2(builder.WriteString(line))
		}
		if i < len(lines)-1 {
			Must(builder.WriteByte('\n'))
		}
	}
	return builder.String()
}
