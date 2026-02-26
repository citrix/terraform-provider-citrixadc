package analyticsglobal_analyticsprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AnalyticsglobalAnalyticsprofileBindingDataSource)(nil)

func ANalyticsglobalAnalyticsprofileBindingDataSource() datasource.DataSource {
	return &AnalyticsglobalAnalyticsprofileBindingDataSource{}
}

type AnalyticsglobalAnalyticsprofileBindingDataSource struct {
	client *service.NitroClient
}

func (d *AnalyticsglobalAnalyticsprofileBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analyticsglobal_analyticsprofile_binding"
}

func (d *AnalyticsglobalAnalyticsprofileBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AnalyticsglobalAnalyticsprofileBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AnalyticsglobalAnalyticsprofileBindingDataSourceSchema()
}

func (d *AnalyticsglobalAnalyticsprofileBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AnalyticsglobalAnalyticsprofileBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	analyticsprofile_Name := data.Analyticsprofile

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Analyticsglobal_analyticsprofile_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read analyticsglobal_analyticsprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "analyticsglobal_analyticsprofile_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check analyticsprofile
		if val, ok := v["analyticsprofile"].(string); ok {
			if analyticsprofile_Name.IsNull() || val != analyticsprofile_Name.ValueString() {
				match = false
				continue
			}
		} else if !analyticsprofile_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("analyticsglobal_analyticsprofile_binding with analyticsprofile %s not found", analyticsprofile_Name))
		return
	}

	analyticsglobal_analyticsprofile_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
