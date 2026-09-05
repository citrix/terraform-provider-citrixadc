package aaauser_vpnsecureprivateaccessprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AaauserVpnsecureprivateaccessprofileBindingDataSource)(nil)

func AAauserVpnsecureprivateaccessprofileBindingDataSource() datasource.DataSource {
	return &AaauserVpnsecureprivateaccessprofileBindingDataSource{}
}

type AaauserVpnsecureprivateaccessprofileBindingDataSource struct {
	client *service.NitroClient
}

func (d *AaauserVpnsecureprivateaccessprofileBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaauser_vpnsecureprivateaccessprofile_binding"
}

func (d *AaauserVpnsecureprivateaccessprofileBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AaauserVpnsecureprivateaccessprofileBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AaauserVpnsecureprivateaccessprofileBindingDataSourceSchema()
}

func (d *AaauserVpnsecureprivateaccessprofileBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AaauserVpnsecureprivateaccessprofileBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	username_Name := data.Username.ValueString()
	secureprivateaccessprofile_Name := data.Secureprivateaccessprofile

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Aaauser_vpnsecureprivateaccessprofile_binding.Type(),
		ResourceName:             username_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read aaauser_vpnsecureprivateaccessprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "aaauser_vpnsecureprivateaccessprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check secureprivateaccessprofile
		if val, ok := v["secureprivateaccessprofile"].(string); ok {
			if secureprivateaccessprofile_Name.IsNull() || val != secureprivateaccessprofile_Name.ValueString() {
				match = false
				continue
			}
		} else if !secureprivateaccessprofile_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("aaauser_vpnsecureprivateaccessprofile_binding with secureprivateaccessprofile %s not found", secureprivateaccessprofile_Name))
		return
	}

	aaauser_vpnsecureprivateaccessprofile_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
