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

var t10 = Tail{10, URGENT, time.Now()}

func TestCompareToSelf(t *testing.T) {

	if t10.Compare(t10) != 0 {
		t.Fail()
	}
}

func TestCompareSameId(t *testing.T) {
	t10_dup := Tail{10, URGENT, time.Now()}

	if t10.Compare(t10_dup) != 0 {
		t.Fail()
	}
}

func TestCompareSameIdDifferentPriority(t *testing.T) {
	t10_dif := Tail{10, DISMISSED, time.Now().Add(time.Minute)}

	if t10.Compare(t10_dif) != 0 {
		t.Fail()
	}
}
