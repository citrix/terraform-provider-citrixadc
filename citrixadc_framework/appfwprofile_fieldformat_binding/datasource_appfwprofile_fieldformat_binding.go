package appfwprofile_fieldformat_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileFieldformatBindingDataSource)(nil)

func APpfwprofileFieldformatBindingDataSource() datasource.DataSource {
	return &AppfwprofileFieldformatBindingDataSource{}
}

type AppfwprofileFieldformatBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileFieldformatBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fieldformat_binding"
}

func (d *AppfwprofileFieldformatBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileFieldformatBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileFieldformatBindingDataSourceSchema()
}

func (d *AppfwprofileFieldformatBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileFieldformatBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	fieldformat_Name := data.Fieldformat
	formactionurlff_Name := data.FormactionurlFf

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_fieldformat_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fieldformat_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_fieldformat_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check fieldformat
		if val, ok := v["fieldformat"].(string); ok {
			if fieldformat_Name.IsNull() || val != fieldformat_Name.ValueString() {
				match = false
				continue
			}
		} else if !fieldformat_Name.IsNull() {
			match = false
			continue
		}

		// Check formactionurl_ff
		if val, ok := v["formactionurl_ff"].(string); ok {
			if formactionurlff_Name.IsNull() || val != formactionurlff_Name.ValueString() {
				match = false
				continue
			}
		} else if !formactionurlff_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_fieldformat_binding with fieldformat %s not found", fieldformat_Name))
		return
	}

	appfwprofile_fieldformat_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
