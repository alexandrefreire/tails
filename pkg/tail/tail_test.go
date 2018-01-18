/* ***************************************************************************
# Copyright (c) 2018, Industrial Logic, Inc., All Rights Reserved.
#
# This code is the exclusive property of Industrial Logic, Inc. It may ONLY be
# used by students during Industrial Logic's workshops or by individuals
# who are being coached by Industrial Logic on a project.
#
# This code may NOT be copied or used for any other purpose without the prior
# written consent of Industrial Logic, Inc.
# ****************************************************************************/
package tail

import (
	"testing"
	"time"
)

var now = time.Now()
var t10 = Tail{10, URGENT, now}
var t10_dup = Tail{10, URGENT, now}
var t10_diff = Tail{10, DISMISSED, time.Now().Add(time.Minute)}
var t9 = Tail{9, IMPORTANT, now}
var t8 = Tail{8, REQUESTING, now}
var t7 = Tail{7, DISMISSED, now}
var t6 = Tail{6, URGENT, now}
var t6_dismissed = Tail{6, DISMISSED, now}
var t5 = Tail{5, URGENT, now.Add(time.Minute)}

var tailTests = []struct {
	t1       Tail
	t2       Tail
	expected int
}{
	{t10, t10, 0},
	{t10, t10_dup, 0},
	{t10, t10_diff, 0},
	{t10, t9, -1},
	{t9, t10, 1},
	{t9, t8, -1},
	{t8, t9, 1},
	{t8, t7, -1},
	{t7, t8, 1},
	{t6, t5, -1},
	{t5, t6, 1},
	{t6_dismissed, t7, -1},
	{t7, t6_dismissed, 1},
}

func TestAllTheThings(t *testing.T) {
	for _, tt := range tailTests {
		actual := tt.t1.Compare(tt.t2)
		if actual != tt.expected {
			t.Errorf("Compare(%v, %v): expected %d, actual %d", tt.t1, tt.t2, tt.expected, actual)
		}
	}
}
