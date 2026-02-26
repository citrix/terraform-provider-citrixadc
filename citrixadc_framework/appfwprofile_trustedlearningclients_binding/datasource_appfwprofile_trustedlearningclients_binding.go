package appfwprofile_trustedlearningclients_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileTrustedlearningclientsBindingDataSource)(nil)

func APpfwprofileTrustedlearningclientsBindingDataSource() datasource.DataSource {
	return &AppfwprofileTrustedlearningclientsBindingDataSource{}
}

type AppfwprofileTrustedlearningclientsBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileTrustedlearningclientsBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_trustedlearningclients_binding"
}

func (d *AppfwprofileTrustedlearningclientsBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileTrustedlearningclientsBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileTrustedlearningclientsBindingDataSourceSchema()
}

func (d *AppfwprofileTrustedlearningclientsBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	trustedlearningclients_Name := data.Trustedlearningclients

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_trustedlearningclients_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_trustedlearningclients_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_trustedlearningclients_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check trustedlearningclients
		if val, ok := v["trustedlearningclients"].(string); ok {
			if trustedlearningclients_Name.IsNull() || val != trustedlearningclients_Name.ValueString() {
				match = false
				continue
			}
		} else if !trustedlearningclients_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_trustedlearningclients_binding with trustedlearningclients %s not found", trustedlearningclients_Name))
		return
	}

	appfwprofile_trustedlearningclients_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
