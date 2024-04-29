// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAllPackagesOnManagedInstancesInCompartmentDetails Provides the information used to update all packages of a specified type on managed instances within the specified compartment.
type UpdateAllPackagesOnManagedInstancesInCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The types of updates to be applied.
	UpdateTypes []UpdateTypesEnum `mandatory:"false" json:"updateTypes,omitempty"`

	WorkRequestDetails *WorkRequestDetails `mandatory:"false" json:"workRequestDetails"`
}

func (m UpdateAllPackagesOnManagedInstancesInCompartmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAllPackagesOnManagedInstancesInCompartmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.UpdateTypes {
		if _, ok := GetMappingUpdateTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateTypes: %s. Supported values are: %s.", val, strings.Join(GetUpdateTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
