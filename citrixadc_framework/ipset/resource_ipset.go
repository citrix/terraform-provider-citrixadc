package ipset

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IpsetResource{}
var _ resource.ResourceWithConfigure = (*IpsetResource)(nil)
var _ resource.ResourceWithImportState = (*IpsetResource)(nil)

func NewIpsetResource() resource.Resource {
	return &IpsetResource{}
}

// IpsetResource defines the resource implementation.
type IpsetResource struct {
	client *service.NitroClient
}

func (r *IpsetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IpsetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipset"
}

func (r *IpsetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IpsetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpsetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ipset resource")

	ipsetName := data.Name.ValueString()

	// Named resource - use AddResource
	ipset := ipsetGetThePayloadFromtheConfig(ctx, &data)
	_, err := r.client.AddResource(service.Ipset.Type(), ipsetName, &ipset)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ipset, got error: %s", err))
		return
	}

	// Set ID for the resource before applying bindings / reading state
	data.Id = types.StringValue(ipsetName)

	// Apply the nsip / nsip6 binding convenience blocks (add all configured
	// bindings; on create there is no prior state so old sets are empty).
	if err := updateIpsetNsipBindings(ctx, r.client, ipsetName, nil, setToStringSlice(ctx, data.Nsipbinding)); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to configure ipset nsip bindings, got error: %s", err))
		return
	}
	if err := updateIpsetNsip6Bindings(ctx, r.client, ipsetName, nil, setToStringSlice(ctx, data.Nsip6binding)); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to configure ipset nsip6 bindings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ipset resource")

	// Read the updated state back
	if !r.readIpsetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ipset not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpsetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ipset resource")

	found := r.readIpsetFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state IpsetResourceModel

	// Read Terraform prior state to compute binding differences and preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is ForceNew so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating ipset resource")

	ipsetName := data.Id.ValueString()

	// The only updateable surface for ipset is the binding convenience blocks;
	// name and td are ForceNew and never reach Update.
	if err := updateIpsetNsipBindings(ctx, r.client, ipsetName, setToStringSlice(ctx, state.Nsipbinding), setToStringSlice(ctx, data.Nsipbinding)); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ipset nsip bindings, got error: %s", err))
		return
	}
	if err := updateIpsetNsip6Bindings(ctx, r.client, ipsetName, setToStringSlice(ctx, state.Nsip6binding), setToStringSlice(ctx, data.Nsip6binding)); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ipset nsip6 bindings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated ipset resource")

	// Read the updated state back
	if !r.readIpsetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ipset not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpsetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ipset resource")

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Ipset.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ipset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted ipset resource")
}

// Helper function to read ipset data from API
func (r *IpsetResource) readIpsetFromApi(ctx context.Context, data *IpsetResourceModel, diags *diag.Diagnostics) bool {
	ipsetName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Ipset.Type(), ipsetName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ipset, got error: %s", err))
		return false
	}

	ipsetSetAttrFromGet(ctx, data, getResponseData)

	// Populate the binding convenience blocks from the ADC
	readIpsetNsipBindings(ctx, r.client, data)
	readIpsetNsip6Bindings(ctx, r.client, data)

	return true
}

// -------------------------------------------------------------------------
// nsip / nsip6 binding convenience-block helpers (ported from SDK v2)
// -------------------------------------------------------------------------

func addSingleIpsetNsipBinding(client *service.NitroClient, name string, nsip string) error {
	binding := network.Ipsetnsipbinding{
		Name:      name,
		Ipaddress: nsip,
	}
	// HTTP PUT semantics for a binding => UpdateResource
	_, err := client.UpdateResource(service.Ipset_nsip_binding.Type(), name, binding)
	return err
}

func deleteSingleIpsetNsipBinding(client *service.NitroClient, name string, nsip string) error {
	args := []string{fmt.Sprintf("ipaddress:%s", nsip)}
	return client.DeleteResourceWithArgs(service.Ipset_nsip_binding.Type(), name, args)
}

func updateIpsetNsipBindings(ctx context.Context, client *service.NitroClient, name string, oldList, newList []string) error {
	toAdd, toRemove := stringSliceDiff(oldList, newList)
	for _, nsip := range toAdd {
		if err := addSingleIpsetNsipBinding(client, name, nsip); err != nil {
			return err
		}
	}
	for _, nsip := range toRemove {
		if err := deleteSingleIpsetNsipBinding(client, name, nsip); err != nil {
			return err
		}
	}
	return nil
}

func readIpsetNsipBindings(ctx context.Context, client *service.NitroClient, data *IpsetResourceModel) {
	bindings, _ := client.FindResourceArray(service.Ipset_nsip_binding.Type(), data.Id.ValueString())
	vals := make([]string, 0, len(bindings))
	for _, b := range bindings {
		if ip, ok := b["ipaddress"]; ok && ip != nil {
			vals = append(vals, ip.(string))
		}
	}
	data.Nsipbinding = stringSliceToSet(vals)
}

func addSingleIpsetNsip6Binding(client *service.NitroClient, name string, nsip6 string) error {
	binding := network.Ipsetnsip6binding{
		Name:      name,
		Ipaddress: nsip6,
	}
	// HTTP PUT semantics for a binding => UpdateResource
	_, err := client.UpdateResource(service.Ipset_nsip6_binding.Type(), name, binding)
	return err
}

func deleteSingleIpsetNsip6Binding(client *service.NitroClient, name string, nsip6 string) error {
	// IPv6 addresses contain ':' and '/', which must be url-escaped in the
	// delete arg (matches SDK v2 behaviour).
	args := []string{fmt.Sprintf("ipaddress:%s", url.QueryEscape(nsip6))}
	return client.DeleteResourceWithArgs(service.Ipset_nsip6_binding.Type(), name, args)
}

func updateIpsetNsip6Bindings(ctx context.Context, client *service.NitroClient, name string, oldList, newList []string) error {
	toAdd, toRemove := stringSliceDiff(oldList, newList)
	for _, nsip6 := range toAdd {
		if err := addSingleIpsetNsip6Binding(client, name, nsip6); err != nil {
			return err
		}
	}
	for _, nsip6 := range toRemove {
		if err := deleteSingleIpsetNsip6Binding(client, name, nsip6); err != nil {
			return err
		}
	}
	return nil
}

func readIpsetNsip6Bindings(ctx context.Context, client *service.NitroClient, data *IpsetResourceModel) {
	bindings, _ := client.FindResourceArray(service.Ipset_nsip6_binding.Type(), data.Id.ValueString())
	vals := make([]string, 0, len(bindings))
	for _, b := range bindings {
		if ip, ok := b["ipaddress"]; ok && ip != nil {
			vals = append(vals, ip.(string))
		}
	}
	data.Nsip6binding = stringSliceToSet(vals)
}

// -------------------------------------------------------------------------
// small set/slice utilities
// -------------------------------------------------------------------------

func setToStringSlice(ctx context.Context, s types.Set) []string {
	if s.IsNull() || s.IsUnknown() {
		return nil
	}
	var out []string
	s.ElementsAs(ctx, &out, false)
	return out
}

func stringSliceToSet(vals []string) types.Set {
	elems := make([]attr.Value, 0, len(vals))
	for _, v := range vals {
		elems = append(elems, types.StringValue(v))
	}
	return types.SetValueMust(types.StringType, elems)
}

func stringSliceDiff(oldList, newList []string) (toAdd, toRemove []string) {
	oldSet := make(map[string]bool, len(oldList))
	newSet := make(map[string]bool, len(newList))
	for _, v := range oldList {
		oldSet[v] = true
	}
	for _, v := range newList {
		newSet[v] = true
	}
	for _, v := range newList {
		if !oldSet[v] {
			toAdd = append(toAdd, v)
		}
	}
	for _, v := range oldList {
		if !newSet[v] {
			toRemove = append(toRemove, v)
		}
	}
	return toAdd, toRemove
}
