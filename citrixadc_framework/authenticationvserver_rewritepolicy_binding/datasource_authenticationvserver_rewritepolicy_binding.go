package authenticationvserver_rewritepolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AuthenticationvserverRewritepolicyBindingDataSource)(nil)

func AUthenticationvserverRewritepolicyBindingDataSource() datasource.DataSource {
	return &AuthenticationvserverRewritepolicyBindingDataSource{}
}

type AuthenticationvserverRewritepolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AuthenticationvserverRewritepolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_rewritepolicy_binding"
}

func (d *AuthenticationvserverRewritepolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AuthenticationvserverRewritepolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AuthenticationvserverRewritepolicyBindingDataSourceSchema()
}

func (d *AuthenticationvserverRewritepolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AuthenticationvserverRewritepolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Lookup keys: parent (name) for the GET, then filter by policy and (if set) bindpoint.
	name_Name := data.Name.ValueString()
	bindpoint_Name := data.Bindpoint
	policy_Name := data.Policy

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Authenticationvserver_rewritepolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_rewritepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "authenticationvserver_rewritepolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching policy (and bindpoint if set).
	foundIndex := -1
	for i, v := range dataArr {
		// Match policy
		if val, ok := v["policy"].(string); ok {
			if policy_Name.IsNull() || val != policy_Name.ValueString() {
				continue
			}
		} else {
			continue
		}

		// Match bindpoint when supplied
		if !bindpoint_Name.IsNull() {
			if val, ok := v["bindpoint"].(string); !ok || val != bindpoint_Name.ValueString() {
				continue
			}
		}

		foundIndex = i
		break
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("authenticationvserver_rewritepolicy_binding with policy %s not found", policy_Name))
		return
	}

	authenticationvserver_rewritepolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
