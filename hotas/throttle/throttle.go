package throttle

import (
	"github.com/rdnt/uinput"
	"github.com/sht/ed-journal/hotas/types"
	"github.com/sht/ed-journal/hotas/util"
	"sync"
)

var state = new(State)

var mux sync.RWMutex

const (
	AxisL uint16 = iota
	AxisR
	AxisF
	AxisTX
	AxisTY
	AxisG
	AxisRTY4
	AxisRTY3
)

const (
	ButtonE uint16 = iota + 0x120
	ButtonF
	ButtonG
	ButtonI
	ButtonH
	ButtonSW1
	ButtonSW2
	ButtonSW3
	ButtonSW4
	ButtonSW5
	ButtonSW6
	ButtonTGL1Up
	ButtonTGL1Down
	ButtonTGL2Up
	ButtonTGL2Down
	ButtonTGL3Up
	ButtonTGL3Down uint16 = iota + 0x2c0
	ButtonTGL4Up
	ButtonTGL4Down
	ButtonH3Up
	ButtonH3Right
	ButtonH3Down
	ButtonH3Left
	ButtonH4Up
	ButtonH4Right
	ButtonH4Down
	ButtonH4Left
	ButtonK1Up
	ButtonK1Down
	ButtonScrollDown
	ButtonScrollUp
	ButtonThumb
	ButtonSLD
	ButtonModeM1
	ButtonModeM2
	ButtonModeS1
)

// State represents the state of the throttle device, 8 axis, 36 buttons
//
// Bit map:
//  0  00000111  LLLLLLLL  ThrottleL
//  1  00000000  RRRRRRLL  ThrottleR, ThrottleL
//  2  00000000  IGFERRRR  I, G, F, E, ThrottleR
//  3  00000000  TSSSSSSH  TGL1Up, SW6, SW5, SW4, SW3, SW2, SW1, H
//  4  00000000  HTTTTTTT  H3(up), TGL4Down, TGL4Up, TGL3Down, TGL3Up, TGL2Down, TGL2Up, TGL1Down
//  5  00000000  KHHHHHHH  K1Up, H4(left), H4(down), H4(right), H4(up), H3(left), H3(down), H3(right)
//  6  00100000  MMMSTSSK  Mode(s1), Mode(m2), Mode(m1), SLD, Thumb(button), ScrollUp, ScrollDown, K1Down
//  7  01111011  FFFFFFFF  F
//  8  10000000  XXXXXXXX  Thumb(x)
//  9  01111111  GGGGGGGG  G
// 10  10000000  YYYYYYYY  Thumb(y)
// 11  00000000  RRRRRRRR  RTY4
// 12  00000000  RRRRRRRR  RTY3
type State struct {
	AxisL            types.Axis
	AxisR            types.Axis
	AxisF            types.Axis
	AxisTX           types.Axis
	AxisTY           types.Axis
	AxisG            types.Axis
	AxisRTY4         types.Axis
	AxisRTY3         types.Axis
	ButtonE          types.Button
	ButtonF          types.Button
	ButtonG          types.Button
	ButtonI          types.Button
	ButtonH          types.Button
	ButtonSW1        types.Button
	ButtonSW2        types.Button
	ButtonSW3        types.Button
	ButtonSW4        types.Button
	ButtonSW5        types.Button
	ButtonSW6        types.Button
	ButtonTGL1Up     types.Button
	ButtonTGL1Down   types.Button
	ButtonTGL2Up     types.Button
	ButtonTGL2Down   types.Button
	ButtonTGL3Up     types.Button
	ButtonTGL3Down   types.Button
	ButtonTGL4Up     types.Button
	ButtonTGL4Down   types.Button
	ButtonH3Up       types.Button
	ButtonH3Right    types.Button
	ButtonH3Down     types.Button
	ButtonH3Left     types.Button
	ButtonH4Up       types.Button
	ButtonH4Right    types.Button
	ButtonH4Down     types.Button
	ButtonH4Left     types.Button
	ButtonK1Up       types.Button
	ButtonK1Down     types.Button
	ButtonScrollDown types.Button
	ButtonScrollUp   types.Button
	ButtonThumb      types.Button
	ButtonSLD        types.Button
	ButtonModeM1     types.Button
	ButtonModeM2     types.Button
	ButtonModeS1     types.Button
}

func GetState() State {
	mux.RLock()
	defer mux.RUnlock()
	return *state
}

var AxesConfig = []uinput.Axis{
	{ID: AxisL, Min: 0, Max: 1023},
	{ID: AxisR, Min: 0, Max: 1023},
	{ID: AxisF, Min: 0, Max: 255},
	{ID: AxisTX, Min: 0, Max: 255},
	{ID: AxisTY, Min: 0, Max: 255},
	{ID: AxisG, Min: 0, Max: 255},
	{ID: AxisRTY4, Min: 0, Max: 255},
	{ID: AxisRTY3, Min: 0, Max: 255},
}

var ButtonsConfig = []uinput.Button{
	{ID: ButtonE},
	{ID: ButtonF},
	{ID: ButtonG},
	{ID: ButtonI},
	{ID: ButtonH},
	{ID: ButtonSW1},
	{ID: ButtonSW2},
	{ID: ButtonSW3},
	{ID: ButtonSW4},
	{ID: ButtonSW5},
	{ID: ButtonSW6},
	{ID: ButtonTGL1Up},
	{ID: ButtonTGL1Down},
	{ID: ButtonTGL2Up},
	{ID: ButtonTGL2Down},
	{ID: ButtonTGL3Up},
	{ID: ButtonTGL3Down},
	{ID: ButtonTGL4Up},
	{ID: ButtonTGL4Down},
	{ID: ButtonH3Up},
	{ID: ButtonH3Right},
	{ID: ButtonH3Down},
	{ID: ButtonH3Left},
	{ID: ButtonH4Up},
	{ID: ButtonH4Right},
	{ID: ButtonH4Down},
	{ID: ButtonH4Left},
	{ID: ButtonK1Up},
	{ID: ButtonK1Down},
	{ID: ButtonScrollDown},
	{ID: ButtonScrollUp},
	{ID: ButtonThumb},
	{ID: ButtonSLD},
	{ID: ButtonModeM1},
	{ID: ButtonModeM2},
	{ID: ButtonModeS1},
}

func (s State) Map() (map[uint16]types.Axis, map[uint16]types.Button) {
	axes := map[uint16]types.Axis{
		AxisL:    s.AxisL,
		AxisR:    s.AxisR,
		AxisF:    s.AxisF,
		AxisTX:   s.AxisTX,
		AxisTY:   s.AxisTY,
		AxisG:    s.AxisG,
		AxisRTY4: s.AxisRTY4,
		AxisRTY3: s.AxisRTY3,
	}

	buttons := map[uint16]types.Button{
		ButtonE:          s.ButtonE,
		ButtonF:          s.ButtonF,
		ButtonG:          s.ButtonG,
		ButtonI:          s.ButtonI,
		ButtonH:          s.ButtonH,
		ButtonSW1:        s.ButtonSW1,
		ButtonSW2:        s.ButtonSW2,
		ButtonSW3:        s.ButtonSW3,
		ButtonSW4:        s.ButtonSW4,
		ButtonSW5:        s.ButtonSW5,
		ButtonSW6:        s.ButtonSW6,
		ButtonTGL1Up:     s.ButtonTGL1Up,
		ButtonTGL1Down:   s.ButtonTGL1Down,
		ButtonTGL2Up:     s.ButtonTGL2Up,
		ButtonTGL2Down:   s.ButtonTGL2Down,
		ButtonTGL3Up:     s.ButtonTGL3Up,
		ButtonTGL3Down:   s.ButtonTGL3Down,
		ButtonTGL4Up:     s.ButtonTGL4Up,
		ButtonTGL4Down:   s.ButtonTGL4Down,
		ButtonH3Up:       s.ButtonH3Up,
		ButtonH3Right:    s.ButtonH3Right,
		ButtonH3Down:     s.ButtonH3Down,
		ButtonH3Left:     s.ButtonH3Left,
		ButtonH4Up:       s.ButtonH4Up,
		ButtonH4Right:    s.ButtonH4Right,
		ButtonH4Down:     s.ButtonH4Down,
		ButtonH4Left:     s.ButtonH4Left,
		ButtonK1Up:       s.ButtonK1Up,
		ButtonK1Down:     s.ButtonK1Down,
		ButtonScrollDown: s.ButtonScrollDown,
		ButtonScrollUp:   s.ButtonScrollUp,
		ButtonThumb:      s.ButtonThumb,
		ButtonSLD:        s.ButtonSLD,
		ButtonModeM1:     s.ButtonModeM1,
		ButtonModeM2:     s.ButtonModeM2,
		ButtonModeS1:     s.ButtonModeS1,
	}

	return axes, buttons
}

func UpdateState(b []byte) {
	mux.Lock()
	defer mux.Unlock()

	b1 := b[1]
	b[1] = b[1] & 0b00000011
	state.AxisL = types.Axis(util.GetAxis(b[0:2]))
	// restore b1
	b[1] = b1

	b2 := b[2]
	// we're doing some fancy stuff here aren't we
	b[1] = b[1] & 0b11111100 >> 2
	b[2] = b[2] & 0b00001111
	b[1] = b[1] | b[2]&0b00000011<<6
	b[2] = b[2] & 0b00001100 >> 2
	state.AxisR = types.Axis(util.GetAxis(b[1:3]))
	// restore b2
	b[2] = b2

	state.ButtonI = types.Button(util.GetBool(b[2], 0))
	state.ButtonG = types.Button(util.GetBool(b[2], 1))
	state.ButtonF = types.Button(util.GetBool(b[2], 2))
	state.ButtonE = types.Button(util.GetBool(b[2], 3))

	// byte 3
	state.ButtonTGL1Up = types.Button(util.GetBool(b[3], 0))
	state.ButtonSW6 = types.Button(util.GetBool(b[3], 1))
	state.ButtonSW5 = types.Button(util.GetBool(b[3], 2))
	state.ButtonSW4 = types.Button(util.GetBool(b[3], 3))
	state.ButtonSW3 = types.Button(util.GetBool(b[3], 4))
	state.ButtonSW2 = types.Button(util.GetBool(b[3], 5))
	state.ButtonSW1 = types.Button(util.GetBool(b[3], 6))
	state.ButtonH = types.Button(util.GetBool(b[3], 7))

	// byte 4
	state.ButtonH3Up = types.Button(util.GetBool(b[4], 0))
	state.ButtonTGL4Down = types.Button(util.GetBool(b[4], 1))
	state.ButtonTGL4Up = types.Button(util.GetBool(b[4], 2))
	state.ButtonTGL3Down = types.Button(util.GetBool(b[4], 3))
	state.ButtonTGL3Up = types.Button(util.GetBool(b[4], 4))
	state.ButtonTGL2Down = types.Button(util.GetBool(b[4], 5))
	state.ButtonTGL2Up = types.Button(util.GetBool(b[4], 6))
	state.ButtonTGL1Down = types.Button(util.GetBool(b[4], 7))

	// byte 5 K1UP, H4LEFT, H4DOWN, H4RIGHT, H4UP, H3LEFT, H3DOWN, H3RIGHT
	state.ButtonK1Up = types.Button(util.GetBool(b[5], 0))
	state.ButtonH4Left = types.Button(util.GetBool(b[5], 1))
	state.ButtonH4Down = types.Button(util.GetBool(b[5], 2))
	state.ButtonH4Right = types.Button(util.GetBool(b[5], 3))
	state.ButtonH4Up = types.Button(util.GetBool(b[5], 4))
	state.ButtonH3Left = types.Button(util.GetBool(b[5], 5))
	state.ButtonH3Down = types.Button(util.GetBool(b[5], 6))
	state.ButtonH3Right = types.Button(util.GetBool(b[5], 7))

	// byte 6
	if util.GetBool(b[6], 0) {
		state.ButtonModeM1 = false
		state.ButtonModeM2 = false
		state.ButtonModeS1 = true
	} else if util.GetBool(b[6], 1) {
		state.ButtonModeM1 = false
		state.ButtonModeM2 = true
		state.ButtonModeS1 = false
	} else if util.GetBool(b[6], 2) {
		state.ButtonModeM1 = true
		state.ButtonModeM2 = false
		state.ButtonModeS1 = false
	}
	state.ButtonSLD = types.Button(util.GetBool(b[6], 3))
	state.ButtonThumb = types.Button(util.GetBool(b[6], 4))
	state.ButtonScrollUp = types.Button(util.GetBool(b[6], 5))
	state.ButtonScrollDown = types.Button(util.GetBool(b[6], 6))
	state.ButtonK1Down = types.Button(util.GetBool(b[6], 7))

	state.AxisF = types.Axis(b[7])
	state.AxisTX = types.Axis(b[8])
	state.AxisG = types.Axis(b[9])
	state.AxisTY = types.Axis(b[10])
	state.AxisRTY4 = types.Axis(b[11])
	state.AxisRTY3 = types.Axis(b[12])
}
