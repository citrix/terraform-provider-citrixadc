package lsnlogprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LsnlogprofileResourceModel describes the resource data model.
type LsnlogprofileResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Analyticsprofile types.String `tfsdk:"analyticsprofile"`
	Logcompact       types.String `tfsdk:"logcompact"`
	Logipfix         types.String `tfsdk:"logipfix"`
	Logprofilename   types.String `tfsdk:"logprofilename"`
	Logsessdeletion  types.String `tfsdk:"logsessdeletion"`
	Logsubscrinfo    types.String `tfsdk:"logsubscrinfo"`
}

func (r *LsnlogprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnlogprofile resource.",
			},
			"analyticsprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Analytics Profile attached to this lsn profile.",
			},
			"logcompact": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Logs in Compact Logging format if option is enabled.",
			},
			"logipfix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Logs in IPFIX  format if option is enabled.",
			},
			"logprofilename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the logging Profile.",
			},
			"logsessdeletion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "LSN Session deletion will not be logged if disabled.",
			},
			"logsubscrinfo": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Subscriber ID information is logged if option is enabled.",
			},
		},
	}
}

func lsnlogprofileGetThePayloadFromthePlan(ctx context.Context, data *LsnlogprofileResourceModel) lsn.Lsnlogprofile {
	tflog.Debug(ctx, "In lsnlogprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lsnlogprofile := lsn.Lsnlogprofile{}
	if !data.Analyticsprofile.IsNull() && !data.Analyticsprofile.IsUnknown() {
		lsnlogprofile.Analyticsprofile = data.Analyticsprofile.ValueString()
	}
	if !data.Logcompact.IsNull() && !data.Logcompact.IsUnknown() {
		lsnlogprofile.Logcompact = data.Logcompact.ValueString()
	}
	if !data.Logipfix.IsNull() && !data.Logipfix.IsUnknown() {
		lsnlogprofile.Logipfix = data.Logipfix.ValueString()
	}
	if !data.Logprofilename.IsNull() && !data.Logprofilename.IsUnknown() {
		lsnlogprofile.Logprofilename = data.Logprofilename.ValueString()
	}
	if !data.Logsessdeletion.IsNull() && !data.Logsessdeletion.IsUnknown() {
		lsnlogprofile.Logsessdeletion = data.Logsessdeletion.ValueString()
	}
	if !data.Logsubscrinfo.IsNull() && !data.Logsubscrinfo.IsUnknown() {
		lsnlogprofile.Logsubscrinfo = data.Logsubscrinfo.ValueString()
	}

	return lsnlogprofile
}

func lsnlogprofileSetAttrFromGet(ctx context.Context, data *LsnlogprofileResourceModel, getResponseData map[string]interface{}) *LsnlogprofileResourceModel {
	tflog.Debug(ctx, "In lsnlogprofileSetAttrFromGet Function")

	// Convert API response to model. Guard the else-branch to only null a value
	// when it is unknown (never clobber a known configured value that NITRO may
	// omit from the GET response - the omit-on-default trap).
	if val, ok := getResponseData["analyticsprofile"]; ok && val != nil {
		data.Analyticsprofile = types.StringValue(val.(string))
	} else if data.Analyticsprofile.IsUnknown() {
		data.Analyticsprofile = types.StringNull()
	}
	if val, ok := getResponseData["logcompact"]; ok && val != nil {
		data.Logcompact = types.StringValue(val.(string))
	} else if data.Logcompact.IsUnknown() {
		data.Logcompact = types.StringNull()
	}
	if val, ok := getResponseData["logipfix"]; ok && val != nil {
		data.Logipfix = types.StringValue(val.(string))
	} else if data.Logipfix.IsUnknown() {
		data.Logipfix = types.StringNull()
	}
	if val, ok := getResponseData["logprofilename"]; ok && val != nil {
		data.Logprofilename = types.StringValue(val.(string))
	} else if data.Logprofilename.IsUnknown() {
		data.Logprofilename = types.StringNull()
	}
	if val, ok := getResponseData["logsessdeletion"]; ok && val != nil {
		data.Logsessdeletion = types.StringValue(val.(string))
	} else if data.Logsessdeletion.IsUnknown() {
		data.Logsessdeletion = types.StringNull()
	}
	if val, ok := getResponseData["logsubscrinfo"]; ok && val != nil {
		data.Logsubscrinfo = types.StringValue(val.(string))
	} else if data.Logsubscrinfo.IsUnknown() {
		data.Logsubscrinfo = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Logprofilename.ValueString())

	return data
}
