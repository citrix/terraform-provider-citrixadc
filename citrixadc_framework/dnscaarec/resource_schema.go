package dnscaarec

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnscaarecResourceModel describes the resource data model.
type DnscaarecResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"`
	Ecssubnet   types.String `tfsdk:"ecssubnet"`
	Flag        types.String `tfsdk:"flag"`
	Recordid    types.Int64  `tfsdk:"recordid"`
	Tag         types.String `tfsdk:"tag"`
	Ttl         types.Int64  `tfsdk:"ttl"`
	Valuestring types.String `tfsdk:"valuestring"`
}

func (r *DnscaarecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnscaarec resource.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name of the CAA record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet for which the cached CAA record need to be removed.",
			},
			"flag": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Flag associated with the CAA record.",
			},
			"recordid": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Unique, internally generated record ID. View the details of the CAA record to obtain its record ID. Records can be removedby either specifying the domain name and record id OR by specifying domain name and all other CAA record attributes as was supplied during the add command.",
			},
			"tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("issue"),
				Description: "String that represents the identifier of the property represented by the CAA record. The RFC currently defines three available tags - issue, issuwild and iodef.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(3600),
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
			"valuestring": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Value associated with the chosen property tag in the CAA resource record. Enclose the string in single or double quotation marks.",
			},
		},
	}
}

func dnscaarecGetThePayloadFromthePlan(ctx context.Context, data *DnscaarecResourceModel) dns.Dnscaarec {
	tflog.Debug(ctx, "In dnscaarecGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// NOTE: type/nodeid are GET-only filter args and ecssubnet/recordid are
	// delete-only args; none are accepted in the add payload, so they are
	// deliberately excluded here.
	dnscaarec := dns.Dnscaarec{}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		dnscaarec.Domain = data.Domain.ValueString()
	}
	if !data.Flag.IsNull() && !data.Flag.IsUnknown() {
		dnscaarec.Flag = data.Flag.ValueString()
	}
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		dnscaarec.Tag = data.Tag.ValueString()
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		dnscaarec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.Valuestring.IsNull() && !data.Valuestring.IsUnknown() {
		dnscaarec.Valuestring = data.Valuestring.ValueString()
	}

	return dnscaarec
}

func dnscaarecSetAttrFromGet(ctx context.Context, data *DnscaarecResourceModel, getResponseData map[string]interface{}) *DnscaarecResourceModel {
	tflog.Debug(ctx, "In dnscaarecSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["flag"]; ok && val != nil {
		data.Flag = types.StringValue(val.(string))
	} else {
		data.Flag = types.StringNull()
	}
	if val, ok := getResponseData["recordid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Recordid = types.Int64Value(intVal)
		}
	} else {
		data.Recordid = types.Int64Null()
	}
	if val, ok := getResponseData["tag"]; ok && val != nil {
		data.Tag = types.StringValue(val.(string))
	} else {
		data.Tag = types.StringNull()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["valuestring"]; ok && val != nil {
		data.Valuestring = types.StringValue(val.(string))
	} else {
		data.Valuestring = types.StringNull()
	}

	// Set ID for the resource
	// Composite key: domain,recordid (multiple CAA records may share a domain).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("domain:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Domain.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("recordid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Recordid.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
