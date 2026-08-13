package botpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotpolicylabelResourceModel describes the resource data model.
type BotpolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
}

func (r *BotpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				// Matches SDK v2 (Optional + Computed + ForceNew). Computed keeps an
				// existing SDK v2 state value (or empty) stable across refresh when the
				// attribute is omitted from config, avoiding a spurious replace on
				// migration. SetAttrFromGet guards the unknown case (see defect 6).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: create-only (has_update_op False) - keep prior state
					// on unknown to avoid spurious replace on upgrade; only replace
					// when the user actually configures a changed value.
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Any comments to preserve information about this bot policy label.",
			},
			"labelname": schema.StringAttribute{
				// SDK v2: Required + ForceNew -> Required + RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the bot policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder policy label\" or my responder policy label').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "New name for the bot policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
		},
	}
}

func botpolicylabelGetThePayloadFromthePlan(ctx context.Context, data *BotpolicylabelResourceModel) bot.Botpolicylabel {
	tflog.Debug(ctx, "In botpolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	botpolicylabel := bot.Botpolicylabel{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		botpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		botpolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.

	return botpolicylabel
}

func botpolicylabelSetAttrFromGet(ctx context.Context, data *BotpolicylabelResourceModel, getResponseData map[string]interface{}) *BotpolicylabelResourceModel {
	tflog.Debug(ctx, "In botpolicylabelSetAttrFromGet Function")

	// Convert API response to model.
	// Defect 6 guard: NITRO omits "comment" from the GET response when it is empty.
	// Only overwrite from GET when present; when absent, preserve a configured value
	// (avoids inconsistent-result) but resolve an unknown Computed value to null
	// (avoids "still indicated an unknown value").
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	// labelname is the user-facing key. After a rename (via newname) the live object
	// name (tracked by data.Id) diverges from the configured labelname, and GET
	// returns the live (new) name. Overwriting labelname from GET would clobber the
	// user's configured value and trigger a spurious RequiresReplace diff. Only adopt
	// the GET value when we don't already have one (e.g. on import, where state carries
	// only the ID); otherwise preserve.
	if data.Labelname.IsNull() || data.Labelname.IsUnknown() || data.Labelname.ValueString() == "" {
		if val, ok := getResponseData["labelname"]; ok && val != nil {
			data.Labelname = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.

	return data
}

// botpolicylabelSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func botpolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *BotpolicylabelResourceModel, getResponseData map[string]interface{}) *BotpolicylabelResourceModel {
	tflog.Debug(ctx, "In botpolicylabelSetAttrFromGetForDatasource Function")

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

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
