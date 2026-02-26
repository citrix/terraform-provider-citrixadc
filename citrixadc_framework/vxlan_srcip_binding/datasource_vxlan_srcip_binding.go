package vxlan_srcip_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VxlanSrcipBindingDataSource)(nil)

func VXlanSrcipBindingDataSource() datasource.DataSource {
	return &VxlanSrcipBindingDataSource{}
}

type VxlanSrcipBindingDataSource struct {
	client *service.NitroClient
}

func (d *VxlanSrcipBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlan_srcip_binding"
}

func (d *VxlanSrcipBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VxlanSrcipBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VxlanSrcipBindingDataSourceSchema()
}

func (d *VxlanSrcipBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VxlanSrcipBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	id_Name := fmt.Sprintf("%d", data.Vxlanid.ValueInt64())
	srcip_Name := data.Srcip

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vxlan_srcip_binding.Type(),
		ResourceName:             id_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vxlan_srcip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vxlan_srcip_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check srcip
		if val, ok := v["srcip"].(string); ok {
			if srcip_Name.IsNull() || val != srcip_Name.ValueString() {
				match = false
				continue
			}
		} else if !srcip_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vxlan_srcip_binding with srcip %s not found", srcip_Name))
		return
	}

	vxlan_srcip_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
