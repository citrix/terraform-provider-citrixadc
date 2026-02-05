package contentinspectionprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ContentinspectionprofileResourceModel describes the resource data model.
type ContentinspectionprofileResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Egressinterface  types.String `tfsdk:"egressinterface"`
	Egressvlan       types.Int64  `tfsdk:"egressvlan"`
	Ingressinterface types.String `tfsdk:"ingressinterface"`
	Ingressvlan      types.Int64  `tfsdk:"ingressvlan"`
	Iptunnel         types.String `tfsdk:"iptunnel"`
	Name             types.String `tfsdk:"name"`
	Type             types.String `tfsdk:"type"`
}

func (r *ContentinspectionprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionprofile resource.",
			},
			"egressinterface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Egress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of type INLINEINSPECTION or MIRROR.",
			},
			"egressvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Egress Vlan for CI",
			},
			"ingressinterface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ingress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of IPS type.",
			},
			"ingressvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Ingress Vlan for CI",
			},
			"iptunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP Tunnel for CI profile. It is used while creating a ContentInspection profile of type MIRROR when the IDS device is in a different network",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of a ContentInspection profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a IPS profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my ips profile\" or 'my ips profile'\\).",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of ContentInspection profile. Following types are available to configure:\n           INLINEINSPECTION : To inspect the packets/requests using IPS.\n	   MIRROR : To forward cloned packets.",
			},
		},
	}
}

func contentinspectionprofileGetThePayloadFromtheConfig(ctx context.Context, data *ContentinspectionprofileResourceModel) contentinspection.Contentinspectionprofile {
	tflog.Debug(ctx, "In contentinspectionprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	contentinspectionprofile := contentinspection.Contentinspectionprofile{}
	if !data.Egressinterface.IsNull() {
		contentinspectionprofile.Egressinterface = data.Egressinterface.ValueString()
	}
	if !data.Egressvlan.IsNull() {
		contentinspectionprofile.Egressvlan = utils.IntPtr(int(data.Egressvlan.ValueInt64()))
	}
	if !data.Ingressinterface.IsNull() {
		contentinspectionprofile.Ingressinterface = data.Ingressinterface.ValueString()
	}
	if !data.Ingressvlan.IsNull() {
		contentinspectionprofile.Ingressvlan = utils.IntPtr(int(data.Ingressvlan.ValueInt64()))
	}
	if !data.Iptunnel.IsNull() {
		contentinspectionprofile.Iptunnel = data.Iptunnel.ValueString()
	}
	if !data.Name.IsNull() {
		contentinspectionprofile.Name = data.Name.ValueString()
	}
	if !data.Type.IsNull() {
		contentinspectionprofile.Type = data.Type.ValueString()
	}

	return contentinspectionprofile
}

func contentinspectionprofileSetAttrFromGet(ctx context.Context, data *ContentinspectionprofileResourceModel, getResponseData map[string]interface{}) *ContentinspectionprofileResourceModel {
	tflog.Debug(ctx, "In contentinspectionprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["egressinterface"]; ok && val != nil {
		data.Egressinterface = types.StringValue(val.(string))
	} else {
		data.Egressinterface = types.StringNull()
	}
	if val, ok := getResponseData["egressvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Egressvlan = types.Int64Value(intVal)
		}
	} else {
		data.Egressvlan = types.Int64Null()
	}
	if val, ok := getResponseData["ingressinterface"]; ok && val != nil {
		data.Ingressinterface = types.StringValue(val.(string))
	} else {
		data.Ingressinterface = types.StringNull()
	}
	if val, ok := getResponseData["ingressvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ingressvlan = types.Int64Value(intVal)
		}
	} else {
		data.Ingressvlan = types.Int64Null()
	}
	if val, ok := getResponseData["iptunnel"]; ok && val != nil {
		data.Iptunnel = types.StringValue(val.(string))
	} else {
		data.Iptunnel = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
