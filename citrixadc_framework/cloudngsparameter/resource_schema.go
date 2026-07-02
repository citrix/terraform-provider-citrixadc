package cloudngsparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

)

// CloudngsparameterResourceModel describes the resource data model.
type CloudngsparameterResourceModel struct {
	Id types.String `tfsdk:"id"`
	Allowdtls12 types.String `tfsdk:"allowdtls12"`
	Allowedudtversion types.String `tfsdk:"allowedudtversion"`
	Blockonallowedngstktprof types.String `tfsdk:"blockonallowedngstktprof"`
	Csvserverticketingdecouple types.String `tfsdk:"csvserverticketingdecouple"`
}

func (r *CloudngsparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudngsparameter resource.",
			},
			"allowdtls12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables DTLS1.2 for client connections on CGS",
			},
			"allowedudtversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables the required UDT version to EDT connections in the CGS deployment",
			},
			"blockonallowedngstktprof": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables blocking connections authenticated with a ticket createdby by an entity not whitelisted in allowedngstktprofile",
			},
			"csvserverticketingdecouple": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables Decoupling CSVSERVER state from Ticketing Service state in the CGS deployment",
			},
		},
	}
}

func cloudngsparameterGetThePayloadFromthePlan(ctx context.Context, data *CloudngsparameterResourceModel) cloud.Cloudngsparameter {
	tflog.Debug(ctx, "In cloudngsparameterGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudngsparameter := cloud.Cloudngsparameter{}
	if !data.Allowdtls12.IsNull() && !data.Allowdtls12.IsUnknown() {
		cloudngsparameter.Allowdtls12 = data.Allowdtls12.ValueString()
	}
	if !data.Allowedudtversion.IsNull() && !data.Allowedudtversion.IsUnknown() {
		cloudngsparameter.Allowedudtversion = data.Allowedudtversion.ValueString()
	}
	if !data.Blockonallowedngstktprof.IsNull() && !data.Blockonallowedngstktprof.IsUnknown() {
		cloudngsparameter.Blockonallowedngstktprof = data.Blockonallowedngstktprof.ValueString()
	}
	if !data.Csvserverticketingdecouple.IsNull() && !data.Csvserverticketingdecouple.IsUnknown() {
		cloudngsparameter.Csvserverticketingdecouple = data.Csvserverticketingdecouple.ValueString()
	}

	return cloudngsparameter
}

func cloudngsparameterSetAttrFromGet(ctx context.Context, data *CloudngsparameterResourceModel, getResponseData map[string]interface{}) *CloudngsparameterResourceModel {
	tflog.Debug(ctx, "In cloudngsparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["allowdtls12"]; ok && val != nil {
		data.Allowdtls12 = types.StringValue(val.(string))
	} else {
		data.Allowdtls12 = types.StringNull()
	}
	if val, ok := getResponseData["allowedudtversion"]; ok && val != nil {
		data.Allowedudtversion = types.StringValue(val.(string))
	} else {
		data.Allowedudtversion = types.StringNull()
	}
	if val, ok := getResponseData["blockonallowedngstktprof"]; ok && val != nil {
		data.Blockonallowedngstktprof = types.StringValue(val.(string))
	} else {
		data.Blockonallowedngstktprof = types.StringNull()
	}
	if val, ok := getResponseData["csvserverticketingdecouple"]; ok && val != nil {
		data.Csvserverticketingdecouple = types.StringValue(val.(string))
	} else {
		data.Csvserverticketingdecouple = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("cloudngsparameter-config")

	return data
}