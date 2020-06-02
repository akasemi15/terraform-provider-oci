// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreDHCPOptionsTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreDHCPOptionsTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
	}`
	s.ResourceName = "data.oci_core_dhcp_options.t"
}

func (s *DatasourceCoreDHCPOptionsTestSuite) TestAccDatasourceCoreDHCPOptions_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
				data "oci_core_dhcp_options" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					
					filter {
						name = "display_name"
						values = ["Default DHCP Options.*"]
						regex = true
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.display_name", "Default DHCP Options for -tf-vcn"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.options.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "options.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "options.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "options.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "options.0.vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.state", string(core.DhcpOptionsLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.options.0.server_type", string(core.DhcpDnsOptionServerTypeVcnlocalplusinternet)),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreDHCPOptionsTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreDHCPOptionsTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreDHCPOptionsTestSuite))
}
