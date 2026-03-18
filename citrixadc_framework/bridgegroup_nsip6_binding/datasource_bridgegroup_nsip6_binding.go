package bridgegroup_nsip6_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BridgegroupNsip6BindingDataSource)(nil)

func BRidgegroupNsip6BindingDataSource() datasource.DataSource {
	return &BridgegroupNsip6BindingDataSource{}
}

type BridgegroupNsip6BindingDataSource struct {
	client *service.NitroClient
}

func (d *BridgegroupNsip6BindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup_nsip6_binding"
}

func (d *BridgegroupNsip6BindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BridgegroupNsip6BindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BridgegroupNsip6BindingDataSourceSchema()
}

func (d *BridgegroupNsip6BindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BridgegroupNsip6BindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	bridgegroup_id_Name := data.Bridgegroupid
	ipaddress_Name := data.Ipaddress
	netmask_Name := data.Netmask
	ownergroup_Name := data.Ownergroup
	td_Name := data.Td

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Bridgegroup_nsip6_binding.Type(),
		ResourceName:             fmt.Sprintf("%d", bridgegroup_id_Name.ValueInt64()),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup_nsip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "bridgegroup_nsip6_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ipaddress
		if val, ok := v["ipaddress"].(string); ok {
			if ipaddress_Name.IsNull() || val != ipaddress_Name.ValueString() {
				match = false
				continue
			}
		} else if !ipaddress_Name.IsNull() {
			match = false
			continue
		}

		// Check netmask
		if val, ok := v["netmask"].(string); ok {
			if netmask_Name.IsNull() || val != netmask_Name.ValueString() {
				match = false
				continue
			}
		} else if !netmask_Name.IsNull() {
			match = false
			continue
		}

		// Check ownergroup
		if val, ok := v["ownergroup"].(string); ok {
			if ownergroup_Name.IsNull() || val != ownergroup_Name.ValueString() {
				match = false
				continue
			}
		} else if !ownergroup_Name.IsNull() {
			match = false
			continue
		}

		// Check td
		if val, ok := v["td"]; ok {
			val, _ = utils.ConvertToInt64(val)
			if td_Name.IsNull() || val != td_Name.ValueInt64() {
				match = false
				continue
			}
		} else if !td_Name.IsNull() {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("bridgegroup_nsip6_binding with ipaddress %s not found", ipaddress_Name))
		return
	}

	bridgegroup_nsip6_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
