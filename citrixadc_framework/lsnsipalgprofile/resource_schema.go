package lsnsipalgprofile

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

// LsnsipalgprofileResourceModel describes the resource data model.
type LsnsipalgprofileResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Datasessionidletimeout types.Int64  `tfsdk:"datasessionidletimeout"`
	Opencontactpinhole     types.String `tfsdk:"opencontactpinhole"`
	Openrecordroutepinhole types.String `tfsdk:"openrecordroutepinhole"`
	Openregisterpinhole    types.String `tfsdk:"openregisterpinhole"`
	Openroutepinhole       types.String `tfsdk:"openroutepinhole"`
	Openviapinhole         types.String `tfsdk:"openviapinhole"`
	Registrationtimeout    types.Int64  `tfsdk:"registrationtimeout"`
	Rport                  types.String `tfsdk:"rport"`
	Sipalgprofilename      types.String `tfsdk:"sipalgprofilename"`
	Sipdstportrange        types.String `tfsdk:"sipdstportrange"`
	Sipsessiontimeout      types.Int64  `tfsdk:"sipsessiontimeout"`
	Sipsrcportrange        types.String `tfsdk:"sipsrcportrange"`
	Siptransportprotocol   types.String `tfsdk:"siptransportprotocol"`
}

func (r *LsnsipalgprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnsipalgprofile resource.",
			},
			"datasessionidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Idle timeout for the data channel sessions in seconds.",
			},
			"opencontactpinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "ENABLE/DISABLE ContactPinhole creation.",
			},
			"openrecordroutepinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "ENABLE/DISABLE RecordRoutePinhole creation.",
			},
			"openregisterpinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "ENABLE/DISABLE RegisterPinhole creation.",
			},
			"openroutepinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "ENABLE/DISABLE RoutePinhole creation.",
			},
			"openviapinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "ENABLE/DISABLE ViaPinhole creation.",
			},
			"registrationtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(60),
				Description: "SIP registration timeout in seconds.",
			},
			"rport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "ENABLE/DISABLE rport.",
			},
			"sipalgprofilename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the SIPALG Profile.",
			},
			"sipdstportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination port range for SIP_UDP and SIP_TCP.",
			},
			"sipsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(600),
				Description: "SIP control channel session timeout in seconds.",
			},
			"sipsrcportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source port range for SIP_UDP and SIP_TCP.",
			},
			"siptransportprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP ALG Profile transport protocol type.",
			},
		},
	}
}

func lsnsipalgprofileGetThePayloadFromtheConfig(ctx context.Context, data *LsnsipalgprofileResourceModel) lsn.Lsnsipalgprofile {
	tflog.Debug(ctx, "In lsnsipalgprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnsipalgprofile := lsn.Lsnsipalgprofile{}
	if !data.Datasessionidletimeout.IsNull() && !data.Datasessionidletimeout.IsUnknown() {
		lsnsipalgprofile.Datasessionidletimeout = utils.IntPtr(int(data.Datasessionidletimeout.ValueInt64()))
	}
	if !data.Opencontactpinhole.IsNull() && !data.Opencontactpinhole.IsUnknown() {
		lsnsipalgprofile.Opencontactpinhole = data.Opencontactpinhole.ValueString()
	}
	if !data.Openrecordroutepinhole.IsNull() && !data.Openrecordroutepinhole.IsUnknown() {
		lsnsipalgprofile.Openrecordroutepinhole = data.Openrecordroutepinhole.ValueString()
	}
	if !data.Openregisterpinhole.IsNull() && !data.Openregisterpinhole.IsUnknown() {
		lsnsipalgprofile.Openregisterpinhole = data.Openregisterpinhole.ValueString()
	}
	if !data.Openroutepinhole.IsNull() && !data.Openroutepinhole.IsUnknown() {
		lsnsipalgprofile.Openroutepinhole = data.Openroutepinhole.ValueString()
	}
	if !data.Openviapinhole.IsNull() && !data.Openviapinhole.IsUnknown() {
		lsnsipalgprofile.Openviapinhole = data.Openviapinhole.ValueString()
	}
	if !data.Registrationtimeout.IsNull() && !data.Registrationtimeout.IsUnknown() {
		lsnsipalgprofile.Registrationtimeout = utils.IntPtr(int(data.Registrationtimeout.ValueInt64()))
	}
	if !data.Rport.IsNull() && !data.Rport.IsUnknown() {
		lsnsipalgprofile.Rport = data.Rport.ValueString()
	}
	if !data.Sipalgprofilename.IsNull() && !data.Sipalgprofilename.IsUnknown() {
		lsnsipalgprofile.Sipalgprofilename = data.Sipalgprofilename.ValueString()
	}
	if !data.Sipdstportrange.IsNull() && !data.Sipdstportrange.IsUnknown() {
		lsnsipalgprofile.Sipdstportrange = data.Sipdstportrange.ValueString()
	}
	if !data.Sipsessiontimeout.IsNull() && !data.Sipsessiontimeout.IsUnknown() {
		lsnsipalgprofile.Sipsessiontimeout = utils.IntPtr(int(data.Sipsessiontimeout.ValueInt64()))
	}
	if !data.Sipsrcportrange.IsNull() && !data.Sipsrcportrange.IsUnknown() {
		lsnsipalgprofile.Sipsrcportrange = data.Sipsrcportrange.ValueString()
	}
	if !data.Siptransportprotocol.IsNull() && !data.Siptransportprotocol.IsUnknown() {
		lsnsipalgprofile.Siptransportprotocol = data.Siptransportprotocol.ValueString()
	}

	return lsnsipalgprofile
}

func lsnsipalgprofileSetAttrFromGet(ctx context.Context, data *LsnsipalgprofileResourceModel, getResponseData map[string]interface{}) *LsnsipalgprofileResourceModel {
	tflog.Debug(ctx, "In lsnsipalgprofileSetAttrFromGet Function")

	// Convert API response to model. In the else-branches we only null a value
	// that is still Unknown (never clobber a known configured/state value that
	// NITRO happened to omit from the GET response - the omit-on-default trap).
	if val, ok := getResponseData["datasessionidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Datasessionidletimeout = types.Int64Value(intVal)
		}
	} else if data.Datasessionidletimeout.IsUnknown() {
		data.Datasessionidletimeout = types.Int64Null()
	}
	if val, ok := getResponseData["opencontactpinhole"]; ok && val != nil {
		data.Opencontactpinhole = types.StringValue(val.(string))
	} else if data.Opencontactpinhole.IsUnknown() {
		data.Opencontactpinhole = types.StringNull()
	}
	if val, ok := getResponseData["openrecordroutepinhole"]; ok && val != nil {
		data.Openrecordroutepinhole = types.StringValue(val.(string))
	} else if data.Openrecordroutepinhole.IsUnknown() {
		data.Openrecordroutepinhole = types.StringNull()
	}
	if val, ok := getResponseData["openregisterpinhole"]; ok && val != nil {
		data.Openregisterpinhole = types.StringValue(val.(string))
	} else if data.Openregisterpinhole.IsUnknown() {
		data.Openregisterpinhole = types.StringNull()
	}
	if val, ok := getResponseData["openroutepinhole"]; ok && val != nil {
		data.Openroutepinhole = types.StringValue(val.(string))
	} else if data.Openroutepinhole.IsUnknown() {
		data.Openroutepinhole = types.StringNull()
	}
	if val, ok := getResponseData["openviapinhole"]; ok && val != nil {
		data.Openviapinhole = types.StringValue(val.(string))
	} else if data.Openviapinhole.IsUnknown() {
		data.Openviapinhole = types.StringNull()
	}
	if val, ok := getResponseData["registrationtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Registrationtimeout = types.Int64Value(intVal)
		}
	} else if data.Registrationtimeout.IsUnknown() {
		data.Registrationtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["rport"]; ok && val != nil {
		data.Rport = types.StringValue(val.(string))
	} else if data.Rport.IsUnknown() {
		data.Rport = types.StringNull()
	}
	if val, ok := getResponseData["sipalgprofilename"]; ok && val != nil {
		data.Sipalgprofilename = types.StringValue(val.(string))
	} else if data.Sipalgprofilename.IsUnknown() {
		data.Sipalgprofilename = types.StringNull()
	}
	if val, ok := getResponseData["sipdstportrange"]; ok && val != nil {
		data.Sipdstportrange = types.StringValue(val.(string))
	} else if data.Sipdstportrange.IsUnknown() {
		data.Sipdstportrange = types.StringNull()
	}
	if val, ok := getResponseData["sipsessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sipsessiontimeout = types.Int64Value(intVal)
		}
	} else if data.Sipsessiontimeout.IsUnknown() {
		data.Sipsessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["sipsrcportrange"]; ok && val != nil {
		data.Sipsrcportrange = types.StringValue(val.(string))
	} else if data.Sipsrcportrange.IsUnknown() {
		data.Sipsrcportrange = types.StringNull()
	}
	if val, ok := getResponseData["siptransportprotocol"]; ok && val != nil {
		data.Siptransportprotocol = types.StringValue(val.(string))
	} else if data.Siptransportprotocol.IsUnknown() {
		data.Siptransportprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Sipalgprofilename.ValueString())

	return data
}
