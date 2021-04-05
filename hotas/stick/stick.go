package stick

import (
	"github.com/rdnt/uinput"
	"github.com/sht/ed-journal/hotas/types"
	"github.com/sht/ed-journal/hotas/util"
	"sync"
)

const (
	AxisX uint16 = iota
	AxisY
	AxisRX
	AxisRY
	AxisRZ
)

const (
	ButtonTrigger uint16 = iota + 0x120
	ButtonA
	ButtonB
	ButtonThumb
	ButtonD
	ButtonPinky
	ButtonH1Up
	ButtonH1Right
	ButtonH1Down
	ButtonH1Left
	ButtonH2Up
	ButtonH2Right
	ButtonH2Down
	ButtonH2Left
	ButtonPOVUp
	ButtonPOVRight
	ButtonPOVDown uint16 = iota + 0x2c0
	ButtonPOVLeft
)

var state = new(State)
var mux sync.RWMutex

var AxesConfig = []uinput.Axis{
	{ID: AxisX, Min: 0, Max: 65535},
	{ID: AxisY, Min: 0, Max: 65535},
	{ID: AxisRX, Min: 0, Max: 255},
	{ID: AxisRY, Min: 0, Max: 255},
	{ID: AxisRZ, Min: 0, Max: 4095},
}

var ButtonsConfig = []uinput.Button{
	{ID: ButtonTrigger},
	{ID: ButtonA},
	{ID: ButtonB},
	{ID: ButtonThumb},
	{ID: ButtonD},
	{ID: ButtonPinky},
	{ID: ButtonH1Up},
	{ID: ButtonH1Right},
	{ID: ButtonH1Down},
	{ID: ButtonH1Left},
	{ID: ButtonH2Up},
	{ID: ButtonH2Right},
	{ID: ButtonH2Down},
	{ID: ButtonH2Left},
	{ID: ButtonPOVUp},
	{ID: ButtonPOVRight},
	{ID: ButtonPOVDown},
	{ID: ButtonPOVLeft},
}

type State struct {
	AxisX  types.Axis
	AxisY  types.Axis
	AxisRX types.Axis
	AxisRY types.Axis
	AxisRZ types.Axis

	ButtonTrigger  types.Button
	ButtonA        types.Button
	ButtonB        types.Button
	ButtonThumb    types.Button
	ButtonD        types.Button
	ButtonPinky    types.Button
	ButtonH1Up     types.Button
	ButtonH1Right  types.Button
	ButtonH1Down   types.Button
	ButtonH1Left   types.Button
	ButtonH2Up     types.Button
	ButtonH2Right  types.Button
	ButtonH2Down   types.Button
	ButtonH2Left   types.Button
	ButtonPOVUp    types.Button
	ButtonPOVRight types.Button
	ButtonPOVDown  types.Button
	ButtonPOVLeft  types.Button

	//POVUp    types.Button
	//POVRight types.Button
	//POVDown  types.Button
	//POVLeft  types.Button
	//H1Up     types.Button
	//H1Right  types.Button
	//H1Down   types.Button
	//H1Left   types.Button
	//H2Up     types.Button
	//H2Right  types.Button
	//H2Down   types.Button
	//H2Left   types.Button
	//A        types.Button
	//B        types.Button
	//C        types.Button
	//D        types.Button // pinky
	//Paddle   types.Button
	//Trigger  types.Button
}

func GetState() State {
	mux.RLock()
	defer mux.RUnlock()
	return *state
}

func (s State) Map() (map[uint16]types.Axis, map[uint16]types.Button) {
	axes := map[uint16]types.Axis{
		AxisX:  s.AxisX,
		AxisY:  s.AxisY,
		AxisRX: s.AxisRX,
		AxisRY: s.AxisRY,
		AxisRZ: s.AxisRZ,
	}

	buttons := map[uint16]types.Button{
		ButtonTrigger:  s.ButtonTrigger,
		ButtonA:        s.ButtonA,
		ButtonB:        s.ButtonB,
		ButtonThumb:    s.ButtonThumb,
		ButtonD:        s.ButtonD,
		ButtonPinky:    s.ButtonPinky,
		ButtonH1Up:     s.ButtonH1Up,
		ButtonH1Right:  s.ButtonH1Right,
		ButtonH1Down:   s.ButtonH1Down,
		ButtonH1Left:   s.ButtonH1Left,
		ButtonH2Up:     s.ButtonH2Up,
		ButtonH2Right:  s.ButtonH2Right,
		ButtonH2Down:   s.ButtonH2Down,
		ButtonH2Left:   s.ButtonH2Left,
		ButtonPOVUp:    s.ButtonPOVUp,
		ButtonPOVRight: s.ButtonPOVRight,
		ButtonPOVDown:  s.ButtonPOVDown,
		ButtonPOVLeft:  s.ButtonPOVLeft,
	}

	return axes, buttons
}

func UpdateState(b []byte) {
	mux.Lock()
	defer mux.Unlock()

	tmp := make([]byte, 2)

	state.AxisX = types.Axis(util.GetAxis(b[0:2]))
	state.AxisY = types.Axis(util.GetAxis(b[2:4]))

	b5 := b[5]
	b[5] = b[5] & 0b00001111
	state.AxisRZ = types.Axis(util.GetAxis(b[4:6]))
	// restore b5
	b[5] = b5

	b[5] = b[5] & 0b11110000 >> 4
	// 0-8, 0 is off, goes top clockwise & corners
	state.ButtonPOVUp = b[5] == 8 || b[5] == 1 || b[5] == 2
	state.ButtonPOVRight = b[5] == 2 || b[5] == 3 || b[5] == 4
	state.ButtonPOVDown = b[5] == 4 || b[5] == 5 || b[5] == 6
	state.ButtonPOVLeft = b[5] == 6 || b[5] == 7 || b[5] == 8

	b6 := b[6]
	b7 := b[7]
	b6 = b6 & 0b11000000 >> 6
	b7 = b7&0b00000011<<2 | b6
	// bitfield
	state.ButtonH1Up = types.Button(util.GetBool(tmp[1], 7))
	state.ButtonH1Right = types.Button(util.GetBool(tmp[1], 6))
	state.ButtonH1Down = types.Button(util.GetBool(tmp[1], 5))
	state.ButtonH1Left = types.Button(util.GetBool(tmp[1], 4))

	copy(tmp, b[7:8])
	tmp[0] = tmp[0] & 0b00111100 >> 2
	// bitfield
	state.ButtonH2Up = types.Button(util.GetBool(tmp[0], 7))
	state.ButtonH2Right = types.Button(util.GetBool(tmp[0], 6))
	state.ButtonH2Down = types.Button(util.GetBool(tmp[0], 5))
	state.ButtonH2Left = types.Button(util.GetBool(tmp[0], 4))

	copy(tmp, b[6:7])
	// bitfield
	state.ButtonPinky = types.Button(util.GetBool(tmp[0], 2))
	state.ButtonD = types.Button(util.GetBool(tmp[0], 3))
	state.ButtonThumb = types.Button(util.GetBool(tmp[0], 4))
	state.ButtonB = types.Button(util.GetBool(tmp[0], 5))
	state.ButtonA = types.Button(util.GetBool(tmp[0], 6))
	state.ButtonTrigger = types.Button(util.GetBool(tmp[0], 7))

	// uint8 axes
	copy(tmp, b[9:11])
	state.AxisRX = types.Axis(tmp[0])
	state.AxisRY = types.Axis(tmp[1])
}
