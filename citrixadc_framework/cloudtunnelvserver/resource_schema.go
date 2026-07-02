package cloudtunnelvserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cloudtunnel"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CloudtunnelvserverResourceModel describes the resource data model.
type CloudtunnelvserverResourceModel struct {
	Id types.String `tfsdk:"id"`
	Listenpolicy types.String `tfsdk:"listenpolicy"`
	Listenpriority types.Int64 `tfsdk:"listenpriority"`
	Name types.String `tfsdk:"name"`
	Servicetype types.String `tfsdk:"servicetype"`
}

func (r *CloudtunnelvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudtunnelvserver resource.",
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the listen policy for the Cloud Tunnel virtual server. Can be either a named expression or an expression. The Cloud Tunnel virtual server processes only the traffic for which the expression evaluates to true.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Cloud Tunnel virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space,colon (:), at (@), equals (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example,\n\"my server\" or 'my server').",
			},
			"servicetype": schema.StringAttribute{
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ServiceType of Listener using which traffic will be tunneled through cloud tunnel server.",
			},
		},
	}
}

func cloudtunnelvserverGetThePayloadFromthePlan(ctx context.Context, data *CloudtunnelvserverResourceModel) cloudtunnel.Cloudtunnelvserver {
	tflog.Debug(ctx, "In cloudtunnelvserverGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudtunnelvserver := cloudtunnel.Cloudtunnelvserver{}
	if !data.Listenpolicy.IsNull() && !data.Listenpolicy.IsUnknown() {
		cloudtunnelvserver.Listenpolicy = data.Listenpolicy.ValueString()
	}
	if !data.Listenpriority.IsNull() && !data.Listenpriority.IsUnknown() {
		cloudtunnelvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		cloudtunnelvserver.Name = data.Name.ValueString()
	}
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		cloudtunnelvserver.Servicetype = data.Servicetype.ValueString()
	}

	return cloudtunnelvserver
}

func cloudtunnelvserverSetAttrFromGet(ctx context.Context, data *CloudtunnelvserverResourceModel, getResponseData map[string]interface{}) *CloudtunnelvserverResourceModel {
	tflog.Debug(ctx, "In cloudtunnelvserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["listenpolicy"]; ok && val != nil {
		data.Listenpolicy = types.StringValue(val.(string))
	} else {
		data.Listenpolicy = types.StringNull()
	}
	if val, ok := getResponseData["listenpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Listenpriority = types.Int64Value(intVal)
		}
	} else {
		data.Listenpriority = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}