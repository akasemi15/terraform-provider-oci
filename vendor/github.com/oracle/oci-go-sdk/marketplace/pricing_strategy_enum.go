// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// PricingStrategyEnumEnum Enum with underlying type: string
type PricingStrategyEnumEnum string

// Set of constants representing the allowable values for PricingStrategyEnumEnum
const (
	PricingStrategyEnumPerOcpuLinear               PricingStrategyEnumEnum = "PER_OCPU_LINEAR"
	PricingStrategyEnumPerOcpuMinBilling           PricingStrategyEnumEnum = "PER_OCPU_MIN_BILLING"
	PricingStrategyEnumPerInstance                 PricingStrategyEnumEnum = "PER_INSTANCE"
	PricingStrategyEnumPerInstanceMonthlyInclusive PricingStrategyEnumEnum = "PER_INSTANCE_MONTHLY_INCLUSIVE"
)

var mappingPricingStrategyEnum = map[string]PricingStrategyEnumEnum{
	"PER_OCPU_LINEAR":                PricingStrategyEnumPerOcpuLinear,
	"PER_OCPU_MIN_BILLING":           PricingStrategyEnumPerOcpuMinBilling,
	"PER_INSTANCE":                   PricingStrategyEnumPerInstance,
	"PER_INSTANCE_MONTHLY_INCLUSIVE": PricingStrategyEnumPerInstanceMonthlyInclusive,
}

// GetPricingStrategyEnumEnumValues Enumerates the set of values for PricingStrategyEnumEnum
func GetPricingStrategyEnumEnumValues() []PricingStrategyEnumEnum {
	values := make([]PricingStrategyEnumEnum, 0)
	for _, v := range mappingPricingStrategyEnum {
		values = append(values, v)
	}
	return values
}
