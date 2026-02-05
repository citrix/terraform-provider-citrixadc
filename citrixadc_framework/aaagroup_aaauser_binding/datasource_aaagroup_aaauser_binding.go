package aaagroup_aaauser_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AaagroupAaauserBindingDataSource)(nil)

func AAagroupAaauserBindingDataSource() datasource.DataSource {
	return &AaagroupAaauserBindingDataSource{}
}

type AaagroupAaauserBindingDataSource struct {
	client *service.NitroClient
}

func (d *AaagroupAaauserBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_aaauser_binding"
}

func (d *AaagroupAaauserBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AaagroupAaauserBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AaagroupAaauserBindingDataSourceSchema()
}

func (d *AaagroupAaauserBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AaagroupAaauserBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	aaagroup_Name := data.Groupname.ValueString()
	aaauser_Name := data.Username.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Aaagroup_aaauser_binding.Type(),
		ResourceName:             aaagroup_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_aaauser_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "aaagroup_aaauser_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["username"].(string) == aaauser_Name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("aaagroup_aaauser_binding with username %s not found", aaauser_Name))
		return
	}

	aaagroup_aaauser_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
