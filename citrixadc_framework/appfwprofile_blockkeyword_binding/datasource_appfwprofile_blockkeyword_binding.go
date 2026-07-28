package appfwprofile_blockkeyword_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileBlockkeywordBindingDataSource)(nil)

func APpfwprofileBlockkeywordBindingDataSource() datasource.DataSource {
	return &AppfwprofileBlockkeywordBindingDataSource{}
}

type AppfwprofileBlockkeywordBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileBlockkeywordBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_blockkeyword_binding"
}

func (d *AppfwprofileBlockkeywordBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileBlockkeywordBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileBlockkeywordBindingDataSourceSchema()
}

func (d *AppfwprofileBlockkeywordBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileBlockkeywordBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asblockkeywordformurl_Name := data.AsBlockkeywordFormurl
	blockkeyword_Name := data.Blockkeyword
	fieldname_Name := data.Fieldname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_blockkeyword_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_blockkeyword_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_blockkeyword_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_blockkeyword_formurl
		if val, ok := v["as_blockkeyword_formurl"].(string); ok {
			if asblockkeywordformurl_Name.IsNull() || val != asblockkeywordformurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !asblockkeywordformurl_Name.IsNull() {
			match = false
			continue
		}

		// Check blockkeyword
		if val, ok := v["blockkeyword"].(string); ok {
			if blockkeyword_Name.IsNull() || val != blockkeyword_Name.ValueString() {
				match = false
				continue
			}
		} else if !blockkeyword_Name.IsNull() {
			match = false
			continue
		}

		// Check fieldname
		if val, ok := v["fieldname"].(string); ok {
			if fieldname_Name.IsNull() || val != fieldname_Name.ValueString() {
				match = false
				continue
			}
		} else if !fieldname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_blockkeyword_binding with as_blockkeyword_formurl %s not found", asblockkeywordformurl_Name))
		return
	}

	appfwprofile_blockkeyword_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
