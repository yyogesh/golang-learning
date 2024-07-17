package game

func GetRoomItems(room string) []string {
	switch room {
	case "entrance":
		return []string{"map"}
	case "hallway":
		return []string{"torch"}
	case "library":
		return []string{"book"}
	case "chamber":
		return []string{"key"}
	default:
		return []string{}
	}
}
