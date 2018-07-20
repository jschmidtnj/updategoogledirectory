package main

import "fmt"

func configuredata() {
	//fmt.Println(worklocationdata[0].Entry1)
	//First sort everything:
	worklocationdata = quick_sort_worklocation(worklocationdata)
	//fmt.Println(worklocationdata[0].Entry1)
	organizationinfodata = quick_sort_organization(organizationinfodata)
	//fmt.Println(organizationinfodata[0].Entry1)
	employeeinfodata = quick_sort_employee(employeeinfodata)
	//fmt.Println(employeeinfodata[0].Entry2)
	fmt.Println("csv files sorted")
	Log.Println("csv files sorted")
}
