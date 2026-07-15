package fis_interface_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FisInterfaceBindingResource{}
var _ resource.ResourceWithConfigure = (*FisInterfaceBindingResource)(nil)

func NewFisInterfaceBindingResource() resource.Resource {
	return &FisInterfaceBindingResource{}
}

// FisInterfaceBindingResource defines the resource implementation.
type FisInterfaceBindingResource struct {
	client *service.NitroClient
}

func (r *FisInterfaceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fis_interface_binding"
}

func (r *FisInterfaceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FisInterfaceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FisInterfaceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating fis_interface_binding resource")
	fis_interface_binding := fis_interface_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Fis_interface_binding.Type(), &fis_interface_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create fis_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created fis_interface_binding resource")

	// Set ID for the resource exactly once here. Composite key: name,ifnum
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// NITRO exposes no get endpoint for fis_interface_binding, and the aggregate
	// fis_binding/<name> endpoint does not surface interface members (verified live
	// on NS14.1: `show fis` lists the bound interface but no NITRO GET returns it).
	// There is therefore nothing to read back; the plan values are authoritative.

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FisInterfaceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FisInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a tolerant no-op for fis_interface_binding: NITRO exposes no get endpoint
	// (doc Operations: only add/delete) and the aggregate fis_binding/<name> endpoint does
	// not surface interface members (verified live on NS14.1 — `show fis` lists the bound
	// interface but no NITRO GET returns it). Drift detection is impossible by definition,
	// so the prior state is re-stored unchanged rather than hard-erroring.
	tflog.Debug(ctx, "Read is a no-op for fis_interface_binding; NITRO exposes no get endpoint")

	// Save unchanged data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FisInterfaceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state FisInterfaceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for fis_interface_binding: NITRO exposes only add (PUT) and delete
	// (no update/change endpoint), and all schema attributes are RequiresReplace, so Terraform
	// recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for fis_interface_binding; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FisInterfaceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FisInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting fis_interface_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent (name) as the
	// resource (URL) name and ifnum (UrlEncoded, contains '/') passed as args.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["ifnum"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Fis_interface_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete fis_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted fis_interface_binding binding")
}
