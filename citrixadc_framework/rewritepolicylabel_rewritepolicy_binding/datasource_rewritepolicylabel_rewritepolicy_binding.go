package rewritepolicylabel_rewritepolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*RewritepolicylabelRewritepolicyBindingDataSource)(nil)

func REwritepolicylabelRewritepolicyBindingDataSource() datasource.DataSource {
	return &RewritepolicylabelRewritepolicyBindingDataSource{}
}

type RewritepolicylabelRewritepolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *RewritepolicylabelRewritepolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewritepolicylabel_rewritepolicy_binding"
}

func (d *RewritepolicylabelRewritepolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *RewritepolicylabelRewritepolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = RewritepolicylabelRewritepolicyBindingDataSourceSchema()
}

func (d *RewritepolicylabelRewritepolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	labelname_Name := data.Labelname.ValueString()
	policyname_Name := data.Policyname
	priority_Name := data.Priority

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Rewritepolicylabel_rewritepolicy_binding.Type(),
		ResourceName:             labelname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read rewritepolicylabel_rewritepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "rewritepolicylabel_rewritepolicy_binding returned empty array.")
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

		// Check priority - only filter when the user supplied a priority value.
		if !priority_Name.IsNull() {
			if val, ok := v["priority"]; ok {
				val, _ = utils.ConvertToInt64(val)
				if val != priority_Name.ValueInt64() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("rewritepolicylabel_rewritepolicy_binding with policyname %s not found", policyname_Name))
		return
	}

	rewritepolicylabel_rewritepolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// Datasource has no Create, so set the ID here (matches the resource ID format:
	// comma-separated key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
