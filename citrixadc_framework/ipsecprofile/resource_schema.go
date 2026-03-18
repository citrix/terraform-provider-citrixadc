package ipsecprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ipsec"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IpsecprofileResourceModel describes the resource data model.
type IpsecprofileResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Encalgo               types.List   `tfsdk:"encalgo"`
	Hashalgo              types.List   `tfsdk:"hashalgo"`
	Ikeretryinterval      types.Int64  `tfsdk:"ikeretryinterval"`
	Ikeversion            types.String `tfsdk:"ikeversion"`
	Lifetime              types.Int64  `tfsdk:"lifetime"`
	Livenesscheckinterval types.Int64  `tfsdk:"livenesscheckinterval"`
	Name                  types.String `tfsdk:"name"`
	Peerpublickey         types.String `tfsdk:"peerpublickey"`
	Perfectforwardsecrecy types.String `tfsdk:"perfectforwardsecrecy"`
	Privatekey            types.String `tfsdk:"privatekey"`
	Psk                   types.String `tfsdk:"psk"`
	Publickey             types.String `tfsdk:"publickey"`
	Replaywindowsize      types.Int64  `tfsdk:"replaywindowsize"`
	Retransmissiontime    types.Int64  `tfsdk:"retransmissiontime"`
}

func (r *IpsecprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipsecprofile resource.",
			},
			"encalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Type of encryption algorithm (Note: Selection of AES enables AES128)",
			},
			"hashalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Type of hashing algorithm",
			},
			"ikeretryinterval": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "IKE retry interval for bringing up the connection",
			},
			"ikeversion": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IKE Protocol Version",
			},
			"lifetime": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8)",
			},
			"livenesscheckinterval": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the ipsec profile",
			},
			"peerpublickey": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Peer public key file path",
			},
			"perfectforwardsecrecy": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enable/Disable PFS.",
			},
			"privatekey": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Private key file path",
			},
			"psk": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pre shared key value",
			},
			"publickey": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Public key file path",
			},
			"replaywindowsize": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "IPSec Replay window size for the data traffic",
			},
			"retransmissiontime": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure.",
			},
		},
	}
}

func ipsecprofileGetThePayloadFromtheConfig(ctx context.Context, data *IpsecprofileResourceModel) ipsec.Ipsecprofile {
	tflog.Debug(ctx, "In ipsecprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ipsecprofile := ipsec.Ipsecprofile{}
	if !data.Ikeretryinterval.IsNull() {
		ipsecprofile.Ikeretryinterval = utils.IntPtr(int(data.Ikeretryinterval.ValueInt64()))
	}
	if !data.Ikeversion.IsNull() {
		ipsecprofile.Ikeversion = data.Ikeversion.ValueString()
	}
	if !data.Lifetime.IsNull() {
		ipsecprofile.Lifetime = utils.IntPtr(int(data.Lifetime.ValueInt64()))
	}
	if !data.Livenesscheckinterval.IsNull() {
		ipsecprofile.Livenesscheckinterval = utils.IntPtr(int(data.Livenesscheckinterval.ValueInt64()))
	}
	if !data.Name.IsNull() {
		ipsecprofile.Name = data.Name.ValueString()
	}
	if !data.Peerpublickey.IsNull() {
		ipsecprofile.Peerpublickey = data.Peerpublickey.ValueString()
	}
	if !data.Perfectforwardsecrecy.IsNull() {
		ipsecprofile.Perfectforwardsecrecy = data.Perfectforwardsecrecy.ValueString()
	}
	if !data.Privatekey.IsNull() {
		ipsecprofile.Privatekey = data.Privatekey.ValueString()
	}
	if !data.Psk.IsNull() {
		ipsecprofile.Psk = data.Psk.ValueString()
	}
	if !data.Publickey.IsNull() {
		ipsecprofile.Publickey = data.Publickey.ValueString()
	}
	if !data.Replaywindowsize.IsNull() {
		ipsecprofile.Replaywindowsize = utils.IntPtr(int(data.Replaywindowsize.ValueInt64()))
	}
	if !data.Retransmissiontime.IsNull() {
		ipsecprofile.Retransmissiontime = utils.IntPtr(int(data.Retransmissiontime.ValueInt64()))
	}

	return ipsecprofile
}

func ipsecprofileSetAttrFromGet(ctx context.Context, data *IpsecprofileResourceModel, getResponseData map[string]interface{}) *IpsecprofileResourceModel {
	tflog.Debug(ctx, "In ipsecprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ikeretryinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ikeretryinterval = types.Int64Value(intVal)
		}
	} else {
		data.Ikeretryinterval = types.Int64Null()
	}
	if val, ok := getResponseData["ikeversion"]; ok && val != nil {
		data.Ikeversion = types.StringValue(val.(string))
	} else {
		data.Ikeversion = types.StringNull()
	}
	if val, ok := getResponseData["lifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Lifetime = types.Int64Value(intVal)
		}
	} else {
		data.Lifetime = types.Int64Null()
	}
	if val, ok := getResponseData["livenesscheckinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Livenesscheckinterval = types.Int64Value(intVal)
		}
	} else {
		data.Livenesscheckinterval = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["peerpublickey"]; ok && val != nil {
		data.Peerpublickey = types.StringValue(val.(string))
	} else {
		data.Peerpublickey = types.StringNull()
	}
	if val, ok := getResponseData["perfectforwardsecrecy"]; ok && val != nil {
		data.Perfectforwardsecrecy = types.StringValue(val.(string))
	} else {
		data.Perfectforwardsecrecy = types.StringNull()
	}
	if val, ok := getResponseData["privatekey"]; ok && val != nil {
		data.Privatekey = types.StringValue(val.(string))
	} else {
		data.Privatekey = types.StringNull()
	}
	if val, ok := getResponseData["psk"]; ok && val != nil {
		data.Psk = types.StringValue(val.(string))
	} else {
		data.Psk = types.StringNull()
	}
	if val, ok := getResponseData["publickey"]; ok && val != nil {
		data.Publickey = types.StringValue(val.(string))
	} else {
		data.Publickey = types.StringNull()
	}
	if val, ok := getResponseData["replaywindowsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Replaywindowsize = types.Int64Value(intVal)
		}
	} else {
		data.Replaywindowsize = types.Int64Null()
	}
	if val, ok := getResponseData["retransmissiontime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retransmissiontime = types.Int64Value(intVal)
		}
	} else {
		data.Retransmissiontime = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
