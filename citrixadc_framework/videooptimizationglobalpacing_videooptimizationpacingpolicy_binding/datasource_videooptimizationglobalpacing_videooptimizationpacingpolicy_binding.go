package videooptimizationglobalpacing_videooptimizationpacingpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource)(nil)

func VIdeooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource() datasource.DataSource {
	return &VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource{}
}

type VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding"
}

func (d *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSourceSchema()
}

func (d *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	policyname_Name := data.Policyname
	priority_Name := data.Priority
	type_Name := data.Type

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	var err error
	if !type_Name.IsNull() && type_Name.ValueString() != "" {
		argsMap["type"] = type_Name.ValueString()
	}

	findParams := service.FindParams{
		ResourceType:             service.Videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationglobalpacing_videooptimizationpacingpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "videooptimizationglobalpacing_videooptimizationpacingpolicy_binding returned empty array")
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
		// Check type_Name
		if !type_Name.IsNull() && type_Name.ValueString() != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != type_Name.ValueString() {
					match = false
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("videooptimizationglobalpacing_videooptimizationpacingpolicy_binding with policyname %s not found", policyname_Name))
		return
	}

	videooptimizationglobalpacing_videooptimizationpacingpolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
