package entities

type DashboardPanel struct {
	Id             int    `json:"id"`
	LayoutName     string `json:"layout_name"`
	ComponentIndex string `json:"i"`
	ComponentKey   int    `json:"key"`
	Xposition      int    `json:"x"`
	Yposition      int    `json:"y"`
	Width          int    `json:"w"`
	Height         int    `json:"h"`
	Component      string `json:"component"`
}
