package vpnnexthopserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnnexthopserverResourceModel describes the resource data model.
type VpnnexthopserverResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Nexthopfqdn    types.String `tfsdk:"nexthopfqdn"`
	Nexthopip      types.String `tfsdk:"nexthopip"`
	Nexthopport    types.Int64  `tfsdk:"nexthopport"`
	Resaddresstype types.String `tfsdk:"resaddresstype"`
	Secure         types.String `tfsdk:"secure"`
}

func (r *VpnnexthopserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnnexthopserver resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Citrix Gateway appliance in the first DMZ.",
			},
			"nexthopfqdn": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "FQDN of the Citrix Gateway proxy in the second DMZ.",
			},
			"nexthopip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the Citrix Gateway proxy in the second DMZ.",
			},
			"nexthopport": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number of the Citrix Gateway proxy in the second DMZ.",
			},
			"resaddresstype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Address Type (IPV4/IPv6) of DNS name of nextHopServer FQDN.",
			},
			"secure": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Use of a secure port, such as 443, for the double-hop configuration.",
			},
		},
	}
}

func vpnnexthopserverGetThePayloadFromtheConfig(ctx context.Context, data *VpnnexthopserverResourceModel) vpn.Vpnnexthopserver {
	tflog.Debug(ctx, "In vpnnexthopserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnnexthopserver := vpn.Vpnnexthopserver{}
	if !data.Name.IsNull() {
		vpnnexthopserver.Name = data.Name.ValueString()
	}
	if !data.Nexthopfqdn.IsNull() {
		vpnnexthopserver.Nexthopfqdn = data.Nexthopfqdn.ValueString()
	}
	if !data.Nexthopip.IsNull() {
		vpnnexthopserver.Nexthopip = data.Nexthopip.ValueString()
	}
	if !data.Nexthopport.IsNull() {
		vpnnexthopserver.Nexthopport = utils.IntPtr(int(data.Nexthopport.ValueInt64()))
	}
	if !data.Resaddresstype.IsNull() {
		vpnnexthopserver.Resaddresstype = data.Resaddresstype.ValueString()
	}
	if !data.Secure.IsNull() {
		vpnnexthopserver.Secure = data.Secure.ValueString()
	}

	return vpnnexthopserver
}

func vpnnexthopserverSetAttrFromGet(ctx context.Context, data *VpnnexthopserverResourceModel, getResponseData map[string]interface{}) *VpnnexthopserverResourceModel {
	tflog.Debug(ctx, "In vpnnexthopserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nexthopfqdn"]; ok && val != nil {
		data.Nexthopfqdn = types.StringValue(val.(string))
	} else {
		data.Nexthopfqdn = types.StringNull()
	}
	if val, ok := getResponseData["nexthopip"]; ok && val != nil {
		data.Nexthopip = types.StringValue(val.(string))
	} else {
		data.Nexthopip = types.StringNull()
	}
	if val, ok := getResponseData["nexthopport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nexthopport = types.Int64Value(intVal)
		}
	} else {
		data.Nexthopport = types.Int64Null()
	}
	if val, ok := getResponseData["resaddresstype"]; ok && val != nil {
		data.Resaddresstype = types.StringValue(val.(string))
	} else {
		data.Resaddresstype = types.StringNull()
	}
	if val, ok := getResponseData["secure"]; ok && val != nil {
		data.Secure = types.StringValue(val.(string))
	} else {
		data.Secure = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
