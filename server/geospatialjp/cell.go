package geospatialjp

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type CellPos struct {
	x int
	y int
}

var cellRegexp = regexp.MustCompile("^([A-Z]+)([0-9]+)$")

const cellAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ParseCellPos(p string) (CellPos, error) {
	m := cellRegexp.FindStringSubmatch(p)
	if len(m) < 3 {
		return CellPos{}, errors.New("invalid cell pos")
	}
	y, err := strconv.Atoi(m[2])
	if err != nil {
		return CellPos{}, errors.New("invalid cell pos")
	}
	return CellPos{x: xCode(strings.ToUpper(m[1])), y: y}, nil
}

func (c CellPos) ShiftX(shift int) CellPos {
	return CellPos{x: c.x + shift, y: c.y}
}

func (c CellPos) ShiftY(n int) CellPos {
	return CellPos{x: c.x, y: c.y + n}
}

func (c CellPos) String() string {
	return fmt.Sprintf("%s%d", fromXCode(c.x), c.y)
}

func xCode(x string) (code int) {
	le := len(cellAlphabet)
	for i := 0; i < len(x); i++ {
		ri := len(x) - i - 1
		if si := strings.Index(cellAlphabet, string(x[ri])); si >= 0 {
			if i == 0 {
				code += si
			} else {
				code += (si + 1) * int(math.Pow(float64(le), float64(i)))
			}
		}
	}
	return
}

func fromXCode(code int) string {
	if code < 0 {
		code = 0
	}

	x := []rune{}
	le := len(cellAlphabet)

	for {
		pos := code % le
		x = append(x, rune(cellAlphabet[pos]))
		next := code / le
		if next < 1 {
			break
		}
		code = next - 1
	}

	return string(lo.Reverse(x))
}
