package auditnslogglobal_auditnslogpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AuditnslogglobalAuditnslogpolicyBindingDataSource)(nil)

func AUditnslogglobalAuditnslogpolicyBindingDataSource() datasource.DataSource {
	return &AuditnslogglobalAuditnslogpolicyBindingDataSource{}
}

type AuditnslogglobalAuditnslogpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AuditnslogglobalAuditnslogpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditnslogglobal_auditnslogpolicy_binding"
}

func (d *AuditnslogglobalAuditnslogpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AuditnslogglobalAuditnslogpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AuditnslogglobalAuditnslogpolicyBindingDataSourceSchema()
}

func (d *AuditnslogglobalAuditnslogpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AuditnslogglobalAuditnslogpolicyBindingResourceModel
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
		ResourceType:             service.Auditnslogglobal_auditnslogpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read auditnslogglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "auditnslogglobal_auditnslogpolicy_binding returned empty array")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("auditnslogglobal_auditnslogpolicy_binding with globalbindtype %s not found", globalbindtype_Name))
		return
	}

	auditnslogglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
