package vars

type ScrollResult struct {
	List    []interface{} `json:"list"`
	MinTime int64         `json:"minTime"`
	Offset  int64         `json:"offset"`
}
