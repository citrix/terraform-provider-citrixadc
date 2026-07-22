package quicprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/quic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// QuicprofileResourceModel describes the resource data model.
type QuicprofileResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Ackdelayexponent               types.Int64  `tfsdk:"ackdelayexponent"`
	Activeconnectionidlimit        types.Int64  `tfsdk:"activeconnectionidlimit"`
	Activeconnectionmigration      types.String `tfsdk:"activeconnectionmigration"`
	Congestionctrlalgorithm        types.String `tfsdk:"congestionctrlalgorithm"`
	Initialmaxdata                 types.Int64  `tfsdk:"initialmaxdata"`
	Initialmaxstreamdatabidilocal  types.Int64  `tfsdk:"initialmaxstreamdatabidilocal"`
	Initialmaxstreamdatabidiremote types.Int64  `tfsdk:"initialmaxstreamdatabidiremote"`
	Initialmaxstreamdatauni        types.Int64  `tfsdk:"initialmaxstreamdatauni"`
	Initialmaxstreamsbidi          types.Int64  `tfsdk:"initialmaxstreamsbidi"`
	Initialmaxstreamsuni           types.Int64  `tfsdk:"initialmaxstreamsuni"`
	Maxackdelay                    types.Int64  `tfsdk:"maxackdelay"`
	Maxidletimeout                 types.Int64  `tfsdk:"maxidletimeout"`
	Maxudpdatagramsperburst        types.Int64  `tfsdk:"maxudpdatagramsperburst"`
	Maxudppayloadsize              types.Int64  `tfsdk:"maxudppayloadsize"`
	Name                           types.String `tfsdk:"name"`
	Newtokenvalidityperiod         types.Int64  `tfsdk:"newtokenvalidityperiod"`
	Retrytokenvalidityperiod       types.Int64  `tfsdk:"retrytokenvalidityperiod"`
	Statelessaddressvalidation     types.String `tfsdk:"statelessaddressvalidation"`
}

func (r *QuicprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the quicprofile resource.",
			},
			"ackdelayexponent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, indicating an exponent that the remote QUIC endpoint should use, to decode the ACK Delay field in QUIC ACK frames sent by the Citrix ADC.",
			},
			"activeconnectionidlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum number of QUIC connection IDs from the remote QUIC endpoint, that the Citrix ADC is willing to store.",
			},
			"activeconnectionmigration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether the Citrix ADC should allow the remote QUIC endpoint to perform active QUIC connection migration.",
			},
			"congestionctrlalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the congestion control algorithm to be used for QUIC connections. The default congestion control algorithm is CUBIC.",
			},
			"initialmaxdata": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial value, in bytes, for the maximum amount of data that can be sent on a QUIC connection.",
			},
			"initialmaxstreamdatabidilocal": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the Citrix ADC.",
			},
			"initialmaxstreamdatabidiremote": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the remote QUIC endpoint.",
			},
			"initialmaxstreamdatauni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for unidirectional streams initiated by the remote QUIC endpoint.",
			},
			"initialmaxstreamsbidi": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of bidirectional streams the remote QUIC endpoint may initiate.",
			},
			"initialmaxstreamsuni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of unidirectional streams the remote QUIC endpoint may initiate.",
			},
			"maxackdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum amount of time, in milliseconds, by which the Citrix ADC will delay sending acknowledgments.",
			},
			"maxidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum idle timeout, in seconds, for a QUIC connection. A QUIC connection will be silently discarded by the Citrix ADC if it remains idle for longer than the minimum of the idle timeout values advertised by the Citrix ADC and the remote QUIC endpoint, and three times the current Probe Timeout (PTO).",
			},
			"maxudpdatagramsperburst": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the maximum number of UDP datagrams that can be transmitted by the Citrix ADC in a single transmission burst on a QUIC connection.",
			},
			"maxudppayloadsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the size of the largest UDP datagram payload, in bytes, that the Citrix ADC is willing to receive on a QUIC connection.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"newtokenvalidityperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC NEW_TOKEN frames sent by the Citrix ADC.",
			},
			"retrytokenvalidityperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC Retry packets sent by the Citrix ADC.",
			},
			"statelessaddressvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether the Citrix ADC should perform stateless address validation for QUIC clients, by sending tokens in QUIC Retry packets during QUIC connection establishment, and by sending tokens in QUIC NEW_TOKEN frames after QUIC connection establishment.",
			},
		},
	}
}

func quicprofileGetThePayloadFromthePlan(ctx context.Context, data *QuicprofileResourceModel) quic.Quicprofile {
	tflog.Debug(ctx, "In quicprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	quicprofile := quic.Quicprofile{}
	if !data.Ackdelayexponent.IsNull() && !data.Ackdelayexponent.IsUnknown() {
		quicprofile.Ackdelayexponent = utils.IntPtr(int(data.Ackdelayexponent.ValueInt64()))
	}
	if !data.Activeconnectionidlimit.IsNull() && !data.Activeconnectionidlimit.IsUnknown() {
		quicprofile.Activeconnectionidlimit = utils.IntPtr(int(data.Activeconnectionidlimit.ValueInt64()))
	}
	if !data.Activeconnectionmigration.IsNull() && !data.Activeconnectionmigration.IsUnknown() {
		quicprofile.Activeconnectionmigration = data.Activeconnectionmigration.ValueString()
	}
	if !data.Congestionctrlalgorithm.IsNull() && !data.Congestionctrlalgorithm.IsUnknown() {
		quicprofile.Congestionctrlalgorithm = data.Congestionctrlalgorithm.ValueString()
	}
	if !data.Initialmaxdata.IsNull() && !data.Initialmaxdata.IsUnknown() {
		quicprofile.Initialmaxdata = utils.IntPtr(int(data.Initialmaxdata.ValueInt64()))
	}
	if !data.Initialmaxstreamdatabidilocal.IsNull() && !data.Initialmaxstreamdatabidilocal.IsUnknown() {
		quicprofile.Initialmaxstreamdatabidilocal = utils.IntPtr(int(data.Initialmaxstreamdatabidilocal.ValueInt64()))
	}
	if !data.Initialmaxstreamdatabidiremote.IsNull() && !data.Initialmaxstreamdatabidiremote.IsUnknown() {
		quicprofile.Initialmaxstreamdatabidiremote = utils.IntPtr(int(data.Initialmaxstreamdatabidiremote.ValueInt64()))
	}
	if !data.Initialmaxstreamdatauni.IsNull() && !data.Initialmaxstreamdatauni.IsUnknown() {
		quicprofile.Initialmaxstreamdatauni = utils.IntPtr(int(data.Initialmaxstreamdatauni.ValueInt64()))
	}
	if !data.Initialmaxstreamsbidi.IsNull() && !data.Initialmaxstreamsbidi.IsUnknown() {
		quicprofile.Initialmaxstreamsbidi = utils.IntPtr(int(data.Initialmaxstreamsbidi.ValueInt64()))
	}
	if !data.Initialmaxstreamsuni.IsNull() && !data.Initialmaxstreamsuni.IsUnknown() {
		quicprofile.Initialmaxstreamsuni = utils.IntPtr(int(data.Initialmaxstreamsuni.ValueInt64()))
	}
	if !data.Maxackdelay.IsNull() && !data.Maxackdelay.IsUnknown() {
		quicprofile.Maxackdelay = utils.IntPtr(int(data.Maxackdelay.ValueInt64()))
	}
	if !data.Maxidletimeout.IsNull() && !data.Maxidletimeout.IsUnknown() {
		quicprofile.Maxidletimeout = utils.IntPtr(int(data.Maxidletimeout.ValueInt64()))
	}
	if !data.Maxudpdatagramsperburst.IsNull() && !data.Maxudpdatagramsperburst.IsUnknown() {
		quicprofile.Maxudpdatagramsperburst = utils.IntPtr(int(data.Maxudpdatagramsperburst.ValueInt64()))
	}
	if !data.Maxudppayloadsize.IsNull() && !data.Maxudppayloadsize.IsUnknown() {
		quicprofile.Maxudppayloadsize = utils.IntPtr(int(data.Maxudppayloadsize.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		quicprofile.Name = data.Name.ValueString()
	}
	if !data.Newtokenvalidityperiod.IsNull() && !data.Newtokenvalidityperiod.IsUnknown() {
		quicprofile.Newtokenvalidityperiod = utils.IntPtr(int(data.Newtokenvalidityperiod.ValueInt64()))
	}
	if !data.Retrytokenvalidityperiod.IsNull() && !data.Retrytokenvalidityperiod.IsUnknown() {
		quicprofile.Retrytokenvalidityperiod = utils.IntPtr(int(data.Retrytokenvalidityperiod.ValueInt64()))
	}
	if !data.Statelessaddressvalidation.IsNull() && !data.Statelessaddressvalidation.IsUnknown() {
		quicprofile.Statelessaddressvalidation = data.Statelessaddressvalidation.ValueString()
	}

	return quicprofile
}

func quicprofileSetAttrFromGet(ctx context.Context, data *QuicprofileResourceModel, getResponseData map[string]interface{}) *QuicprofileResourceModel {
	tflog.Debug(ctx, "In quicprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ackdelayexponent"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ackdelayexponent = types.Int64Value(intVal)
		}
	} else {
		data.Ackdelayexponent = types.Int64Null()
	}
	if val, ok := getResponseData["activeconnectionidlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Activeconnectionidlimit = types.Int64Value(intVal)
		}
	} else {
		data.Activeconnectionidlimit = types.Int64Null()
	}
	if val, ok := getResponseData["activeconnectionmigration"]; ok && val != nil {
		data.Activeconnectionmigration = types.StringValue(val.(string))
	} else {
		data.Activeconnectionmigration = types.StringNull()
	}
	if val, ok := getResponseData["congestionctrlalgorithm"]; ok && val != nil {
		data.Congestionctrlalgorithm = types.StringValue(val.(string))
	} else {
		data.Congestionctrlalgorithm = types.StringNull()
	}
	if val, ok := getResponseData["initialmaxdata"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialmaxdata = types.Int64Value(intVal)
		}
	} else {
		data.Initialmaxdata = types.Int64Null()
	}
	if val, ok := getResponseData["initialmaxstreamdatabidilocal"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialmaxstreamdatabidilocal = types.Int64Value(intVal)
		}
	} else {
		data.Initialmaxstreamdatabidilocal = types.Int64Null()
	}
	if val, ok := getResponseData["initialmaxstreamdatabidiremote"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialmaxstreamdatabidiremote = types.Int64Value(intVal)
		}
	} else {
		data.Initialmaxstreamdatabidiremote = types.Int64Null()
	}
	if val, ok := getResponseData["initialmaxstreamdatauni"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialmaxstreamdatauni = types.Int64Value(intVal)
		}
	} else {
		data.Initialmaxstreamdatauni = types.Int64Null()
	}
	if val, ok := getResponseData["initialmaxstreamsbidi"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialmaxstreamsbidi = types.Int64Value(intVal)
		}
	} else {
		data.Initialmaxstreamsbidi = types.Int64Null()
	}
	if val, ok := getResponseData["initialmaxstreamsuni"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialmaxstreamsuni = types.Int64Value(intVal)
		}
	} else {
		data.Initialmaxstreamsuni = types.Int64Null()
	}
	if val, ok := getResponseData["maxackdelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxackdelay = types.Int64Value(intVal)
		}
	} else {
		data.Maxackdelay = types.Int64Null()
	}
	if val, ok := getResponseData["maxidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxidletimeout = types.Int64Value(intVal)
		}
	} else {
		data.Maxidletimeout = types.Int64Null()
	}
	if val, ok := getResponseData["maxudpdatagramsperburst"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxudpdatagramsperburst = types.Int64Value(intVal)
		}
	} else {
		data.Maxudpdatagramsperburst = types.Int64Null()
	}
	if val, ok := getResponseData["maxudppayloadsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxudppayloadsize = types.Int64Value(intVal)
		}
	} else {
		data.Maxudppayloadsize = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newtokenvalidityperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Newtokenvalidityperiod = types.Int64Value(intVal)
		}
	} else {
		data.Newtokenvalidityperiod = types.Int64Null()
	}
	if val, ok := getResponseData["retrytokenvalidityperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retrytokenvalidityperiod = types.Int64Value(intVal)
		}
	} else {
		data.Retrytokenvalidityperiod = types.Int64Null()
	}
	if val, ok := getResponseData["statelessaddressvalidation"]; ok && val != nil {
		data.Statelessaddressvalidation = types.StringValue(val.(string))
	} else {
		data.Statelessaddressvalidation = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
