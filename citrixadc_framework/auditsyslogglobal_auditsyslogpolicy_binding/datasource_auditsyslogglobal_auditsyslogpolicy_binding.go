package auditsyslogglobal_auditsyslogpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AuditsyslogglobalAuditsyslogpolicyBindingDataSource)(nil)

func AUditsyslogglobalAuditsyslogpolicyBindingDataSource() datasource.DataSource {
	return &AuditsyslogglobalAuditsyslogpolicyBindingDataSource{}
}

type AuditsyslogglobalAuditsyslogpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AuditsyslogglobalAuditsyslogpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditsyslogglobal_auditsyslogpolicy_binding"
}

func (d *AuditsyslogglobalAuditsyslogpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AuditsyslogglobalAuditsyslogpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AuditsyslogglobalAuditsyslogpolicyBindingDataSourceSchema()
}

func (d *AuditsyslogglobalAuditsyslogpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AuditsyslogglobalAuditsyslogpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	globalbindtype_Name := data.Globalbindtype
	policyname_Name := data.Policyname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Auditsyslogglobal_auditsyslogpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read auditsyslogglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "auditsyslogglobal_auditsyslogpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check globalbindtype
		if val, ok := v["globalbindtype"].(string); ok {
			if globalbindtype_Name.IsNull() || val != globalbindtype_Name.ValueString() {
				match = false
				continue
			}
		} else if !globalbindtype_Name.IsNull() {
			match = false
			continue
		}

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("auditsyslogglobal_auditsyslogpolicy_binding with globalbindtype %s not found", globalbindtype_Name))
		return
	}

	auditsyslogglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
