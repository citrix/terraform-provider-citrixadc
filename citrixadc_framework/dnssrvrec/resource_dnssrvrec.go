package dnssrvrec

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnssrvrecResource{}
var _ resource.ResourceWithConfigure = (*DnssrvrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnssrvrecResource)(nil)

func NewDnssrvrecResource() resource.Resource {
	return &DnssrvrecResource{}
}

// DnssrvrecResource defines the resource implementation.
type DnssrvrecResource struct {
	client *service.NitroClient
}

// legacyIdAttrOrder mirrors the SDK v2 comma-separated ID order (domain,target)
// recorded in resource_id_mapping.json so imported SDK v2 state parses correctly.
var legacyIdAttrOrder = []string{"domain", "target"}

func (r *DnssrvrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnssrvrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssrvrec"
}

func (r *DnssrvrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnssrvrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnssrvrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnssrvrec resource")

	dnssrvrec := dnssrvrecGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (POST /nitro/v1/config/dnssrvrec).
	// The name argument mirrors the SDK v2 "domain,target" composite.
	dnssrvrecName := fmt.Sprintf("%s,%s", data.Domain.ValueString(), data.Target.ValueString())
	_, err := r.client.AddResource(service.Dnssrvrec.Type(), dnssrvrecName, &dnssrvrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnssrvrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnssrvrec resource")

	// Set ID for the resource before reading state (SDK v2 "domain,target" format)
	data.Id = types.StringValue(dnssrvrecName)

	// Read the updated state back
	if !r.readDnssrvrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssrvrec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssrvrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnssrvrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnssrvrec resource")

	found := r.readDnssrvrecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnssrvrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnssrvrecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (domain/target are RequiresReplace so it is stable)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnssrvrec resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Ecssubnet.Equal(state.Ecssubnet) {
		tflog.Debug(ctx, "ecssubnet has changed for dnssrvrec")
		hasChange = true
	}
	if !data.Nodeid.Equal(state.Nodeid) {
		tflog.Debug(ctx, "nodeid has changed for dnssrvrec")
		hasChange = true
	}
	if !data.Port.Equal(state.Port) {
		tflog.Debug(ctx, "port has changed for dnssrvrec")
		hasChange = true
	}
	if !data.Priority.Equal(state.Priority) {
		tflog.Debug(ctx, "priority has changed for dnssrvrec")
		hasChange = true
	}
	if !data.Ttl.Equal(state.Ttl) {
		tflog.Debug(ctx, "ttl has changed for dnssrvrec")
		hasChange = true
	}
	if !data.Weight.Equal(state.Weight) {
		tflog.Debug(ctx, "weight has changed for dnssrvrec")
		hasChange = true
	}

	if hasChange {
		// domain and target are always included so NITRO can identify the record.
		dnssrvrec := dnssrvrecGetThePayloadFromthePlan(ctx, &data)
		// Update uses PUT /nitro/v1/config/dnssrvrec (UpdateUnnamedResource),
		// matching the SDK v2 implementation.
		err := r.client.UpdateUnnamedResource(service.Dnssrvrec.Type(), &dnssrvrec)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnssrvrec, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated dnssrvrec resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnssrvrec resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnssrvrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssrvrec not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssrvrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnssrvrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnssrvrec resource")

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), legacyIdAttrOrder, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse dnssrvrec ID, got error: %s", err))
		return
	}
	domain := idMap["domain"]
	target := idMap["target"]

	// DELETE /nitro/v1/config/dnssrvrec/<domain>?args=target:<target>,ecssubnet:<ecssubnet>
	// mirrors the SDK v2 delete (values are URL-encoded before being placed in the args map).
	argsMap := make(map[string]string)
	argsMap["target"] = url.QueryEscape(target)
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() && data.Ecssubnet.ValueString() != "" {
		argsMap["ecssubnet"] = url.QueryEscape(data.Ecssubnet.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Dnssrvrec.Type(), domain, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnssrvrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnssrvrec resource")
}

// Helper function to read dnssrvrec data from API. Returns false if the record
// no longer exists on the appliance (so Read can drop it from state).
func (r *DnssrvrecResource) readDnssrvrecFromApi(ctx context.Context, data *DnssrvrecResourceModel, diags *diag.Diagnostics) bool {
	// dnssrvrec exposes only a "get (all)" endpoint, so fetch the array and filter
	// by the domain+target identity (same approach as the SDK v2 resource).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), legacyIdAttrOrder, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse dnssrvrec ID, got error: %s", err))
		return false
	}
	domain := idMap["domain"]
	target := idMap["target"]

	findParams := service.FindParams{
		ResourceType: service.Dnssrvrec.Type(),
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnssrvrec, got error: %s", err))
		return false
	}

	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if d, ok := v["domain"].(string); !ok || d != domain {
			continue
		}
		if t, ok := v["target"].(string); !ok || t != target {
			continue
		}
		foundIndex = i
		break
	}
	if foundIndex == -1 {
		return false
	}

	dnssrvrecSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
