package pcpprofile

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/resource/config/pcp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PcpprofileResourceModel describes the resource data model.
type PcpprofileResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Announcemulticount types.String `tfsdk:"announcemulticount"`
	Mapping            types.String `tfsdk:"mapping"`
	Maxmaplife         types.Int64  `tfsdk:"maxmaplife"`
	Minmaplife         types.Int64  `tfsdk:"minmaplife"`
	Name               types.String `tfsdk:"name"`
	Peer               types.String `tfsdk:"peer"`
	Thirdparty         types.String `tfsdk:"thirdparty"`
}

func (r *PcpprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the pcpprofile resource.",
			},
			// SDK v2 typed announcemulticount as TypeString (Optional+Computed, no default).
			// Preserve that contract: StringAttribute, no Default.
			"announcemulticount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10"),
				Description: "Integer value that identify the number announce message to be send.",
			},
			"mapping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "This argument is for enabling/disabling the MAP opcode  of current PCP Profile",
			},
			"maxmaplife": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(86400),
				Description: "Integer value that identify the maximum mapping lifetime (in seconds) for a pcp profile. default(86400s = 24Hours).",
			},
			"minmaplife": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Integer value that identify the minimum mapping lifetime (in seconds) for a pcp profile. default(120s)",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2 name is ForceNew -> RequiresReplace.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the PCP Profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my pcpProfile\" or my pcpProfile).",
			},
			"peer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "This argument is for enabling/disabling the PEER opcode of current PCP Profile",
			},
			"thirdparty": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This argument is for enabling/disabling the THIRD PARTY opcode of current PCP Profile",
			},
		},
	}
}

func pcpprofileGetThePayloadFromtheConfig(ctx context.Context, data *PcpprofileResourceModel) pcp.Pcpprofile {
	tflog.Debug(ctx, "In pcpprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	pcpprofile := pcp.Pcpprofile{}
	if !data.Announcemulticount.IsNull() && !data.Announcemulticount.IsUnknown() {
		// announcemulticount is a string in the schema (SDK v2 parity) but an int in NITRO.
		if intVal, err := strconv.Atoi(data.Announcemulticount.ValueString()); err == nil {
			pcpprofile.Announcemulticount = utils.IntPtr(intVal)
		}
	}
	if !data.Mapping.IsNull() && !data.Mapping.IsUnknown() {
		pcpprofile.Mapping = data.Mapping.ValueString()
	}
	if !data.Maxmaplife.IsNull() && !data.Maxmaplife.IsUnknown() {
		pcpprofile.Maxmaplife = utils.IntPtr(int(data.Maxmaplife.ValueInt64()))
	}
	if !data.Minmaplife.IsNull() && !data.Minmaplife.IsUnknown() {
		pcpprofile.Minmaplife = utils.IntPtr(int(data.Minmaplife.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		pcpprofile.Name = data.Name.ValueString()
	}
	if !data.Peer.IsNull() && !data.Peer.IsUnknown() {
		pcpprofile.Peer = data.Peer.ValueString()
	}
	if !data.Thirdparty.IsNull() && !data.Thirdparty.IsUnknown() {
		pcpprofile.Thirdparty = data.Thirdparty.ValueString()
	}

	return pcpprofile
}

func pcpprofileSetAttrFromGet(ctx context.Context, data *PcpprofileResourceModel, getResponseData map[string]interface{}) *PcpprofileResourceModel {
	tflog.Debug(ctx, "In pcpprofileSetAttrFromGet Function")

	// Convert API response to model.
	// Else-branches only null a value that is still Unknown, never clobber a known
	// configured value that NITRO happens to omit (omit-on-default guard).
	if val, ok := getResponseData["announcemulticount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Announcemulticount = types.StringValue(strconv.FormatInt(intVal, 10))
		} else {
			data.Announcemulticount = types.StringValue(fmt.Sprintf("%v", val))
		}
	} else if data.Announcemulticount.IsUnknown() {
		data.Announcemulticount = types.StringNull()
	}
	if val, ok := getResponseData["mapping"]; ok && val != nil {
		data.Mapping = types.StringValue(val.(string))
	} else if data.Mapping.IsUnknown() {
		data.Mapping = types.StringNull()
	}
	if val, ok := getResponseData["maxmaplife"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxmaplife = types.Int64Value(intVal)
		}
	} else if data.Maxmaplife.IsUnknown() {
		data.Maxmaplife = types.Int64Null()
	}
	if val, ok := getResponseData["minmaplife"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minmaplife = types.Int64Value(intVal)
		}
	} else if data.Minmaplife.IsUnknown() {
		data.Minmaplife = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["peer"]; ok && val != nil {
		data.Peer = types.StringValue(val.(string))
	} else if data.Peer.IsUnknown() {
		data.Peer = types.StringNull()
	}
	if val, ok := getResponseData["thirdparty"]; ok && val != nil {
		data.Thirdparty = types.StringValue(val.(string))
	} else if data.Thirdparty.IsUnknown() {
		data.Thirdparty = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute (name) - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
