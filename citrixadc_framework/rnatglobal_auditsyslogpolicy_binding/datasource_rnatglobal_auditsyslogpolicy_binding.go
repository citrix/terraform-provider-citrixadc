package rnatglobal_auditsyslogpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*RnatglobalAuditsyslogpolicyBindingDataSource)(nil)

func RNatglobalAuditsyslogpolicyBindingDataSource() datasource.DataSource {
	return &RnatglobalAuditsyslogpolicyBindingDataSource{}
}

type RnatglobalAuditsyslogpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *RnatglobalAuditsyslogpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnatglobal_auditsyslogpolicy_binding"
}

func (d *RnatglobalAuditsyslogpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *RnatglobalAuditsyslogpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = RnatglobalAuditsyslogpolicyBindingDataSourceSchema()
}

func (d *RnatglobalAuditsyslogpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	policy_Name := data.Policy

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Rnatglobal_auditsyslogpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "rnatglobal_auditsyslogpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policy
		if val, ok := v["policy"].(string); ok {
			if policy_Name.IsNull() || val != policy_Name.ValueString() {
				match = false
				continue
			}
		} else if !policy_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("rnatglobal_auditsyslogpolicy_binding with policy %s not found", policy_Name))
		return
	}

	rnatglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
