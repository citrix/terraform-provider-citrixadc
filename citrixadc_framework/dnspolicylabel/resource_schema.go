package dnspolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnspolicylabelResourceModel describes the resource data model.
type DnspolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
	Transform types.String `tfsdk:"transform"`
}

func (r *DnspolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnspolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew (primary key).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the dns policy label.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "The new name of the dns policylabel.",
			},
			"transform": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The type of transformations allowed by the policies bound to the label.",
			},
		},
	}
}

func dnspolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *DnspolicylabelResourceModel) dns.Dnspolicylabel {
	tflog.Debug(ctx, "In dnspolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnspolicylabel := dns.Dnspolicylabel{}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		dnspolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Transform.IsNull() && !data.Transform.IsUnknown() {
		dnspolicylabel.Transform = data.Transform.ValueString()
	}

	return dnspolicylabel
}

func dnspolicylabelSetAttrFromGet(ctx context.Context, data *DnspolicylabelResourceModel, getResponseData map[string]interface{}) *DnspolicylabelResourceModel {
	tflog.Debug(ctx, "In dnspolicylabelSetAttrFromGet Function")

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
	if val, ok := getResponseData["transform"]; ok && val != nil {
		data.Transform = types.StringValue(val.(string))
	} else {
		data.Transform = types.StringNull()
	}

	return data
}

// dnspolicylabelSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func dnspolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *DnspolicylabelResourceModel, getResponseData map[string]interface{}) *DnspolicylabelResourceModel {
	tflog.Debug(ctx, "In dnspolicylabelSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["transform"]; ok && val != nil {
		data.Transform = types.StringValue(val.(string))
	} else {
		data.Transform = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
