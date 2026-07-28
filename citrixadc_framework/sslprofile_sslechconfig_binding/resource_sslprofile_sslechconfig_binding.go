package sslprofile_sslechconfig_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslprofileSslechconfigBindingResource{}
var _ resource.ResourceWithConfigure = (*SslprofileSslechconfigBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileSslechconfigBindingResource)(nil)

func NewSslprofileSslechconfigBindingResource() resource.Resource {
	return &SslprofileSslechconfigBindingResource{}
}

// SslprofileSslechconfigBindingResource defines the resource implementation.
type SslprofileSslechconfigBindingResource struct {
	client *service.NitroClient
}

func (r *SslprofileSslechconfigBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileSslechconfigBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_sslechconfig_binding"
}

func (r *SslprofileSslechconfigBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslprofileSslechconfigBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslprofileSslechconfigBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile_sslechconfig_binding resource")
	sslprofile_sslechconfig_binding := sslprofile_sslechconfig_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslprofile_sslechconfig_binding.Type(), &sslprofile_sslechconfig_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile_sslechconfig_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslprofile_sslechconfig_binding resource")

	// Set ID for the resource before reading state
	// Composite ID = name,echconfigname
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("echconfigname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Echconfigname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslprofileSslechconfigBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslechconfigBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileSslechconfigBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile_sslechconfig_binding resource")

	r.readSslprofileSslechconfigBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the binding is gone out-of-band, remove it from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslechconfigBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslprofileSslechconfigBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: no NITRO update endpoint for this binding; all attributes are
	// RequiresReplace, so Update is a no-op that just reconciles state.
	tflog.Debug(ctx, "Update is a no-op for sslprofile_sslechconfig_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslprofileSslechconfigBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslechconfigBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileSslechconfigBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile_sslechconfig_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs (parent name + member arg)
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "echconfigname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	args := make([]string, 0, 1)
	if val, ok := idMap["echconfigname"]; ok && val != "" {
		args = append(args, "echconfigname:"+utils.UrlEncode(val))
	}

	err = r.client.DeleteResourceWithArgs(service.Sslprofile_sslechconfig_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslprofile_sslechconfig_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslprofile_sslechconfig_binding binding")
}

// Helper function to read sslprofile_sslechconfig_binding data from API
func (r *SslprofileSslechconfigBindingResource) readSslprofileSslechconfigBindingFromApi(ctx context.Context, data *SslprofileSslechconfigBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "echconfigname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslprofile_sslechconfig_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_sslechconfig_binding, got error: %s", err))
		return
	}

	// Resource is missing - deleted out-of-band; signal "gone" so Read removes it from state.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check echconfigname
		if idVal, ok := idMap["echconfigname"]; ok {
			if val, ok := v["echconfigname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["echconfigname"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - deleted out-of-band; signal "gone" so Read removes it from state.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	sslprofile_sslechconfig_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
