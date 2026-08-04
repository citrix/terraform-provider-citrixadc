package appqoecustomresp

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppqoecustomrespResourceModel describes the resource data model.
type AppqoecustomrespResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *AppqoecustomrespResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appqoecustomresp resource.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates name of the custom response HTML page to import/update.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			// NOTE: not Computed. The NITRO GET returns src stripped of the
			// "local://" scheme prefix used at import time, so src is preserved
			// from config on read (see appqoecustomrespSetAttrFromGet) to avoid
			// an "inconsistent result after apply" on this Required attribute.
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Local file name of the custom response HTML page to import (e.g. local://<file>).",
			},
		},
	}
}

func appqoecustomrespGetThePayloadFromtheConfig(ctx context.Context, data *AppqoecustomrespResourceModel) appqoe.Appqoecustomresp {
	tflog.Debug(ctx, "In appqoecustomrespGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appqoecustomresp := appqoe.Appqoecustomresp{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appqoecustomresp.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appqoecustomresp.Src = data.Src.ValueString()
	}

	return appqoecustomresp
}

// appqoecustomrespSetAttrFromGet is the RESOURCE mapping of the NITRO GET
// response into the model. It intentionally does NOT overwrite src: the NITRO
// GET returns src without the "local://" scheme prefix supplied at import time
// (config "local://foo.html" -> GET "foo.html"), and src is a Required (not
// Computed) attribute, so overwriting it would trigger an "inconsistent result
// after apply". This mirrors the SDK v2 resource, which deliberately did not
// d.Set("src", ...) on read.
func appqoecustomrespSetAttrFromGet(ctx context.Context, data *AppqoecustomrespResourceModel, getResponseData map[string]interface{}) *AppqoecustomrespResourceModel {
	tflog.Debug(ctx, "In appqoecustomrespSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// src is intentionally NOT set from the API response (preserved from config/state).

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// appqoecustomrespSetAttrFromGetForDatasource is the DATASOURCE mapping. Unlike
// the resource variant, it copies every attribute from the NITRO GET response
// (including src, as returned by the API) and sets data.Id so the datasource
// exposes the live server-side values.
func appqoecustomrespSetAttrFromGetForDatasource(ctx context.Context, data *AppqoecustomrespResourceModel, getResponseData map[string]interface{}) *AppqoecustomrespResourceModel {
	tflog.Debug(ctx, "In appqoecustomrespSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}

	// Set ID for the datasource
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
