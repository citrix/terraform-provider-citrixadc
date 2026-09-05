package locationfile6

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			// SDK v2: locationfile is Required + ForceNew.
			"locationfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the IPv6 location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			// SDK v2: format is Optional + ForceNew (no Default). Made Optional+Computed
			// so the value can be read back from the ADC; because it was ForceNew, use
			// UseStateForUnknown (avoid known-after-apply churn) + RequiresReplaceIfConfigured
			// (replace only when the user changes a configured value, matching ForceNew).
			"format": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Format of the IPv6 location file. Required for the NetScaler to identify how to read the location file.",
			},
			// SDK v2: src is Optional + ForceNew. Same treatment as format.
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

// locationfile6GetThePayloadFromtheConfig builds the body for the add operation.
// Mirroring SDK v2 createLocationfile6Func, only Locationfile and Format are sent
// (src is used by the separate import action, not the add operation).
func locationfile6GetThePayloadFromtheConfig(ctx context.Context, data *Locationfile6ResourceModel) basic.Locationfile6 {
	tflog.Debug(ctx, "In locationfile6GetThePayloadFromtheConfig Function")

	locationfile6 := basic.Locationfile6{}
	if !data.Locationfile.IsNull() && !data.Locationfile.IsUnknown() {
		locationfile6.Locationfile = data.Locationfile.ValueString()
	}
	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		locationfile6.Format = data.Format.ValueString()
	}

	return locationfile6
}

// locationfile6SetAttrFromGet maps a GET response onto the resource model. The
// else-branches only null a value when it is Unknown, never clobbering a known
// (configured/state) value that NITRO omits from the GET response.
func locationfile6SetAttrFromGet(ctx context.Context, data *Locationfile6ResourceModel, getResponseData map[string]interface{}) *Locationfile6ResourceModel {
	tflog.Debug(ctx, "In locationfile6SetAttrFromGet Function")

	if val, ok := getResponseData["Locationfile"]; ok && val != nil {
		data.Locationfile = types.StringValue(val.(string))
	} else if data.Locationfile.IsUnknown() {
		data.Locationfile = types.StringNull()
	}
	if val, ok := getResponseData["format"]; ok && val != nil {
		data.Format = types.StringValue(val.(string))
	} else if data.Format.IsUnknown() {
		data.Format = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else if data.Src.IsUnknown() {
		data.Src = types.StringNull()
	}

	// Backward-compatible ID: SDK v2 used d.SetId(locationfile). Keep the ID
	// aligned with the locationfile value.
	if !data.Locationfile.IsNull() && !data.Locationfile.IsUnknown() {
		data.Id = types.StringValue(data.Locationfile.ValueString())
	}

	return data
}

// locationfile6SetAttrFromGetForDatasource copies every attribute from the GET
// response into the model and sets the datasource ID.
func locationfile6SetAttrFromGetForDatasource(ctx context.Context, data *Locationfile6ResourceModel, getResponseData map[string]interface{}) *Locationfile6ResourceModel {
	tflog.Debug(ctx, "In locationfile6SetAttrFromGetForDatasource Function")

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

	if !data.Locationfile.IsNull() {
		data.Id = types.StringValue(data.Locationfile.ValueString())
	} else {
		data.Id = types.StringValue("locationfile6-config")
	}

	return data
}
