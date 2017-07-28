package main
// this doesn't work
func main() {
  sm := newStateMachine()
  sm.send(1) // "state A+1 => state B"
  sm.send(0) // "state B+0 => state C"
  println(sm.state()) // "state C"
}

type stateMachine struct{
  state string
}

func newStateMachine() stateMachine {
  return stateMachine{}
}

func (sm *stateMachine) send(action int) string {
  state := sm.state()
  newState := state
  return println("state %s + %d => %s", state, action, newState)
}
