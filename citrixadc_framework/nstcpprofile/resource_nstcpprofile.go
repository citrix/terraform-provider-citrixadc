package nstcpprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NstcpprofileResource{}
var _ resource.ResourceWithConfigure = (*NstcpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NstcpprofileResource)(nil)

func NewNstcpprofileResource() resource.Resource {
	return &NstcpprofileResource{}
}

// NstcpprofileResource defines the resource implementation.
type NstcpprofileResource struct {
	client *service.NitroClient
}

func (r *NstcpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstcpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstcpprofile"
}

func (r *NstcpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstcpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstcpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstcpprofile resource")

	// Create API request body from the model
	nstcpprofile := nstcpprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	nstcpprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nstcpprofile.Type(), nstcpprofileName, &nstcpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nstcpprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(nstcpprofileName)

	// Read the updated state back
	if !r.readNstcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nstcpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstcpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstcpprofile resource")

	found := r.readNstcpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NstcpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NstcpprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nstcpprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Ackaggregation.Equal(state.Ackaggregation) {
		if config.Ackaggregation.IsNull() {
			attributesToUnset = append(attributesToUnset, "ackaggregation")
		} else {
			hasChange = true
		}
	}
	if !data.Ackonpush.Equal(state.Ackonpush) {
		if config.Ackonpush.IsNull() {
			attributesToUnset = append(attributesToUnset, "ackonpush")
		} else {
			hasChange = true
		}
	}
	if !data.Applyadaptivetcp.Equal(state.Applyadaptivetcp) {
		if config.Applyadaptivetcp.IsNull() {
			attributesToUnset = append(attributesToUnset, "applyadaptivetcp")
		} else {
			hasChange = true
		}
	}
	if !data.Buffersize.Equal(state.Buffersize) {
		if config.Buffersize.IsNull() {
			attributesToUnset = append(attributesToUnset, "buffersize")
		} else {
			hasChange = true
		}
	}
	if !data.Burstratecontrol.Equal(state.Burstratecontrol) {
		if config.Burstratecontrol.IsNull() {
			attributesToUnset = append(attributesToUnset, "burstratecontrol")
		} else {
			hasChange = true
		}
	}
	if !data.Clientiptcpoption.Equal(state.Clientiptcpoption) {
		if config.Clientiptcpoption.IsNull() {
			attributesToUnset = append(attributesToUnset, "clientiptcpoption")
		} else {
			hasChange = true
		}
	}
	if !data.Clientiptcpoptionnumber.Equal(state.Clientiptcpoptionnumber) {
		if config.Clientiptcpoptionnumber.IsNull() {
			attributesToUnset = append(attributesToUnset, "clientiptcpoptionnumber")
		} else {
			hasChange = true
		}
	}
	if !data.Delayedack.Equal(state.Delayedack) {
		if config.Delayedack.IsNull() {
			attributesToUnset = append(attributesToUnset, "delayedack")
		} else {
			hasChange = true
		}
	}
	if !data.Dropestconnontimeout.Equal(state.Dropestconnontimeout) {
		if config.Dropestconnontimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "dropestconnontimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Drophalfclosedconnontimeout.Equal(state.Drophalfclosedconnontimeout) {
		if config.Drophalfclosedconnontimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "drophalfclosedconnontimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Dsack.Equal(state.Dsack) {
		if config.Dsack.IsNull() {
			attributesToUnset = append(attributesToUnset, "dsack")
		} else {
			hasChange = true
		}
	}
	if !data.Dupackthresh.Equal(state.Dupackthresh) {
		if config.Dupackthresh.IsNull() {
			attributesToUnset = append(attributesToUnset, "dupackthresh")
		} else {
			hasChange = true
		}
	}
	if !data.Dynamicreceivebuffering.Equal(state.Dynamicreceivebuffering) {
		if config.Dynamicreceivebuffering.IsNull() {
			attributesToUnset = append(attributesToUnset, "dynamicreceivebuffering")
		} else {
			hasChange = true
		}
	}
	if !data.Ecn.Equal(state.Ecn) {
		if config.Ecn.IsNull() {
			attributesToUnset = append(attributesToUnset, "ecn")
		} else {
			hasChange = true
		}
	}
	if !data.Establishclientconn.Equal(state.Establishclientconn) {
		if config.Establishclientconn.IsNull() {
			attributesToUnset = append(attributesToUnset, "establishclientconn")
		} else {
			hasChange = true
		}
	}
	if !data.Fack.Equal(state.Fack) {
		if config.Fack.IsNull() {
			attributesToUnset = append(attributesToUnset, "fack")
		} else {
			hasChange = true
		}
	}
	if !data.Flavor.Equal(state.Flavor) {
		if config.Flavor.IsNull() {
			attributesToUnset = append(attributesToUnset, "flavor")
		} else {
			hasChange = true
		}
	}
	if !data.Frto.Equal(state.Frto) {
		if config.Frto.IsNull() {
			attributesToUnset = append(attributesToUnset, "frto")
		} else {
			hasChange = true
		}
	}
	if !data.Hystart.Equal(state.Hystart) {
		if config.Hystart.IsNull() {
			attributesToUnset = append(attributesToUnset, "hystart")
		} else {
			hasChange = true
		}
	}
	if !data.Initialcwnd.Equal(state.Initialcwnd) {
		if config.Initialcwnd.IsNull() {
			attributesToUnset = append(attributesToUnset, "initialcwnd")
		} else {
			hasChange = true
		}
	}
	if !data.Ka.Equal(state.Ka) {
		if config.Ka.IsNull() {
			attributesToUnset = append(attributesToUnset, "ka")
		} else {
			hasChange = true
		}
	}
	if !data.Kaconnidletime.Equal(state.Kaconnidletime) {
		if config.Kaconnidletime.IsNull() {
			attributesToUnset = append(attributesToUnset, "kaconnidletime")
		} else {
			hasChange = true
		}
	}
	if !data.Kamaxprobes.Equal(state.Kamaxprobes) {
		if config.Kamaxprobes.IsNull() {
			attributesToUnset = append(attributesToUnset, "kamaxprobes")
		} else {
			hasChange = true
		}
	}
	if !data.Kaprobeinterval.Equal(state.Kaprobeinterval) {
		if config.Kaprobeinterval.IsNull() {
			attributesToUnset = append(attributesToUnset, "kaprobeinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Kaprobeupdatelastactivity.Equal(state.Kaprobeupdatelastactivity) {
		if config.Kaprobeupdatelastactivity.IsNull() {
			attributesToUnset = append(attributesToUnset, "kaprobeupdatelastactivity")
		} else {
			hasChange = true
		}
	}
	if !data.Maxburst.Equal(state.Maxburst) {
		if config.Maxburst.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxburst")
		} else {
			hasChange = true
		}
	}
	if !data.Maxcwnd.Equal(state.Maxcwnd) {
		if config.Maxcwnd.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxcwnd")
		} else {
			hasChange = true
		}
	}
	if !data.Maxpktpermss.Equal(state.Maxpktpermss) {
		if config.Maxpktpermss.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxpktpermss")
		} else {
			hasChange = true
		}
	}
	if !data.Minrto.Equal(state.Minrto) {
		if config.Minrto.IsNull() {
			attributesToUnset = append(attributesToUnset, "minrto")
		} else {
			hasChange = true
		}
	}
	if !data.Mpcapablecbit.Equal(state.Mpcapablecbit) {
		if config.Mpcapablecbit.IsNull() {
			attributesToUnset = append(attributesToUnset, "mpcapablecbit")
		} else {
			hasChange = true
		}
	}
	if !data.Mptcp.Equal(state.Mptcp) {
		if config.Mptcp.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcp")
		} else {
			hasChange = true
		}
	}
	if !data.Mptcpdropdataonpreestsf.Equal(state.Mptcpdropdataonpreestsf) {
		if config.Mptcpdropdataonpreestsf.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcpdropdataonpreestsf")
		} else {
			hasChange = true
		}
	}
	if !data.Mptcpfastopen.Equal(state.Mptcpfastopen) {
		if config.Mptcpfastopen.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcpfastopen")
		} else {
			hasChange = true
		}
	}
	if !data.Mptcpsessiontimeout.Equal(state.Mptcpsessiontimeout) {
		if config.Mptcpsessiontimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcpsessiontimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Mss.Equal(state.Mss) {
		if config.Mss.IsNull() {
			attributesToUnset = append(attributesToUnset, "mss")
		} else {
			hasChange = true
		}
	}
	if !data.Nagle.Equal(state.Nagle) {
		if config.Nagle.IsNull() {
			attributesToUnset = append(attributesToUnset, "nagle")
		} else {
			hasChange = true
		}
	}
	if !data.Oooqsize.Equal(state.Oooqsize) {
		if config.Oooqsize.IsNull() {
			attributesToUnset = append(attributesToUnset, "oooqsize")
		} else {
			hasChange = true
		}
	}
	if !data.Pktperretx.Equal(state.Pktperretx) {
		if config.Pktperretx.IsNull() {
			attributesToUnset = append(attributesToUnset, "pktperretx")
		} else {
			hasChange = true
		}
	}
	if !data.Rateqmax.Equal(state.Rateqmax) {
		if config.Rateqmax.IsNull() {
			attributesToUnset = append(attributesToUnset, "rateqmax")
		} else {
			hasChange = true
		}
	}
	if !data.Rfc5961compliance.Equal(state.Rfc5961compliance) {
		if config.Rfc5961compliance.IsNull() {
			attributesToUnset = append(attributesToUnset, "rfc5961compliance")
		} else {
			hasChange = true
		}
	}
	if !data.Rstmaxack.Equal(state.Rstmaxack) {
		if config.Rstmaxack.IsNull() {
			attributesToUnset = append(attributesToUnset, "rstmaxack")
		} else {
			hasChange = true
		}
	}
	if !data.Rstwindowattenuate.Equal(state.Rstwindowattenuate) {
		if config.Rstwindowattenuate.IsNull() {
			attributesToUnset = append(attributesToUnset, "rstwindowattenuate")
		} else {
			hasChange = true
		}
	}
	if !data.Sack.Equal(state.Sack) {
		if config.Sack.IsNull() {
			attributesToUnset = append(attributesToUnset, "sack")
		} else {
			hasChange = true
		}
	}
	if !data.Sendbuffsize.Equal(state.Sendbuffsize) {
		if config.Sendbuffsize.IsNull() {
			attributesToUnset = append(attributesToUnset, "sendbuffsize")
		} else {
			hasChange = true
		}
	}
	if !data.Sendclientportintcpoption.Equal(state.Sendclientportintcpoption) {
		if config.Sendclientportintcpoption.IsNull() {
			attributesToUnset = append(attributesToUnset, "sendclientportintcpoption")
		} else {
			hasChange = true
		}
	}
	if !data.Slowstartincr.Equal(state.Slowstartincr) {
		if config.Slowstartincr.IsNull() {
			attributesToUnset = append(attributesToUnset, "slowstartincr")
		} else {
			hasChange = true
		}
	}
	if !data.Slowstartthreshold.Equal(state.Slowstartthreshold) {
		if config.Slowstartthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "slowstartthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Spoofsyndrop.Equal(state.Spoofsyndrop) {
		if config.Spoofsyndrop.IsNull() {
			attributesToUnset = append(attributesToUnset, "spoofsyndrop")
		} else {
			hasChange = true
		}
	}
	if !data.Syncookie.Equal(state.Syncookie) {
		if config.Syncookie.IsNull() {
			attributesToUnset = append(attributesToUnset, "syncookie")
		} else {
			hasChange = true
		}
	}
	if !data.Taillossprobe.Equal(state.Taillossprobe) {
		if config.Taillossprobe.IsNull() {
			attributesToUnset = append(attributesToUnset, "taillossprobe")
		} else {
			hasChange = true
		}
	}
	if !data.Tcpfastopen.Equal(state.Tcpfastopen) {
		if config.Tcpfastopen.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcpfastopen")
		} else {
			hasChange = true
		}
	}
	if !data.Tcpfastopencookiesize.Equal(state.Tcpfastopencookiesize) {
		if config.Tcpfastopencookiesize.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcpfastopencookiesize")
		} else {
			hasChange = true
		}
	}
	if !data.Tcpmode.Equal(state.Tcpmode) {
		if config.Tcpmode.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcpmode")
		} else {
			hasChange = true
		}
	}
	if !data.Tcprate.Equal(state.Tcprate) {
		if config.Tcprate.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcprate")
		} else {
			hasChange = true
		}
	}
	if !data.Tcpsegoffload.Equal(state.Tcpsegoffload) {
		if config.Tcpsegoffload.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcpsegoffload")
		} else {
			hasChange = true
		}
	}
	if !data.Timestamp.Equal(state.Timestamp) {
		if config.Timestamp.IsNull() {
			attributesToUnset = append(attributesToUnset, "timestamp")
		} else {
			hasChange = true
		}
	}
	if !data.Ws.Equal(state.Ws) {
		if config.Ws.IsNull() {
			attributesToUnset = append(attributesToUnset, "ws")
		} else {
			hasChange = true
		}
	}
	if !data.Wsval.Equal(state.Wsval) {
		if config.Wsval.IsNull() {
			attributesToUnset = append(attributesToUnset, "wsval")
		} else {
			hasChange = true
		}
	}

	nstcpprofileName := data.Name.ValueString()
	if hasChange {
		// Build the payload from the plan (name identifies the resource)
		nstcpprofile := nstcpprofileGetThePayloadFromthePlan(ctx, &data)

		// Make API call
		// Named resource - use UpdateResource
		_, err := r.client.UpdateResource(service.Nstcpprofile.Type(), nstcpprofileName, &nstcpprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstcpprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nstcpprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nstcpprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nstcpprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nstcpprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNstcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nstcpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstcpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstcpprofile resource")
	// Named resource - delete using DeleteResource
	nstcpprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nstcpprofile.Type(), nstcpprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nstcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nstcpprofile resource")
}

// Helper function to read nstcpprofile data from API
func (r *NstcpprofileResource) readNstcpprofileFromApi(ctx context.Context, data *NstcpprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	nstcpprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nstcpprofile.Type(), nstcpprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstcpprofile, got error: %s", err))
		return false
	}

	nstcpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
