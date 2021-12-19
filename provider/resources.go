// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vultr

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/vincentbernat/pulumi-vultr/provider/pkg/version"
	"github.com/vultr/terraform-provider-vultr/vultr"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "vultr"
	// modules:
	mainMod = "index" // the vultr module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(vultr.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "vultr",
		Description:          "A Pulumi package for creating and managing vultr cloud resources.",
		Keywords:             []string{"pulumi", "vultr"},
		License:              "Apache-2.0",
		Homepage:             "https://pulumi.io",
		Repository:           "https://github.com/vincentbernat/pulumi-vultr",
		GitHubOrg:            "vultr",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"vultr_bare_metal_server": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BareMetalServer")},
			"vultr_block_storage":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BlockStorage")},
			"vultr_dns_domain":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsDomain")},
			"vultr_dns_record":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsRecord")},
			"vultr_firewall_group":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FirewallGroup")},
			"vultr_firewall_rule":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FirewallRule")},
			"vultr_instance":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Instance")},
			"vultr_instance_ipv4":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceIPv4")},
			"vultr_iso_private":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IsoPrivate")},
			"vultr_load_balancer":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LoadBalancer")},
			"vultr_object_storage":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectStorage")},
			"vultr_private_network":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PrivateNetwork")},
			"vultr_reserved_ip":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ReservedIP")},
			"vultr_reverse_ipv4":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ReverseIpv4")},
			"vultr_reverse_ipv6":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ReverseIpv6")},
			"vultr_snapshot":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Snapshot")},
			"vultr_snapshot_from_url": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SnapshotFromUrl")},
			"vultr_ssh_key": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "SshKey"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ssh_key": {
						CSharpName: "Key",
					},
				},
			},
			"vultr_startup_script": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StartupScript")},
			"vultr_user":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vultr_account":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAccount")},
			"vultr_application":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApplication")},
			"vultr_backup":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBackup")},
			"vultr_bare_metal_plan":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBareMetalPlan")},
			"vultr_bare_metal_server": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBareMetalServer")},
			"vultr_block_storage":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBlockStorage")},
			"vultr_dns_domain":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDnsDomain")},
			"vultr_firewall_group":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFirewallGroup")},
			"vultr_instance":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstance")},
			"vultr_instance_ipv4":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceIpv4")},
			"vultr_iso_private":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIsoPrivate")},
			"vultr_iso_public":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIsoPublic")},
			"vultr_load_balancer":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLoadBalancer")},
			"vultr_object_storage":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getObjectStorage")},
			"vultr_os":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOs")},
			"vultr_plan":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPlan")},
			"vultr_private_network":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPrivateNetwork")},
			"vultr_region":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegion")},
			"vultr_reserved_ip":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getReservedIp")},
			"vultr_reverse_ipv4":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getReverseIpv4")},
			"vultr_reverse_ipv6":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getReverseIpv6")},
			"vultr_snapshot":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSnapshot")},
			"vultr_ssh_key": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSshKey"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ssh_key": {
						CSharpName: "Key",
					},
				},
			},
			"vultr_startup_script": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStartupScript")},
			"vultr_user":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/vincentbernat/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
