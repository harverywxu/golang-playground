package combination

import (
	"fmt"
	"testing"
)

type Base struct {
	name string
}

func (b Base) GetName() string {
	return b.name
}

func (b *Base) SetName(new string) {
	b.name = new
}

type StudentS struct {
	Base
	game string
}

func (s StudentS) GetGame() string {
	return s.game
}

func (s *StudentS) SetGame(new string) {
	s.game = new
}

type ConInterface interface {
	GetName() string
	SetName(name string)
	GetGame() string
	SetGame(game string)
}

func ConInter(c ConInterface)  {
	c.SetName("hello name")
	// hello name
	fmt.Println(c.GetName())
	// hello game
	c.SetGame("hello game")
	fmt.Println(c.GetGame())

}

func TestConFun(t *testing.T) {
	b := Base{name:"base"}
	// base
	t.Logf(b.GetName())
	b.SetName("new base")
	// new base
	t.Logf(b.GetName())

	s := StudentS{
		Base{name:"base"},
		"game",
	}
	s.SetName("new s name")
	// new s name
	t.Logf(s.GetName())
	s.SetGame("new s game")
	// new s game
	t.Logf(s.GetGame())

	// ConInter(s)
	// Cannot use 's' (type StudentS) as type ConInterface, Type does not implement 'ConInterface' as 'SetName' method has a pointer receiver
	// Type StudentS dose not implement method SetName/SetGame
	// Type *StudentS implements method SetName/SetGame
	ConInter(&s)
}
