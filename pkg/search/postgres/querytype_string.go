// Code generated by "stringer -type=QueryType"; DO NOT EDIT.

package postgres

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SEARCH-0]
	_ = x[GET-1]
	_ = x[COUNT-2]
	_ = x[DELETE-3]
	_ = x[SELECT-4]
	_ = x[DELETERETURNINGIDS-5]
}

const _QueryType_name = "SEARCHGETCOUNTDELETESELECTDELETERETURNINGIDS"

var _QueryType_index = [...]uint8{0, 6, 9, 14, 20, 26, 44}

func (i QueryType) String() string {
	if i < 0 || i >= QueryType(len(_QueryType_index)-1) {
		return "QueryType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _QueryType_name[_QueryType_index[i]:_QueryType_index[i+1]]
}
