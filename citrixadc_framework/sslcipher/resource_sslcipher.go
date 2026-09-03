package sslcipher

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslcipherResource{}
var _ resource.ResourceWithConfigure = (*SslcipherResource)(nil)
var _ resource.ResourceWithImportState = (*SslcipherResource)(nil)

func NewSslcipherResource() resource.Resource {
	return &SslcipherResource{}
}

// SslcipherResource defines the resource implementation.
type SslcipherResource struct {
	client *service.NitroClient
}

func (r *SslcipherResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcipherResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcipher"
}

func (r *SslcipherResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcipherResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcipherResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcipher resource")

	ciphergroupname := data.Ciphergroupname.ValueString()

	// Named resource - create the empty cipher group first.
	sslcipher := ssl.Sslcipher{
		Ciphergroupname: ciphergroupname,
	}
	_, err := r.client.AddResource(service.Sslcipher.Type(), ciphergroupname, &sslcipher)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcipher, got error: %s", err))
		return
	}

	// Set ID for the resource (single unique attr -> plain value, SDK v2 compatible).
	data.Id = types.StringValue(ciphergroupname)

	// Ignore bindings unless there is an explicit configuration for it.
	if !data.Ciphersuitebinding.IsNull() && !data.Ciphersuitebinding.IsUnknown() && len(data.Ciphersuitebinding.Elements()) > 0 {
		bindings, d := expandCiphersuitebindings(ctx, data.Ciphersuitebinding)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		if err := updateSslcipherCiphersuiteBindings(r.client, ciphergroupname, bindings); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to bind ciphersuites to sslcipher, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Created sslcipher resource")

	// Confirm existence / refresh name. Preserve the planned ciphersuitebinding
	// value (bindings are pushed to the appliance; state mirrors config).
	if !r.readSslcipherFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslcipher not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcipherResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcipher resource")

	found := r.readSslcipherFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Refresh the ciphersuitebinding set from the appliance ONLY when the block is
	// explicitly configured (state value is non-null), so out-of-band changes to a
	// block this resource manages are detected. This mirrors the SDK v2 read, which
	// gated the refresh on d.GetOk("ciphersuitebinding"): when no block is configured
	// the resource ignores bindings entirely, so it does not adopt (and then plan to
	// remove) ciphers owned by a separate citrixadc_sslcipher_sslciphersuite_binding
	// resource on the same cipher group. The refresh is done ONLY here in Read, not in
	// the shared readSslcipherFromApi that Create/Update also call: overwriting the
	// planned bindings there could make the post-apply state differ from the plan
	// ("inconsistent result after apply"). Refresh has no such consistency check.
	if !data.Ciphersuitebinding.IsNull() {
		set, d := readSslcipherCiphersuiteBindings(ctx, r.client, data.Id.ValueString())
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		data.Ciphersuitebinding = set
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcipherResourceModel

	// Read prior state (for change detection and ID) and the plan.
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (ciphergroupname is ForceNew).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslcipher resource")

	ciphergroupname := data.Ciphergroupname.ValueString()

	// Only the ciphersuitebinding set is updateable; ciphergroupname is ForceNew.
	if !data.Ciphersuitebinding.Equal(state.Ciphersuitebinding) {
		tflog.Debug(ctx, fmt.Sprintf("ciphersuitebinding has changed for sslcipher %s, starting update", ciphergroupname))

		var bindings []CiphersuitebindingModel
		if !data.Ciphersuitebinding.IsNull() && !data.Ciphersuitebinding.IsUnknown() {
			var d diag.Diagnostics
			bindings, d = expandCiphersuitebindings(ctx, data.Ciphersuitebinding)
			resp.Diagnostics.Append(d...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
		if err := updateSslcipherCiphersuiteBindings(r.client, ciphergroupname, bindings); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcipher ciphersuite bindings, got error: %s", err))
			return
		}
	} else {
		tflog.Debug(ctx, "No changes detected for sslcipher resource, skipping update")
	}

	tflog.Trace(ctx, "Updated sslcipher resource")

	// Confirm existence / refresh name; preserve planned ciphersuitebinding value.
	if !r.readSslcipherFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslcipher not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcipherResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcipherResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcipher resource")

	// Named resource - delete using DeleteResource (ID is ciphergroupname).
	err := r.client.DeleteResource(service.Sslcipher.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcipher, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcipher resource")
}

// readSslcipherFromApi confirms the cipher group exists on the appliance and
// refreshes ciphergroupname + id. It intentionally does NOT overwrite the
// ciphersuitebinding set: Create/Update call this after pushing the bindings, and
// preserving the planned value there guarantees plan-consistency for the
// (non-Computed) set. Read refreshes the binding set separately (see Read) so
// out-of-band drift is still detected. Returns false if the cipher group no
// longer exists.
func (r *SslcipherResource) readSslcipherFromApi(ctx context.Context, data *SslcipherResourceModel, diags *diag.Diagnostics) bool {
	ciphergroupname := data.Id.ValueString()

	// Mirror SDK v2: some NetScaler versions do not support the per-name GET,
	// so use FindAllResources and filter by ciphergroupname.
	dataArr, err := r.client.FindAllResources(service.Sslcipher.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcipher, got error: %s", err))
		return false
	}

	found := false
	for _, v := range dataArr {
		if name, ok := v["ciphergroupname"].(string); ok && name == ciphergroupname {
			data.Ciphergroupname = types.StringValue(name)
			found = true
			break
		}
	}
	if !found {
		return false
	}

	data.Id = types.StringValue(ciphergroupname)
	return true
}

// expandCiphersuitebindings converts the ciphersuitebinding set into a slice of
// binding models.
func expandCiphersuitebindings(ctx context.Context, set types.Set) ([]CiphersuitebindingModel, diag.Diagnostics) {
	var bindings []CiphersuitebindingModel
	diags := set.ElementsAs(ctx, &bindings, false)
	return bindings, diags
}

// updateSslcipherCiphersuiteBindings mirrors the SDK v2 behavior: delete every
// existing binding on the cipher group, then (re-)add the configured bindings in
// ascending cipherpriority order. Re-creating all bindings is required because
// adding a ciphersuite with a lower priority than an existing one bumps the
// existing priority by one on the appliance.
func updateSslcipherCiphersuiteBindings(client *service.NitroClient, ciphergroupname string, bindings []CiphersuitebindingModel) error {
	// Fetch and delete all existing bindings.
	findParams := service.FindParams{
		ResourceType: service.Sslcipher_sslciphersuite_binding.Type(),
		ResourceName: ciphergroupname,
	}
	existing, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}
	for _, b := range existing {
		ciphername, _ := b["ciphername"].(string)
		args := []string{fmt.Sprintf("ciphername:%v", ciphername)}
		if err := client.DeleteResourceWithArgs(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname, args); err != nil {
			return err
		}
	}

	// Add all configured bindings, sorted by ascending priority.
	sorted := make([]CiphersuitebindingModel, len(bindings))
	copy(sorted, bindings)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Cipherpriority.ValueInt64() < sorted[j].Cipherpriority.ValueInt64()
	})

	for _, b := range sorted {
		bindingStruct := ssl.Sslcipherciphersuitebinding{
			Ciphergroupname: ciphergroupname,
			Ciphername:      b.Ciphername.ValueString(),
			Cipherpriority:  uint32(b.Cipherpriority.ValueInt64()),
		}
		// HTTP PUT (UpdateResource) as in SDK v2.
		if _, err := client.UpdateResource(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname, bindingStruct); err != nil {
			return err
		}
	}
	return nil
}

// readSslcipherCiphersuiteBindings reads the ciphersuite bindings for a cipher
// group from the appliance and returns them as a framework Set. Used by the
// datasource (which has no prior state to preserve).
func readSslcipherCiphersuiteBindings(ctx context.Context, client *service.NitroClient, ciphergroupname string) (types.Set, diag.Diagnostics) {
	bindings, _ := client.FindResourceArray(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname)

	elems := make([]CiphersuitebindingModel, 0, len(bindings))
	for _, b := range bindings {
		ciphername, _ := b["ciphername"].(string)
		priority := int64(0)
		if p, ok := b["cipherpriority"].(string); ok {
			if iv, err := strconv.Atoi(p); err == nil {
				priority = int64(iv)
			}
		}
		elems = append(elems, CiphersuitebindingModel{
			Ciphername:     types.StringValue(ciphername),
			Cipherpriority: types.Int64Value(priority),
		})
	}

	return types.SetValueFrom(ctx, ciphersuitebindingObjectType(), elems)
}
