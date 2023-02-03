package icinga2api

import (
	"fmt"

	"github.com/cnanaaron/terraform-provider-icinga2api/iapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIcinga2Endpoint() *schema.Resource {

	return &schema.Resource{
		Create: resourceIcinga2EndpointCreate,
		Read:   resourceIcinga2EndpointRead,
		Delete: resourceIcinga2EndpointDelete,
		Schema: map[string]*schema.Schema{
			"endpointname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Endpoint name",
				ForceNew:    true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"log_duration": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIcinga2EndpointCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	endpointname := d.Get("endpointname").(string)
	host := d.Get("host").(string)
	port := d.Get("port").(string)
	log_duration := d.Get("log_duration").(int)

	// Call CreateEndpoint with normalized data
	endpoints, err := client.CreateEndpoint(endpointname, host, port, log_duration)
	if err != nil {
		return err
	}

	found := false
	for _, endpoint := range endpoints {
		if endpoint.Name == endpointname {
			d.SetId(endpointname)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Failed to Create Endpoint %s : %s", endpointname, err)
	}

	return nil
}

func resourceIcinga2EndpointRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	endpointname := d.Get("endpointname").(string)

	endpoints, err := client.GetEndpoint(endpointname)
	if err != nil {
		return err
	}

	found := false
	for _, endpoint := range endpoints {
		if endpoint.Name == endpointname {
			d.SetId(endpointname)
			_ = d.Set("endpoint", endpoint.Name)
			_ = d.Set("host", endpoint.Attrs.Host)
			_ = d.Set("port", endpoint.Attrs.Port)
			_ = d.Set("log_duration", endpoint.Attrs.LogDuration)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Failed to Read Endpoint %s : %s", endpointname, err)
	}

	return nil
}

func resourceIcinga2EndpointDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	endpointname := d.Get("endpointname").(string)

	err := client.DeleteEndpoint(endpointname)
	if err != nil {
		return fmt.Errorf("Failed to Delete Endpoint %s : %s", endpointname, err)
	}

	return nil

}
