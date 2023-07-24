package main

func getUniqueOnly(schedules *[]CAUSchedule) []CAUSchedule {
	unique := make([]CAUSchedule, len(*schedules))
	count := 0

	for i := 0; i < len(*schedules); i++ {
		schedule := (*schedules)[i]
		duplicate := false
		for j := 0; j < len(unique); j++ {
			if unique[j].Title == schedule.Title &&
				unique[j].StartDate.Equal(schedule.StartDate) &&
				unique[j].EndDate.Equal(schedule.EndDate) {
				duplicate = true
				break
			}
		}

		if duplicate {
			continue
		} else {
			unique[count] = schedule
			count++
		}
	}

	return unique[:count]
}
