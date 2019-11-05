package fabenc

import "fmt"

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/11/02     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/11/2 5:06 PM
 * @date 2019/11/2 5:06 PM
 * @since 1.0.0
 */

type Color uint8

const ColorNone Color = 0

const (
	ColorBlack Color = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

func (c Color) Normal() string {
	return fmt.Sprintf("\x1b[%dm", c)
}

func (c Color) Bold() string {
	if c == ColorNone {
		return c.Normal()
	}
	return fmt.Sprintf("\x1b[%d;1m", c)
}

func ResetColor() string { return ColorNone.Normal() }
