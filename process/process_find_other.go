// +build !darwin

package process

//PidExists check is pid exist  change by feng ignore Z status
func PidExists(pid int32) (bool, error) {
	pids, err := Pids()
	if err != nil {
		return false, err
	}

	for _, i := range pids {
		if i == pid {
			return true, err
		}
	}

	return false, err
}
