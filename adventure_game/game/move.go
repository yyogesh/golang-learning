package game

func Move(currentRoom, direction string) string {
	switch currentRoom {
	case "entrance":
		if direction == "north" {
			return "hallway"
		} else if direction == "east" {
			return "library"
		}
	case "hallway":
		if direction == "south" {
			return "entrance"
		} else if direction == "north" {
			return "chamber"
		}
	case "library":
		if direction == "west" {
			return "entrance"
		}
	case "chamber":
		if direction == "south" {
			return "hallway"
		} else if direction == "north" {
			return "treasure room"
		}
	}
	return ""
}
