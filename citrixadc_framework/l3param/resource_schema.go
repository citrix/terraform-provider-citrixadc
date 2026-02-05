package l3param

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// L3paramResourceModel describes the resource data model.
type L3paramResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acllogtime           types.Int64  `tfsdk:"acllogtime"`
	Allowclasseipv4      types.String `tfsdk:"allowclasseipv4"`
	Dropdfflag           types.String `tfsdk:"dropdfflag"`
	Dropipfragments      types.String `tfsdk:"dropipfragments"`
	Dynamicrouting       types.String `tfsdk:"dynamicrouting"`
	Externalloopback     types.String `tfsdk:"externalloopback"`
	Forwardicmpfragments types.String `tfsdk:"forwardicmpfragments"`
	Icmpgenratethreshold types.Int64  `tfsdk:"icmpgenratethreshold"`
	Implicitaclallow     types.String `tfsdk:"implicitaclallow"`
	Implicitpbr          types.String `tfsdk:"implicitpbr"`
	Ipv6dynamicrouting   types.String `tfsdk:"ipv6dynamicrouting"`
	Miproundrobin        types.String `tfsdk:"miproundrobin"`
	Overridernat         types.String `tfsdk:"overridernat"`
	Srcnat               types.String `tfsdk:"srcnat"`
	Tnlpmtuwoconn        types.String `tfsdk:"tnlpmtuwoconn"`
	Usipserverstraypkt   types.String `tfsdk:"usipserverstraypkt"`
}

func (r *L3paramResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the l3param resource.",
			},
			"acllogtime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5000),
				Description: "Parameter to tune acl logging time",
			},
			"allowclasseipv4": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable IPv4 Class E address clients",
			},
			"dropdfflag": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable dropping the IP DF flag.",
			},
			"dropipfragments": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable dropping of IP fragments.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable Dynamic routing on partition. This configuration is not applicable to default partition",
			},
			"externalloopback": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable external loopback.",
			},
			"forwardicmpfragments": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable forwarding of ICMP fragments.",
			},
			"icmpgenratethreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "NS generated ICMP pkts per 10ms rate threshold",
			},
			"implicitaclallow": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Do not apply ACLs for internal ports",
			},
			"implicitpbr": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable Policy Based Routing for control packets",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable IPv6 Dynamic routing",
			},
			"miproundrobin": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable round robin usage of mapped IPs.",
			},
			"overridernat": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "USNIP/USIP settings override RNAT settings for configured\n              service/virtual server traffic..",
			},
			"srcnat": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Perform NAT if only the source is in the private network",
			},
			"tnlpmtuwoconn": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable/Disable learning PMTU of IP tunnel when ICMP error does not contain connection information.",
			},
			"usipserverstraypkt": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable detection of stray server side pkts in USIP mode.",
			},
		},
	}
}

func l3paramGetThePayloadFromtheConfig(ctx context.Context, data *L3paramResourceModel) network.L3param {
	tflog.Debug(ctx, "In l3paramGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	l3param := network.L3param{}
	if !data.Acllogtime.IsNull() {
		l3param.Acllogtime = utils.IntPtr(int(data.Acllogtime.ValueInt64()))
	}
	if !data.Allowclasseipv4.IsNull() {
		l3param.Allowclasseipv4 = data.Allowclasseipv4.ValueString()
	}
	if !data.Dropdfflag.IsNull() {
		l3param.Dropdfflag = data.Dropdfflag.ValueString()
	}
	if !data.Dropipfragments.IsNull() {
		l3param.Dropipfragments = data.Dropipfragments.ValueString()
	}
	if !data.Dynamicrouting.IsNull() {
		l3param.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Externalloopback.IsNull() {
		l3param.Externalloopback = data.Externalloopback.ValueString()
	}
	if !data.Forwardicmpfragments.IsNull() {
		l3param.Forwardicmpfragments = data.Forwardicmpfragments.ValueString()
	}
	if !data.Icmpgenratethreshold.IsNull() {
		l3param.Icmpgenratethreshold = utils.IntPtr(int(data.Icmpgenratethreshold.ValueInt64()))
	}
	if !data.Implicitaclallow.IsNull() {
		l3param.Implicitaclallow = data.Implicitaclallow.ValueString()
	}
	if !data.Implicitpbr.IsNull() {
		l3param.Implicitpbr = data.Implicitpbr.ValueString()
	}
	if !data.Ipv6dynamicrouting.IsNull() {
		l3param.Ipv6dynamicrouting = data.Ipv6dynamicrouting.ValueString()
	}
	if !data.Miproundrobin.IsNull() {
		l3param.Miproundrobin = data.Miproundrobin.ValueString()
	}
	if !data.Overridernat.IsNull() {
		l3param.Overridernat = data.Overridernat.ValueString()
	}
	if !data.Srcnat.IsNull() {
		l3param.Srcnat = data.Srcnat.ValueString()
	}
	if !data.Tnlpmtuwoconn.IsNull() {
		l3param.Tnlpmtuwoconn = data.Tnlpmtuwoconn.ValueString()
	}
	if !data.Usipserverstraypkt.IsNull() {
		l3param.Usipserverstraypkt = data.Usipserverstraypkt.ValueString()
	}

	return l3param
}

func l3paramSetAttrFromGet(ctx context.Context, data *L3paramResourceModel, getResponseData map[string]interface{}) *L3paramResourceModel {
	tflog.Debug(ctx, "In l3paramSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["acllogtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Acllogtime = types.Int64Value(intVal)
		}
	} else {
		data.Acllogtime = types.Int64Null()
	}
	if val, ok := getResponseData["allowclasseipv4"]; ok && val != nil {
		data.Allowclasseipv4 = types.StringValue(val.(string))
	} else {
		data.Allowclasseipv4 = types.StringNull()
	}
	if val, ok := getResponseData["dropdfflag"]; ok && val != nil {
		data.Dropdfflag = types.StringValue(val.(string))
	} else {
		data.Dropdfflag = types.StringNull()
	}
	if val, ok := getResponseData["dropipfragments"]; ok && val != nil {
		data.Dropipfragments = types.StringValue(val.(string))
	} else {
		data.Dropipfragments = types.StringNull()
	}
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["externalloopback"]; ok && val != nil {
		data.Externalloopback = types.StringValue(val.(string))
	} else {
		data.Externalloopback = types.StringNull()
	}
	if val, ok := getResponseData["forwardicmpfragments"]; ok && val != nil {
		data.Forwardicmpfragments = types.StringValue(val.(string))
	} else {
		data.Forwardicmpfragments = types.StringNull()
	}
	if val, ok := getResponseData["icmpgenratethreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Icmpgenratethreshold = types.Int64Value(intVal)
		}
	} else {
		data.Icmpgenratethreshold = types.Int64Null()
	}
	if val, ok := getResponseData["implicitaclallow"]; ok && val != nil {
		data.Implicitaclallow = types.StringValue(val.(string))
	} else {
		data.Implicitaclallow = types.StringNull()
	}
	if val, ok := getResponseData["implicitpbr"]; ok && val != nil {
		data.Implicitpbr = types.StringValue(val.(string))
	} else {
		data.Implicitpbr = types.StringNull()
	}
	if val, ok := getResponseData["ipv6dynamicrouting"]; ok && val != nil {
		data.Ipv6dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Ipv6dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["miproundrobin"]; ok && val != nil {
		data.Miproundrobin = types.StringValue(val.(string))
	} else {
		data.Miproundrobin = types.StringNull()
	}
	if val, ok := getResponseData["overridernat"]; ok && val != nil {
		data.Overridernat = types.StringValue(val.(string))
	} else {
		data.Overridernat = types.StringNull()
	}
	if val, ok := getResponseData["srcnat"]; ok && val != nil {
		data.Srcnat = types.StringValue(val.(string))
	} else {
		data.Srcnat = types.StringNull()
	}
	if val, ok := getResponseData["tnlpmtuwoconn"]; ok && val != nil {
		data.Tnlpmtuwoconn = types.StringValue(val.(string))
	} else {
		data.Tnlpmtuwoconn = types.StringNull()
	}
	if val, ok := getResponseData["usipserverstraypkt"]; ok && val != nil {
		data.Usipserverstraypkt = types.StringValue(val.(string))
	} else {
		data.Usipserverstraypkt = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("l3param-config")

	return data
}
