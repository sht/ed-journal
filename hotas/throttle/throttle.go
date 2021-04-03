package throttle

import (
	"github.com/sht/ed-journal/hotas/types"
	"github.com/sht/ed-journal/hotas/util"
	"sync"
)

var state = new(State)
var mux sync.RWMutex

// State represents the state of the throttle device
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
	ThrottleL types.Axis // 0-1023, deadzone at 511, if connected with ThrottleR then ThrottleR is synced with ThrottleL
	ThrottleR types.Axis // 0-1023
	Thumb     struct {
		types.Button
		X types.Axis // 0-255, resets at 128
		Y types.Axis // 0-255, resets at 128
	}
	RTY3 types.Axis // 0-255
	RTY4 types.Axis // 0-255
	E    types.Button
	F    struct {
		types.Axis // 0-255, floats around 127 on reset (+-5)
		types.Button
	}
	G struct {
		types.Axis // 0-255, floats around 127 on reset (+-5)
		types.Button
	}
	H3         types.Hat
	H4         types.Hat
	Mode       types.Mode
	SLD        types.Switch
	H          types.Button
	I          types.Button
	SW1        types.Button
	SW2        types.Button
	SW3        types.Button
	SW4        types.Button
	SW5        types.Button
	SW6        types.Button
	TGL1Up     types.Button
	TGL1Down   types.Button
	TGL2Up     types.Button
	TGL2Down   types.Button
	TGL3Up     types.Button
	TGL3Down   types.Button
	TGL4Up     types.Button
	TGL4Down   types.Button
	K1Down     types.Button
	K1Up       types.Button
	ScrollUp   types.Button
	ScrollDown types.Button
}

func GetState() State {
	mux.RLock()
	defer mux.RUnlock()
	return *state
}

func UpdateState(b []byte) {
	mux.Lock()
	defer mux.Unlock()

	b1 := b[1]
	b[1] = b[1] & 0b00000011
	state.ThrottleL = types.Axis(util.GetAxis(b[0:2]))
	// restore b1
	b[1] = b1

	b2 := b[2]
	// we're doing some fancy stuff here aren't we
	b[1] = b[1] & 0b11111100 >> 2
	b[2] = b[2] & 0b00001111
	b[1] = b[1] | b[2]&0b00000011<<6
	b[2] = b[2] & 0b00001100 >> 2
	state.ThrottleR = types.Axis(util.GetAxis(b[1:3]))
	// restore b2
	b[2] = b2

	state.I = types.Button(util.GetBool(b[2], 0))
	state.G.Button = types.Button(util.GetBool(b[2], 1))
	state.F.Button = types.Button(util.GetBool(b[2], 2))
	state.E = types.Button(util.GetBool(b[2], 3))

	// byte 3
	state.TGL1Up = types.Button(util.GetBool(b[3], 0))
	state.SW6 = types.Button(util.GetBool(b[3], 1))
	state.SW5 = types.Button(util.GetBool(b[3], 2))
	state.SW4 = types.Button(util.GetBool(b[3], 3))
	state.SW3 = types.Button(util.GetBool(b[3], 4))
	state.SW2 = types.Button(util.GetBool(b[3], 5))
	state.SW1 = types.Button(util.GetBool(b[3], 6))
	state.H = types.Button(util.GetBool(b[3], 7))

	// byte 4
	state.H3.Up = util.GetBool(b[4], 0)
	state.TGL4Down = types.Button(util.GetBool(b[4], 1))
	state.TGL4Up = types.Button(util.GetBool(b[4], 2))
	state.TGL3Down = types.Button(util.GetBool(b[4], 3))
	state.TGL3Up = types.Button(util.GetBool(b[4], 4))
	state.TGL2Down = types.Button(util.GetBool(b[4], 5))
	state.TGL2Up = types.Button(util.GetBool(b[4], 6))
	state.TGL1Down = types.Button(util.GetBool(b[4], 7))

	// byte 5 K1UP, H4LEFT, H4DOWN, H4RIGHT, H4UP, H3LEFT, H3DOWN, H3RIGHT
	state.K1Up = types.Button(util.GetBool(b[5], 0))
	state.H4.Left = util.GetBool(b[5], 1)
	state.H4.Down = util.GetBool(b[5], 2)
	state.H4.Right = util.GetBool(b[5], 3)
	state.H4.Up = util.GetBool(b[5], 4)
	state.H3.Left = util.GetBool(b[5], 5)
	state.H3.Down = util.GetBool(b[5], 6)
	state.H3.Right = util.GetBool(b[5], 7)

	// byte 6
	if util.GetBool(b[6], 0) {
		state.Mode = types.ModeS1
	} else if util.GetBool(b[6], 1) {
		state.Mode = types.ModeM2
	} else if util.GetBool(b[6], 2) {
		state.Mode = types.ModeM1
	}
	state.SLD = types.Switch(util.GetBool(b[6], 3))
	state.Thumb.Button = types.Button(util.GetBool(b[6], 4))
	state.ScrollUp = types.Button(util.GetBool(b[6], 5))
	state.ScrollDown = types.Button(util.GetBool(b[6], 6))
	state.K1Down = types.Button(util.GetBool(b[6], 7))

	state.F.Axis = types.Axis(b[7])
	state.Thumb.X = types.Axis(b[8])
	state.G.Axis = types.Axis(b[9])
	state.Thumb.Y = types.Axis(b[10])
	state.RTY4 = types.Axis(b[11])
	state.RTY3 = types.Axis(b[12])
}
