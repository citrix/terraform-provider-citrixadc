package dnsaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnsactionResourceModel describes the resource data model.
type DnsactionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Actionname       types.String `tfsdk:"actionname"`
	Actiontype       types.String `tfsdk:"actiontype"`
	Dnsprofilename   types.String `tfsdk:"dnsprofilename"`
	Ipaddress        types.List   `tfsdk:"ipaddress"`
	Preferredloclist types.List   `tfsdk:"preferredloclist"`
	Ttl              types.Int64  `tfsdk:"ttl"`
	Viewname         types.String `tfsdk:"viewname"`
}

func (r *DnsactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsaction resource.",
			},
			"actionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dns action.",
			},
			"actiontype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The type of DNS action that is being configured.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the transaction for which the action is chosen",
			},
			"ipaddress": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of IP address to be returned in case of rewrite_response actiontype. They can be of IPV4 or IPV6 type.\n        In case of set command We will remove all the IP address previously present in the action and will add new once given in set dns action command.",
			},
			"preferredloclist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The location list in priority order used for the given action.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Time to live, in seconds.",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The view name that must be used for the given action.",
			},
		},
	}
}

func dnsactionGetThePayloadFromtheConfig(ctx context.Context, data *DnsactionResourceModel) dns.Dnsaction {
	tflog.Debug(ctx, "In dnsactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsaction := dns.Dnsaction{}
	if !data.Actionname.IsNull() {
		dnsaction.Actionname = data.Actionname.ValueString()
	}
	if !data.Actiontype.IsNull() {
		dnsaction.Actiontype = data.Actiontype.ValueString()
	}
	if !data.Dnsprofilename.IsNull() {
		dnsaction.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Ttl.IsNull() {
		dnsaction.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.Viewname.IsNull() {
		dnsaction.Viewname = data.Viewname.ValueString()
	}

	return dnsaction
}

func dnsactionSetAttrFromGet(ctx context.Context, data *DnsactionResourceModel, getResponseData map[string]interface{}) *DnsactionResourceModel {
	tflog.Debug(ctx, "In dnsactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["actionname"]; ok && val != nil {
		data.Actionname = types.StringValue(val.(string))
	} else {
		data.Actionname = types.StringNull()
	}
	if val, ok := getResponseData["actiontype"]; ok && val != nil {
		data.Actiontype = types.StringValue(val.(string))
	} else {
		data.Actiontype = types.StringNull()
	}
	if val, ok := getResponseData["dnsprofilename"]; ok && val != nil {
		data.Dnsprofilename = types.StringValue(val.(string))
	} else {
		data.Dnsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["viewname"]; ok && val != nil {
		data.Viewname = types.StringValue(val.(string))
	} else {
		data.Viewname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Actionname.ValueString())

	return data
}
