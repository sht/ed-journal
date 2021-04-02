package types

const (
	ModeM1 Mode = iota
	ModeM2
	ModeS1
)

type Mode int

type Axis int

type Hat struct {
	Up    bool
	Right bool
	Down  bool
	Left  bool
}

type Button bool

type Switch bool
