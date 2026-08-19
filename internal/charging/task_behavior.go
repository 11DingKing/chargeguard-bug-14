package charging

type LegacyStation struct {
	ID          string
	Assignments map[string]string
}

func RestoreLegacyStation(id string) *LegacyStation {
	station := &LegacyStation{ID: id}
	return station
}
func (s *LegacyStation) Assign(role, owner string) { s.Assignments[role] = owner }
func (s *LegacyStation) Owner(role string) string  { return s.Assignments[role] }
