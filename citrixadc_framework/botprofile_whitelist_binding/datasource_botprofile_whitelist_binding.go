package botprofile_whitelist_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BotprofileWhitelistBindingDataSource)(nil)

func BOtprofileWhitelistBindingDataSource() datasource.DataSource {
	return &BotprofileWhitelistBindingDataSource{}
}

type BotprofileWhitelistBindingDataSource struct {
	client *service.NitroClient
}

func (d *BotprofileWhitelistBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_whitelist_binding"
}

func (d *BotprofileWhitelistBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BotprofileWhitelistBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BotprofileWhitelistBindingDataSourceSchema()
}

func (d *BotprofileWhitelistBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BotprofileWhitelistBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	botwhitelistvalue_Name := data.BotWhitelistValue

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Botprofile_whitelist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_whitelist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "botprofile_whitelist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_whitelist_value
		if val, ok := v["bot_whitelist_value"].(string); ok {
			if botwhitelistvalue_Name.IsNull() || val != botwhitelistvalue_Name.ValueString() {
				match = false
				continue
			}
		} else if !botwhitelistvalue_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("botprofile_whitelist_binding with bot_whitelist_value %s not found", botwhitelistvalue_Name))
		return
	}

	botprofile_whitelist_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
