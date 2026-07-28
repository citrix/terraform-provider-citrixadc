package botprofile_kmdetectionexpr_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BotprofileKmdetectionexprBindingDataSource)(nil)

func BOtprofileKmdetectionexprBindingDataSource() datasource.DataSource {
	return &BotprofileKmdetectionexprBindingDataSource{}
}

type BotprofileKmdetectionexprBindingDataSource struct {
	client *service.NitroClient
}

func (d *BotprofileKmdetectionexprBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_kmdetectionexpr_binding"
}

func (d *BotprofileKmdetectionexprBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BotprofileKmdetectionexprBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BotprofileKmdetectionexprBindingDataSourceSchema()
}

func (d *BotprofileKmdetectionexprBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BotprofileKmdetectionexprBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	botkmexpressionname_Name := data.BotKmExpressionName
	kmdetectionexpr_Name := data.Kmdetectionexpr

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Botprofile_kmdetectionexpr_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_kmdetectionexpr_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "botprofile_kmdetectionexpr_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_km_expression_name
		if val, ok := v["bot_km_expression_name"].(string); ok {
			if botkmexpressionname_Name.IsNull() || val != botkmexpressionname_Name.ValueString() {
				match = false
				continue
			}
		} else if !botkmexpressionname_Name.IsNull() {
			match = false
			continue
		}

		// Check kmdetectionexpr
		if val, ok := v["kmdetectionexpr"].(bool); ok {
			if kmdetectionexpr_Name.IsNull() || val != kmdetectionexpr_Name.ValueBool() {
				match = false
				continue
			}
		} else if !kmdetectionexpr_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("botprofile_kmdetectionexpr_binding with bot_km_expression_name %s not found", botkmexpressionname_Name))
		return
	}

	botprofile_kmdetectionexpr_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
