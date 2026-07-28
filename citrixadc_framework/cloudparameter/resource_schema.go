package cloudparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CloudparameterResourceModel describes the resource data model.
type CloudparameterResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Activationcode     types.String `tfsdk:"activationcode"`
	Connectorresidence types.String `tfsdk:"connectorresidence"`
	Controllerfqdn     types.String `tfsdk:"controllerfqdn"`
	Controllerport     types.Int64  `tfsdk:"controllerport"`
	Customerid         types.String `tfsdk:"customerid"`
	Deployment         types.String `tfsdk:"deployment"`
	Instanceid         types.String `tfsdk:"instanceid"`
	Resourcelocation   types.String `tfsdk:"resourcelocation"`
}

func (r *CloudparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudparameter resource.",
			},
			"activationcode": schema.StringAttribute{
				Optional:    true,
				Description: "Activation code for the NGS Connector instance",
			},
			"connectorresidence": schema.StringAttribute{
				Optional:    true,
				Description: "Identifies whether the connector is located Onprem, Aws or Azure",
			},
			"controllerfqdn": schema.StringAttribute{
				Optional:    true,
				Description: "FQDN of the controller to which the Citrix ADC SDProxy Connects",
			},
			"controllerport": schema.Int64Attribute{
				Optional:    true,
				Description: "Port number of the controller to which the Citrix ADC SDProxy connects",
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Description: "Customer ID of the citrix cloud customer",
			},
			"deployment": schema.StringAttribute{
				Optional:    true,
				Description: "Describes if the customer is a Staging/Production or Dev Citrix Cloud customer",
			},
			"instanceid": schema.StringAttribute{
				Optional:    true,
				Description: "Instance ID of the customer provided by Trust",
			},
			"resourcelocation": schema.StringAttribute{
				Optional:    true,
				Description: "Resource Location of the customer provided by Trust",
			},
		},
	}
}

func cloudparameterGetThePayloadFromthePlan(ctx context.Context, data *CloudparameterResourceModel) cloud.Cloudparameter {
	tflog.Debug(ctx, "In cloudparameterGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudparameter := cloud.Cloudparameter{}
	if !data.Activationcode.IsNull() && !data.Activationcode.IsUnknown() {
		cloudparameter.Activationcode = data.Activationcode.ValueString()
	}
	if !data.Connectorresidence.IsNull() && !data.Connectorresidence.IsUnknown() {
		cloudparameter.Connectorresidence = data.Connectorresidence.ValueString()
	}
	if !data.Controllerfqdn.IsNull() && !data.Controllerfqdn.IsUnknown() {
		cloudparameter.Controllerfqdn = data.Controllerfqdn.ValueString()
	}
	if !data.Controllerport.IsNull() && !data.Controllerport.IsUnknown() {
		cloudparameter.Controllerport = utils.IntPtr(int(data.Controllerport.ValueInt64()))
	}
	if !data.Customerid.IsNull() && !data.Customerid.IsUnknown() {
		cloudparameter.Customerid = data.Customerid.ValueString()
	}
	if !data.Deployment.IsNull() && !data.Deployment.IsUnknown() {
		cloudparameter.Deployment = data.Deployment.ValueString()
	}
	if !data.Instanceid.IsNull() && !data.Instanceid.IsUnknown() {
		cloudparameter.Instanceid = data.Instanceid.ValueString()
	}
	if !data.Resourcelocation.IsNull() && !data.Resourcelocation.IsUnknown() {
		cloudparameter.Resourcelocation = data.Resourcelocation.ValueString()
	}

	return cloudparameter
}

// cloudparameterSetAttrFromGet is the RESOURCE state setter.
//
// cloudparameter is a singleton SET-GET resource whose GET view on the live
// appliance is unreliable for round-tripping user config:
//   - activationcode is WRITE-ONLY: it is never returned by GET, so refreshing
//     from the response would clobber the user's value to null (Pattern 7,
//     write-only-field variant) and produce perpetual drift.
//   - The testbed ADC overrides some enum attrs (deployment, connectorresidence)
//     with server-side defaults and drops other configured fields to null in the
//     GET response (Pattern 7, server-overrides variant). Because these attrs are
//     Optional-only (not Computed) in the schema, Terraform requires the
//     post-apply state to match config exactly; adopting the server view would
//     fail with "inconsistent result after apply".
//
// Therefore the resource setter PRESERVES the plan/state values for every
// user-facing attribute and only (re)affirms the static ID. The datasource,
// which has no prior state to preserve, uses the separate faithful setter below.
func cloudparameterSetAttrFromGet(ctx context.Context, data *CloudparameterResourceModel, getResponseData map[string]interface{}) *CloudparameterResourceModel {
	tflog.Debug(ctx, "In cloudparameterSetAttrFromGet Function (preserving plan/state values)")

	// Do not overwrite any user-facing attribute from the GET response:
	// activationcode is write-only and the other attrs may be overridden/dropped
	// by the appliance. Preserve the existing model values verbatim.

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("cloudparameter-config")

	return data
}

// cloudparameterSetAttrFromGetForDatasource is the DATASOURCE state setter.
//
// A datasource has no prior plan/state to preserve, so it faithfully copies
// every attribute the GET response returns. activationcode is intentionally NOT
// read here: the NITRO GET/show never returns it, so it is left null rather than
// advertised as a readable (and misleading) value.
func cloudparameterSetAttrFromGetForDatasource(ctx context.Context, data *CloudparameterResourceModel, getResponseData map[string]interface{}) *CloudparameterResourceModel {
	tflog.Debug(ctx, "In cloudparameterSetAttrFromGetForDatasource Function")

	// activationcode is write-only (never returned by GET) - leave it null.
	data.Activationcode = types.StringNull()
	if val, ok := getResponseData["connectorresidence"]; ok && val != nil {
		data.Connectorresidence = types.StringValue(val.(string))
	} else {
		data.Connectorresidence = types.StringNull()
	}
	if val, ok := getResponseData["controllerfqdn"]; ok && val != nil {
		data.Controllerfqdn = types.StringValue(val.(string))
	} else {
		data.Controllerfqdn = types.StringNull()
	}
	if val, ok := getResponseData["controllerport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Controllerport = types.Int64Value(intVal)
		}
	} else {
		data.Controllerport = types.Int64Null()
	}
	if val, ok := getResponseData["customerid"]; ok && val != nil {
		data.Customerid = types.StringValue(val.(string))
	} else {
		data.Customerid = types.StringNull()
	}
	if val, ok := getResponseData["deployment"]; ok && val != nil {
		data.Deployment = types.StringValue(val.(string))
	} else {
		data.Deployment = types.StringNull()
	}
	if val, ok := getResponseData["instanceid"]; ok && val != nil {
		data.Instanceid = types.StringValue(val.(string))
	} else {
		data.Instanceid = types.StringNull()
	}
	if val, ok := getResponseData["resourcelocation"]; ok && val != nil {
		data.Resourcelocation = types.StringValue(val.(string))
	} else {
		data.Resourcelocation = types.StringNull()
	}

	// Set ID for the datasource (no Create step to set it)
	data.Id = types.StringValue("cloudparameter-config")

	return data
}
