package cmpaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmpactionResourceModel describes the resource data model.
type CmpactionResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Addvaryheader   types.String `tfsdk:"addvaryheader"`
	Cmptype         types.String `tfsdk:"cmptype"`
	Deltatype       types.String `tfsdk:"deltatype"`
	Name            types.String `tfsdk:"name"`
	Newname         types.String `tfsdk:"newname"`
	Varyheadervalue types.String `tfsdk:"varyheadervalue"`
}

func (r *CmpactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cmpaction resource.",
			},
			"addvaryheader": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("GLOBAL"),
				Description: "Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.",
			},
			"cmptype": schema.StringAttribute{
				Required:    true,
				Description: "Type of compression performed by this action.\nAvailable settings function as follows:\n* COMPRESS - Apply GZIP or DEFLATE compression to the response, depending on the request header. Prefer GZIP.\n* GZIP - Apply GZIP compression.\n* DEFLATE - Apply DEFLATE compression.\n* NOCOMPRESS - Do not compress the response if the request matches a policy that uses this action.",
			},
			"deltatype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PERURL"),
				Description: "The type of delta action (if delta type compression action is defined).",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp action\" or 'my cmp action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\nChoose a name that can be correlated with the function that the action performs.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp action\" or 'my cmp action').",
			},
			"varyheadervalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The value of the HTTP Vary header for compressed responses.",
			},
		},
	}
}

func cmpactionGetThePayloadFromtheConfig(ctx context.Context, data *CmpactionResourceModel) cmp.Cmpaction {
	tflog.Debug(ctx, "In cmpactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cmpaction := cmp.Cmpaction{}
	if !data.Addvaryheader.IsNull() {
		cmpaction.Addvaryheader = data.Addvaryheader.ValueString()
	}
	if !data.Cmptype.IsNull() {
		cmpaction.Cmptype = data.Cmptype.ValueString()
	}
	if !data.Deltatype.IsNull() {
		cmpaction.Deltatype = data.Deltatype.ValueString()
	}
	if !data.Name.IsNull() {
		cmpaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		cmpaction.Newname = data.Newname.ValueString()
	}
	if !data.Varyheadervalue.IsNull() {
		cmpaction.Varyheadervalue = data.Varyheadervalue.ValueString()
	}

	return cmpaction
}

func cmpactionSetAttrFromGet(ctx context.Context, data *CmpactionResourceModel, getResponseData map[string]interface{}) *CmpactionResourceModel {
	tflog.Debug(ctx, "In cmpactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["addvaryheader"]; ok && val != nil {
		data.Addvaryheader = types.StringValue(val.(string))
	} else {
		data.Addvaryheader = types.StringNull()
	}
	if val, ok := getResponseData["cmptype"]; ok && val != nil {
		data.Cmptype = types.StringValue(val.(string))
	} else {
		data.Cmptype = types.StringNull()
	}
	if val, ok := getResponseData["deltatype"]; ok && val != nil {
		data.Deltatype = types.StringValue(val.(string))
	} else {
		data.Deltatype = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["varyheadervalue"]; ok && val != nil {
		data.Varyheadervalue = types.StringValue(val.(string))
	} else {
		data.Varyheadervalue = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
