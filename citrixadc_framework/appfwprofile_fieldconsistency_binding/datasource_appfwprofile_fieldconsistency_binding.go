package appfwprofile_fieldconsistency_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileFieldconsistencyBindingDataSource)(nil)

func APpfwprofileFieldconsistencyBindingDataSource() datasource.DataSource {
	return &AppfwprofileFieldconsistencyBindingDataSource{}
}

type AppfwprofileFieldconsistencyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileFieldconsistencyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fieldconsistency_binding"
}

func (d *AppfwprofileFieldconsistencyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileFieldconsistencyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileFieldconsistencyBindingDataSourceSchema()
}

func (d *AppfwprofileFieldconsistencyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	fieldconsistency_Name := data.Fieldconsistency
	formactionurlffc_Name := data.FormactionurlFfc

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_fieldconsistency_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fieldconsistency_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_fieldconsistency_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check fieldconsistency
		if val, ok := v["fieldconsistency"].(string); ok {
			if fieldconsistency_Name.IsNull() || val != fieldconsistency_Name.ValueString() {
				match = false
				continue
			}
		} else if !fieldconsistency_Name.IsNull() {
			match = false
			continue
		}

		// Check formactionurl_ffc
		if val, ok := v["formactionurl_ffc"].(string); ok {
			if formactionurlffc_Name.IsNull() || val != formactionurlffc_Name.ValueString() {
				match = false
				continue
			}
		} else if !formactionurlffc_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_fieldconsistency_binding with fieldconsistency %s not found", fieldconsistency_Name))
		return
	}

	appfwprofile_fieldconsistency_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
