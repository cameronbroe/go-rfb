// Code generated by "stringer -type=TightFilter"; DO NOT EDIT.

package rfb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TightFilterCopy-0]
	_ = x[TightFilterPalette-1]
	_ = x[TightFilterGradient-2]
}

const _TightFilter_name = "TightFilterCopyTightFilterPaletteTightFilterGradient"

var _TightFilter_index = [...]uint8{0, 15, 33, 52}

func (i TightFilter) String() string {
	if i >= TightFilter(len(_TightFilter_index)-1) {
		return "TightFilter(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TightFilter_name[_TightFilter_index[i]:_TightFilter_index[i+1]]
}
