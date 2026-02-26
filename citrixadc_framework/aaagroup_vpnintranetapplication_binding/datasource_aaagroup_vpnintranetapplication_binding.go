package aaagroup_vpnintranetapplication_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AaagroupVpnintranetapplicationBindingDataSource)(nil)

func AAagroupVpnintranetapplicationBindingDataSource() datasource.DataSource {
	return &AaagroupVpnintranetapplicationBindingDataSource{}
}

type AaagroupVpnintranetapplicationBindingDataSource struct {
	client *service.NitroClient
}

func (d *AaagroupVpnintranetapplicationBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_vpnintranetapplication_binding"
}

func (d *AaagroupVpnintranetapplicationBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AaagroupVpnintranetapplicationBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AaagroupVpnintranetapplicationBindingDataSourceSchema()
}

func (d *AaagroupVpnintranetapplicationBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	groupname_Name := data.Groupname.ValueString()
	intranetapplication_Name := data.Intranetapplication

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Aaagroup_vpnintranetapplication_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_vpnintranetapplication_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "aaagroup_vpnintranetapplication_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check intranetapplication
		if val, ok := v["intranetapplication"].(string); ok {
			if intranetapplication_Name.IsNull() || val != intranetapplication_Name.ValueString() {
				match = false
				continue
			}
		} else if !intranetapplication_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("aaagroup_vpnintranetapplication_binding with intranetapplication %s not found", intranetapplication_Name))
		return
	}

	aaagroup_vpnintranetapplication_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
