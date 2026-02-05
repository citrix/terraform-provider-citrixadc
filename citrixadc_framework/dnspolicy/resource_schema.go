package dnspolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnspolicyResourceModel describes the resource data model.
type DnspolicyResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Actionname        types.String `tfsdk:"actionname"`
	Cachebypass       types.String `tfsdk:"cachebypass"`
	Drop              types.String `tfsdk:"drop"`
	Logaction         types.String `tfsdk:"logaction"`
	Name              types.String `tfsdk:"name"`
	Preferredlocation types.String `tfsdk:"preferredlocation"`
	Preferredloclist  types.List   `tfsdk:"preferredloclist"`
	Rule              types.String `tfsdk:"rule"`
	Viewname          types.String `tfsdk:"viewname"`
}

func (r *DnspolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnspolicy resource.",
			},
			"actionname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS action to perform when the rule evaluates to TRUE. The built in actions function as follows:\n* dns_default_act_Drop. Drop the DNS request.\n* dns_default_act_Cachebypass. Bypass the DNS cache and forward the request to the name server.\nYou can create custom actions by using the add dns action command in the CLI or the DNS > Actions > Create DNS Action dialog box in the Citrix ADC configuration utility.",
			},
			"cachebypass": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By pass dns cache for this.",
			},
			"drop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The dns packet must be dropped.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DNS policy.",
			},
			"preferredlocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The location used for the given policy. This is deprecated attribute. Please use -prefLocList",
			},
			"preferredloclist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The location list in priority order used for the given policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression against which DNS traffic is evaluated.\nNote:\n* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n* If the expression itself includes double quotation marks, you must escape the quotations by using the  character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.\nExample: CLIENT.UDP.DNS.DOMAIN.EQ(\"domainname\")",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The view name that must be used for the given policy.",
			},
		},
	}
}

func dnspolicyGetThePayloadFromtheConfig(ctx context.Context, data *DnspolicyResourceModel) dns.Dnspolicy {
	tflog.Debug(ctx, "In dnspolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnspolicy := dns.Dnspolicy{}
	if !data.Actionname.IsNull() {
		dnspolicy.Actionname = data.Actionname.ValueString()
	}
	if !data.Cachebypass.IsNull() {
		dnspolicy.Cachebypass = data.Cachebypass.ValueString()
	}
	if !data.Drop.IsNull() {
		dnspolicy.Drop = data.Drop.ValueString()
	}
	if !data.Logaction.IsNull() {
		dnspolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		dnspolicy.Name = data.Name.ValueString()
	}
	if !data.Preferredlocation.IsNull() {
		dnspolicy.Preferredlocation = data.Preferredlocation.ValueString()
	}
	if !data.Rule.IsNull() {
		dnspolicy.Rule = data.Rule.ValueString()
	}
	if !data.Viewname.IsNull() {
		dnspolicy.Viewname = data.Viewname.ValueString()
	}

	return dnspolicy
}

func dnspolicySetAttrFromGet(ctx context.Context, data *DnspolicyResourceModel, getResponseData map[string]interface{}) *DnspolicyResourceModel {
	tflog.Debug(ctx, "In dnspolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["actionname"]; ok && val != nil {
		data.Actionname = types.StringValue(val.(string))
	} else {
		data.Actionname = types.StringNull()
	}
	if val, ok := getResponseData["cachebypass"]; ok && val != nil {
		data.Cachebypass = types.StringValue(val.(string))
	} else {
		data.Cachebypass = types.StringNull()
	}
	if val, ok := getResponseData["drop"]; ok && val != nil {
		data.Drop = types.StringValue(val.(string))
	} else {
		data.Drop = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["preferredlocation"]; ok && val != nil {
		data.Preferredlocation = types.StringValue(val.(string))
	} else {
		data.Preferredlocation = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["viewname"]; ok && val != nil {
		data.Viewname = types.StringValue(val.(string))
	} else {
		data.Viewname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
