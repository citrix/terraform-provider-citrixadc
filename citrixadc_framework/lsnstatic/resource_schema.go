package lsnstatic

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnstaticResourceModel describes the resource data model.
type LsnstaticResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Destip            types.String `tfsdk:"destip"`
	Dsttd             types.Int64  `tfsdk:"dsttd"`
	Name              types.String `tfsdk:"name"`
	Natip             types.String `tfsdk:"natip"`
	Natport           types.Int64  `tfsdk:"natport"`
	Network6          types.String `tfsdk:"network6"`
	Subscrip          types.String `tfsdk:"subscrip"`
	Subscrport        types.Int64  `tfsdk:"subscrport"`
	Td                types.Int64  `tfsdk:"td"`
	Transportprotocol types.String `tfsdk:"transportprotocol"`
}

func (r *LsnstaticResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnstatic resource.",
			},
			"destip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination IP address for the LSN mapping entry.",
			},
			"dsttd": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the traffic domain through which the destination IP address for this LSN mapping entry is reachable from the Citrix ADC.\n\nIf you do not specify an ID, the destination IP address is assumed to be reachable through the default traffic domain, which has an ID of 0.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LSN static mapping entry. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn static1\" or 'lsn static1').",
			},
			"natip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 address, already existing on the Citrix ADC as type LSN, to be used as NAT IP address for this mapping entry.",
			},
			"natport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "NAT port for this LSN mapping entry. * represents all ports being used. Used in case of static wildcard",
			},
			"network6": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "B4 address in DS-Lite setup",
			},
			"subscrip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4(NAT44 & DS-Lite)/IPv6(NAT64) address of an LSN subscriber for the LSN static mapping entry.",
			},
			"subscrport": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port of the LSN subscriber for the LSN mapping entry. * represents all ports being used. Used in case of static wildcard",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the traffic domain to which the subscriber belongs. \n\nIf you do not specify an ID, the subscriber is assumed to be a part of the default traffic domain.",
			},
			"transportprotocol": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol for the LSN mapping entry.",
			},
		},
	}
}

func lsnstaticGetThePayloadFromtheConfig(ctx context.Context, data *LsnstaticResourceModel) lsn.Lsnstatic {
	tflog.Debug(ctx, "In lsnstaticGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnstatic := lsn.Lsnstatic{}
	if !data.Destip.IsNull() {
		lsnstatic.Destip = data.Destip.ValueString()
	}
	if !data.Dsttd.IsNull() {
		lsnstatic.Dsttd = utils.IntPtr(int(data.Dsttd.ValueInt64()))
	}
	if !data.Name.IsNull() {
		lsnstatic.Name = data.Name.ValueString()
	}
	if !data.Natip.IsNull() {
		lsnstatic.Natip = data.Natip.ValueString()
	}
	if !data.Natport.IsNull() {
		lsnstatic.Natport = utils.IntPtr(int(data.Natport.ValueInt64()))
	}
	if !data.Network6.IsNull() {
		lsnstatic.Network6 = data.Network6.ValueString()
	}
	if !data.Subscrip.IsNull() {
		lsnstatic.Subscrip = data.Subscrip.ValueString()
	}
	if !data.Subscrport.IsNull() {
		lsnstatic.Subscrport = utils.IntPtr(int(data.Subscrport.ValueInt64()))
	}
	if !data.Td.IsNull() {
		lsnstatic.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Transportprotocol.IsNull() {
		lsnstatic.Transportprotocol = data.Transportprotocol.ValueString()
	}

	return lsnstatic
}

func lsnstaticSetAttrFromGet(ctx context.Context, data *LsnstaticResourceModel, getResponseData map[string]interface{}) *LsnstaticResourceModel {
	tflog.Debug(ctx, "In lsnstaticSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["destip"]; ok && val != nil {
		data.Destip = types.StringValue(val.(string))
	} else {
		data.Destip = types.StringNull()
	}
	if val, ok := getResponseData["dsttd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dsttd = types.Int64Value(intVal)
		}
	} else {
		data.Dsttd = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["natip"]; ok && val != nil {
		data.Natip = types.StringValue(val.(string))
	} else {
		data.Natip = types.StringNull()
	}
	if val, ok := getResponseData["natport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Natport = types.Int64Value(intVal)
		}
	} else {
		data.Natport = types.Int64Null()
	}
	if val, ok := getResponseData["network6"]; ok && val != nil {
		data.Network6 = types.StringValue(val.(string))
	} else {
		data.Network6 = types.StringNull()
	}
	if val, ok := getResponseData["subscrip"]; ok && val != nil {
		data.Subscrip = types.StringValue(val.(string))
	} else {
		data.Subscrip = types.StringNull()
	}
	if val, ok := getResponseData["subscrport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Subscrport = types.Int64Value(intVal)
		}
	} else {
		data.Subscrport = types.Int64Null()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["transportprotocol"]; ok && val != nil {
		data.Transportprotocol = types.StringValue(val.(string))
	} else {
		data.Transportprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Name.ValueString()))

	return data
}
