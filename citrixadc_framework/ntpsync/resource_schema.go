package ntpsync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NtpsyncResourceModel describes the resource data model.
//
// ntpsync is an action-only singleton: NITRO exposes only enable/disable
// actions plus a read-only "state" property (ENABLED/DISABLED). The SDK v2
// resource had a single Required "state" attribute driving the enable/disable
// action, so the Framework model preserves that exact contract.
type NtpsyncResourceModel struct {
	Id    types.String `tfsdk:"id"`
	State types.String `tfsdk:"state"`
}

func (r *NtpsyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ntpsync resource.",
			},
			// SDK v2 parity: "state" is Required (Type: schema.TypeString,
			// Required: true) with no Default and no ForceNew. It is the
			// user's desired state which drives the NITRO enable/disable
			// action, and is also echoed back by GET.
			"state": schema.StringAttribute{
				Required:    true,
				Description: "Show NTP status. Possible values = ENABLED, DISABLED",
			},
		},
	}
}

// ntpsyncSetAttrFromGet maps the NITRO GET response onto the model.
//
// Only "state" is read back (SDK v2 read exactly "state"). The internal
// read-only "_nextgenapiresource" NITRO field is intentionally not exposed —
// SDK v2 never surfaced it, so adding it would be a backward-compat break.
// We only overwrite "state" when the GET actually returns it; we never null a
// Required attribute the API omits (omit-on-default guard).
func ntpsyncSetAttrFromGet(ctx context.Context, data *NtpsyncResourceModel, getResponseData map[string]interface{}) *NtpsyncResourceModel {
	tflog.Debug(ctx, "In ntpsyncSetAttrFromGet Function")

	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	}
	// else: preserve the existing configured/state value — never clobber the
	// Required "state" attribute with null if NITRO omits it from GET.

	return data
}
