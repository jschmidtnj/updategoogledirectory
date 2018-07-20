package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hako/durafmt" //github.com/hako/durafmt
	"github.com/stretchr/objx"
	"google.golang.org/api/admin/directory/v1"
)

func getManagerUserList(adminSvc *admin.Service, nextPageToken string, managerID string) (*admin.Users, error) {
	//Drive api Example see https://github.com/google/google-api-go-client/issues/227
	var query string
	query = "externalId=" + "GP" + managerID
	//fmt.Println(query)
	return adminSvc.Users.List().Customer("my_customer").Query(query).MaxResults(500).OrderBy("email").PageToken(nextPageToken).Do()
}

func getDirectoryUserList(adminSvc *admin.Service, nextPageToken string) (*admin.Users, error) {
	//Drive api Example see https://github.com/google/google-api-go-client/issues/227
	return adminSvc.Users.List().Customer("my_customer").MaxResults(500).OrderBy("email").PageToken(nextPageToken).Do()
}

func apirequest() {
	fmt.Println("starting api requests")
	Log.Println("starting api requests")
	//fmt.Println(worklocationdata[0].Entry1)
	//complete api request call with configured data
	nextPageToken := ""
	userList, err := getDirectoryUserList(directoryapiservice, nextPageToken)
	if err != nil {
		panic(err)
	}
	//fmt.Println(userList.NextPageToken)
	totalcount := len(employeeinfodata)
	var count = 0
	var push_error_count = 0
	var get_data_error_count = 0
	var get_employee_data_error_count = 0
	var get_worklocation_data_error_count = 0
	var get_organization_data_error_count = 0
	var get_managers_error_count = 0
	var durations []time.Duration
	var sum time.Duration
	total_start_time := time.Now()
	for userList.NextPageToken != "" {
		if err != nil {
			Log.Printf("Unable to retrieve users: %v", err)
			fmt.Printf("Unable to retrieve users: %v", err)
			panic(err)
		}
		for _, user := range userList.Users {
			start_time := time.Now()
			count++
			if user.ExternalIds != nil {

				var externalIdValue = ""
				//get external Id data
				externalIdData := user.ExternalIds.([]interface{})[0]
				externalId := objx.New(externalIdData)["value"].(string)
				//fmt.Println(user.PrimaryEmail, externalId["value"])

				if externalId != "" {
					externalIdValue = externalId[2:]
				} else {
					continue
				}

				//if externalIdValue == "00049571" { //           COMMENT THIS OUT
				//fmt.Printf("%v\n\n", user)
				//fmt.Printf("%v\n\n", user.CustomSchemas)
				//fmt.Printf(externalId)
				//get address data
				var AddressDataArray []objx.Map
				var address objx.Map
				if user.Addresses != nil {
					AddressDataArrayInterface := user.Addresses.([]interface{})
					//fmt.Println(len(AddressDataArrayInterface))
					for _, addr := range AddressDataArrayInterface {
						address = objx.New(addr)
						if address["postalCode"] != nil {
							//fmt.Println("current postalCode " + address["postalCode"].(string))
						}
						AddressDataArray = append(AddressDataArray, address)
					}
					//fmt.Println(len(AddressDataArray))
				} else {
					//fmt.Println("no addresses")
					address = objx.MustFromJSON(`{}`) //{"type": "", "customType": "", "formatted": "", "poBox": "", "extendedAddress": "", "streetAddress": "", "locality": "", "region": "", "postalCode": "", "country": "", "countryCode": ""}
					//fmt.Printf("%v", address)
					AddressDataArray = append(AddressDataArray, address)
				}

				//get organization data
				var OrganizationDataArray []objx.Map
				var organization objx.Map
				if user.Organizations != nil {
					OrganizationDataArrayInterface := user.Organizations.([]interface{})
					//fmt.Println(len(OrganizationDataArrayInterface))
					for _, org := range OrganizationDataArrayInterface {
						organization = objx.New(org)
						if organization["location"] != nil {
							//fmt.Println("current location " + organization["location"].(string))
						}
						if organization["name"] != nil {
							//fmt.Println("name: " + organization["name"].(string))
						}
						OrganizationDataArray = append(OrganizationDataArray, organization)
					}
				} else {
					//fmt.Println("no organizations")
					organization = objx.MustFromJSON(`{}`) //{"name": "", "title": "", "type": "", "customType": "", "department": "", "symbol": "", "location": "", "description": "", "domain": "", "costCenter": ""}
					//fmt.Printf("%v", organization)
					OrganizationDataArray = append(OrganizationDataArray, organization)
				}

				//get location data
				var LocationDataArray []objx.Map
				var location objx.Map
				if user.Locations != nil {
					LocationDataArrayInterface := user.Locations.([]interface{})
					//fmt.Println(len(LocationDataArrayInterface))
					for _, loc := range LocationDataArrayInterface {
						location = objx.New(loc)
						if location["area"] != nil {
							//fmt.Println("current location area " + location["area"].(string))
						}
						if location["type"] != nil {
							//fmt.Println("location type: " + location["type"].(string))
						}
						if location["customType"] != nil {
							//fmt.Println("custom type: " + location["customType"].(string))
						}
						LocationDataArray = append(LocationDataArray, location)
					}
				} else {
					//fmt.Println("no locations")
					location = objx.MustFromJSON(`{}`) //{"type": "", "customType": "", "area": "", "buildingId": "", "floorName": "", "floorSection": "", "deskCode": ""}
					//fmt.Printf("%v", location)
					//location["deskCode"] = "12345"
					//fmt.Println("testing deskcode " + location["deskCode"].(string))
					LocationDataArray = append(LocationDataArray, location)
				}

				//get relations data
				var RelationsDataArray []objx.Map
				var relation objx.Map

				if user.Relations != nil {
					RelationsDataArrayInterface := user.Relations.([]interface{})
					//fmt.Println(len(RelationsDataArrayInterface))
					//fmt.Printf("%v", RelationsDataArrayInterface)
					for _, rela := range RelationsDataArrayInterface {
						relation = objx.New(rela)
						if relation["value"] != nil {
							//fmt.Println("current location " + relation["value"].(string))
						}
						if organization["type"] != nil {
							//fmt.Println("type: " + organization["type"].(string))
						}
						if organization["customType"] != nil {
							//fmt.Println("custom type: " + organization["customType"].(string))
						}
						RelationsDataArray = append(RelationsDataArray, relation)
					}
				} else {
					//fmt.Println("no relations")
					relation = objx.MustFromJSON(`{}`) //{"value": "", "type": "", "customType": ""}
					//fmt.Printf("%v", relation)
					RelationsDataArray = append(RelationsDataArray, relation)
				}

				//get phone data
				var PhonesDataArray []objx.Map
				var phone objx.Map
				/*
					if user.Phones != nil {
						PhonesDataArrayInterface := user.Phones.([]interface{})
						//fmt.Println(len(RelationsDataArrayInterface))
						//fmt.Printf("%v", RelationsDataArrayInterface)
						for _, pho := range PhonesDataArrayInterface {
							phone = objx.New(pho)
							if phone["value"] != nil {
								//fmt.Println("main phone " + phone["value"].(string))
							}
							PhonesDataArray = append(PhonesDataArray, phone)
						}
					} else {
						//fmt.Println("no phones")
						phone = objx.MustFromJSON(`{}`) //{"value": "", "primary": false, "type": "", "customType": ""}
						//fmt.Printf("%v", relation)
						PhonesDataArray = append(PhonesDataArray, relation)
					}
				*/

				var addressFormatted string
				var addressExtended string
				employee_index := binary_search_employee(employeeinfodata, externalIdValue)
				if employee_index == -1 {
					get_employee_data_error_count++
					get_data_error_count++
					//fmt.Println("employee index not found count " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					Log.Println("employee index not found user count " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					//panic("employee index not found user count " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					continue
				}
				//fmt.Println(employee_index)
				employee := employeeinfodata[employee_index]
				work_index := binary_search_worklocation(worklocationdata, employee.Entry15)
				if work_index == -1 {
					get_worklocation_data_error_count++
					get_data_error_count++
					//fmt.Println("work index not found user " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					Log.Println("work index not found user " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					//panic("work index not found user " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					continue
				}
				//fmt.Println(work_index)
				workLocation := worklocationdata[work_index]
				organization_index := binary_search_organization(organizationinfodata, employee.Entry5)
				if organization_index == -1 {
					get_organization_data_error_count++
					get_data_error_count++
					//fmt.Println("organization index not found user " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					Log.Println("organization index not found user " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					//panic("organization index not found user " + strconv.Itoa(count) + " with current error count " + strconv.Itoa(get_data_error_count))
					continue
				}
				//fmt.Println(organization_index)
				organizationDatapoint := organizationinfodata[organization_index]
				//for _, employee := range employeeinfodata {
				//if employee.Entry2 == externalIdValue {
				//update from employee data
				//for _, workLocation := range worklocationdata {
				//update from worklocation data
				//fmt.Println(workLocation.Entry1, employee.Entry15)
				//if workLocation.Entry1 == employee.Entry15 {
				//fmt.Println(len(organizationinfodata))
				//for _, organizationDatapoint := range organizationinfodata {
				//fmt.Println(organizationDatapoint.Entry1, employee.Entry5)
				//if organizationDatapoint.Entry1 == employee.Entry5 {
				//update the organizationinfo data
				addressFormatted = workLocation.Entry3 + ", " + workLocation.Entry5 + ", " + workLocation.Entry6 + " " + employee.Entry13 + " " + workLocation.Entry7
				addressExtended = addressFormatted

				for i, addr := range AddressDataArray {
					address = addr
					add_to_array := false
					if workLocation.Entry2 != "" {
						address["extendedAddress"] = workLocation.Entry2
						add_to_array = true
					}
					if workLocation.Entry3 != "" {
						address["streetAddress"] = workLocation.Entry3
						add_to_array = true
					}
					if workLocation.Entry5 != "" {
						address["locality"] = workLocation.Entry5
						add_to_array = true
					}
					if workLocation.Entry6 != "" {
						address["region"] = workLocation.Entry6
						add_to_array = true
					}
					if employee.Entry13 != "" {
						address["countryCode"] = employee.Entry13
						add_to_array = true
					}
					if workLocation.Entry7 != "" {
						address["postalCode"] = workLocation.Entry7
						add_to_array = true
					}
					if addressFormatted != "" {
						address["formatted"] = addressFormatted
						add_to_array = true
					}
					if workLocation.Entry4 != "" {
						address["poBox"] = workLocation.Entry4
						add_to_array = true
					}
					if addressExtended != "" {
						address["extendedAddress"] = addressExtended
						add_to_array = true
					}
					if add_to_array {
						address["type"] = "work"
						AddressDataArray[i] = address
					} else {
						AddressDataArray = append(AddressDataArray[:i], AddressDataArray[i+1:]...)
					}
				}
				if len(AddressDataArray) > 0 {
					user.Addresses = AddressDataArray
				}

				for i, org := range OrganizationDataArray {
					organization = org
					add_to_array := false
					if employee.Entry6 != "" {
						organization["title"] = employee.Entry6
						add_to_array = true
					}
					if organizationDatapoint.Entry2 != "" {
						organization["department"] = organizationDatapoint.Entry2
						add_to_array = true
					}
					if addressFormatted != "" {
						organization["location"] = addressFormatted
						organization["formatted"] = addressFormatted
						add_to_array = true
					}
					if getConfig.Test_String != "" {
						organization["description"] = "" + getConfig.Test_String
						add_to_array = true
					}
					if add_to_array {
						OrganizationDataArray[i] = organization
					} else {
						OrganizationDataArray = append(OrganizationDataArray[:i], OrganizationDataArray[i+1:]...)
					}
				}

				if len(OrganizationDataArray) > 0 {
					user.Organizations = OrganizationDataArray
				}

				for i, loc := range LocationDataArray {
					location = loc
					add_to_array := false
					if employee.Entry11 != "" {
						location["deskCode"] = employee.Entry11
					}
					if add_to_array {
						location["type"] = "desk"
						LocationDataArray[i] = location
					} else {
						LocationDataArray = append(LocationDataArray[:i], LocationDataArray[i+1:]...)
					}
				}
				if len(LocationDataArray) > 0 {
					user.Locations = LocationDataArray
				}
				var managers []string
				//emp is the manager information
				//this for loop is not necessary, unless more information about the manager from the csv file is needed
				//for _, emp := range employeeinfodata {
				//	if employee.Entry10 == emp.Entry2 {
				nextPageTokenManagers := ""
				ManagerUserList, err := getManagerUserList(directoryapiservice, nextPageTokenManagers, employee.Entry10)
				if err != nil {
					get_managers_error_count++
					Log.Printf("Unable to retrieve manager user[s]: %v", err)
					//fmt.Printf("Unable to retrieve manager user[s]: %v", err)
					//panic(err)
				}
				for _, manageruser := range ManagerUserList.Users {
					if manageruser.ExternalIds != "" {
						//fmt.Println("got manager " + manageruser.Name.FullName)
						managers = append(managers, manageruser.Name.FullName)
					}
				}
				//fmt.Printf("manager data returned: %v", ManagerUserList.Users)
				//fmt.Printf("page token for manager data: %v", ManagerUserList.NextPageToken)
				for ManagerUserList.NextPageToken != "" {
					if err != nil {
						get_managers_error_count++
						get_data_error_count++
						Log.Printf("Unable to retrieve manager user[s]: %v", err)
						//fmt.Printf("Unable to retrieve manager user[s]: %v", err)
						//panic(err)
					}
					for _, manageruser := range ManagerUserList.Users {
						//fmt.Println("got manager " + manageruser.Name.FullName)
						managers = append(managers, manageruser.Name.FullName)
					}
					fmt.Printf("%v", ManagerUserList.Users)
					ManagerUserList, err = getManagerUserList(directoryapiservice, nextPageTokenManagers, employee.Entry10)
				}
				//	}
				//}
				for i, manager := range managers {
					relation["value"] = manager
					relation["type"] = "manager"
					if i >= len(RelationsDataArray) {
						RelationsDataArray = append(RelationsDataArray, relation)
					} else {
						RelationsDataArray[i] = relation
					}
				}
				user.Relations = RelationsDataArray
				for i, phon := range PhonesDataArray {
					phone = phon
					add_to_array := false
					if employee.Entry12 != "" {
						phone["value"] = employee.Entry12
						add_to_array = true
					}
					if add_to_array {
						if i == 0 {
							phone["primary"] = true //put first phone number as primary
						} else {
							phone["primary"] = false
						}
						phone["type"] = "work"
						PhonesDataArray[i] = phone
					} else {
						PhonesDataArray = append(PhonesDataArray[:i], PhonesDataArray[i+1:]...)
					}
				}
				if len(PhonesDataArray) > 0 {
					user.Phones = PhonesDataArray
				}
				//change name data
				//user.Name.GivenName = employee.Entry4
				//user.Name.FullName = employee.Entry3
				//break
				//}
				//}
				//break
				//}
				//}
				//break
				//}
				//}
				//refresh address data
				//address["postalCode"] = ""
				//var updatedaddress interface{}
				//updatedaddress = address
				//AddressDataArray[0] = updatedaddress
				//user.Addresses = AddressDataArray

				//refresh organization data
				//organization["location"] = "asdf"
				//var updatedorganization interface{}
				//updatedorganization = organization
				//OrganizationDataArray[1] = updatedorganization
				//user.Organizations = OrganizationDataArray

				//user.Locations = LocationDataArray

				//user.Relations = RelationsDataArray

				//user.Phones = PhonesDataArray

				id := user.Id
				//fmt.Printf("%v\n\nold\n\n", user)
				user, err := directoryapiservice.Users.Update(id, user).Do()
				//fmt.Printf("%v\n\nnew\n\n", user)
				if err != nil {
					push_error_count++
					Log.Printf("problem with put request: %s %s. Current error count %s\n", err.Error(), strconv.Itoa(count), strconv.Itoa(push_error_count))
					//fmt.Printf("problem with put request: %s %s. Current error count %s\n", err.Error(), strconv.Itoa(count), strconv.Itoa(push_error_count))
					//panic(err) //uncomment to stop function at error
				} else {
					//fmt.Printf("directory updated: %v\n", user)
					Log.Printf("directory updated: %v\n", user)
				}
				if organization["location"] != nil {
					//fmt.Println("new address: " + organization["location"].(string))
				}
				//} //               COMMENT THIS OUT
			}
			//fmt.Printf("%s (%s)%s\n", user.PrimaryEmail, user.Name.FullName, user.Id)
			//fmt.Println(user.Emails)

			duration := time.Since(start_time)
			durations = append(durations, duration)
			//fmt.Printf("User # " + strconv.Itoa(count) + " complete in " + durafmt.Parse(duration).String())
			sum += duration //sum milliseconds to value
			avg := sum / time.Duration(len(durations))
			time_left := avg * time.Duration((totalcount - count))
			if count%getConfig.Print_Status_Every == 0 {
				fmt.Println("average time: " + durafmt.Parse(avg).String())
				Log.Println("average time: " + durafmt.Parse(avg).String())
				fmt.Printf("User # " + strconv.Itoa(count) + ". Time projected: " + durafmt.Parse(time_left).String() + ", Current Error count " + strconv.Itoa(push_error_count+get_data_error_count) + "\n")
				Log.Printf("User # " + strconv.Itoa(count) + ". Time projected: " + durafmt.Parse(time_left).String() + ", Current Error count " + strconv.Itoa(push_error_count+get_data_error_count) + "\n")
				if getConfig.Debug_Mode {
					fmt.Println("ERRORS= from push errors: " + strconv.Itoa(push_error_count) + ", from employee data: " + strconv.Itoa(get_employee_data_error_count) + ", from work data: " + strconv.Itoa(get_worklocation_data_error_count) + ", from org data: " + strconv.Itoa(get_organization_data_error_count) + ", from manager data: " + strconv.Itoa(get_managers_error_count) + ", error rate: " + strconv.FormatFloat(float64(push_error_count+get_data_error_count)/float64(count)*100, 'f', 2, 64) + "%")
					Log.Println("ERRORS= from push errors: " + strconv.Itoa(push_error_count) + ", from employee data: " + strconv.Itoa(get_employee_data_error_count) + ", from work data: " + strconv.Itoa(get_worklocation_data_error_count) + ", from org data: " + strconv.Itoa(get_organization_data_error_count) + ", from manager data: " + strconv.Itoa(get_managers_error_count) + ", error rate: " + strconv.FormatFloat(float64(push_error_count+get_data_error_count)/float64(count)*100, 'f', 2, 64) + "%")
				}
			}
		}
		t := time.Now()
		nextPageToken = userList.NextPageToken
		userList, err = getDirectoryUserList(directoryapiservice, nextPageToken)
		fmt.Printf("Next page - users completed: " + strconv.Itoa(count) + ", errors: " + strconv.Itoa(push_error_count+get_data_error_count+get_managers_error_count) + ", " + strconv.FormatFloat(float64(count)/float64(totalcount)*100, 'f', 2, 64) + "%% complete, error rate: " + strconv.FormatFloat(float64(push_error_count+get_data_error_count+get_managers_error_count)/float64(count)*100, 'f', 2, 64) + "%%, " + string(t.Format("2006-01-02 15:04:05")) + "\n")
		Log.Printf("Next page - users completed: " + strconv.Itoa(count) + ", errors: " + strconv.Itoa(push_error_count+get_data_error_count+get_managers_error_count) + ", " + strconv.FormatFloat(float64(count)/float64(totalcount)*100, 'f', 2, 64) + "%% complete, error rate: " + strconv.FormatFloat(float64(push_error_count+get_data_error_count+get_managers_error_count)/float64(count)*100, 'f', 2, 64) + "%%, " + string(t.Format("2006-01-02 15:04:05")) + "\n")
	}
	t := time.Now()
	var main_error string
	if push_error_count+get_data_error_count+get_managers_error_count == 0 {
		main_error = "no errors!"
	} else if push_error_count > get_employee_data_error_count+get_managers_error_count+get_organization_data_error_count {
		main_error = "pushing the data"
	} else if get_managers_error_count > get_employee_data_error_count+get_managers_error_count+get_organization_data_error_count {
		main_error = "getting manager data"
	} else {
		main_error = "getting data from csv files"
	}
	fmt.Println("All done. It took " + durafmt.Parse(time.Since(total_start_time)).String() + "to process " + strconv.Itoa(count) + " users " + strconv.Itoa(count-push_error_count-get_data_error_count-get_managers_error_count) + " successful and " + strconv.Itoa(push_error_count) + " errors. error percent: " + strconv.FormatFloat(float64(push_error_count+get_data_error_count+get_managers_error_count)/float64(count)*100, 'f', 2, 64) + "%%, mainly due to " + main_error + ", " + string(t.Format("2006-01-02 15:04:05")))
	Log.Println("All done. It took " + durafmt.Parse(time.Since(total_start_time)).String() + "to process " + strconv.Itoa(count) + " users " + strconv.Itoa(count-push_error_count-get_data_error_count-get_managers_error_count) + " successful and " + strconv.Itoa(push_error_count) + " errors. error percent: " + strconv.FormatFloat(float64(push_error_count+get_data_error_count+get_managers_error_count)/float64(count)*100, 'f', 2, 64) + "%%, mainly due to " + main_error + ", " + string(t.Format("2006-01-02 15:04:05")))
	fmt.Println(strconv.Itoa(push_error_count) + " push errors and " + strconv.Itoa(get_data_error_count) + " getting data from csv errors and " + strconv.Itoa(get_managers_error_count) + " manager data errors")
	Log.Println(strconv.Itoa(push_error_count) + " push errors and " + strconv.Itoa(get_data_error_count) + " getting data from csv errors and " + strconv.Itoa(get_managers_error_count) + " manager data errors")

	/*
		//original method
		r, err := directoryapiservice.Users.List().Customer("my_customer").MaxResults(500).
			OrderBy("email").Do()
		//fmt.Println(reflect.TypeOf(r))
		if err != nil {
			Log.Printf("Unable to retrieve users in domain.", err)
			panic(err)
		}

		if len(r.Users) == 0 {
			fmt.Print("No users found.\n")
		} else {
			fmt.Print("Users:\n")
			var count = 0
			for _, u := range r.Users {
				count++
				if count > 250 {
					data := u.ExternalIds.([]interface{})
					var newdata = data[0]
					fmt.Println(newdata)
					data1 := objx.New(newdata)
					fmt.Println(u.Emails, data1["value"])
					fmt.Printf("%s (%s)%s\n", u.PrimaryEmail, u.Name.FullName, u.Id)
					break
				}
			}
		}
	*/

}
