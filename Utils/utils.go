package Utils


func CheckWordExistance(word string, list []string)bool {
	/*Checks wether element is in the given list */

	for _, excw := range(list){
		if excw == word{
			return true
		}
	}
	return false
}