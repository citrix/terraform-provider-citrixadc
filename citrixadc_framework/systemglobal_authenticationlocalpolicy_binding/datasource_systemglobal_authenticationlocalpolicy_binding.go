package systemglobal_authenticationlocalpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SystemglobalAuthenticationlocalpolicyBindingDataSource)(nil)

func SYstemglobalAuthenticationlocalpolicyBindingDataSource() datasource.DataSource {
	return &SystemglobalAuthenticationlocalpolicyBindingDataSource{}
}

type SystemglobalAuthenticationlocalpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *SystemglobalAuthenticationlocalpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_authenticationlocalpolicy_binding"
}

func (d *SystemglobalAuthenticationlocalpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SystemglobalAuthenticationlocalpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SystemglobalAuthenticationlocalpolicyBindingDataSourceSchema()
}

func (d *SystemglobalAuthenticationlocalpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SystemglobalAuthenticationlocalpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	policyname_Name := data.Policyname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Systemglobal_authenticationlocalpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_authenticationlocalpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "systemglobal_authenticationlocalpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if val, ok := v["policyname"].(string); ok {
			if policyname_Name.IsNull() || val != policyname_Name.ValueString() {
				match = false
				continue
			}
		} else if !policyname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("systemglobal_authenticationlocalpolicy_binding with policyname %s not found", policyname_Name))
		return
	}

	systemglobal_authenticationlocalpolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
