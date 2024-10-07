// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"strings"
)

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted           WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress         WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusFailed             WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusSucceeded          WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusPartiallySucceeded WorkRequestStatusEnum = "PARTIALLY_SUCCEEDED"
	WorkRequestStatusCanceling          WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled           WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":            WorkRequestStatusAccepted,
	"IN_PROGRESS":         WorkRequestStatusInProgress,
	"FAILED":              WorkRequestStatusFailed,
	"SUCCEEDED":           WorkRequestStatusSucceeded,
	"PARTIALLY_SUCCEEDED": WorkRequestStatusPartiallySucceeded,
	"CANCELING":           WorkRequestStatusCanceling,
	"CANCELED":            WorkRequestStatusCanceled,
}

var mappingWorkRequestStatusEnumLowerCase = map[string]WorkRequestStatusEnum{
	"accepted":            WorkRequestStatusAccepted,
	"in_progress":         WorkRequestStatusInProgress,
	"failed":              WorkRequestStatusFailed,
	"succeeded":           WorkRequestStatusSucceeded,
	"partially_succeeded": WorkRequestStatusPartiallySucceeded,
	"canceling":           WorkRequestStatusCanceling,
	"canceled":            WorkRequestStatusCanceled,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestStatusEnumStringValues Enumerates the set of values in String for WorkRequestStatusEnum
func GetWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"PARTIALLY_SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	enum, ok := mappingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
