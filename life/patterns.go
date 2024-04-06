package life

var patterns = map[string][][]uint8{
	"glider": [][]uint8{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	},
	"lwss": [][]uint8{
		{0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
	},
	"bee-hive": [][]uint8{
		{0, 1, 0, 0},
		{1, 0, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 0, 0},
	},
}