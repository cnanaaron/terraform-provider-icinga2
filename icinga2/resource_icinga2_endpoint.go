package icinga2

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/cnanaaron/go-icinga2-api/iapi"
)

func resourceIcinga2Endpoint() *schema.Resource {

	return &schema.Resource{
		Create: resourceIcinga2EndpointCreate,
		Read:   resourceIcinga2EndpointRead,
		Delete: resourceIcinga2EndpointDelete,
		Schema: map[string]*schema.Schema{
			"name": {
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
				Type:     schema.TypeNumber,
				Optional: true,
				ForceNew: true,
			},
			"log_duration": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIcinga2EndpointCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	endpoint := d.Get("endpoint").(string)
	host := d.Get("host").(string)
	port := d.Get("port").(int)
	log_duration := d.Get("log_duration").(string)

	// Call CreateEndpoint with normalized data
	endpoints, err := client.CreateEndpoint(endpoint, host, port, log_duration)
	if err != nil {
		return err
	}

	found := false
	for _, endpoint := range endpoints {
		if endpoint.Name == endpoint {
			d.SetId(endpoint)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Failed to Create Endpoint %s : %s", endpoint, err)
	}

	return nil
}

func resourceIcinga2EndpointRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	endpoint := d.Get("endpoint").(string)

	endpoints, err := client.GetEndpoint(endpoint)
	if err != nil {
		return err
	}

	found := false
	for _, endpoint := range endpoints {
		if endpoint.Name == endpoint {
			d.SetId(endpoint)
			_ = d.Set("endpoint", endpoint.Name)
			_ = d.Set("host", endpoint.Attrs.host)
			_ = d.Set("port", endpoint.Attrs.port)
			_ = d.Set("log_duration", endpoint.Attrs.log_duration)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Failed to Read Endpoint %s : %s", endpoint, err)
	}

	return nil
}

func resourceIcinga2EndpointDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	endpoint := d.Get("endpoint").(string)

	err := client.DeleteEndpoint(endpoint)
	if err != nil {
		return fmt.Errorf("Failed to Delete Endpoint %s : %s", endpoint, err)
	}

	return nil

}
