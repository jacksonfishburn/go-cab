package files

type Manager struct {

}


func (m *Manager) Ping() bool {
	return true	
}

func (m *Manager) Add() string {
	return "add"	
}

func (m *Manager) Grab() string {
	return "grab"	
}

func (m *Manager) Del() string {
	return "delete"	
}

func (m *Manager) Peek() string {
	return "peek"	
}