package botprofile_ratelimit_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BotprofileRatelimitBindingDataSource)(nil)

func BOtprofileRatelimitBindingDataSource() datasource.DataSource {
	return &BotprofileRatelimitBindingDataSource{}
}

type BotprofileRatelimitBindingDataSource struct {
	client *service.NitroClient
}

func (d *BotprofileRatelimitBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_ratelimit_binding"
}

func (d *BotprofileRatelimitBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BotprofileRatelimitBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BotprofileRatelimitBindingDataSourceSchema()
}

func (d *BotprofileRatelimitBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BotprofileRatelimitBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	botratelimittype_Name := data.BotRateLimitType
	botratelimiturl_Name := data.BotRateLimitUrl
	condition_Name := data.Condition
	cookiename_Name := data.Cookiename
	countrycode_Name := data.Countrycode

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Botprofile_ratelimit_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_ratelimit_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "botprofile_ratelimit_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_rate_limit_type
		if val, ok := v["bot_rate_limit_type"].(string); ok {
			if botratelimittype_Name.IsNull() || val != botratelimittype_Name.ValueString() {
				match = false
				continue
			}
		} else if !botratelimittype_Name.IsNull() {
			match = false
			continue
		}

		// Check bot_rate_limit_url
		if val, ok := v["bot_rate_limit_url"].(string); ok {
			if botratelimiturl_Name.IsNull() || val != botratelimiturl_Name.ValueString() {
				match = false
				continue
			}
		} else if !botratelimiturl_Name.IsNull() {
			match = false
			continue
		}

		// Check condition
		if val, ok := v["condition"].(string); ok {
			if condition_Name.IsNull() || val != condition_Name.ValueString() {
				match = false
				continue
			}
		} else if !condition_Name.IsNull() {
			match = false
			continue
		}

		// Check cookiename
		if val, ok := v["cookiename"].(string); ok {
			if cookiename_Name.IsNull() || val != cookiename_Name.ValueString() {
				match = false
				continue
			}
		} else if !cookiename_Name.IsNull() {
			match = false
			continue
		}

		// Check countrycode
		if val, ok := v["countrycode"].(string); ok {
			if countrycode_Name.IsNull() || val != countrycode_Name.ValueString() {
				match = false
				continue
			}
		} else if !countrycode_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("botprofile_ratelimit_binding with bot_rate_limit_type %s not found", botratelimittype_Name))
		return
	}

	botprofile_ratelimit_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
