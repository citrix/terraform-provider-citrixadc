# Citrix ADC configuration examples

> For individual resource examples, please refer to [Citrix ADC docs in Terraform Registry](https://registry.terraform.io/providers/citrix/citrixadc/latest/docs)

> :round_pushpin:For deploying Citrix ADC in Public Cloud - AWS and Azure, check out cloud scripts in github repo [terraform-cloud-scripts](https://github.com/citrix/terraform-cloud-scripts).

> :envelope: For any immediate issues or help , reach out to us at NetScaler-AutomationToolkit@cloud.com !

Below is the table showing the usecase you would like to achieve and where you can find the respective example.

|**Folder**|**ADC Usecase**|**Terraform Example Folder**|
|--|--|--|
|**Basic ADC Operations**|Ping Citrix ADC|[HERE](./basic_adc_operations/pinger/)|
||Change Citrix ADC Password|[HERE](./basic_adc_operations/password_resetter/)|
||Save Citrix ADC Config|[HERE](./saveconfig/)|
||Reboot Citrix ADC|[HERE](./basic_adc_operations/rebooter/)|
||Save/Update/Clear Citrix ADC configs|[HERE](./basic_adc_operations/nsconfig_save_update_clear_configs/)|
||Copy a file from local to Citrix ADC|[HERE](./basic_adc_operations/systemfile_copy_a_file_from_local_to_citrixadc/)|
||Upgrade Citrix ADC|[HERE](./basic_adc_operations/upgrade_citrixadc/)|
||Upgrade HA Pair|[HERE](./basic_adc_operations/upgrade_ha_pair/)|
||Take Citrix ADC system backup|[HERE](./basic_adc_operations/systembackup/)|
|**ADC Usecases**|Configure a simple loadbalancer|[HERE](./adc_usecases/simple_lb/)|
||Get SSL A+ Certified apps using Citrix ADC|[HERE](./adc_usecases/aplus-certified-via-citrix-adc/)|
||Configure a secure content switching server in Citrix ADC|[HERE](./adc_usecases/secure_cs_server/)|
||Redirect External URL to internal URL using Rewrite/Responder policies|[HERE](./adc_usecases/redirect_external_url_to_internal_url/)|
||Redirect client URL to New URL using Reponder and Rewrite Polices|[HERE](./adc_usecases/redirecting_client_to_new_url/)|
|**Special ADC Resources**|Config Citrix ADC if there are no citrixadc resources (using local-exec)|[HERE](./special_adc_resources/using_local_exec_to_configure_citrixadc_as_the_last_resort/)|
||How to get data of various NITRO objects|[HERE](./special_adc_resources/nitro_info_get_information_of_various_nitro_objects/)|
||Create NITRO resources in generic way|[HERE](./special_adc_resources/nitro_resource_generically_create_nitro_resources/)|
