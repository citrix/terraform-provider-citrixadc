package locationfile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LocationfileResourceModel describes the resource data model.
type LocationfileResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Locationfile types.String `tfsdk:"locationfile"`
	Format       types.String `tfsdk:"format"`
	Src          types.String `tfsdk:"src"`
}

func (r *LocationfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the locationfile resource.",
			},
			"locationfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			"format": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("netscaler"),
				Description: "Format of the location file. Required for the NetScaler to identify how to read the location file.",
			},
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func locationfileGetThePayloadFromtheConfig(ctx context.Context, data *LocationfileResourceModel) basic.Locationfile {
	tflog.Debug(ctx, "In locationfileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	locationfile := basic.Locationfile{}
	if !data.Locationfile.IsNull() {
		locationfile.Locationfile = data.Locationfile.ValueString()
	}
	if !data.Format.IsNull() {
		locationfile.Format = data.Format.ValueString()
	}
	if !data.Src.IsNull() {
		locationfile.Src = data.Src.ValueString()
	}

	return locationfile
}

func locationfileSetAttrFromGet(ctx context.Context, data *LocationfileResourceModel, getResponseData map[string]interface{}) *LocationfileResourceModel {
	tflog.Debug(ctx, "In locationfileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Locationfile"]; ok && val != nil {
		data.Locationfile = types.StringValue(val.(string))
	} else {
		data.Locationfile = types.StringNull()
	}
	if val, ok := getResponseData["format"]; ok && val != nil {
		data.Format = types.StringValue(val.(string))
	} else {
		data.Format = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("locationfile-config")

	return data
}
