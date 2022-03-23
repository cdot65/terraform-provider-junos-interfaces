
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlInterfacesInterfaceName struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_interface  struct {
			XMLName xml.Name `xml:"interface"`
			V_name  *string  `xml:"name,omitempty"`
		} `xml:"interfaces>interface"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosInterfacesInterfaceNameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlInterfacesInterfaceName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosInterfacesInterfaceNameRead(d,m)
}

func junosInterfacesInterfaceNameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlInterfacesInterfaceName{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_interface.V_name)

	return nil
}

func junosInterfacesInterfaceNameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlInterfacesInterfaceName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosInterfacesInterfaceNameRead(d,m)
}

func junosInterfacesInterfaceNameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosInterfacesInterfaceName() *schema.Resource {
	return &schema.Resource{
		Create: junosInterfacesInterfaceNameCreate,
		Read: junosInterfacesInterfaceNameRead,
		Update: junosInterfacesInterfaceNameUpdate,
		Delete: junosInterfacesInterfaceNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface. ",
			},
		},
	}
}