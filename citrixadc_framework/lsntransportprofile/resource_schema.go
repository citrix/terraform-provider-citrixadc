package lsntransportprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

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

// LsntransportprofileResourceModel describes the resource data model.
type LsntransportprofileResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Finrsttimeout        types.Int64  `tfsdk:"finrsttimeout"`
	Groupsessionlimit    types.Int64  `tfsdk:"groupsessionlimit"`
	Portpreserveparity   types.String `tfsdk:"portpreserveparity"`
	Portpreserverange    types.String `tfsdk:"portpreserverange"`
	Portquota            types.Int64  `tfsdk:"portquota"`
	Sessionquota         types.Int64  `tfsdk:"sessionquota"`
	Sessiontimeout       types.Int64  `tfsdk:"sessiontimeout"`
	Stuntimeout          types.Int64  `tfsdk:"stuntimeout"`
	Syncheck             types.String `tfsdk:"syncheck"`
	Synidletimeout       types.Int64  `tfsdk:"synidletimeout"`
	Transportprofilename types.String `tfsdk:"transportprofilename"`
	Transportprotocol    types.String `tfsdk:"transportprotocol"`
}

func (r *LsntransportprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsntransportprofile resource.",
			},
			"finrsttimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Timeout, in seconds, for a TCP LSN session after a FIN or RST message is received from one of the endpoints.\n\nIf a TCP LSN session is idle (after the Citrix ADC receives a FIN or RST message) for a time that exceeds this value, the Citrix ADC ADC removes the session.\n\nSince the LSN feature of the Citrix ADC does not maintain state information of any TCP LSN sessions, this timeout accommodates the transmission of the FIN or RST, and ACK messages from the other endpoint so that both endpoints can properly close the connection.",
			},
			"groupsessionlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent LSN sessions(for the specified protocol) allowed for all subscriber of a group to which this profile has bound. This limit will get split across the Citrix ADCs packet engines and rounded down. When the number of LSN sessions reaches the limit for a group in packet engine, the Citrix ADC does not allow the subscriber of that group to open additional sessions through that packet engine.",
			},
			"portpreserveparity": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable port parity between a subscriber port and its mapped LSN NAT port. For example, if a subscriber initiates a connection from an odd numbered port, the Citrix ADC allocates an odd numbered LSN NAT port for this connection. \nYou must set this parameter for proper functioning of protocols that require the source port to be even or odd numbered, for example, in peer-to-peer applications that use RTP or RTCP protocol.",
			},
			"portpreserverange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If a subscriber initiates a connection from a well-known port (0-1023), allocate a NAT port from the well-known port range (0-1023) for this connection. For example, if a subscriber initiates a connection from port 80, the Citrix ADC can allocate port 100 as the NAT port for this connection.\n\nThis parameter applies to dynamic NAT without port block allocation. It also applies to Deterministic NAT if the range of ports allocated includes well-known ports.\n\nWhen all the well-known ports of all the available NAT IP addresses are used in different subscriber's connections (LSN sessions), and a subscriber initiates a connection from a well-known port, the Citrix ADC drops this connection.",
			},
			"portquota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of LSN NAT ports to be used at a time by each subscriber for the specified protocol. For example, each subscriber can be limited to a maximum of 500 TCP NAT ports. When the LSN NAT mappings for a subscriber reach the limit, the Citrix ADC does not allocate additional NAT ports for that subscriber.",
			},
			"sessionquota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent LSN sessions allowed for each subscriber for the specified protocol. \nWhen the number of LSN sessions reaches the limit for a subscriber, the Citrix ADC does not allow the subscriber to open additional sessions.",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Timeout, in seconds, for an idle LSN session. If an LSN session is idle for a time that exceeds this value, the Citrix ADC removes the session.\n\nThis timeout does not apply for a TCP LSN session when a FIN or RST message is received from either of the endpoints.",
			},
			"stuntimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(600),
				Description: "STUN protocol timeout",
			},
			"syncheck": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Silently drop any non-SYN packets for connections for which there is no LSN-NAT session present on the Citrix ADC. \n\nIf you disable this parameter, the Citrix ADC accepts any non-SYN packets and creates a new LSN session entry for this connection. \n\nFollowing are some reasons for the Citrix ADC to receive such packets:\n\n* LSN session for a connection existed but the Citrix ADC removed this session because the LSN session was idle for a time that exceeded the configured session timeout.\n* Such packets can be a part of a DoS attack.",
			},
			"synidletimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "SYN Idle timeout",
			},
			"transportprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN transport profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN transport profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn transport profile1\" or 'lsn transport profile1').",
			},
			"transportprotocol": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol for which to set the LSN transport profile parameters.",
			},
		},
	}
}

func lsntransportprofileGetThePayloadFromtheConfig(ctx context.Context, data *LsntransportprofileResourceModel) lsn.Lsntransportprofile {
	tflog.Debug(ctx, "In lsntransportprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsntransportprofile := lsn.Lsntransportprofile{}
	if !data.Finrsttimeout.IsNull() {
		lsntransportprofile.Finrsttimeout = utils.IntPtr(int(data.Finrsttimeout.ValueInt64()))
	}
	if !data.Groupsessionlimit.IsNull() {
		lsntransportprofile.Groupsessionlimit = utils.IntPtr(int(data.Groupsessionlimit.ValueInt64()))
	}
	if !data.Portpreserveparity.IsNull() {
		lsntransportprofile.Portpreserveparity = data.Portpreserveparity.ValueString()
	}
	if !data.Portpreserverange.IsNull() {
		lsntransportprofile.Portpreserverange = data.Portpreserverange.ValueString()
	}
	if !data.Portquota.IsNull() {
		lsntransportprofile.Portquota = utils.IntPtr(int(data.Portquota.ValueInt64()))
	}
	if !data.Sessionquota.IsNull() {
		lsntransportprofile.Sessionquota = utils.IntPtr(int(data.Sessionquota.ValueInt64()))
	}
	if !data.Sessiontimeout.IsNull() {
		lsntransportprofile.Sessiontimeout = utils.IntPtr(int(data.Sessiontimeout.ValueInt64()))
	}
	if !data.Stuntimeout.IsNull() {
		lsntransportprofile.Stuntimeout = utils.IntPtr(int(data.Stuntimeout.ValueInt64()))
	}
	if !data.Syncheck.IsNull() {
		lsntransportprofile.Syncheck = data.Syncheck.ValueString()
	}
	if !data.Synidletimeout.IsNull() {
		lsntransportprofile.Synidletimeout = utils.IntPtr(int(data.Synidletimeout.ValueInt64()))
	}
	if !data.Transportprofilename.IsNull() {
		lsntransportprofile.Transportprofilename = data.Transportprofilename.ValueString()
	}
	if !data.Transportprotocol.IsNull() {
		lsntransportprofile.Transportprotocol = data.Transportprotocol.ValueString()
	}

	return lsntransportprofile
}

func lsntransportprofileSetAttrFromGet(ctx context.Context, data *LsntransportprofileResourceModel, getResponseData map[string]interface{}) *LsntransportprofileResourceModel {
	tflog.Debug(ctx, "In lsntransportprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["finrsttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Finrsttimeout = types.Int64Value(intVal)
		}
	} else {
		data.Finrsttimeout = types.Int64Null()
	}
	if val, ok := getResponseData["groupsessionlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Groupsessionlimit = types.Int64Value(intVal)
		}
	} else {
		data.Groupsessionlimit = types.Int64Null()
	}
	if val, ok := getResponseData["portpreserveparity"]; ok && val != nil {
		data.Portpreserveparity = types.StringValue(val.(string))
	} else {
		data.Portpreserveparity = types.StringNull()
	}
	if val, ok := getResponseData["portpreserverange"]; ok && val != nil {
		data.Portpreserverange = types.StringValue(val.(string))
	} else {
		data.Portpreserverange = types.StringNull()
	}
	if val, ok := getResponseData["portquota"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Portquota = types.Int64Value(intVal)
		}
	} else {
		data.Portquota = types.Int64Null()
	}
	if val, ok := getResponseData["sessionquota"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessionquota = types.Int64Value(intVal)
		}
	} else {
		data.Sessionquota = types.Int64Null()
	}
	if val, ok := getResponseData["sessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessiontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["stuntimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Stuntimeout = types.Int64Value(intVal)
		}
	} else {
		data.Stuntimeout = types.Int64Null()
	}
	if val, ok := getResponseData["syncheck"]; ok && val != nil {
		data.Syncheck = types.StringValue(val.(string))
	} else {
		data.Syncheck = types.StringNull()
	}
	if val, ok := getResponseData["synidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Synidletimeout = types.Int64Value(intVal)
		}
	} else {
		data.Synidletimeout = types.Int64Null()
	}
	if val, ok := getResponseData["transportprofilename"]; ok && val != nil {
		data.Transportprofilename = types.StringValue(val.(string))
	} else {
		data.Transportprofilename = types.StringNull()
	}
	if val, ok := getResponseData["transportprotocol"]; ok && val != nil {
		data.Transportprotocol = types.StringValue(val.(string))
	} else {
		data.Transportprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Transportprofilename.ValueString())

	return data
}
