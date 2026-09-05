package dnstxtrec

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
var _ resource.Resource = &DnstxtrecResource{}
var _ resource.ResourceWithConfigure = (*DnstxtrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnstxtrecResource)(nil)

func NewDnstxtrecResource() resource.Resource {
	return &DnstxtrecResource{}
}

// DnstxtrecResource defines the resource implementation.
type DnstxtrecResource struct {
	client *service.NitroClient
}

func (r *DnstxtrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnstxtrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnstxtrec"
}

func (r *DnstxtrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnstxtrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnstxtrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnstxtrec resource")

	dnstxtrec := dnstxtrecGetThePayloadFromthePlan(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Named resource - use AddResource
	domainValue := data.Domain.ValueString()
	_, err := r.client.AddResource(service.Dnstxtrec.Type(), domainValue, &dnstxtrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnstxtrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnstxtrec resource")

	// Set ID for the resource before reading state (single unique attribute: domain)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Domain.ValueString()))

	// Read the updated state back
	if !r.readDnstxtrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnstxtrec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnstxtrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnstxtrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnstxtrec resource")

	found := r.readDnstxtrecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnstxtrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnstxtrecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnstxtrec resource")

	// dnstxtrec exposes no NITRO "update" verb and every configurable attribute is
	// RequiresReplace, so any attribute change triggers a destroy/recreate rather
	// than reaching this path. This method therefore performs no API write and only
	// refreshes state (belt-and-suspenders for computed-only diffs).
	if !r.readDnstxtrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnstxtrec not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnstxtrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnstxtrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnstxtrec resource")

	dnstxtrecName := data.Id.ValueString()

	// The NITRO delete requires the server-generated recordid as a query argument,
	// so fetch the current record first (mirrors the SDK v2 delete flow).
	getResponseData, err := r.client.FindResource(service.Dnstxtrec.Type(), dnstxtrecName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Already gone - nothing to delete.
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnstxtrec before delete, got error: %s", err))
		return
	}

	argsMap := make(map[string]string)
	if v, ok := getResponseData["recordid"]; ok && v != nil {
		argsMap["recordid"] = fmt.Sprintf("%v", v)
	}
	argsMap["domain"] = url.QueryEscape(dnstxtrecName)

	err = r.client.DeleteResourceWithArgsMap(service.Dnstxtrec.Type(), dnstxtrecName, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnstxtrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnstxtrec resource")
}

// Helper function to read dnstxtrec data from API. Returns false when the resource
// no longer exists on the ADC so callers can remove it from state.
func (r *DnstxtrecResource) readDnstxtrecFromApi(ctx context.Context, data *DnstxtrecResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain domain value
	dnstxtrecName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Dnstxtrec.Type(), dnstxtrecName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnstxtrec, got error: %s", err))
		return false
	}

	dnstxtrecSetAttrFromGet(ctx, data, getResponseData)

	return true
}
