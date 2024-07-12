// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type DeviceSubinterface struct {
	Id                       types.String                      `tfsdk:"id"`
	Domain                   types.String                      `tfsdk:"domain"`
	DeviceId                 types.String                      `tfsdk:"device_id"`
	InterfaceId              types.String                      `tfsdk:"interface_id"`
	Enabled                  types.Bool                        `tfsdk:"enabled"`
	InterfaceName            types.String                      `tfsdk:"interface_name"`
	Index                    types.Int64                       `tfsdk:"index"`
	VlanId                   types.Int64                       `tfsdk:"vlan_id"`
	LogicalName              types.String                      `tfsdk:"logical_name"`
	Description              types.String                      `tfsdk:"description"`
	ManagementOnly           types.Bool                        `tfsdk:"management_only"`
	Mtu                      types.Int64                       `tfsdk:"mtu"`
	Priority                 types.Int64                       `tfsdk:"priority"`
	SecurityZoneId           types.String                      `tfsdk:"security_zone_id"`
	Ipv4StaticAddress        types.String                      `tfsdk:"ipv4_static_address"`
	Ipv4StaticNetmask        types.String                      `tfsdk:"ipv4_static_netmask"`
	Ipv4DhcpObtainRoute      types.Bool                        `tfsdk:"ipv4_dhcp_obtain_route"`
	Ipv4DhcpRouteMetric      types.Int64                       `tfsdk:"ipv4_dhcp_route_metric"`
	Ipv6Enable               types.Bool                        `tfsdk:"ipv6_enable"`
	Ipv6EnforceEui           types.Bool                        `tfsdk:"ipv6_enforce_eui"`
	Ipv6EnableAutoConfig     types.Bool                        `tfsdk:"ipv6_enable_auto_config"`
	Ipv6EnableDhcpAddress    types.Bool                        `tfsdk:"ipv6_enable_dhcp_address"`
	Ipv6EnableDhcpNonaddress types.Bool                        `tfsdk:"ipv6_enable_dhcp_nonaddress"`
	Ipv6EnableRa             types.Bool                        `tfsdk:"ipv6_enable_ra"`
	Ipv6Addresses            []DeviceSubinterfaceIpv6Addresses `tfsdk:"ipv6_addresses"`
}

type DeviceSubinterfaceIpv6Addresses struct {
	Address    types.String `tfsdk:"address"`
	Prefix     types.String `tfsdk:"prefix"`
	EnforceEui types.Bool   `tfsdk:"enforce_eui"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DeviceSubinterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/subinterfaces", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DeviceSubinterface) toBody(ctx context.Context, state DeviceSubinterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "links.parent", data.InterfaceId.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if state.InterfaceName.ValueString() != "" {
		body, _ = sjson.Set(body, "name", state.InterfaceName.ValueString())
	}
	if !data.Index.IsNull() {
		body, _ = sjson.Set(body, "subIntfId", data.Index.ValueInt64())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "ifname", data.LogicalName.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.ManagementOnly.IsNull() {
		body, _ = sjson.Set(body, "managementOnly", data.ManagementOnly.ValueBool())
	}
	if !data.Mtu.IsNull() {
		body, _ = sjson.Set(body, "MTU", data.Mtu.ValueInt64())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, "priority", data.Priority.ValueInt64())
	}
	if !data.SecurityZoneId.IsNull() {
		body, _ = sjson.Set(body, "securityZone.id", data.SecurityZoneId.ValueString())
	}
	body, _ = sjson.Set(body, "securityZone.type", "SecurityZone")
	if !data.Ipv4StaticAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.address", data.Ipv4StaticAddress.ValueString())
	}
	if !data.Ipv4StaticNetmask.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.netmask", data.Ipv4StaticNetmask.ValueString())
	}
	if !data.Ipv4DhcpObtainRoute.IsNull() {
		body, _ = sjson.Set(body, "ipv4.dhcp.enableDefaultRouteDHCP", data.Ipv4DhcpObtainRoute.ValueBool())
	}
	if !data.Ipv4DhcpRouteMetric.IsNull() {
		body, _ = sjson.Set(body, "ipv4.dhcp.dhcpRouteMetric", data.Ipv4DhcpRouteMetric.ValueInt64())
	}
	if !data.Ipv6Enable.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableIPV6", data.Ipv6Enable.ValueBool())
	}
	if !data.Ipv6EnforceEui.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enforceEUI64", data.Ipv6EnforceEui.ValueBool())
	}
	if !data.Ipv6EnableAutoConfig.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableAutoConfig", data.Ipv6EnableAutoConfig.ValueBool())
	}
	if !data.Ipv6EnableDhcpAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDHCPAddrConfig", data.Ipv6EnableDhcpAddress.ValueBool())
	}
	if !data.Ipv6EnableDhcpNonaddress.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDHCPNonAddrConfig", data.Ipv6EnableDhcpNonaddress.ValueBool())
	}
	if !data.Ipv6EnableRa.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableRA", data.Ipv6EnableRa.ValueBool())
	}
	if len(data.Ipv6Addresses) > 0 {
		body, _ = sjson.Set(body, "ipv6.addresses", []interface{}{})
		for _, item := range data.Ipv6Addresses {
			itemBody := ""
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.Address.ValueString())
			}
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "prefix", item.Prefix.ValueString())
			}
			if !item.EnforceEui.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enforceEUI64", item.EnforceEui.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "ipv6.addresses.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DeviceSubinterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolValue(true)
	}
	if value := res.Get("name"); value.Exists() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("subIntfId"); value.Exists() {
		data.Index = types.Int64Value(value.Int())
	} else {
		data.Index = types.Int64Null()
	}
	if value := res.Get("vlanId"); value.Exists() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("ifname"); value.Exists() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("managementOnly"); value.Exists() {
		data.ManagementOnly = types.BoolValue(value.Bool())
	} else {
		data.ManagementOnly = types.BoolNull()
	}
	if value := res.Get("MTU"); value.Exists() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("securityZone.id"); value.Exists() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	if value := res.Get("ipv4.dhcp.dhcpRouteMetric"); value.Exists() {
		data.Ipv4DhcpRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4DhcpRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv6.enableIPV6"); value.Exists() {
		data.Ipv6Enable = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	if value := res.Get("ipv6.enforceEUI64"); value.Exists() {
		data.Ipv6EnforceEui = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnforceEui = types.BoolNull()
	}
	if value := res.Get("ipv6.enableAutoConfig"); value.Exists() {
		data.Ipv6EnableAutoConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableAutoConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPAddrConfig"); value.Exists() {
		data.Ipv6EnableDhcpAddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpAddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPNonAddrConfig"); value.Exists() {
		data.Ipv6EnableDhcpNonaddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpNonaddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableRA"); value.Exists() {
		data.Ipv6EnableRa = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableRa = types.BoolNull()
	}
	if value := res.Get("ipv6.addresses"); value.Exists() {
		data.Ipv6Addresses = make([]DeviceSubinterfaceIpv6Addresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DeviceSubinterfaceIpv6Addresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			} else {
				item.Address = types.StringNull()
			}
			if cValue := v.Get("prefix"); cValue.Exists() {
				item.Prefix = types.StringValue(cValue.String())
			} else {
				item.Prefix = types.StringNull()
			}
			if cValue := v.Get("enforceEUI64"); cValue.Exists() {
				item.EnforceEui = types.BoolValue(cValue.Bool())
			} else {
				item.EnforceEui = types.BoolNull()
			}
			data.Ipv6Addresses = append(data.Ipv6Addresses, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *DeviceSubinterface) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else if data.Enabled.ValueBool() != true {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("subIntfId"); value.Exists() && !data.Index.IsNull() {
		data.Index = types.Int64Value(value.Int())
	} else {
		data.Index = types.Int64Null()
	}
	if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("ifname"); value.Exists() && !data.LogicalName.IsNull() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("managementOnly"); value.Exists() && !data.ManagementOnly.IsNull() {
		data.ManagementOnly = types.BoolValue(value.Bool())
	} else {
		data.ManagementOnly = types.BoolNull()
	}
	if value := res.Get("MTU"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("securityZone.id"); value.Exists() && !data.SecurityZoneId.IsNull() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() && !data.Ipv4StaticAddress.IsNull() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() && !data.Ipv4StaticNetmask.IsNull() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() && !data.Ipv4DhcpObtainRoute.IsNull() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	if value := res.Get("ipv4.dhcp.dhcpRouteMetric"); value.Exists() && !data.Ipv4DhcpRouteMetric.IsNull() {
		data.Ipv4DhcpRouteMetric = types.Int64Value(value.Int())
	} else {
		data.Ipv4DhcpRouteMetric = types.Int64Null()
	}
	if value := res.Get("ipv6.enableIPV6"); value.Exists() && !data.Ipv6Enable.IsNull() {
		data.Ipv6Enable = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	if value := res.Get("ipv6.enforceEUI64"); value.Exists() && !data.Ipv6EnforceEui.IsNull() {
		data.Ipv6EnforceEui = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnforceEui = types.BoolNull()
	}
	if value := res.Get("ipv6.enableAutoConfig"); value.Exists() && !data.Ipv6EnableAutoConfig.IsNull() {
		data.Ipv6EnableAutoConfig = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableAutoConfig = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPAddrConfig"); value.Exists() && !data.Ipv6EnableDhcpAddress.IsNull() {
		data.Ipv6EnableDhcpAddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpAddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableDHCPNonAddrConfig"); value.Exists() && !data.Ipv6EnableDhcpNonaddress.IsNull() {
		data.Ipv6EnableDhcpNonaddress = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDhcpNonaddress = types.BoolNull()
	}
	if value := res.Get("ipv6.enableRA"); value.Exists() && !data.Ipv6EnableRa.IsNull() {
		data.Ipv6EnableRa = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableRa = types.BoolNull()
	}
	for i := range data.Ipv6Addresses {
		keys := [...]string{"address", "prefix"}
		keyValues := [...]string{data.Ipv6Addresses[i].Address.ValueString(), data.Ipv6Addresses[i].Prefix.ValueString()}

		var r gjson.Result
		res.Get("ipv6.addresses").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("address"); value.Exists() && !data.Ipv6Addresses[i].Address.IsNull() {
			data.Ipv6Addresses[i].Address = types.StringValue(value.String())
		} else {
			data.Ipv6Addresses[i].Address = types.StringNull()
		}
		if value := r.Get("prefix"); value.Exists() && !data.Ipv6Addresses[i].Prefix.IsNull() {
			data.Ipv6Addresses[i].Prefix = types.StringValue(value.String())
		} else {
			data.Ipv6Addresses[i].Prefix = types.StringNull()
		}
		if value := r.Get("enforceEUI64"); value.Exists() && !data.Ipv6Addresses[i].EnforceEui.IsNull() {
			data.Ipv6Addresses[i].EnforceEui = types.BoolValue(value.Bool())
		} else {
			data.Ipv6Addresses[i].EnforceEui = types.BoolNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *DeviceSubinterface) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceId.IsNull() {
		return false
	}
	if !data.InterfaceId.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.InterfaceName.IsNull() {
		return false
	}
	if !data.Index.IsNull() {
		return false
	}
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.LogicalName.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.ManagementOnly.IsNull() {
		return false
	}
	if !data.Mtu.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.SecurityZoneId.IsNull() {
		return false
	}
	if !data.Ipv4StaticAddress.IsNull() {
		return false
	}
	if !data.Ipv4StaticNetmask.IsNull() {
		return false
	}
	if !data.Ipv4DhcpObtainRoute.IsNull() {
		return false
	}
	if !data.Ipv4DhcpRouteMetric.IsNull() {
		return false
	}
	if !data.Ipv6Enable.IsNull() {
		return false
	}
	if !data.Ipv6EnforceEui.IsNull() {
		return false
	}
	if !data.Ipv6EnableAutoConfig.IsNull() {
		return false
	}
	if !data.Ipv6EnableDhcpAddress.IsNull() {
		return false
	}
	if !data.Ipv6EnableDhcpNonaddress.IsNull() {
		return false
	}
	if !data.Ipv6EnableRa.IsNull() {
		return false
	}
	if len(data.Ipv6Addresses) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull