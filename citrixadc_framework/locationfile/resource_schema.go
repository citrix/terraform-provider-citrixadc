package locationfile

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
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			"locationfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			// SDK v2: Optional + ForceNew (NO Default). Made Optional+Computed so the
			// ADC-side default ("netscaler") is read back without a perpetual diff.
			// A Default is invalid without Computed, and SDK v2 had none — so no Default.
			// UseStateForUnknown() keeps the computed value stable across plans;
			// RequiresReplaceIfConfigured() reproduces ForceNew only when the user
			// actually configures the attribute (avoids computed-churn replacements).
			"format": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Format of the location file. Required for the NetScaler to identify how to read the location file.",
			},
			// SDK v2: Optional + ForceNew (NOT Computed). The NITRO add/get operations
			// do not carry src (only the separate Import action does), so src is never
			// returned by GET — keep it Optional-only + RequiresReplace, matching SDK v2.
			"src": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

// locationfileGetThePayloadFromthePlan builds the body for the `add` operation.
// Mirroring SDK v2 (and the NITRO `add` doc, whose payload lists only Locationfile
// + format), ONLY Locationfile and format are sent; src is excluded from the add
// payload (src is used only by the separate Import action).
func locationfileGetThePayloadFromthePlan(ctx context.Context, data *LocationfileResourceModel) basic.Locationfile {
	tflog.Debug(ctx, "In locationfileGetThePayloadFromthePlan Function")

	locationfile := basic.Locationfile{}
	if !data.Locationfile.IsNull() && !data.Locationfile.IsUnknown() {
		locationfile.Locationfile = data.Locationfile.ValueString()
	}
	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		locationfile.Format = data.Format.ValueString()
	}

	return locationfile
}

// locationfileSetAttrFromGet is the RESOURCE state setter. It preserves the
// user-configured identity/inputs and only reads back computed values, so it never
// clobbers a configured value with something the ADC omits from GET.
func locationfileSetAttrFromGet(ctx context.Context, data *LocationfileResourceModel, getResponseData map[string]interface{}) *LocationfileResourceModel {
	tflog.Debug(ctx, "In locationfileSetAttrFromGet Function")

	// locationfile is Required + RequiresReplace. Adopt the GET value only when the
	// model has no value yet (import, where state carries only the ID); otherwise
	// preserve the configured value to avoid clobbering it / inconsistent-result errors.
	if data.Locationfile.IsNull() || data.Locationfile.IsUnknown() || data.Locationfile.ValueString() == "" {
		if val, ok := getResponseData["Locationfile"]; ok && val != nil {
			data.Locationfile = types.StringValue(val.(string))
		}
	}

	// format is Optional+Computed: read the live value from the ADC. Guard the
	// else-branch so a value the ADC omits from GET is only nulled when unknown,
	// never clobbering a known configured value (omit-on-default trap).
	if val, ok := getResponseData["format"]; ok && val != nil {
		data.Format = types.StringValue(val.(string))
	} else if data.Format.IsUnknown() {
		data.Format = types.StringNull()
	}

	// src is Optional-only and never returned by GET (the add/get ops do not carry
	// it). Preserve the configured/state value; do NOT clobber it.

	// Id is set in Create/Read (SDK v2 ID = locationfile name); do not overwrite here.
	return data
}

// locationfileSetAttrFromGetForDatasource is the DATASOURCE state setter. Unlike
// the resource setter, it copies every value straight from the GET response and
// assigns the datasource ID.
func locationfileSetAttrFromGetForDatasource(ctx context.Context, data *LocationfileResourceModel, getResponseData map[string]interface{}) *LocationfileResourceModel {
	tflog.Debug(ctx, "In locationfileSetAttrFromGetForDatasource Function")

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
	// src is not returned by GET.
	data.Src = types.StringNull()

	// Datasource ID mirrors the SDK v2 resource ID scheme (the location file name),
	// falling back to a static handle if the ADC returns no name.
	if val, ok := getResponseData["Locationfile"]; ok && val != nil {
		data.Id = types.StringValue(val.(string))
	} else {
		data.Id = types.StringValue("locationfile-config")
	}
	return data
}
