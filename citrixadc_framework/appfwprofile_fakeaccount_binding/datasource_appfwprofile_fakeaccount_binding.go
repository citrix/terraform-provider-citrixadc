package appfwprofile_fakeaccount_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileFakeaccountBindingDataSource)(nil)

func APpfwprofileFakeaccountBindingDataSource() datasource.DataSource {
	return &AppfwprofileFakeaccountBindingDataSource{}
}

type AppfwprofileFakeaccountBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileFakeaccountBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fakeaccount_binding"
}

func (d *AppfwprofileFakeaccountBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileFakeaccountBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileFakeaccountBindingDataSourceSchema()
}

func (d *AppfwprofileFakeaccountBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileFakeaccountBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	fakeaccount_Name := data.Fakeaccount
	formexpression_Name := data.Formexpression
	formurlfad_Name := data.FormurlFad
	tag_Name := data.Tag

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_fakeaccount_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fakeaccount_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_fakeaccount_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check fakeaccount
		if val, ok := v["fakeaccount"].(string); ok {
			if fakeaccount_Name.IsNull() || val != fakeaccount_Name.ValueString() {
				match = false
				continue
			}
		} else if !fakeaccount_Name.IsNull() {
			match = false
			continue
		}

		// Check formexpression (at-most-one arm). Only filter when the caller
		// supplied a non-empty value; an empty-string or null arm is tolerated
		// (the empty-string workaround in older configs is harmless). If the
		// caller did NOT specify this arm but the record carries a non-empty
		// value for it, this is not their record.
		if !formexpression_Name.IsNull() && formexpression_Name.ValueString() != "" {
			if val, ok := v["formexpression"].(string); !ok || val != formexpression_Name.ValueString() {
				match = false
				continue
			}
		} else if val, ok := v["formexpression"].(string); ok && val != "" {
			match = false
			continue
		}

		// Check formurl_fad (at-most-one arm). Same tolerant logic as above.
		if !formurlfad_Name.IsNull() && formurlfad_Name.ValueString() != "" {
			if val, ok := v["formurl_fad"].(string); !ok || val != formurlfad_Name.ValueString() {
				match = false
				continue
			}
		} else if val, ok := v["formurl_fad"].(string); ok && val != "" {
			match = false
			continue
		}

		// Check tag
		if val, ok := v["tag"].(string); ok {
			if tag_Name.IsNull() || val != tag_Name.ValueString() {
				match = false
				continue
			}
		} else if !tag_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_fakeaccount_binding with fakeaccount %s not found", fakeaccount_Name))
		return
	}

	appfwprofile_fakeaccount_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
