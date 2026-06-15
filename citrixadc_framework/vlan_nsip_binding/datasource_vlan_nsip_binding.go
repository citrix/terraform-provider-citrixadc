package vlan_nsip_binding

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VlanNsipBindingDataSource)(nil)

func VLanNsipBindingDataSource() datasource.DataSource {
	return &VlanNsipBindingDataSource{}
}

type VlanNsipBindingDataSource struct {
	client *service.NitroClient
}

func (d *VlanNsipBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_nsip_binding"
}

func (d *VlanNsipBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VlanNsipBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VlanNsipBindingDataSourceSchema()
}

func (d *VlanNsipBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VlanNsipBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// The parent name for the binding GET is the vlanid.
	parentName := strconv.FormatInt(data.Vlanid.ValueInt64(), 10)
	ipaddress_Name := data.Ipaddress
	netmask_Name := data.Netmask
	ownergroup_Name := data.Ownergroup
	td_Name := data.Td

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vlan_nsip_binding.Type(),
		ResourceName:             parentName,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vlan_nsip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vlan_nsip_binding returned empty array.")
		return
	}

	// Iterate through results to find the matching binding
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ipaddress (always provided by the datasource user)
		if val, ok := v["ipaddress"].(string); ok {
			if ipaddress_Name.IsNull() || val != ipaddress_Name.ValueString() {
				match = false
				continue
			}
		} else if !ipaddress_Name.IsNull() {
			match = false
			continue
		}

		// Check netmask only when supplied
		if !netmask_Name.IsNull() {
			if val, ok := v["netmask"].(string); ok {
				if val != netmask_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check ownergroup only when supplied
		if !ownergroup_Name.IsNull() {
			if val, ok := v["ownergroup"].(string); ok {
				if val != ownergroup_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check td only when supplied
		if !td_Name.IsNull() {
			if val, ok := v["td"]; ok {
				valInt, _ := utils.ConvertToInt64(val)
				if valInt != td_Name.ValueInt64() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vlan_nsip_binding with ipaddress %s not found", ipaddress_Name))
		return
	}

	vlan_nsip_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
