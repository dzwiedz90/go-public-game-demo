package states

type GameState struct {
	StateName string
	OnEnter   func()
	OnUpdate  func()
	OnExit    func()
}

var (
	CurrentGameState  GameState
	PreviousGameState GameState
)
