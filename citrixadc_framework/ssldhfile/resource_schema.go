package ssldhfile

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

// SsldhfileResourceModel describes the resource data model.
type SsldhfileResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Src  types.String `tfsdk:"src"`
}

func (r *SsldhfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssldhfile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported DH file.  Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL specifying the protocol, host, and path, including file name, to the DH file to be imported. For example, http://www.example.com/dh_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}

func ssldhfileGetThePayloadFromthePlan(ctx context.Context, data *SsldhfileResourceModel) ssl.Ssldhfile {
	tflog.Debug(ctx, "In ssldhfileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	ssldhfile := ssl.Ssldhfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		ssldhfile.Name = data.Name.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		ssldhfile.Src = data.Src.ValueString()
	}

	return ssldhfile
}

func ssldhfileSetAttrFromGet(ctx context.Context, data *SsldhfileResourceModel, getResponseData map[string]interface{}) *SsldhfileResourceModel {
	tflog.Debug(ctx, "In ssldhfileSetAttrFromGet Function")

	// Resource setter: preserves plan/state values. `name` is the key and `src`
	// is a write-only Import input that NITRO does not faithfully echo back, so
	// do not overwrite either from the GET response. ID is set once in Create.
	if data.Name.IsNull() || data.Name.IsUnknown() {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}

	return data
}

func ssldhfileSetAttrFromGetForDatasource(ctx context.Context, data *SsldhfileResourceModel, getResponseData map[string]interface{}) *SsldhfileResourceModel {
	tflog.Debug(ctx, "In ssldhfileSetAttrFromGetForDatasource Function")

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
