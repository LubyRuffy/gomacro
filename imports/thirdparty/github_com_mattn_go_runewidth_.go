// this file was generated by gomacro command: import "github.com/mattn/go-runewidth"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package thirdparty

import (
	. "reflect"

	go_runewidth_ "github.com/mattn/go-runewidth"
)

// reflection: allow interpreted code to import "github.com/mattn/go-runewidth"
func init() {
	Packages["github.com/mattn/go-runewidth"] = Package{
		Binds: map[string]Value{
			"DefaultCondition": ValueOf(&go_runewidth_.DefaultCondition).Elem(),
			"EastAsianWidth":   ValueOf(&go_runewidth_.EastAsianWidth).Elem(),
			"FillLeft":         ValueOf(go_runewidth_.FillLeft),
			"FillRight":        ValueOf(go_runewidth_.FillRight),
			"IsAmbiguousWidth": ValueOf(go_runewidth_.IsAmbiguousWidth),
			"IsEastAsian":      ValueOf(go_runewidth_.IsEastAsian),
			"IsNeutralWidth":   ValueOf(go_runewidth_.IsNeutralWidth),
			"NewCondition":     ValueOf(go_runewidth_.NewCondition),
			"RuneWidth":        ValueOf(go_runewidth_.RuneWidth),
			"StringWidth":      ValueOf(go_runewidth_.StringWidth),
			"Truncate":         ValueOf(go_runewidth_.Truncate),
			"Wrap":             ValueOf(go_runewidth_.Wrap),
		}, Types: map[string]Type{
			"Condition": TypeOf((*go_runewidth_.Condition)(nil)).Elem(),
		},
	}
}
