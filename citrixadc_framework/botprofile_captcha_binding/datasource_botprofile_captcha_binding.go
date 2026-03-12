package botprofile_captcha_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BotprofileCaptchaBindingDataSource)(nil)

func BOtprofileCaptchaBindingDataSource() datasource.DataSource {
	return &BotprofileCaptchaBindingDataSource{}
}

type BotprofileCaptchaBindingDataSource struct {
	client *service.NitroClient
}

func (d *BotprofileCaptchaBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_captcha_binding"
}

func (d *BotprofileCaptchaBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BotprofileCaptchaBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BotprofileCaptchaBindingDataSourceSchema()
}

func (d *BotprofileCaptchaBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BotprofileCaptchaBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	botcaptchaurl_Name := data.BotCaptchaUrl

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Botprofile_captcha_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_captcha_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "botprofile_captcha_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_captcha_url
		if val, ok := v["bot_captcha_url"].(string); ok {
			if botcaptchaurl_Name.IsNull() || val != botcaptchaurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !botcaptchaurl_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("botprofile_captcha_binding with bot_captcha_url %s not found", botcaptchaurl_Name))
		return
	}

	botprofile_captcha_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
