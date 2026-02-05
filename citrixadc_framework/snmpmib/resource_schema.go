package snmpmib

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SnmpmibResourceModel describes the resource data model.
type SnmpmibResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Contact   types.String `tfsdk:"contact"`
	Customid  types.String `tfsdk:"customid"`
	Location  types.String `tfsdk:"location"`
	Name      types.String `tfsdk:"name"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
}

func (r *SnmpmibResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpmib resource.",
			},
			"contact": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("WebMaster (default)"),
				Description: "Name of the administrator for this Citrix ADC. Along with the name, you can include information on how to contact this person, such as a phone number or an email address. Can consist of 1 to 127 characters that include uppercase and  lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the information includes one or more spaces, enclose it in double or single quotation marks (for example, \"my contact\" or 'my contact').",
			},
			"customid": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Default"),
				Description: "Custom identification number for the Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a custom identification that helps identify the Citrix ADC appliance.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the ID includes one or more spaces, enclose it in double or single quotation marks (for example, \"my ID\" or 'my ID').",
			},
			"location": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("POP (default)"),
				Description: "Physical location of the Citrix ADC. For example, you can specify building name, lab number, and rack number. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the location includes one or more spaces, enclose it in double or single quotation marks (for example, \"my location\" or 'my location').",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NetScaler"),
				Description: "Name for this Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the Citrix ADC appliance.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(-1),
				Description: "ID of the cluster node for which we are setting the mib. This is a mandatory argument to set snmp mib on CLIP.",
			},
		},
	}
}

func snmpmibGetThePayloadFromtheConfig(ctx context.Context, data *SnmpmibResourceModel) snmp.Snmpmib {
	tflog.Debug(ctx, "In snmpmibGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmpmib := snmp.Snmpmib{}
	if !data.Contact.IsNull() {
		snmpmib.Contact = data.Contact.ValueString()
	}
	if !data.Customid.IsNull() {
		snmpmib.Customid = data.Customid.ValueString()
	}
	if !data.Location.IsNull() {
		snmpmib.Location = data.Location.ValueString()
	}
	if !data.Name.IsNull() {
		snmpmib.Name = data.Name.ValueString()
	}
	if !data.Ownernode.IsNull() {
		snmpmib.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return snmpmib
}

func snmpmibSetAttrFromGet(ctx context.Context, data *SnmpmibResourceModel, getResponseData map[string]interface{}) *SnmpmibResourceModel {
	tflog.Debug(ctx, "In snmpmibSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["contact"]; ok && val != nil {
		data.Contact = types.StringValue(val.(string))
	} else {
		data.Contact = types.StringNull()
	}
	if val, ok := getResponseData["customid"]; ok && val != nil {
		data.Customid = types.StringValue(val.(string))
	} else {
		data.Customid = types.StringNull()
	}
	if val, ok := getResponseData["location"]; ok && val != nil {
		data.Location = types.StringValue(val.(string))
	} else {
		data.Location = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	return data
}
