package nsmode

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsmodeResource{}
var _ resource.ResourceWithConfigure = (*NsmodeResource)(nil)
var _ resource.ResourceWithImportState = (*NsmodeResource)(nil)

func NewNsmodeResource() resource.Resource {
	return &NsmodeResource{}
}

// NsmodeResource defines the resource implementation.
type NsmodeResource struct {
	client *service.NitroClient
}

// nsmodeFeature is the payload used with the enable/disable NITRO actions.
// It mirrors the SDK v2 resource, which pushes the set of modes to toggle via
// ActOnResource("nsmode", {"mode": [...]}, "enable"|"disable").
type nsmodeFeature struct {
	Mode []string `json:"mode"`
}

func (r *NsmodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsmodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsmode"
}

func (r *NsmodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsmodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsmodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsmode resource")

	// Apply the configured modes via the enable/disable NITRO actions.
	if err := r.syncNsmode(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsmode, got error: %s", err))
		return
	}

	// nsmode is a singleton (no unique attributes) - static ID.
	data.Id = types.StringValue("nsmode-config")

	tflog.Trace(ctx, "Created nsmode resource")

	// Read the updated state back
	r.readNsmodeFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsmodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsmode resource")

	r.readNsmodeFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsmodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsmode resource")

	// Every mode attribute is RequiresReplaceIfConfigured (matching the SDK v2
	// ForceNew contract), so a configured change is handled via Delete+Create.
	// Update is only reachable when nothing that maps to an appliance action
	// changed; re-apply the modes defensively to keep state in sync.
	if err := r.syncNsmode(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsmode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nsmode resource")

	// Read the updated state back
	r.readNsmodeFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsmodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsmodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsmode resource")

	// nsmode modes cannot be "deleted"; the SDK v2 resource only removed the
	// object from Terraform state. Mirror that behaviour (state removal is
	// handled by the framework once Delete returns without error).
	tflog.Trace(ctx, "Deleted nsmode resource from state")
}

// syncNsmode toggles the configured modes on the appliance using the enable and
// disable NITRO actions, mirroring the SDK v2 syncNsmode. Only modes that are
// known and non-null (i.e. actually set in configuration, the framework
// equivalent of GetOkExists) are pushed; Computed/unconfigured modes are left
// untouched and are read back from the appliance.
func (r *NsmodeResource) syncNsmode(ctx context.Context, data *NsmodeResourceModel) error {
	tflog.Debug(ctx, "In syncNsmode Function")

	modeValues := []struct {
		name string
		val  types.Bool
	}{
		{"fr", data.Fr},
		{"l2", data.L2},
		{"usip", data.Usip},
		{"cka", data.Cka},
		{"tcpb", data.Tcpb},
		{"mbf", data.Mbf},
		{"edge", data.Edge},
		{"usnip", data.Usnip},
		{"l3", data.L3},
		{"pmtud", data.Pmtud},
		{"mediaclassification", data.Mediaclassification},
		{"sradv", data.Sradv},
		{"dradv", data.Dradv},
		{"iradv", data.Iradv},
		{"sradv6", data.Sradv6},
		{"dradv6", data.Dradv6},
		{"bridgebpdus", data.Bridgebpdus},
		{"ulfd", data.Ulfd},
	}

	enableList := make([]string, 0, len(modeValues))
	disableList := make([]string, 0, len(modeValues))

	for _, m := range modeValues {
		if m.val.IsNull() || m.val.IsUnknown() {
			continue
		}
		if m.val.ValueBool() {
			enableList = append(enableList, m.name)
		} else {
			disableList = append(disableList, m.name)
		}
	}

	if len(enableList) > 0 {
		if err := r.client.ActOnResource(service.Nsmode.Type(), &nsmodeFeature{Mode: enableList}, "enable"); err != nil {
			return err
		}
	}

	if len(disableList) > 0 {
		if err := r.client.ActOnResource(service.Nsmode.Type(), &nsmodeFeature{Mode: disableList}, "disable"); err != nil {
			return err
		}
	}

	return nil
}

// Helper function to read nsmode data from API
func (r *NsmodeResource) readNsmodeFromApi(ctx context.Context, data *NsmodeResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsmode.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsmode, got error: %s", err))
		return
	}

	nsmodeSetAttrFromGet(ctx, data, getResponseData)
}
