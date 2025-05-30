package models

import "encoding/json"

type EZWGetMenusByUser struct {
	MenuIdList json.RawMessage // key: menu_id_list
}
