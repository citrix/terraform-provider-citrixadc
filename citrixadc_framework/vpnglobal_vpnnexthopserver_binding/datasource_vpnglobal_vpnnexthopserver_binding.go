package vpnglobal_vpnnexthopserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnglobalVpnnexthopserverBindingDataSource)(nil)

func VPnglobalVpnnexthopserverBindingDataSource() datasource.DataSource {
	return &VpnglobalVpnnexthopserverBindingDataSource{}
}

type VpnglobalVpnnexthopserverBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnglobalVpnnexthopserverBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnnexthopserver_binding"
}

func (d *VpnglobalVpnnexthopserverBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnglobalVpnnexthopserverBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnglobalVpnnexthopserverBindingDataSourceSchema()
}

func (d *VpnglobalVpnnexthopserverBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnglobalVpnnexthopserverBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	nexthopserver_Name := data.Nexthopserver

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_vpnnexthopserver_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnnexthopserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_vpnnexthopserver_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check nexthopserver
		if val, ok := v["nexthopserver"].(string); ok {
			if nexthopserver_Name.IsNull() || val != nexthopserver_Name.ValueString() {
				match = false
				continue
			}
		} else if !nexthopserver_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnglobal_vpnnexthopserver_binding with nexthopserver %s not found", nexthopserver_Name))
		return
	}

	vpnglobal_vpnnexthopserver_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
