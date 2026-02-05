package aaaglobal_aaapreauthenticationpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AaaglobalAaapreauthenticationpolicyBindingDataSource)(nil)

func AAaglobalAaapreauthenticationpolicyBindingDataSource() datasource.DataSource {
	return &AaaglobalAaapreauthenticationpolicyBindingDataSource{}
}

type AaaglobalAaapreauthenticationpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AaaglobalAaapreauthenticationpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaglobal_aaapreauthenticationpolicy_binding"
}

func (d *AaaglobalAaapreauthenticationpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AaaglobalAaapreauthenticationpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AaaglobalAaapreauthenticationpolicyBindingDataSourceSchema()
}

func (d *AaaglobalAaapreauthenticationpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AaaglobalAaapreauthenticationpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	policyName := data.Policy.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "aaaglobal_aaapreauthenticationpolicy_binding",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read aaaglobal_aaapreauthenticationpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "aaaglobal_aaapreauthenticationpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policy"].(string) == policyName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("aaaglobal_aaapreauthenticationpolicy_binding with policy %s not found", policyName))
		return
	}

	// Fallthrough

	aaaglobal_aaapreauthenticationpolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
