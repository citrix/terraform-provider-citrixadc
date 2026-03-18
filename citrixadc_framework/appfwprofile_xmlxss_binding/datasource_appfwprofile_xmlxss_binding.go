package appfwprofile_xmlxss_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileXmlxssBindingDataSource)(nil)

func APpfwprofileXmlxssBindingDataSource() datasource.DataSource {
	return &AppfwprofileXmlxssBindingDataSource{}
}

type AppfwprofileXmlxssBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileXmlxssBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_xmlxss_binding"
}

func (d *AppfwprofileXmlxssBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileXmlxssBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileXmlxssBindingDataSourceSchema()
}

func (d *AppfwprofileXmlxssBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileXmlxssBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asscanlocationxmlxss_Name := data.AsScanLocationXmlxss
	xmlxss_Name := data.Xmlxss

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_xmlxss_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_xmlxss_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_xmlxss_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_scan_location_xmlxss
		if val, ok := v["as_scan_location_xmlxss"].(string); ok {
			if asscanlocationxmlxss_Name.IsNull() || val != asscanlocationxmlxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !asscanlocationxmlxss_Name.IsNull() {
			match = false
			continue
		}

		// Check xmlxss
		if val, ok := v["xmlxss"].(string); ok {
			if xmlxss_Name.IsNull() || val != xmlxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !xmlxss_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_xmlxss_binding with as_scan_location_xmlxss %s not found", asscanlocationxmlxss_Name))
		return
	}

	appfwprofile_xmlxss_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
