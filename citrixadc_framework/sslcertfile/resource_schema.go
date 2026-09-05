package sslcertfile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcertfileResourceModel describes the resource data model.
type SslcertfileResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *SslcertfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertfile resource.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported certificate file. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the certificate file to be imported. For example, http://www.example.com/cert_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func sslcertfileGetThePayloadFromthePlan(ctx context.Context, data *SslcertfileResourceModel) ssl.Sslcertfile {
	tflog.Debug(ctx, "In sslcertfileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcertfile := ssl.Sslcertfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslcertfile.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		sslcertfile.Src = data.Src.ValueString()
	}

	return sslcertfile
}

// sslcertfileSetAttrFromGet is the RESOURCE-side setter. It preserves plan/state
// values: `name` is the key and `src` is a write-only Import input that NITRO
// does not echo back faithfully (the server strips the "local://" scheme prefix,
// e.g. config "local://certificate1.crt" is returned as "certificate1.crt").
// Overwriting `src` from the GET response would trigger an "inconsistent result
// after apply" error (Pattern 7 normalized-form). So do not touch `src`, and only
// adopt `name` from GET when it is unset (import case). ID is set in Create /
// preserved in Read, never here.
func sslcertfileSetAttrFromGet(ctx context.Context, data *SslcertfileResourceModel, getResponseData map[string]interface{}) *SslcertfileResourceModel {
	tflog.Debug(ctx, "In sslcertfileSetAttrFromGet Function")

	if data.Name.IsNull() || data.Name.IsUnknown() {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}

	return data
}

// sslcertfileSetAttrFromGetForDatasource is the DATASOURCE-side setter. A
// datasource has no prior plan/state to preserve, so it faithfully copies every
// field from the GET response and sets the ID (single unique key = name).
func sslcertfileSetAttrFromGetForDatasource(ctx context.Context, data *SslcertfileResourceModel, getResponseData map[string]interface{}) *SslcertfileResourceModel {
	tflog.Debug(ctx, "In sslcertfileSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (single unique attribute = name).
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
