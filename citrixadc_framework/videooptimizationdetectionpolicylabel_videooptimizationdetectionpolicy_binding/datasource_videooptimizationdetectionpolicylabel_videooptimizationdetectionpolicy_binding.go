package videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

var _ datasource.DataSource = (*VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource)(nil)

func VIdeooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource() datasource.DataSource {
	return &VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource{}
}

type VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding"
}

func (d *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSourceSchema()
}

func (d *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	labelname_Name := data.Labelname.ValueString()
	policyname_Name := data.Policyname
	priority_Name := data.Priority

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.Type(),
		ResourceName:             labelname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

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

		// Check priority
		if val, ok := v["priority"]; ok {
			val, _ = utils.ConvertToInt64(val)
			if priority_Name.IsNull() || val != priority_Name.ValueInt64() {
				match = false
				continue
			}
		} else if !priority_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding with policyname %s not found", policyname_Name))
		return
	}

	videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}