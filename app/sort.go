package main

import (
	"fmt"
	"strconv"
)

func quick_sort_worklocation(data []Worklocation) []Worklocation {
	array := data
	var less []Worklocation
	var equal []Worklocation
	var greater []Worklocation
	if len(array) > 1 {
		pivot := array[0].Entry1
		for _, item := range array {
			item_index := item.Entry1
			if item_index < pivot {
				less = append(less, item)
			} else if item_index == pivot {
				equal = append(equal, item)
			} else {
				greater = append(greater, item)
			}
		}
		//join lists with +
		return append(append(quick_sort_worklocation(less), equal...), quick_sort_worklocation(greater)...)
	} else {
		return array
	}
}

func quick_sort_organization(data []Organizationinfo) []Organizationinfo {
	array := data
	var less []Organizationinfo
	var equal []Organizationinfo
	var greater []Organizationinfo
	if len(array) > 1 {
		pivot, err := strconv.Atoi(array[0].Entry1)
		if err != nil {
			fmt.Printf("int convert error for organization array: %v\n", err)
			Log.Printf("int convert error for organization array: %v\n", err)
			panic(err)
		}
		for _, item := range array {
			item_index, err := strconv.Atoi(item.Entry1)
			if err != nil {
				fmt.Printf("int convert error for organization item: %v\n", err)
				Log.Printf("int convert error for organization item: %v\n", err)
				panic(err)
			}
			if item_index < pivot {
				less = append(less, item)
			} else if item_index == pivot {
				equal = append(equal, item)
			} else {
				greater = append(greater, item)
			}
		}
		//join lists with +
		return append(append(quick_sort_organization(less), equal...), quick_sort_organization(greater)...)
	} else {
		return array
	}
}

func quick_sort_employee(data []Employeeinfo) []Employeeinfo {
	array := data
	var less []Employeeinfo
	var equal []Employeeinfo
	var greater []Employeeinfo
	if len(array) > 1 {
		pivot, err := strconv.Atoi(array[0].Entry2)
		if err != nil {
			fmt.Printf("int convert error for employee array: %v\n", err)
			Log.Printf("int convert error for employee array: %v\n", err)
			panic(err)
		}
		for _, item := range array {
			item_index, err := strconv.Atoi(item.Entry2)
			if err != nil {
				fmt.Printf("int convert error for employee item: %v\n")
				Log.Printf("int convert error for employee item: %v\n")
				panic(err)
			}
			if item_index < pivot {
				less = append(less, item)
			} else if item_index == pivot {
				equal = append(equal, item)
			} else {
				greater = append(greater, item)
			}
		}
		//join lists with +
		return append(append(quick_sort_employee(less), equal...), quick_sort_employee(greater)...)
	} else {
		return array
	}
}

func quick_sort_country(data []country) []country {
	array := data
	var less []country
	var equal []country
	var greater []country
	if len(array) > 1 {
		pivot := array[0].Alpha2code
		for _, item := range array {
			item_index := item.Alpha2code
			if item_index < pivot {
				less = append(less, item)
			} else if item_index == pivot {
				equal = append(equal, item)
			} else {
				greater = append(greater, item)
			}
		}
		//join lists with +
		return append(append(quick_sort_country(less), equal...), quick_sort_country(greater)...)
	} else {
		return array
	}
}
