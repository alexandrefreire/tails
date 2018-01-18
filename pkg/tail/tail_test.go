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
	equals(t10, t10, t)
}

func TestCompareSameId(t *testing.T) {
	equals(t10, Tail{10, URGENT, time.Now()}, t)
}

func TestCompareSameIdDifferentPriority(t *testing.T) {
	equals(t10, Tail{10, DISMISSED, time.Now().Add(time.Minute)}, t)
}

func equals(t1, t2 Tail, t *testing.T) {
	if t1.Compare(t2) != 0 {
		t.Fail()
	}
	if t2.Compare(t1) != 0 {
		t.Fail()
	}
}
