// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OptimizerStatisticsAdvisorExecution The summary of the Optimizer Statistics Advisor execution, which includes information about the Managed Database
// and a comprehensive execution report.
type OptimizerStatisticsAdvisorExecution struct {

	// The name of the Optimizer Statistics Advisor task.
	TaskName *string `mandatory:"true" json:"taskName"`

	// The name of the Optimizer Statistics Advisor execution.
	ExecutionName *string `mandatory:"true" json:"executionName"`

	// The start time of the time range to retrieve the Optimizer Statistics Advisor execution of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// The end time of the time range to retrieve the Optimizer Statistics Advisor execution of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`

	// The status of the Optimizer Statistics Advisor execution.
	Status OptimizerStatisticsAdvisorExecutionStatusEnum `mandatory:"true" json:"status"`

	// The Optimizer Statistics Advisor execution status message, if any.
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The errors in the Optimizer Statistics Advisor execution, if any.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The number of findings generated by the Optimizer Statistics Advisor execution.
	Findings *int `mandatory:"false" json:"findings"`

	Database *OptimizerDatabase `mandatory:"false" json:"database"`

	Report *OptimizerStatisticsAdvisorExecutionReport `mandatory:"false" json:"report"`
}

func (m OptimizerStatisticsAdvisorExecution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OptimizerStatisticsAdvisorExecution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOptimizerStatisticsAdvisorExecutionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOptimizerStatisticsAdvisorExecutionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OptimizerStatisticsAdvisorExecutionStatusEnum Enum with underlying type: string
type OptimizerStatisticsAdvisorExecutionStatusEnum string

// Set of constants representing the allowable values for OptimizerStatisticsAdvisorExecutionStatusEnum
const (
	OptimizerStatisticsAdvisorExecutionStatusExecuting   OptimizerStatisticsAdvisorExecutionStatusEnum = "EXECUTING"
	OptimizerStatisticsAdvisorExecutionStatusCompleted   OptimizerStatisticsAdvisorExecutionStatusEnum = "COMPLETED"
	OptimizerStatisticsAdvisorExecutionStatusInterrupted OptimizerStatisticsAdvisorExecutionStatusEnum = "INTERRUPTED"
	OptimizerStatisticsAdvisorExecutionStatusCancelled   OptimizerStatisticsAdvisorExecutionStatusEnum = "CANCELLED"
	OptimizerStatisticsAdvisorExecutionStatusFatalError  OptimizerStatisticsAdvisorExecutionStatusEnum = "FATAL_ERROR"
)

var mappingOptimizerStatisticsAdvisorExecutionStatusEnum = map[string]OptimizerStatisticsAdvisorExecutionStatusEnum{
	"EXECUTING":   OptimizerStatisticsAdvisorExecutionStatusExecuting,
	"COMPLETED":   OptimizerStatisticsAdvisorExecutionStatusCompleted,
	"INTERRUPTED": OptimizerStatisticsAdvisorExecutionStatusInterrupted,
	"CANCELLED":   OptimizerStatisticsAdvisorExecutionStatusCancelled,
	"FATAL_ERROR": OptimizerStatisticsAdvisorExecutionStatusFatalError,
}

var mappingOptimizerStatisticsAdvisorExecutionStatusEnumLowerCase = map[string]OptimizerStatisticsAdvisorExecutionStatusEnum{
	"executing":   OptimizerStatisticsAdvisorExecutionStatusExecuting,
	"completed":   OptimizerStatisticsAdvisorExecutionStatusCompleted,
	"interrupted": OptimizerStatisticsAdvisorExecutionStatusInterrupted,
	"cancelled":   OptimizerStatisticsAdvisorExecutionStatusCancelled,
	"fatal_error": OptimizerStatisticsAdvisorExecutionStatusFatalError,
}

// GetOptimizerStatisticsAdvisorExecutionStatusEnumValues Enumerates the set of values for OptimizerStatisticsAdvisorExecutionStatusEnum
func GetOptimizerStatisticsAdvisorExecutionStatusEnumValues() []OptimizerStatisticsAdvisorExecutionStatusEnum {
	values := make([]OptimizerStatisticsAdvisorExecutionStatusEnum, 0)
	for _, v := range mappingOptimizerStatisticsAdvisorExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsAdvisorExecutionStatusEnumStringValues Enumerates the set of values in String for OptimizerStatisticsAdvisorExecutionStatusEnum
func GetOptimizerStatisticsAdvisorExecutionStatusEnumStringValues() []string {
	return []string{
		"EXECUTING",
		"COMPLETED",
		"INTERRUPTED",
		"CANCELLED",
		"FATAL_ERROR",
	}
}

// GetMappingOptimizerStatisticsAdvisorExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsAdvisorExecutionStatusEnum(val string) (OptimizerStatisticsAdvisorExecutionStatusEnum, bool) {
	enum, ok := mappingOptimizerStatisticsAdvisorExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
