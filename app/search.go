package main

func binary_search_employee(data []Employeeinfo, target string) int {
	start_index := 0
	end_index := len(data) - 1

	for start_index <= end_index {
		median := (start_index + end_index) / 2
		if data[median].Entry2 < target {
			start_index = median + 1
		} else {
			end_index = median - 1
		}
	}

	if start_index == len(data) || data[start_index].Entry2 != target {
		return -1
	} else {
		return start_index
	}
}

func binary_search_worklocation(data []Worklocation, target string) int {
	start_index := 0
	end_index := len(data) - 1

	for start_index <= end_index {
		median := (start_index + end_index) / 2
		if data[median].Entry1 < target {
			start_index = median + 1
		} else {
			end_index = median - 1
		}
	}

	if start_index == len(data) || data[start_index].Entry1 != target {
		return -1
	} else {
		return start_index
	}
}

func binary_search_organization(data []Organizationinfo, target string) int {
	start_index := 0
	end_index := len(data) - 1

	for start_index <= end_index {
		median := (start_index + end_index) / 2
		if data[median].Entry1 < target {
			start_index = median + 1
		} else {
			end_index = median - 1
		}
	}

	if start_index == len(data) || data[start_index].Entry1 != target {
		return -1
	} else {
		return start_index
	}
}
