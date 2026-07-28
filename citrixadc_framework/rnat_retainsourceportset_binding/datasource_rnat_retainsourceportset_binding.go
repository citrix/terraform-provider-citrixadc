package rnat_retainsourceportset_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*RnatRetainsourceportsetBindingDataSource)(nil)

func RNatRetainsourceportsetBindingDataSource() datasource.DataSource {
	return &RnatRetainsourceportsetBindingDataSource{}
}

type RnatRetainsourceportsetBindingDataSource struct {
	client *service.NitroClient
}

func (d *RnatRetainsourceportsetBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat_retainsourceportset_binding"
}

func (d *RnatRetainsourceportsetBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *RnatRetainsourceportsetBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = RnatRetainsourceportsetBindingDataSourceSchema()
}

func (d *RnatRetainsourceportsetBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RnatRetainsourceportsetBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	retainsourceportrange_Name := data.Retainsourceportrange

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Rnat_retainsourceportset_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read rnat_retainsourceportset_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "rnat_retainsourceportset_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check retainsourceportrange
		if val, ok := v["retainsourceportrange"].(string); ok {
			if retainsourceportrange_Name.IsNull() || val != retainsourceportrange_Name.ValueString() {
				match = false
				continue
			}
		} else if !retainsourceportrange_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("rnat_retainsourceportset_binding with retainsourceportrange %s not found", retainsourceportrange_Name))
		return
	}

	rnat_retainsourceportset_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
