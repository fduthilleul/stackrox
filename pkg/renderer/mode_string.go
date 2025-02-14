// Code generated by "stringer -type=mode"; DO NOT EDIT.

package renderer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[renderAll-0]
	_ = x[scannerOnly-1]
	_ = x[centralTLSOnly-2]
	_ = x[centralDBTLSOnly-3]
	_ = x[scannerTLSOnly-4]
	_ = x[centralDBOnly-5]
	_ = x[scannerV4TLSOnly-6]
}

const _mode_name = "renderAllscannerOnlycentralTLSOnlycentralDBTLSOnlyscannerTLSOnlycentralDBOnlyscannerV4TLSOnly"

var _mode_index = [...]uint8{0, 9, 20, 34, 50, 64, 77, 93}

func (i mode) String() string {
	if i < 0 || i >= mode(len(_mode_index)-1) {
		return "mode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _mode_name[_mode_index[i]:_mode_index[i+1]]
}
