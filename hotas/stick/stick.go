package stick

import (
	"github.com/sht/ed-journal/hotas/types"
	"github.com/sht/ed-journal/hotas/util"
	"sync"
)

var state = new(State)
var mux sync.RWMutex

type State struct {
	AxisX   types.Axis
	AxisY   types.Axis
	AxisZ   types.Axis
	AxisRX  types.Axis
	AxisRY  types.Axis
	POV     types.Hat
	H1      types.Hat
	H2      types.Hat
	A       types.Button
	B       types.Button
	C       types.Button
	D       types.Button // pinky
	Paddle  types.Button
	Trigger types.Button
}

func GetState() State {
	mux.RLock()
	defer mux.RUnlock()
	return *state
}

func UpdateState(b []byte) {
	mux.Lock()
	defer mux.Unlock()

	tmp := make([]byte, 2)

	state.AxisX = types.Axis(util.GetAxis(b[0:2]))
	state.AxisY = types.Axis(util.GetAxis(b[2:4]))

	b5 := b[5]
	b[5] = b[5] & 0b00001111
	state.AxisZ = types.Axis(util.GetAxis(b[4:6]))
	// restore b5
	b[5] = b5

	b[5] = b[5] & 0b11110000 >> 4
	// 0-8, 0 is off, goes top clockwise & corners
	state.POV.Up = b[5] == 8 || b[5] == 1 || b[5] == 2
	state.POV.Right = b[5] == 2 || b[5] == 3 || b[5] == 4
	state.POV.Down = b[5] == 4 || b[5] == 5 || b[5] == 6
	state.POV.Left = b[5] == 6 || b[5] == 7 || b[5] == 8

	b6 := b[6]
	b7 := b[7]
	b6 = b6 & 0b11000000 >> 6
	b7 = b7&0b00000011<<2 | b6
	// bitfield
	state.H1.Up = util.GetBool(tmp[1], 7)
	state.H1.Right = util.GetBool(tmp[1], 6)
	state.H1.Down = util.GetBool(tmp[1], 5)
	state.H1.Left = util.GetBool(tmp[1], 4)

	copy(tmp, b[7:8])
	tmp[0] = tmp[0] & 0b00111100 >> 2
	// bitfield
	state.H2.Up = util.GetBool(tmp[0], 7)
	state.H2.Right = util.GetBool(tmp[0], 6)
	state.H2.Down = util.GetBool(tmp[0], 5)
	state.H2.Left = util.GetBool(tmp[0], 4)

	copy(tmp, b[6:7])
	// bitfield
	state.Paddle = types.Button(util.GetBool(tmp[0], 2))
	state.D = types.Button(util.GetBool(tmp[0], 3))
	state.C = types.Button(util.GetBool(tmp[0], 4))
	state.B = types.Button(util.GetBool(tmp[0], 5))
	state.A = types.Button(util.GetBool(tmp[0], 6))
	state.Trigger = types.Button(util.GetBool(tmp[0], 7))

	// uint8 axes
	copy(tmp, b[9:11])
	state.AxisRX = types.Axis(tmp[0])
	state.AxisRY = types.Axis(tmp[1])
}
