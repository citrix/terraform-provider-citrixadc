package lbvserver_analyticsprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LbvserverAnalyticsprofileBindingDataSource)(nil)

func LBvserverAnalyticsprofileBindingDataSource() datasource.DataSource {
	return &LbvserverAnalyticsprofileBindingDataSource{}
}

type LbvserverAnalyticsprofileBindingDataSource struct {
	client *service.NitroClient
}

func (d *LbvserverAnalyticsprofileBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_analyticsprofile_binding"
}

func (d *LbvserverAnalyticsprofileBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LbvserverAnalyticsprofileBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LbvserverAnalyticsprofileBindingDataSourceSchema()
}

func (d *LbvserverAnalyticsprofileBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LbvserverAnalyticsprofileBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	analyticsprofile_Name := data.Analyticsprofile

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lbvserver_analyticsprofile_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_analyticsprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lbvserver_analyticsprofile_binding returned empty array.")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lbvserver_analyticsprofile_binding with analyticsprofile %s not found", analyticsprofile_Name))
		return
	}

	lbvserver_analyticsprofile_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
