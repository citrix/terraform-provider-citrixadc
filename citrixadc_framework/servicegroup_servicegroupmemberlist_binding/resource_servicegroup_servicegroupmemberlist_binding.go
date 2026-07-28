package servicegroup_servicegroupmemberlist_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ServicegroupServicegroupmemberlistBindingResource{}
var _ resource.ResourceWithConfigure = (*ServicegroupServicegroupmemberlistBindingResource)(nil)

func NewServicegroupServicegroupmemberlistBindingResource() resource.Resource {
	return &ServicegroupServicegroupmemberlistBindingResource{}
}

// ServicegroupServicegroupmemberlistBindingResource defines the resource implementation.
type ServicegroupServicegroupmemberlistBindingResource struct {
	client *service.NitroClient
}

func (r *ServicegroupServicegroupmemberlistBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup_servicegroupmemberlist_binding"
}

func (r *ServicegroupServicegroupmemberlistBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServicegroupServicegroupmemberlistBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServicegroupServicegroupmemberlistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating servicegroup_servicegroupmemberlist_binding resource")
	servicegroup_servicegroupmemberlist_binding := servicegroup_servicegroupmemberlist_bindingGetThePayloadFromthePlan(ctx, &data)

	// NITRO add for this binding is HTTP PUT - use UpdateUnnamedResource (PUT replaces the whole member set)
	err := r.client.UpdateUnnamedResource(service.Servicegroup_servicegroupmemberlist_binding.Type(), &servicegroup_servicegroupmemberlist_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create servicegroup_servicegroupmemberlist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created servicegroup_servicegroupmemberlist_binding resource")

	// Set ID for the resource. The resource manages the entire member list atomically;
	// the ID is the parent servicegroupname (plain value).
	data.Id = types.StringValue(data.Servicegroupname.ValueString())

	// No GET endpoint exists for this resource - state is what was applied.
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberlistBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServicegroupServicegroupmemberlistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Pattern 13: NITRO exposes no get/get(all)/count endpoint for this binding.
	// Read is a no-op preserving prior state; drift detection is impossible by design.
	tflog.Debug(ctx, "Read is a no-op for servicegroup_servicegroupmemberlist_binding; NITRO has no GET endpoint")

	// Save (unchanged) prior state back
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberlistBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ServicegroupServicegroupmemberlistBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: all attributes are RequiresReplace and the resource is not updateable
	// (is_updateable=0 in metadata). Terraform never reaches Update with a real diff;
	// this is a documented no-op that preserves state.
	tflog.Debug(ctx, "Update is a no-op for servicegroup_servicegroupmemberlist_binding; all attributes are RequiresReplace")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberlistBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServicegroupServicegroupmemberlistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting servicegroup_servicegroupmemberlist_binding resource")

	// NITRO delete is DELETE by URL path key servicegroupname (no args).
	err := r.client.DeleteResource(service.Servicegroup_servicegroupmemberlist_binding.Type(), data.Servicegroupname.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete servicegroup_servicegroupmemberlist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Removed servicegroup_servicegroupmemberlist_binding from Terraform state")
}
