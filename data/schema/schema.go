package schema

// define a struct to hold the json output from hyprctl clients -j
type Client []struct {
	Address        string        `json:"address"`
	At             []int         `json:"at"`
	Size           []int         `json:"size"`
	Workspaces     Workspace     `json:"workspace"`
	Floating       bool          `json:"floating"`
	Monitor        int           `json:"monitor"`
	Class          string        `json:"class"`
	Title          string        `json:"title"`
	Pid            int           `json:"pid"`
	Xwayland       bool          `json:"xwayland"`
	Pinned         bool          `json:"pinned"`
	Fullscreen     bool          `json:"fullscreen"`
	FullscreenMode int           `json:"fullscreenMode"`
	FakeFullscreen bool          `json:"fakeFullscreen"`
	Grouped        []interface{} `json:"grouped"`
	Swallowing     interface{}   `json:"swallowing"`
}

type Workspace struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
