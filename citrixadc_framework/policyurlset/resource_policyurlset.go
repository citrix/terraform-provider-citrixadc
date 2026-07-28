package policyurlset

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicyurlsetResource{}
var _ resource.ResourceWithConfigure = (*PolicyurlsetResource)(nil)
var _ resource.ResourceWithImportState = (*PolicyurlsetResource)(nil)
var _ resource.ResourceWithValidateConfig = (*PolicyurlsetResource)(nil)

func NewPolicyurlsetResource() resource.Resource {
	return &PolicyurlsetResource{}
}

// PolicyurlsetResource defines the resource implementation.
type PolicyurlsetResource struct {
	client *service.NitroClient
}

func (r *PolicyurlsetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicyurlsetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyurlset"
}

func (r *PolicyurlsetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the mandatory secret attribute (Pattern 17): url is
// x-secret-attr + required, expanded into the url / url_wo / url_wo_version
// triple whose value attributes are both Optional. At least one must be set.
func (r *PolicyurlsetResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data PolicyurlsetResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Url.IsNull() && data.UrlWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Missing Required Attribute",
			"Either \"url\" or \"url_wo\" must be specified.",
		)
	}
}

func (r *PolicyurlsetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config PolicyurlsetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policyurlset resource")
	// Get payload from plan (regular attributes)
	policyurlset := policyurlsetGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	policyurlsetGetThePayloadFromtheConfig(ctx, &config, &policyurlset)

	// Make API call
	// Create is the NITRO "Import" action: POST /policyurlset?action=Import
	// (capital "I" — NITRO action names are case-sensitive). Not AddResource.
	err := r.client.ActOnResource(service.Policyurlset.Type(), &policyurlset, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policyurlset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policyurlset resource")

	// Set ID for the resource before reading state (single key: name)
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	r.readPolicyurlsetFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicyurlsetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policyurlset resource")

	r.readPolicyurlsetFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// If the resource was deleted out-of-band, remove it from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PolicyurlsetResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for policyurlset; all configurable attributes are
	// RequiresReplace (the NITRO ?action=update path is not used by this
	// import-as-create resource). Preserve ID and refresh from API.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for policyurlset; all attributes are RequiresReplace")

	r.readPolicyurlsetFromApi(ctx, &data, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicyurlsetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policyurlset resource")
	// Plain DELETE /policyurlset/<name> — no ?action, no args.
	// NITRO quirk: this DELETE actually removes the imported urlset but reports a
	// spurious errorcode 258 "No such resource [name, ...]". Treat that specific
	// response as a successful delete (the vendored DeleteResource also special-
	// cases policyurlset to skip the GET-by-name pre-check and suppress this).
	err := r.client.DeleteResource(service.Policyurlset.Type(), data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "No such resource") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policyurlset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policyurlset resource")
}

// Helper function to read policyurlset data from API.
func (r *PolicyurlsetResource) readPolicyurlsetFromApi(ctx context.Context, data *PolicyurlsetResourceModel, diags *diag.Diagnostics) {
	// An imported urlset is NOT retrievable via a plain GET /policyurlset/<name>
	// (that returns errorcode 258 "No such resource"). It is only listed by the
	// filtered GET /policyurlset?args=imported:true (Pattern 15: "imported" is a
	// GET-only filter param). Fetch that list and match by name.
	policyurlsetName := data.Id.ValueString()

	getResponseData, err := findImportedPolicyurlsetByName(r.client, policyurlsetName)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policyurlset, got error: %s", err))
		return
	}

	// Resource is missing
	if getResponseData == nil {
		tflog.Warn(ctx, fmt.Sprintf("policyurlset %s not found, removing from state", policyurlsetName))
		data.Id = types.StringNull()
		return
	}

	policyurlsetSetAttrFromGet(ctx, data, getResponseData)

	// Category (a) identity backfill: the composite ID is the single key `name`.
	// On import there is no prior plan/state, so recover `name` directly from the
	// parsed ID. This is done AFTER the found/not-found self-heal above (so a
	// null-Id on not-found is preserved) and after SetAttrFromGet, and it always
	// yields the SAME ID (Id == name), keeping the datasource path unaffected.
	if !data.Id.IsNull() {
		data.Name = types.StringValue(policyurlsetName)
	}
}

// findImportedPolicyurlsetByName lists imported urlsets
// (GET /policyurlset?args=imported:true) and returns the entry whose "name"
// matches the supplied name, or (nil, nil) if no such entry exists.
func findImportedPolicyurlsetByName(client *service.NitroClient, name string) (map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType: service.Policyurlset.Type(),
		ArgsMap:      map[string]string{"imported": "true"},
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}
	for _, item := range dataArr {
		if val, ok := item["name"]; ok && val != nil {
			if val.(string) == name {
				return item, nil
			}
		}
	}
	return nil, nil
}
