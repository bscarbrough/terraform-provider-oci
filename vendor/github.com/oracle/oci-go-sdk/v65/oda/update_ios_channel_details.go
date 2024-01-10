// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateIosChannelDetails Properties to update an iOS channel.
type UpdateIosChannelDetails struct {

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	Name *string `mandatory:"false" json:"name"`

	// A short description of the Channel.
	Description *string `mandatory:"false" json:"description"`

	// The number of milliseconds before a session expires.
	SessionExpiryDurationInMilliseconds *int64 `mandatory:"false" json:"sessionExpiryDurationInMilliseconds"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The maximum time until the token expires (in minutes).
	MaxTokenExpirationTimeInMinutes *int64 `mandatory:"false" json:"maxTokenExpirationTimeInMinutes"`

	// Whether client authentication is enabled or not.
	IsClientAuthenticationEnabled *bool `mandatory:"false" json:"isClientAuthenticationEnabled"`

	// The ID of the Skill or Digital Assistant that the Channel is routed to.
	BotId *string `mandatory:"false" json:"botId"`
}

// GetName returns Name
func (m UpdateIosChannelDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m UpdateIosChannelDetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m UpdateIosChannelDetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m UpdateIosChannelDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateIosChannelDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateIosChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateIosChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateIosChannelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateIosChannelDetails UpdateIosChannelDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateIosChannelDetails
	}{
		"IOS",
		(MarshalTypeUpdateIosChannelDetails)(m),
	}

	return json.Marshal(&s)
}
