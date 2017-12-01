package main

import "fmt"

type State struct {
	WizardAlive, WarriorAlive, PriestAlive, ThiefAlive bool
	WizardHere, WarriorHere, PriestHere, ThiefHere     bool
	NumGoblins                                         []int
	RoomNum                                            int
	WizardHasResurrected                               bool
}

func NewState() *State {
	s := &State{true, true, true, true, true, true, true, true, make([]int, 101), 1, false}
	for i := 1; i <= 100; i++ {
		s.NumGoblins[i] = i
	}
	return s
}

func (s *State) Kill(numGoblins int) int {
	if numGoblins > s.NumGoblins[s.RoomNum] {
		numGoblins = s.NumGoblins[s.RoomNum]
	}
	s.NumGoblins[s.RoomNum] -= numGoblins
	return numGoblins
}

func (s *State) AdventurersAlive() int {
	n := 0
	if s.WizardAlive {
		n++
	}
	if s.WarriorAlive {
		n++
	}
	if s.PriestAlive {
		n++
	}
	if s.ThiefAlive {
		n++
	}
	return n
}

func (s *State) GoblinsAlive() int {
	return s.NumGoblins[s.RoomNum]
}

func (s *State) AdventurersOutnumbered() bool {
	return s.GoblinsAlive() >= 10*s.AdventurersAlive()
}

func (s *State) NextRoom() {
	s.WizardHasResurrected = false
	if !s.WizardAlive {
		s.WizardHere = false
	}
	if !s.WarriorAlive {
		s.WarriorHere = false
	}
	if !s.PriestAlive {
		s.PriestHere = false
	}
	if !s.ThiefAlive {
		s.ThiefHere = false
	}
	s.RoomNum++
}

func (s *State) SimulateRoom() {
	fmt.Println("--- Start of room", s.RoomNum, "---")
	fmt.Printf("%#v\n", s)
	if s.ThiefAlive {
		fmt.Println("Thief kills", s.Kill(1), "goblins")
	}
	if s.WizardAlive {
		fmt.Println("Wizard kills", s.Kill(10), "goblins")
	}
	if s.WarriorAlive {
		fmt.Println("Warrior kills", s.Kill(1), "goblins")
	}
	if s.PriestAlive && s.WarriorHere && !s.WarriorAlive && !s.WizardHasResurrected {
		fmt.Println("Priest resurrected warrior")
		s.WarriorAlive = true
		s.WizardHasResurrected = true
	}
	if s.PriestAlive && s.WizardHere && !s.WizardAlive && !s.WizardHasResurrected {
		fmt.Println("Priest resurrected wizard")
		s.WarriorAlive = true
		s.WizardHasResurrected = true
	}
	if !s.WizardAlive && !s.WarriorAlive && !s.PriestAlive && s.ThiefAlive {
		fmt.Println("The others are dead so the thief sneaks along to the next room")
		s.NextRoom()
		return
	}
	if (s.WizardAlive || s.WarriorAlive || s.PriestAlive || s.ThiefAlive) && s.AdventurersOutnumbered() {
		fmt.Println("The adventurers are outnumbered")
		if s.WarriorAlive {
			fmt.Println("Goblins kill warrior")
			s.WarriorAlive = false
		} else if s.WizardAlive {
			fmt.Println("Goblins kill wizard")
			s.WizardAlive = false
		} else if s.PriestAlive {
			fmt.Println("Goblins kill priest")
			s.PriestAlive = false
		}
	}
	if s.AdventurersAlive() >= 1 && s.GoblinsAlive() >= 1 {
		fmt.Println("The battle isn't done yet")
		return
	}
	if s.AdventurersAlive() >= 1 && s.GoblinsAlive() == 0 {
		fmt.Println("Room cleared, on to the next one!")
		s.NextRoom()
		return
	}
	return
}

func main() {
	s := NewState()
	for s.AdventurersAlive() > 0 && s.RoomNum <= 100 {
		s.SimulateRoom()
	}
	fmt.Printf("%#v\n", s)
	n := s.AdventurersAlive()
	if s.RoomNum > 100 {
		n += 17
	}
	for i := 1; i <= 100; i++ {
		n += s.NumGoblins[i]
	}
	fmt.Println("Creatures who survived:", n)
}
