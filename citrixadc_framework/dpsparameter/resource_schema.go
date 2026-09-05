package dpsparameter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DpsparameterResourceModel describes the resource data model.
//
// NOTE: The adc-nitro-go SDK has NO typed struct for "dpsparameter", so the
// NITRO payload is built with a local untyped struct (dpsparameterPayload,
// defined below) carrying json tags. Fields were derived directly from the NITRO
// doc (nitro-rest-73x/dps/dpsparameter.html). FLAG FOR REVIEW if the SDK later
// ships a dps.Dpsparameter type.
type DpsparameterResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Customerid types.String `tfsdk:"customerid"`
	Deployment types.String `tfsdk:"deployment"`
	Serviceurl types.String `tfsdk:"serviceurl"`
}

// dpsparameterPayload is the untyped NITRO payload for dpsparameter. It is
// marshalled by the generic NitroClient calls as {"dpsparameter": {...}}. Only
// writable attributes are included; read-only props (builtin, feature,
// _nextgenapiresource) are never sent. omitempty keeps unset fields out of the
// request body.
type dpsparameterPayload struct {
	Customerid string `json:"customerid,omitempty"`
	Deployment string `json:"deployment,omitempty"`
	Serviceurl string `json:"serviceurl,omitempty"`
}

func (r *DpsparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dpsparameter resource.",
			},
			// customerid/deployment/serviceurl each revert to a KNOWN non-null
			// appliance default on unset (customerid="None", deployment="COMMERCIAL",
			// serviceurl="https://device-posture-controller.cloud.com") and NITRO
			// always returns them on GET (never omitted). A matching schema Default
			// therefore keeps the plan stable (config-omit -> default == state) and
			// drives the revert-on-removal without the perpetual "known after apply"
			// churn an unset-on-remove/unknown-marking modifier would cause. See the
			// systemautosaveparam resource for the same pattern and rationale.
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("None"),
				Description: "Customer ID of the Citrix Cloud customer.",
			},
			"deployment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("COMMERCIAL"),
				Description: "Describes if the customer is connecting to Commerical/JapanCloud/Gov Citrix Cloud customer. Possible values = COMMERCIAL, GOV, JAPANCLOUD",
			},
			"serviceurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("https://device-posture-controller.cloud.com"),
				Description: "Service URL of the Citrix Cloud customer.",
			},
		},
	}
}

func dpsparameterGetThePayloadFromthePlan(ctx context.Context, data *DpsparameterResourceModel) dpsparameterPayload {
	tflog.Debug(ctx, "In dpsparameterGetThePayloadFromthePlan Function")

	// Create API request body from the model (untyped payload; no SDK struct)
	dpsparameter := dpsparameterPayload{}
	if !data.Customerid.IsNull() && !data.Customerid.IsUnknown() {
		dpsparameter.Customerid = data.Customerid.ValueString()
	}
	if !data.Deployment.IsNull() && !data.Deployment.IsUnknown() {
		dpsparameter.Deployment = data.Deployment.ValueString()
	}
	if !data.Serviceurl.IsNull() && !data.Serviceurl.IsUnknown() {
		dpsparameter.Serviceurl = data.Serviceurl.ValueString()
	}

	return dpsparameter
}

func dpsparameterSetAttrFromGet(ctx context.Context, data *DpsparameterResourceModel, getResponseData map[string]interface{}) *DpsparameterResourceModel {
	tflog.Debug(ctx, "In dpsparameterSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["serviceurl"]; ok && val != nil {
		data.Serviceurl = types.StringValue(val.(string))
	} else {
		data.Serviceurl = types.StringNull()
	}
	// Read-only NITRO props (builtin, feature, _nextgenapiresource) are omitted
	// from the schema and intentionally not mapped.

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("dpsparameter-config")

	return data
}
