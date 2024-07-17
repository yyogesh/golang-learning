package game

func TakeItem(room string, invetory *[]string, item string) bool {
	items := GetRoomItems(room)
	for _, v := range items {
		if v == item {
			*invetory = append(*invetory, item)
			return true
		}
	}
	return false
}
