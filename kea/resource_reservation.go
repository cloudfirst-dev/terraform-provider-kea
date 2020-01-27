package kea

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceReservationCreate,
		Read:   resourceReservationRead,
		Update: resourceReservationUpdate,
		Delete: resourceReservationDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hw_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceReservationCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	saveHost := &SaveHost{
		Address:        d.Get("ip_address").(string),
		Identifier:     d.Get("hw_address").(string),
		Hostname:       d.Get("hostname").(string),
		SubnetID:       int64(d.Get("subnet_id").(int)),
		IdentifierType: HwAddress,
	}

	log.Printf("[DEBUG] saveHost : %v", saveHost)

	host, err := client.CreateReservation(saveHost)

	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(host.ID, 10))
	return resourceReservationRead(d, m)
}

func resourceReservationRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	log.Printf("[DEBUG] getting host : %s", d.Id())

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	_, err := client.GetReservationById(id)
	if err != nil {
		return err
	}

	return nil
}

func resourceReservationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceReservationRead(d, m)
}

func resourceReservationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
