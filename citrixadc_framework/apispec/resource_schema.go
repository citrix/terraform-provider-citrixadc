package apispec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/api"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ApispecResourceModel describes the resource data model.
type ApispecResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Encrypted      types.Bool   `tfsdk:"encrypted"`
	File           types.String `tfsdk:"file"`
	Name           types.String `tfsdk:"name"`
	Skipvalidation types.String `tfsdk:"skipvalidation"`
	Type           types.String `tfsdk:"type"`
}

func (r *ApispecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the apispec resource.",
			},
			"encrypted": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Specify the encrypted API spec. Must be in NetScaler format",
			},
			"file": schema.StringAttribute{
				Required:    true,
				Description: "Name of and, optionally, path to the api spec file. The spec file should be present on the appliance's hard-disk drive or solid-state drive. Storing a spec file in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/apispec/ is the default path.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the spec. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the spec is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my spec\" or 'my spec').",
			},
			"skipvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disabling openapi spec validation while adding it",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the spec file. The three formats supported by the appliance are:\nPROTO \nOAS/Swagger\nGRAPHQL",
			},
		},
	}
}

// apispecGetThePayloadFromthePlan builds the add payload. NITRO accepts
// `encrypted` only on the add action, not on the change (update) action.
func apispecGetThePayloadFromthePlan(ctx context.Context, data *ApispecResourceModel) api.Apispec {
	tflog.Debug(ctx, "In apispecGetThePayloadFromthePlan Function")

	apispec := api.Apispec{}
	if !data.Encrypted.IsNull() && !data.Encrypted.IsUnknown() {
		apispec.Encrypted = data.Encrypted.ValueBool()
	}
	if !data.File.IsNull() && !data.File.IsUnknown() {
		apispec.File = data.File.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		apispec.Name = data.Name.ValueString()
	}
	if !data.Skipvalidation.IsNull() && !data.Skipvalidation.IsUnknown() {
		apispec.Skipvalidation = data.Skipvalidation.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		apispec.Type = data.Type.ValueString()
	}

	return apispec
}

// apispecGetTheUpdatePayloadFromthePlan builds the change/update payload.
// Per the NITRO doc, the change action does NOT accept `encrypted`. The
// struct field has `omitempty` so setting it to false drops it from the
// marshalled JSON.
func apispecGetTheUpdatePayloadFromthePlan(ctx context.Context, data *ApispecResourceModel) api.Apispec {
	tflog.Debug(ctx, "In apispecGetTheUpdatePayloadFromthePlan Function")

	apispec := apispecGetThePayloadFromthePlan(ctx, data)
	apispec.Encrypted = false
	return apispec
}

func apispecSetAttrFromGet(ctx context.Context, data *ApispecResourceModel, getResponseData map[string]interface{}) *ApispecResourceModel {
	tflog.Debug(ctx, "In apispecSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["encrypted"]; ok && val != nil {
		data.Encrypted = types.BoolValue(val.(bool))
	} else {
		data.Encrypted = types.BoolNull()
	}
	if val, ok := getResponseData["file"]; ok && val != nil {
		data.File = types.StringValue(val.(string))
	} else {
		data.File = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["skipvalidation"]; ok && val != nil {
		data.Skipvalidation = types.StringValue(val.(string))
	} else {
		data.Skipvalidation = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
