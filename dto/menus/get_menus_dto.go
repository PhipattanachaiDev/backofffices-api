package menus

import "encoding/json"

type GetMenusResponse struct {
	Menus json.RawMessage `json:"menus"`
}
