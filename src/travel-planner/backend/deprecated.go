package backend

// func (backend *MySQLBackend) GetRoutes(sites []uint32) (int32, []model.Activity, []model.Transportation) {
// 	// get activities
// 	var activitiesResult []model.Activity
// 	result := backend.db.Table("Activities").Find(&activitiesResult)
// 	fmt.Println(activitiesResult, result)
// 	if result.Error != nil {
// 		return -1, nil, nil
// 	}
	
// 	// get transportations
// 	var tranportationsResult []model.Transportation
// 	result = backend.db.Table("Transportations").Find(&tranportationsResult)
// 	fmt.Println(tranportationsResult, result)
// 	if result.Error != nil {
// 		return -1, activitiesResult, nil
// 	}

// 	return 8, activitiesResult, tranportationsResult
// }