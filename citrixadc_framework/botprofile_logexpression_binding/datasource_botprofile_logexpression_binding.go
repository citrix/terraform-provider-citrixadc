package botprofile_logexpression_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BotprofileLogexpressionBindingDataSource)(nil)

func BOtprofileLogexpressionBindingDataSource() datasource.DataSource {
	return &BotprofileLogexpressionBindingDataSource{}
}

type BotprofileLogexpressionBindingDataSource struct {
	client *service.NitroClient
}

func (d *BotprofileLogexpressionBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_logexpression_binding"
}

func (d *BotprofileLogexpressionBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BotprofileLogexpressionBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BotprofileLogexpressionBindingDataSourceSchema()
}

func (d *BotprofileLogexpressionBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BotprofileLogexpressionBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	botlogexpressionname_Name := data.BotLogExpressionName

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Botprofile_logexpression_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_logexpression_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "botprofile_logexpression_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_log_expression_name
		if val, ok := v["bot_log_expression_name"].(string); ok {
			if botlogexpressionname_Name.IsNull() || val != botlogexpressionname_Name.ValueString() {
				match = false
				continue
			}
		} else if !botlogexpressionname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("botprofile_logexpression_binding with bot_log_expression_name %s not found", botlogexpressionname_Name))
		return
	}

	botprofile_logexpression_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
