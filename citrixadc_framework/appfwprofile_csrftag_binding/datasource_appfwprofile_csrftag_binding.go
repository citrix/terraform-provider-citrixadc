package appfwprofile_csrftag_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileCsrftagBindingDataSource)(nil)

func APpfwprofileCsrftagBindingDataSource() datasource.DataSource {
	return &AppfwprofileCsrftagBindingDataSource{}
}

type AppfwprofileCsrftagBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileCsrftagBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_csrftag_binding"
}

func (d *AppfwprofileCsrftagBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileCsrftagBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileCsrftagBindingDataSourceSchema()
}

func (d *AppfwprofileCsrftagBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileCsrftagBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	csrfformactionurl_Name := data.Csrfformactionurl
	csrftag_Name := data.Csrftag

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_csrftag_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_csrftag_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_csrftag_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check csrfformactionurl
		if val, ok := v["csrfformactionurl"].(string); ok {
			if csrfformactionurl_Name.IsNull() || val != csrfformactionurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !csrfformactionurl_Name.IsNull() {
			match = false
			continue
		}

		// Check csrftag
		if val, ok := v["csrftag"].(string); ok {
			if csrftag_Name.IsNull() || val != csrftag_Name.ValueString() {
				match = false
				continue
			}
		} else if !csrftag_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_csrftag_binding with csrfformactionurl %s not found", csrfformactionurl_Name))
		return
	}

	appfwprofile_csrftag_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
