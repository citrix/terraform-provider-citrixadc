package dnspolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

func dnspolicyGetThePayloadFromthePlan(ctx context.Context, data *DnspolicyResourceModel) dns.Dnspolicy {
	tflog.Debug(ctx, "In dnspolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnspolicy := dns.Dnspolicy{}
	if !data.Actionname.IsNull() && !data.Actionname.IsUnknown() {
		dnspolicy.Actionname = data.Actionname.ValueString()
	}
	if !data.Cachebypass.IsNull() && !data.Cachebypass.IsUnknown() {
		dnspolicy.Cachebypass = data.Cachebypass.ValueString()
	}
	if !data.Drop.IsNull() && !data.Drop.IsUnknown() {
		dnspolicy.Drop = data.Drop.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		dnspolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		dnspolicy.Name = data.Name.ValueString()
	}
	if !data.Preferredlocation.IsNull() && !data.Preferredlocation.IsUnknown() {
		dnspolicy.Preferredlocation = data.Preferredlocation.ValueString()
	}
	if !data.Preferredloclist.IsNull() && !data.Preferredloclist.IsUnknown() {
		var preferredloclist []string
		data.Preferredloclist.ElementsAs(ctx, &preferredloclist, false)
		dnspolicy.Preferredloclist = preferredloclist
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		dnspolicy.Rule = data.Rule.ValueString()
	}
	if !data.Viewname.IsNull() && !data.Viewname.IsUnknown() {
		dnspolicy.Viewname = data.Viewname.ValueString()
	}

	return dnspolicy
}

func dnspolicySetAttrFromGet(ctx context.Context, data *DnspolicyResourceModel, getResponseData map[string]interface{}) *DnspolicyResourceModel {
	tflog.Debug(ctx, "In dnspolicySetAttrFromGet Function")

	// Convert API response to model.
	// Optional+Computed attributes are guarded: when NITRO omits the value from
	// the GET response, a configured value is preserved (only a fresh unknown is
	// nulled) so a configured value that is not echoed back does not cause an
	// "inconsistent result after apply" error.
	if val, ok := getResponseData["actionname"]; ok && val != nil {
		data.Actionname = types.StringValue(val.(string))
	} else if data.Actionname.IsUnknown() {
		data.Actionname = types.StringNull()
	}
	if val, ok := getResponseData["cachebypass"]; ok && val != nil {
		data.Cachebypass = types.StringValue(val.(string))
	} else if data.Cachebypass.IsUnknown() {
		data.Cachebypass = types.StringNull()
	}
	if val, ok := getResponseData["drop"]; ok && val != nil {
		data.Drop = types.StringValue(val.(string))
	} else if data.Drop.IsUnknown() {
		data.Drop = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else if data.Logaction.IsUnknown() {
		data.Logaction = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["preferredlocation"]; ok && val != nil {
		data.Preferredlocation = types.StringValue(val.(string))
	} else if data.Preferredlocation.IsUnknown() {
		data.Preferredlocation = types.StringNull()
	}
	if val, ok := getResponseData["preferredloclist"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Preferredloclist = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Preferredloclist = listValue
		default:
			data.Preferredloclist = types.ListNull(types.StringType)
		}
	} else if data.Preferredloclist.IsUnknown() {
		data.Preferredloclist = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["viewname"]; ok && val != nil {
		data.Viewname = types.StringValue(val.(string))
	} else if data.Viewname.IsUnknown() {
		data.Viewname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
