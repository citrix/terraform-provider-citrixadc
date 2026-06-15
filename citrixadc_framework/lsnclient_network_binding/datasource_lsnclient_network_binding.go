package lsnclient_network_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsnclientNetworkBindingDataSource)(nil)

func LSnclientNetworkBindingDataSource() datasource.DataSource {
	return &LsnclientNetworkBindingDataSource{}
}

type LsnclientNetworkBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsnclientNetworkBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_network_binding"
}

func (d *LsnclientNetworkBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnclientNetworkBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnclientNetworkBindingDataSourceSchema()
}

func (d *LsnclientNetworkBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnclientNetworkBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	clientname_Name := data.Clientname.ValueString()
	netmask_Name := data.Netmask
	network_Name := data.Network
	td_Name := data.Td

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsnclient_network_binding.Type(),
		ResourceName:             clientname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_network_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsnclient_network_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

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

		// Check network
		if val, ok := v["network"].(string); ok {
			if network_Name.IsNull() || val != network_Name.ValueString() {
				match = false
				continue
			}
		} else if !network_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsnclient_network_binding with netmask %s not found", netmask_Name))
		return
	}

	lsnclient_network_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
