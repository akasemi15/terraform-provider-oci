// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	letterOfAuthoritySingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	LetterOfAuthorityResourceConfig = generateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", Required, Create, crossConnectLocationDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectRepresentation)
)

func TestCoreLetterOfAuthorityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreLetterOfAuthorityResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_letter_of_authority.test_letter_of_authority"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_letter_of_authority", "test_letter_of_authority", Required, Create, letterOfAuthoritySingularDataSourceRepresentation) +
					compartmentIdVariableStr + LetterOfAuthorityResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "authorized_entity_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "circuit_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "facility_location"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "port_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expires"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_issued"),
				),
			},
		},
	})
}
