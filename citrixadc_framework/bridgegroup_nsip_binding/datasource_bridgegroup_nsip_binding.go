package bridgegroup_nsip_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BridgegroupNsipBindingDataSource)(nil)

func BRidgegroupNsipBindingDataSource() datasource.DataSource {
	return &BridgegroupNsipBindingDataSource{}
}

type BridgegroupNsipBindingDataSource struct {
	client *service.NitroClient
}

func (d *BridgegroupNsipBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup_nsip_binding"
}

func (d *BridgegroupNsipBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BridgegroupNsipBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BridgegroupNsipBindingDataSourceSchema()
}

func (d *BridgegroupNsipBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BridgegroupNsipBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	bridgegroup_id_Name := data.Bridgegroupid
	ipaddress_Name := data.Ipaddress
	ownergroup_Name := data.Ownergroup
	td_Name := data.Td

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Bridgegroup_nsip_binding.Type(),
		ResourceName:             fmt.Sprintf("%d", bridgegroup_id_Name.ValueInt64()),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup_nsip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "bridgegroup_nsip_binding returned empty array.")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("bridgegroup_nsip_binding with ipaddress %s not found", ipaddress_Name))
		return
	}

	bridgegroup_nsip_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
