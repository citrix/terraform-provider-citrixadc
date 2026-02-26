package appfwprofile_xmlwsiurl_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileXmlwsiurlBindingDataSource)(nil)

func APpfwprofileXmlwsiurlBindingDataSource() datasource.DataSource {
	return &AppfwprofileXmlwsiurlBindingDataSource{}
}

type AppfwprofileXmlwsiurlBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileXmlwsiurlBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_xmlwsiurl_binding"
}

func (d *AppfwprofileXmlwsiurlBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileXmlwsiurlBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileXmlwsiurlBindingDataSourceSchema()
}

func (d *AppfwprofileXmlwsiurlBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	xmlwsiurl_Name := data.Xmlwsiurl

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_xmlwsiurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_xmlwsiurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_xmlwsiurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check xmlwsiurl
		if val, ok := v["xmlwsiurl"].(string); ok {
			if xmlwsiurl_Name.IsNull() || val != xmlwsiurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !xmlwsiurl_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_xmlwsiurl_binding with xmlwsiurl %s not found", xmlwsiurl_Name))
		return
	}

	appfwprofile_xmlwsiurl_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
