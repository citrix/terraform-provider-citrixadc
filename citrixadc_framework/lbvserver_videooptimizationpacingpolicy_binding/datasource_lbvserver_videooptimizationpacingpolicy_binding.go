package lbvserver_videooptimizationpacingpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LbvserverVideooptimizationpacingpolicyBindingDataSource)(nil)

func LBvserverVideooptimizationpacingpolicyBindingDataSource() datasource.DataSource {
	return &LbvserverVideooptimizationpacingpolicyBindingDataSource{}
}

type LbvserverVideooptimizationpacingpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *LbvserverVideooptimizationpacingpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_videooptimizationpacingpolicy_binding"
}

func (d *LbvserverVideooptimizationpacingpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LbvserverVideooptimizationpacingpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LbvserverVideooptimizationpacingpolicyBindingDataSourceSchema()
}

func (d *LbvserverVideooptimizationpacingpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LbvserverVideooptimizationpacingpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	bindpoint_Name := data.Bindpoint
	policyname_Name := data.Policyname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lbvserver_videooptimizationpacingpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_videooptimizationpacingpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lbvserver_videooptimizationpacingpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bindpoint
		if val, ok := v["bindpoint"].(string); ok {
			if bindpoint_Name.IsNull() || val != bindpoint_Name.ValueString() {
				match = false
				continue
			}
		} else if !bindpoint_Name.IsNull() {
			match = false
			continue
		}

		// Check policyname
		if val, ok := v["policyname"].(string); ok {
			if policyname_Name.IsNull() || val != policyname_Name.ValueString() {
				match = false
				continue
			}
		} else if !policyname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lbvserver_videooptimizationpacingpolicy_binding with bindpoint %s not found", bindpoint_Name))
		return
	}

	lbvserver_videooptimizationpacingpolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
