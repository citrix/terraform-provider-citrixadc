package lbpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbpolicylabelResourceModel describes the resource data model.
type LbpolicylabelResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Comment         types.String `tfsdk:"comment"`
	Labelname       types.String `tfsdk:"labelname"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`
}

func (r *LbpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				// NOT Computed: NITRO has no server-side default for comment and the GET
				// response omits it entirely when empty. An Optional+Computed attribute
				// that the server neither defaults nor echoes stays UNKNOWN after apply,
				// which the framework rejects ("still indicated an unknown value").
				// Optional-only -> unset yields a known null; when set it is echoed back
				// by GET and read into state by SetAttrFromGet.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this LB policy label.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy label\" or 'my lb policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
				Description: "New name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocols supported by the policylabel. Available Types are :\n* HTTP - HTTP requests.\n* DNS - DNS request.\n* OTHERTCP - OTHERTCP request.\n* SIP_UDP - SIP_UDP request.\n* SIP_TCP - SIP_TCP request.\n* MYSQL - MYSQL request.\n* MSSQL - MSSQL request.\n* ORACLE - ORACLE request.\n* NAT - NAT request.\n* DIAMETER - DIAMETER request.\n* RADIUS - RADIUS request.\n* MQTT - MQTT request.\n* QUIC_BRIDGE - QUIC_BRIDGE request.\n* HTTP_QUIC - HTTP_QUIC request.",
			},
		},
	}
}

func lbpolicylabelGetThePayloadFromthePlan(ctx context.Context, data *LbpolicylabelResourceModel) lb.Lbpolicylabel {
	tflog.Debug(ctx, "In lbpolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbpolicylabel := lb.Lbpolicylabel{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		lbpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		lbpolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Policylabeltype.IsNull() && !data.Policylabeltype.IsUnknown() {
		lbpolicylabel.Policylabeltype = data.Policylabeltype.ValueString()
	}

	return lbpolicylabel
}

func lbpolicylabelSetAttrFromGet(ctx context.Context, data *LbpolicylabelResourceModel, getResponseData map[string]interface{}) *LbpolicylabelResourceModel {
	tflog.Debug(ctx, "In lbpolicylabelSetAttrFromGet Function")

	// Convert API response to model.
	// Pattern 7: NITRO omits "comment" from the GET response when it is empty, so
	// only overwrite the model value when the field is actually present. Otherwise
	// preserve the plan/state value to avoid a perpetual diff / nulled user input.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	}
	// labelname is the user-facing key. Once a rename has happened (via newname),
	// the live object name (tracked by data.Id) diverges from the configured
	// labelname, and GET returns the live (new) name. Overwriting labelname from
	// GET would clobber the user's configured value and trigger a spurious
	// RequiresReplace diff. So only adopt the GET value when we don't already have
	// one (e.g. on import, where state carries only the ID); otherwise preserve.
	if data.Labelname.IsNull() || data.Labelname.IsUnknown() || data.Labelname.ValueString() == "" {
		if val, ok := getResponseData["labelname"]; ok && val != nil {
			data.Labelname = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["policylabeltype"]; ok && val != nil {
		data.Policylabeltype = types.StringValue(val.(string))
	}

	return data
}

// lbpolicylabelSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func lbpolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *LbpolicylabelResourceModel, getResponseData map[string]interface{}) *LbpolicylabelResourceModel {
	tflog.Debug(ctx, "In lbpolicylabelSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["policylabeltype"]; ok && val != nil {
		data.Policylabeltype = types.StringValue(val.(string))
	} else {
		data.Policylabeltype = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
