package ssldtlsprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

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

// SsldtlsprofileResourceModel describes the resource data model.
type SsldtlsprofileResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Helloverifyrequest   types.String `tfsdk:"helloverifyrequest"`
	Initialretrytimeout  types.Int64  `tfsdk:"initialretrytimeout"`
	Maxbadmacignorecount types.Int64  `tfsdk:"maxbadmacignorecount"`
	Maxholdqlen          types.Int64  `tfsdk:"maxholdqlen"`
	Maxpacketsize        types.Int64  `tfsdk:"maxpacketsize"`
	Maxrecordsize        types.Int64  `tfsdk:"maxrecordsize"`
	Maxretrytime         types.Int64  `tfsdk:"maxretrytime"`
	Name                 types.String `tfsdk:"name"`
	Pmtudiscovery        types.String `tfsdk:"pmtudiscovery"`
	Terminatesession     types.String `tfsdk:"terminatesession"`
}

func (r *SsldtlsprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssldtlsprofile resource.",
			},
			"helloverifyrequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Send a Hello Verify request to validate the client.",
			},
			"initialretrytimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Initial time out value to retransmit the last flight sent from the NetScaler.",
			},
			"maxbadmacignorecount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Maximum number of bad MAC errors to ignore for a connection prior disconnect. Disabling parameter terminateSession terminates session immediately when bad MAC is detected in the connection.",
			},
			"maxholdqlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(32),
				Description: "Maximum number of datagrams that can be queued at DTLS layer for processing",
			},
			"maxpacketsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Maximum number of packets to reassemble. This value helps protect against a fragmented packet attack.",
			},
			"maxrecordsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1459),
				Description: "Maximum size of records that can be sent if PMTU is disabled.",
			},
			"maxretrytime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Wait for the specified time, in seconds, before resending the request.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the DTLS profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"pmtudiscovery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Source for the maximum record size value. If ENABLED, the value is taken from the PMTU table. If DISABLED, the value is taken from the profile.",
			},
			"terminatesession": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Terminate the session if the message authentication code (MAC) of the client and server do not match.",
			},
		},
	}
}

func ssldtlsprofileGetThePayloadFromthePlan(ctx context.Context, data *SsldtlsprofileResourceModel) ssl.Ssldtlsprofile {
	tflog.Debug(ctx, "In ssldtlsprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	ssldtlsprofile := ssl.Ssldtlsprofile{}
	if !data.Helloverifyrequest.IsNull() && !data.Helloverifyrequest.IsUnknown() {
		ssldtlsprofile.Helloverifyrequest = data.Helloverifyrequest.ValueString()
	}
	if !data.Initialretrytimeout.IsNull() && !data.Initialretrytimeout.IsUnknown() {
		ssldtlsprofile.Initialretrytimeout = utils.IntPtr(int(data.Initialretrytimeout.ValueInt64()))
	}
	if !data.Maxbadmacignorecount.IsNull() && !data.Maxbadmacignorecount.IsUnknown() {
		ssldtlsprofile.Maxbadmacignorecount = utils.IntPtr(int(data.Maxbadmacignorecount.ValueInt64()))
	}
	if !data.Maxholdqlen.IsNull() && !data.Maxholdqlen.IsUnknown() {
		ssldtlsprofile.Maxholdqlen = utils.IntPtr(int(data.Maxholdqlen.ValueInt64()))
	}
	if !data.Maxpacketsize.IsNull() && !data.Maxpacketsize.IsUnknown() {
		ssldtlsprofile.Maxpacketsize = utils.IntPtr(int(data.Maxpacketsize.ValueInt64()))
	}
	if !data.Maxrecordsize.IsNull() && !data.Maxrecordsize.IsUnknown() {
		ssldtlsprofile.Maxrecordsize = utils.IntPtr(int(data.Maxrecordsize.ValueInt64()))
	}
	if !data.Maxretrytime.IsNull() && !data.Maxretrytime.IsUnknown() {
		ssldtlsprofile.Maxretrytime = utils.IntPtr(int(data.Maxretrytime.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		ssldtlsprofile.Name = data.Name.ValueString()
	}
	if !data.Pmtudiscovery.IsNull() && !data.Pmtudiscovery.IsUnknown() {
		ssldtlsprofile.Pmtudiscovery = data.Pmtudiscovery.ValueString()
	}
	if !data.Terminatesession.IsNull() && !data.Terminatesession.IsUnknown() {
		ssldtlsprofile.Terminatesession = data.Terminatesession.ValueString()
	}

	return ssldtlsprofile
}

// ssldtlsprofileSetAttrFromGet populates the resource model from a NITRO GET
// response while preserving state: an else-branch only nulls an attribute when
// the current model value is unknown (never clobbers a known/configured value
// that NITRO omits from GET at its default - the omit-on-default trap).
func ssldtlsprofileSetAttrFromGet(ctx context.Context, data *SsldtlsprofileResourceModel, getResponseData map[string]interface{}) *SsldtlsprofileResourceModel {
	tflog.Debug(ctx, "In ssldtlsprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["helloverifyrequest"]; ok && val != nil {
		data.Helloverifyrequest = types.StringValue(val.(string))
	} else if data.Helloverifyrequest.IsUnknown() {
		data.Helloverifyrequest = types.StringNull()
	}
	if val, ok := getResponseData["initialretrytimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialretrytimeout = types.Int64Value(intVal)
		}
	} else if data.Initialretrytimeout.IsUnknown() {
		data.Initialretrytimeout = types.Int64Null()
	}
	if val, ok := getResponseData["maxbadmacignorecount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbadmacignorecount = types.Int64Value(intVal)
		}
	} else if data.Maxbadmacignorecount.IsUnknown() {
		data.Maxbadmacignorecount = types.Int64Null()
	}
	if val, ok := getResponseData["maxholdqlen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxholdqlen = types.Int64Value(intVal)
		}
	} else if data.Maxholdqlen.IsUnknown() {
		data.Maxholdqlen = types.Int64Null()
	}
	if val, ok := getResponseData["maxpacketsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpacketsize = types.Int64Value(intVal)
		}
	} else if data.Maxpacketsize.IsUnknown() {
		data.Maxpacketsize = types.Int64Null()
	}
	if val, ok := getResponseData["maxrecordsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxrecordsize = types.Int64Value(intVal)
		}
	} else if data.Maxrecordsize.IsUnknown() {
		data.Maxrecordsize = types.Int64Null()
	}
	if val, ok := getResponseData["maxretrytime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxretrytime = types.Int64Value(intVal)
		}
	} else if data.Maxretrytime.IsUnknown() {
		data.Maxretrytime = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["pmtudiscovery"]; ok && val != nil {
		data.Pmtudiscovery = types.StringValue(val.(string))
	} else if data.Pmtudiscovery.IsUnknown() {
		data.Pmtudiscovery = types.StringNull()
	}
	if val, ok := getResponseData["terminatesession"]; ok && val != nil {
		data.Terminatesession = types.StringValue(val.(string))
	} else if data.Terminatesession.IsUnknown() {
		data.Terminatesession = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}

// ssldtlsprofileSetAttrFromGetForDatasource populates the model from a NITRO GET
// response for the datasource: it copies every attribute directly (nulling when
// absent) and sets the ID. State-preservation guards are unnecessary for reads.
func ssldtlsprofileSetAttrFromGetForDatasource(ctx context.Context, data *SsldtlsprofileResourceModel, getResponseData map[string]interface{}) *SsldtlsprofileResourceModel {
	tflog.Debug(ctx, "In ssldtlsprofileSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["helloverifyrequest"]; ok && val != nil {
		data.Helloverifyrequest = types.StringValue(val.(string))
	} else {
		data.Helloverifyrequest = types.StringNull()
	}
	if val, ok := getResponseData["initialretrytimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialretrytimeout = types.Int64Value(intVal)
		}
	} else {
		data.Initialretrytimeout = types.Int64Null()
	}
	if val, ok := getResponseData["maxbadmacignorecount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbadmacignorecount = types.Int64Value(intVal)
		}
	} else {
		data.Maxbadmacignorecount = types.Int64Null()
	}
	if val, ok := getResponseData["maxholdqlen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxholdqlen = types.Int64Value(intVal)
		}
	} else {
		data.Maxholdqlen = types.Int64Null()
	}
	if val, ok := getResponseData["maxpacketsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpacketsize = types.Int64Value(intVal)
		}
	} else {
		data.Maxpacketsize = types.Int64Null()
	}
	if val, ok := getResponseData["maxrecordsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxrecordsize = types.Int64Value(intVal)
		}
	} else {
		data.Maxrecordsize = types.Int64Null()
	}
	if val, ok := getResponseData["maxretrytime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxretrytime = types.Int64Value(intVal)
		}
	} else {
		data.Maxretrytime = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["pmtudiscovery"]; ok && val != nil {
		data.Pmtudiscovery = types.StringValue(val.(string))
	} else {
		data.Pmtudiscovery = types.StringNull()
	}
	if val, ok := getResponseData["terminatesession"]; ok && val != nil {
		data.Terminatesession = types.StringValue(val.(string))
	} else {
		data.Terminatesession = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
