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
	"time"
)

var URGENT = 0
var IMPORTANT = 1
var REQUESTING = 2
var DISMISSED = 3

type Tail struct {
	Id          int
	Priority    int
	LastContact time.Time
}

func (me Tail) Compare(other Tail) int {
	if me.Id == other.Id {
		return 0
	}
	if me.Priority < other.Priority {
		return -1
	}
	if me.Priority > other.Priority {
		return 1
	}
	if me.LastContact.Before(other.LastContact) {
		return -1
	}
	if me.LastContact.After(other.LastContact) {
		return 1
	}
	if me.Id < other.Id {
		return -1
	}
	return 1
}
