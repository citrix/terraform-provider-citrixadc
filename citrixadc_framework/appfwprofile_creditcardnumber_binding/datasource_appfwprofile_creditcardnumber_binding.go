package appfwprofile_creditcardnumber_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileCreditcardnumberBindingDataSource)(nil)

func APpfwprofileCreditcardnumberBindingDataSource() datasource.DataSource {
	return &AppfwprofileCreditcardnumberBindingDataSource{}
}

type AppfwprofileCreditcardnumberBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileCreditcardnumberBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_creditcardnumber_binding"
}

func (d *AppfwprofileCreditcardnumberBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileCreditcardnumberBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileCreditcardnumberBindingDataSourceSchema()
}

func (d *AppfwprofileCreditcardnumberBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileCreditcardnumberBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	creditcardnumber_Name := data.Creditcardnumber
	creditcardnumberurl_Name := data.Creditcardnumberurl

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_creditcardnumber_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_creditcardnumber_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_creditcardnumber_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check creditcardnumber
		if val, ok := v["creditcardnumber"].(string); ok {
			if creditcardnumber_Name.IsNull() || val != creditcardnumber_Name.ValueString() {
				match = false
				continue
			}
		} else if !creditcardnumber_Name.IsNull() {
			match = false
			continue
		}

		// Check creditcardnumberurl
		if val, ok := v["creditcardnumberurl"].(string); ok {
			if creditcardnumberurl_Name.IsNull() || val != creditcardnumberurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !creditcardnumberurl_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_creditcardnumber_binding with creditcardnumber %s not found", creditcardnumber_Name))
		return
	}

	appfwprofile_creditcardnumber_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
