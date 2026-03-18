package appflowcollector

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppflowcollectorResourceModel describes the resource data model.
type AppflowcollectorResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Ipaddress  types.String `tfsdk:"ipaddress"`
	Name       types.String `tfsdk:"name"`
	Netprofile types.String `tfsdk:"netprofile"`
	Newname    types.String `tfsdk:"newname"`
	Port       types.Int64  `tfsdk:"port"`
	Transport  types.String `tfsdk:"transport"`
}

func (r *AppflowcollectorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appflowcollector resource.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 address of the collector.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n Only four collectors can be configured.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow collector\" or 'my appflow collector').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Netprofile to associate with the collector. The IP address defined in the profile is used as the source IP address for AppFlow traffic for this collector.  If you do not set this parameter, the Citrix ADC IP (NSIP) address is used as the source IP address.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must\ncontain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow coll\" or 'my appflow coll').",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the collector listens.",
			},
			"transport": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ipfix"),
				Description: "Type of collector: either logstream or ipfix or rest.",
			},
		},
	}
}

func appflowcollectorGetThePayloadFromtheConfig(ctx context.Context, data *AppflowcollectorResourceModel) appflow.Appflowcollector {
	tflog.Debug(ctx, "In appflowcollectorGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appflowcollector := appflow.Appflowcollector{}
	if !data.Ipaddress.IsNull() {
		appflowcollector.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		appflowcollector.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() {
		appflowcollector.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Newname.IsNull() {
		appflowcollector.Newname = data.Newname.ValueString()
	}
	if !data.Port.IsNull() {
		appflowcollector.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Transport.IsNull() {
		appflowcollector.Transport = data.Transport.ValueString()
	}

	return appflowcollector
}

func appflowcollectorSetAttrFromGet(ctx context.Context, data *AppflowcollectorResourceModel, getResponseData map[string]interface{}) *AppflowcollectorResourceModel {
	tflog.Debug(ctx, "In appflowcollectorSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netprofile"]; ok && val != nil {
		data.Netprofile = types.StringValue(val.(string))
	} else {
		data.Netprofile = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["transport"]; ok && val != nil {
		data.Transport = types.StringValue(val.(string))
	} else {
		data.Transport = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
