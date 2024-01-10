// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ConfigurationListItemTypeEnum Enum with underlying type: string
type ConfigurationListItemTypeEnum string

// Set of constants representing the allowable values for ConfigurationListItemTypeEnum
const (
	ConfigurationListItemTypeManaged ConfigurationListItemTypeEnum = "MANAGED"
	ConfigurationListItemTypeCustom  ConfigurationListItemTypeEnum = "CUSTOM"
)

var mappingConfigurationListItemTypeEnum = map[string]ConfigurationListItemTypeEnum{
	"MANAGED": ConfigurationListItemTypeManaged,
	"CUSTOM":  ConfigurationListItemTypeCustom,
}

var mappingConfigurationListItemTypeEnumLowerCase = map[string]ConfigurationListItemTypeEnum{
	"managed": ConfigurationListItemTypeManaged,
	"custom":  ConfigurationListItemTypeCustom,
}

// GetConfigurationListItemTypeEnumValues Enumerates the set of values for ConfigurationListItemTypeEnum
func GetConfigurationListItemTypeEnumValues() []ConfigurationListItemTypeEnum {
	values := make([]ConfigurationListItemTypeEnum, 0)
	for _, v := range mappingConfigurationListItemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationListItemTypeEnumStringValues Enumerates the set of values in String for ConfigurationListItemTypeEnum
func GetConfigurationListItemTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingConfigurationListItemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationListItemTypeEnum(val string) (ConfigurationListItemTypeEnum, bool) {
	enum, ok := mappingConfigurationListItemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
