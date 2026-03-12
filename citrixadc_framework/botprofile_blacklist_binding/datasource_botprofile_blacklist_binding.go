package botprofile_blacklist_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BotprofileBlacklistBindingDataSource)(nil)

func BOtprofileBlacklistBindingDataSource() datasource.DataSource {
	return &BotprofileBlacklistBindingDataSource{}
}

type BotprofileBlacklistBindingDataSource struct {
	client *service.NitroClient
}

func (d *BotprofileBlacklistBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_blacklist_binding"
}

func (d *BotprofileBlacklistBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BotprofileBlacklistBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BotprofileBlacklistBindingDataSourceSchema()
}

func (d *BotprofileBlacklistBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BotprofileBlacklistBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	botblacklist_Name := data.BotBlacklist
	botblacklistvalue_Name := data.BotBlacklistValue

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Botprofile_blacklist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_blacklist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "botprofile_blacklist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_blacklist_value
		if val, ok := v["bot_blacklist_value"].(string); ok {
			if botblacklistvalue_Name.IsNull() || val != botblacklistvalue_Name.ValueString() {
				match = false
				continue
			}
		} else if !botblacklistvalue_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("botprofile_blacklist_binding with bot_blacklist %s not found", botblacklist_Name))
		return
	}

	botprofile_blacklist_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
