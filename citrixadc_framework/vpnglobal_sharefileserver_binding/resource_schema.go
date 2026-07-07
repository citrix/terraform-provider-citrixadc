package vpnglobal_sharefileserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalSharefileserverBindingResourceModel describes the resource data model.
type VpnglobalSharefileserverBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Sharefile              types.String `tfsdk:"sharefile"`
}

func (r *VpnglobalSharefileserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_sharefileserver_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Pattern 13: gotopriorityexpression is a pure user input that the NITRO
				// GET for this binding does not echo back. Keeping Computed would force
				// "known after apply"/inconsistent-result churn; it is Optional only,
				// matching the SDK v2 user-facing contract (omitting it leaves it unset).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"sharefile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Configured Sharefile server, in the format IP:PORT / FQDN:PORT",
			},
		},
	}
}

func vpnglobal_sharefileserver_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalSharefileserverBindingResourceModel) vpn.Vpnglobalsharefileserverbinding {
	tflog.Debug(ctx, "In vpnglobal_sharefileserver_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_sharefileserver_binding := vpn.Vpnglobalsharefileserverbinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_sharefileserver_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Sharefile.IsNull() && !data.Sharefile.IsUnknown() {
		vpnglobal_sharefileserver_binding.Sharefile = data.Sharefile.ValueString()
	}

	return vpnglobal_sharefileserver_binding
}

// vpnglobal_sharefileserver_bindingSetAttrFromGet is the RESOURCE setter: it preserves
// user-supplied inputs that the NITRO GET does not echo back. The binding GET only
// returns "sharefile" (and stateflag); it never echoes "gotopriorityexpression", so we
// MUST NOT null it here or Terraform reports "inconsistent result after apply"
// (Pattern 7 — server-non-echoed input). The id is set once in Create, not here.
func vpnglobal_sharefileserver_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalSharefileserverBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalSharefileserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_sharefileserver_bindingSetAttrFromGet Function")

	// gotopriorityexpression is a write-only input not echoed by GET — preserve the
	// existing plan/state value; do not overwrite or null it.

	if val, ok := getResponseData["sharefile"]; ok && val != nil {
		data.Sharefile = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Sharefile.ValueString()))

	return data
}

// vpnglobal_sharefileserver_bindingSetAttrFromGetForDatasource is the DATASOURCE setter
// (Pattern 7 datasource split): a datasource has no prior plan/state, so it faithfully
// copies every field from the GET response and sets its own id (the datasource never
// calls Create). gotopriorityexpression is not echoed by GET, so it is set to null.
func vpnglobal_sharefileserver_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalSharefileserverBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalSharefileserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_sharefileserver_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["sharefile"]; ok && val != nil {
		data.Sharefile = types.StringValue(val.(string))
	} else {
		data.Sharefile = types.StringNull()
	}

	// Set ID for the datasource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Sharefile.ValueString()))

	return data
}
