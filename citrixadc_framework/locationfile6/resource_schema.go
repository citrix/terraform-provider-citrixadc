package locationfile6

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

// Locationfile6ResourceModel describes the resource data model.
type Locationfile6ResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Locationfile types.String `tfsdk:"locationfile"`
	Format       types.String `tfsdk:"format"`
	Src          types.String `tfsdk:"src"`
}

func (r *Locationfile6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the locationfile6 resource.",
			},
			"locationfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the IPv6 location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			"format": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("netscaler6"),
				Description: "Format of the IPv6 location file. Required for the NetScaler to identify how to read the location file.",
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

func locationfile6GetThePayloadFromtheConfig(ctx context.Context, data *Locationfile6ResourceModel) basic.Locationfile6 {
	tflog.Debug(ctx, "In locationfile6GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	locationfile6 := basic.Locationfile6{}
	if !data.Locationfile.IsNull() {
		locationfile6.Locationfile = data.Locationfile.ValueString()
	}
	if !data.Format.IsNull() {
		locationfile6.Format = data.Format.ValueString()
	}
	if !data.Src.IsNull() {
		locationfile6.Src = data.Src.ValueString()
	}

	return locationfile6
}

func locationfile6SetAttrFromGet(ctx context.Context, data *Locationfile6ResourceModel, getResponseData map[string]interface{}) *Locationfile6ResourceModel {
	tflog.Debug(ctx, "In locationfile6SetAttrFromGet Function")

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
	data.Id = types.StringValue("locationfile6-config")

	return data
}
