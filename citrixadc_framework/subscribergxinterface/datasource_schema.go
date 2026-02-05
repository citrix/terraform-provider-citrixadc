package subscribergxinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SubscribergxinterfaceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cerrequesttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Healthcheck request timeout, in seconds, after which the Citrix ADC considers that no CCA packet received to the initiated CCR. After this time Citrix ADC should send again CCR to PCRF server. !",
			},
			"healthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Set this setting to yes if Citrix ADC should send DWR packets to PCRF server. When the session is idle, healthcheck timer expires and DWR packets are initiated in order to check that PCRF server is active. By default set to No. !",
			},
			"healthcheckttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Healthcheck timeout, in seconds, after which the DWR will be sent in order to ensure the state of the PCRF server. Any CCR, CCA, RAR or RRA message resets the timer. !",
			},
			"holdonsubscriberabsence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set this setting to yes if Citrix ADC needs to Hold pakcets till subscriber session is fetched from PCRF. Else set to NO. By default set to yes. If this setting is set to NO, then till Citrix ADC fetches subscriber from PCRF, default subscriber profile will be applied to this subscriber if configured. If default subscriber profile is also not configured an undef would be raised to expressions which use Subscriber attributes.",
			},
			"idlettl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Idle Time, in seconds, after which the Gx CCR-U request will be sent after any PCRF activity on a session. Any RAR or CCA message resets the timer.\nZero value disables the idle timeout. !",
			},
			"negativettl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Negative TTL, in seconds, after which the Gx CCR-I request will be resent for sessions that have not been resolved by PCRF due to server being down or no response or failed response. Instead of polling the PCRF server constantly, negative-TTL makes Citrix ADC stick to un-resolved session. Meanwhile Citrix ADC installs a negative session to avoid going to PCRF.\nFor Negative Sessions, Netcaler inherits the attributes from default subscriber profile if default subscriber is configured. A default subscriber could be configured as 'add subscriber profile *'. Or these attributes can be inherited from Radius as well if Radius is configued.\nZero value disables the Negative Sessions. And Citrix ADC does not install Negative sessions even if subscriber session could not be fetched. !",
			},
			"negativettllimitedsuccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set this to YES if Citrix ADC should create negative session for Result-Code DIAMETER_LIMITED_SUCCESS (2002) received in CCA-I. If set to NO, regular session is created.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"pcrfrealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRF realm is of type DiameterIdentity and contains the realm of PCRF to which the message is to be routed. This is the realm used in Destination-Realm AVP by Citrix ADC Gx client (as a Diameter node).",
			},
			"purgesdbongxfailure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set this setting to YES if needed to purge Subscriber Database in case of Gx failure. By default set to NO.",
			},
			"requestretryattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "If the request does not complete within requestTimeout time, the request is retransmitted for requestRetryAttempts time.",
			},
			"requesttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Time, in seconds, within which the Gx CCR request must complete. If the request does not complete within this time, the request is retransmitted for requestRetryAttempts time. If still reuqest is not complete then default subscriber profile will be applied to this subscriber if configured. If default subscriber profile is also not configured an undef would be raised to expressions which use Subscriber attributes.\nZero disables the timeout. !",
			},
			"revalidationtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Revalidation Timeout, in seconds, after which the Gx CCR-U request will be sent after any PCRF activity on a session. Any RAR or CCA message resets the timer.\nZero value disables the idle timeout. !",
			},
			"service": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of DIAMETER/SSL_DIAMETER service corresponding to PCRF to which the Gx connection is established. The service type of the service must be DIAMETER/SSL_DIAMETER. Mutually exclusive with vserver parameter. Therefore, you cannot set both Service and the Virtual Server in the Gx Interface.",
			},
			"servicepathavp": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The AVP code in which PCRF sends service path applicable for subscriber.",
			},
			"servicepathvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The vendorid of the AVP in which PCRF sends service path for subscriber.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing, or content switching vserver to which the Gx connections are established. The service type of the virtual server must be DIAMETER/SSL_DIAMETER. Mutually exclusive with the service parameter. Therefore, you cannot set both service and the Virtual Server in the Gx Interface.",
			},
		},
	}
}
