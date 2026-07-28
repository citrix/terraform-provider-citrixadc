package sslcacertbundle

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

// SslcacertbundleResourceModel describes the resource data model.
type SslcacertbundleResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Bundlefile       types.String `tfsdk:"bundlefile"`
	Cacertbundlename types.String `tfsdk:"cacertbundlename"`
}

func (r *SslcacertbundleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcacertbundle resource.",
			},
			"bundlefile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the X509 CA certificate bundle file that is used to form cacertbundle entity. The CA certificate bundle file should be present on the appliance's hard-disk drive or solid-state drive. /nsconfig/ssl/ is the default path. The CA certificate bundle file consists of list of certificates.",
			},
			"cacertbundlename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name given to the CA certbundle. The name will be used for bind/unbind/update operations. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
		},
	}
}

func sslcacertbundleGetThePayloadFromthePlan(ctx context.Context, data *SslcacertbundleResourceModel) ssl.Sslcacertbundle {
	tflog.Debug(ctx, "In sslcacertbundleGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcacertbundle := ssl.Sslcacertbundle{}
	if !data.Bundlefile.IsNull() && !data.Bundlefile.IsUnknown() {
		sslcacertbundle.Bundlefile = data.Bundlefile.ValueString()
	}
	if !data.Cacertbundlename.IsNull() && !data.Cacertbundlename.IsUnknown() {
		sslcacertbundle.Cacertbundlename = data.Cacertbundlename.ValueString()
	}

	return sslcacertbundle
}

func sslcacertbundleSetAttrFromGet(ctx context.Context, data *SslcacertbundleResourceModel, getResponseData map[string]interface{}) *SslcacertbundleResourceModel {
	tflog.Debug(ctx, "In sslcacertbundleSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bundlefile"]; ok && val != nil {
		data.Bundlefile = types.StringValue(val.(string))
	} else {
		data.Bundlefile = types.StringNull()
	}
	if val, ok := getResponseData["cacertbundlename"]; ok && val != nil {
		data.Cacertbundlename = types.StringValue(val.(string))
	} else {
		data.Cacertbundlename = types.StringNull()
	}

	// ID is set once in Create; do not recompute here (Pattern 6).

	return data
}

func sslcacertbundleSetAttrFromGetForDatasource(ctx context.Context, data *SslcacertbundleResourceModel, getResponseData map[string]interface{}) *SslcacertbundleResourceModel {
	tflog.Debug(ctx, "In sslcacertbundleSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["bundlefile"]; ok && val != nil {
		data.Bundlefile = types.StringValue(val.(string))
	} else {
		data.Bundlefile = types.StringNull()
	}
	if val, ok := getResponseData["cacertbundlename"]; ok && val != nil {
		data.Cacertbundlename = types.StringValue(val.(string))
	} else {
		data.Cacertbundlename = types.StringNull()
	}

	// Datasource has no Create; set ID here.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Cacertbundlename.ValueString()))

	return data
}
