// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Address == nil || len(r.Address) == 0 || len(r.Name) == 0 {
		return false
	}

	for key, val := range r.Address {
		if len(key) == 0 || len(val) == 0 || key != "street" {
			return false
		}
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Age = 0
	r.Name = ""
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var counter int
	for idx := range residents {
		if resident := residents[idx]; resident.HasRequiredInfo() {
			counter++
		}
	}
	return counter
}
