package vxlanvlanmap_vxlan_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VxlanvlanmapVxlanBindingDataSource)(nil)

func VXlanvlanmapVxlanBindingDataSource() datasource.DataSource {
	return &VxlanvlanmapVxlanBindingDataSource{}
}

type VxlanvlanmapVxlanBindingDataSource struct {
	client *service.NitroClient
}

func (d *VxlanvlanmapVxlanBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlanvlanmap_vxlan_binding"
}

func (d *VxlanvlanmapVxlanBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VxlanvlanmapVxlanBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VxlanvlanmapVxlanBindingDataSourceSchema()
}

func (d *VxlanvlanmapVxlanBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	vxlan_Name := data.Vxlan

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vxlanvlanmap_vxlan_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vxlanvlanmap_vxlan_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vxlanvlanmap_vxlan_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check vxlan
		if val, ok := v["vxlan"]; ok {
			val, _ = utils.ConvertToInt64(val)
			if vxlan_Name.IsNull() || val != vxlan_Name.ValueInt64() {
				match = false
				continue
			}
		} else if !vxlan_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vxlanvlanmap_vxlan_binding with vxlan %s not found", vxlan_Name))
		return
	}

	vxlanvlanmap_vxlan_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
