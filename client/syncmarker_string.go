// Code generated by "stringer -type=syncMarker -trimprefix=syncMarker"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[syncMarkerNone-0]
	_ = x[syncMarkerFriendList-1]
	_ = x[syncMarkerFriendInfo-2]
	_ = x[syncMarkerGroupList-3]
	_ = x[syncMarkerGroupInfo-4]
	_ = x[syncMarkerGroupMemberList-5]
	_ = x[syncMarkerGroupMemberInfo-6]
}

const _syncMarker_name = "NoneFriendListFriendInfoGroupListGroupInfoGroupMemberListGroupMemberInfo"

var _syncMarker_index = [...]uint8{0, 4, 14, 24, 33, 42, 57, 72}

func (i syncMarker) String() string {
	if i < 0 || i >= syncMarker(len(_syncMarker_index)-1) {
		return "syncMarker(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _syncMarker_name[_syncMarker_index[i]:_syncMarker_index[i+1]]
}
