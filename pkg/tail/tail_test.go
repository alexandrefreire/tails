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

func TestCompareToSelf(t *testing.T) {
	tail := Tail{10, URGENT, time.Now()}

	if tail.Compare(tail) != 0 {
		t.Fail()
	}
}

func TestCompareSameId(t *testing.T) {
	t1 := Tail{10, URGENT, time.Now()}
	t2 := Tail{10, URGENT, time.Now()}

	if t1.Compare(t2) != 0 {
		t.Fail()
	}
}

func TestCompareSameIdDifferentPriority(t *testing.T) {
	t1 := Tail{10, URGENT, time.Now()}
	t2 := Tail{10, DISMISSED, time.Now()}

	if t1.Compare(t2) != 0 {
		t.Fail()
	}
}
