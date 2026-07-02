package sslcrlfile

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

// SslcrlfileResourceModel describes the resource data model.
type SslcrlfileResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *SslcrlfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcrlfile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported CRL file. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name to the CRL file to be imported. For example, http://www.example.com/crl_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func sslcrlfileGetThePayloadFromthePlan(ctx context.Context, data *SslcrlfileResourceModel) ssl.Sslcrlfile {
	tflog.Debug(ctx, "In sslcrlfileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcrlfile := ssl.Sslcrlfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslcrlfile.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		sslcrlfile.Src = data.Src.ValueString()
	}

	return sslcrlfile
}

func sslcrlfileSetAttrFromGet(ctx context.Context, data *SslcrlfileResourceModel, getResponseData map[string]interface{}) *SslcrlfileResourceModel {
	tflog.Debug(ctx, "In sslcrlfileSetAttrFromGet Function")

	// Resource setter: preserves plan/state values. `name` is the key and `src`
	// is a write-only Import input that NITRO does not faithfully echo back
	// (it may be absent or normalized), so do not overwrite either from the GET
	// response. ID is set once in Create / preserved across Read.
	if data.Name.IsNull() || data.Name.IsUnknown() {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}

	return data
}

func sslcrlfileSetAttrFromGetForDatasource(ctx context.Context, data *SslcrlfileResourceModel, getResponseData map[string]interface{}) *SslcrlfileResourceModel {
	tflog.Debug(ctx, "In sslcrlfileSetAttrFromGetForDatasource Function")

	// Datasource setter: faithfully copy the GET response (no prior state).
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

	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
