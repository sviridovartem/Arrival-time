package backend

// IArrivalTimeManager interface of ArrivalTimeManager
type IArrivalTimeManager interface {
	GetLunarArrivalTime(string) (string, error)
}

// ArrivalTimeManager arrival time manager instance
type ArrivalTimeManager struct {
	core *InstanceCore
}

// NewArrivalTimeManager create a new arrival time manager  instance
func NewArrivalTimeManager(core *InstanceCore) IArrivalTimeManager {
	return &ArrivalTimeManager{
		core: core,
	}
}

// GetLunarArrivalTime returns lunar arrival time
func (atm *ArrivalTimeManager) GetLunarArrivalTime(userInfo string) (string, error) {

	return "GetLunarArrivalTime", nil
}
