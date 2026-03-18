package appfwprofile_cookieconsistency_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileCookieconsistencyBindingDataSource)(nil)

func APpfwprofileCookieconsistencyBindingDataSource() datasource.DataSource {
	return &AppfwprofileCookieconsistencyBindingDataSource{}
}

type AppfwprofileCookieconsistencyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileCookieconsistencyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_cookieconsistency_binding"
}

func (d *AppfwprofileCookieconsistencyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileCookieconsistencyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileCookieconsistencyBindingDataSourceSchema()
}

func (d *AppfwprofileCookieconsistencyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileCookieconsistencyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	cookieconsistency_Name := data.Cookieconsistency

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_cookieconsistency_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_cookieconsistency_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_cookieconsistency_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check cookieconsistency
		if val, ok := v["cookieconsistency"].(string); ok {
			if cookieconsistency_Name.IsNull() || val != cookieconsistency_Name.ValueString() {
				match = false
				continue
			}
		} else if !cookieconsistency_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_cookieconsistency_binding with cookieconsistency %s not found", cookieconsistency_Name))
		return
	}

	appfwprofile_cookieconsistency_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
