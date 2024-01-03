package myecspkg

type System interface {
	Update()
	Priority() int
}

func AddSystem(newSystem System) {
	insertIdx := len(systems)
	for i, system := range systems {
		if newSystem.Priority() > system.Priority() {
			insertIdx = i
			break
		}
	}

	if len(systems) == insertIdx {
		systems = append(systems, newSystem)
	} else {
		systems = append(systems[:insertIdx+1], systems[insertIdx:]...)
		systems[insertIdx] = newSystem
	}
}

func ClearSystems() {
    systems = make([]System, 0)
}

func Update() {
    for _, system := range systems {
        system.Update()
    }
}