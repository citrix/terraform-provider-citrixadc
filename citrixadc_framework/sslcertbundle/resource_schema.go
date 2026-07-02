package sslcertbundle

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcertbundleResourceModel describes the resource data model.
type SslcertbundleResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *SslcertbundleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertbundle resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported certificate bundle. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the certificate bundle to be imported or exported. For example, http://www.example.com/cert_bundle_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func sslcertbundleGetThePayloadFromthePlan(ctx context.Context, data *SslcertbundleResourceModel) ssl.Sslcertbundle {
	tflog.Debug(ctx, "In sslcertbundleGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcertbundle := ssl.Sslcertbundle{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslcertbundle.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		sslcertbundle.Src = data.Src.ValueString()
	}

	return sslcertbundle
}

func sslcertbundleSetAttrFromGet(ctx context.Context, data *SslcertbundleResourceModel, getResponseData map[string]interface{}) *SslcertbundleResourceModel {
	tflog.Debug(ctx, "In sslcertbundleSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// src is an Import-only input; NITRO GET does not echo it back. Preserve plan/state value (Pattern 7).

	// ID is set once in Create; do not recompute here (Pattern 6).

	return data
}

func sslcertbundleSetAttrFromGetForDatasource(ctx context.Context, data *SslcertbundleResourceModel, getResponseData map[string]interface{}) *SslcertbundleResourceModel {
	tflog.Debug(ctx, "In sslcertbundleSetAttrFromGetForDatasource Function")

	// Convert API response to model
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

	// Datasource has no Create; set ID here.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
