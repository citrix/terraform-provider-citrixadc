package appfwprofile_excluderescontenttype_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileExcluderescontenttypeBindingDataSource)(nil)

func APpfwprofileExcluderescontenttypeBindingDataSource() datasource.DataSource {
	return &AppfwprofileExcluderescontenttypeBindingDataSource{}
}

type AppfwprofileExcluderescontenttypeBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileExcluderescontenttypeBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_excluderescontenttype_binding"
}

func (d *AppfwprofileExcluderescontenttypeBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileExcluderescontenttypeBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileExcluderescontenttypeBindingDataSourceSchema()
}

func (d *AppfwprofileExcluderescontenttypeBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileExcluderescontenttypeBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	excluderescontenttype_Name := data.Excluderescontenttype

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_excluderescontenttype_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_excluderescontenttype_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_excluderescontenttype_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check excluderescontenttype
		if val, ok := v["excluderescontenttype"].(string); ok {
			if excluderescontenttype_Name.IsNull() || val != excluderescontenttype_Name.ValueString() {
				match = false
				continue
			}
		} else if !excluderescontenttype_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_excluderescontenttype_binding with excluderescontenttype %s not found", excluderescontenttype_Name))
		return
	}

	appfwprofile_excluderescontenttype_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
