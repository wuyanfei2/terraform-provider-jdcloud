package jdcloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/client"
	"time"
)

func resourceJDCloudVpc() *schema.Resource {

	return &schema.Resource{

		Create: resourceVpcCreate,
		Read:   resourceVpcRead,
		Update: resourceVpcUpdate,
		Delete: resourceVpcDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"vpc_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"cidr_block": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpcCreate(d *schema.ResourceData, m interface{}) error {

	config := m.(*JDCloudConfig)
	req := apis.NewCreateVpcRequest(config.Region, d.Get("vpc_name").(string))
	conn := client.NewVpcClient(config.Credential)

	if _, ok := d.GetOk("cidr_block"); ok {
		req.AddressPrefix = GetStringAddr(d, "cidr_block")
	}
	if _, ok := d.GetOk("description"); ok {
		req.Description = GetStringAddr(d, "description")
	}

	e := resource.Retry(20*time.Second, func() *resource.RetryError {

		resp, err := conn.CreateVpc(req)

		if err == nil && resp.Error.Code == REQUEST_COMPLETED {
			d.SetId(resp.Result.VpcId)
			return nil
		}

		if connectionError(err) {
			return resource.RetryableError(formatConnectionErrorMessage())
		} else {
			return resource.NonRetryableError(formatErrorMessage(resp.Error, err))
		}
	})
	if e != nil {
		return e
	}
	return resourceVpcRead(d, m)
}

func resourceVpcRead(d *schema.ResourceData, m interface{}) error {

	config := m.(*JDCloudConfig)
	vpcClient := client.NewVpcClient(config.Credential)

	req := apis.NewDescribeVpcRequest(config.Region, d.Id())
	resp, err := vpcClient.DescribeVpc(req)

	if err != nil {
		return fmt.Errorf("[ERROR] resourceVpcRead failed %s ", err.Error())
	}

	if resp.Error.Code == RESOURCE_NOT_FOUND {
		d.SetId("")
		return nil
	}

	if resp.Error.Code != REQUEST_COMPLETED {
		return fmt.Errorf("[ERROR] resourceVpcRead failed  code:%d staus:%s message:%s ", resp.Error.Code, resp.Error.Status, resp.Error.Message)
	}

	d.Set("vpc_name", resp.Result.Vpc.VpcName)
	d.Set("cidr_block", resp.Result.Vpc.AddressPrefix)
	d.Set("description", resp.Result.Vpc.Description)
	return nil
}

func resourceVpcUpdate(d *schema.ResourceData, m interface{}) error {

	if d.HasChange("vpc_name") || d.HasChange("description") {

		config := m.(*JDCloudConfig)
		vpcClient := client.NewVpcClient(config.Credential)
		req := apis.NewModifyVpcRequestWithAllParams(
			config.Region,
			d.Id(),
			GetStringAddr(d, "vpc_name"),
			GetStringAddr(d, "description"),
		)
		resp, err := vpcClient.ModifyVpc(req)

		if err != nil {
			return fmt.Errorf("[ERROR] resourceVpcUpdate failed %s ", err.Error())
		}

		if resp.Error.Code != REQUEST_COMPLETED {
			return fmt.Errorf("[ERROR] resourceVpcUpdate failed  code:%d staus:%s message:%s ", resp.Error.Code, resp.Error.Status, resp.Error.Message)
		}

	}

	return resourceVpcRead(d, m)
}

func resourceVpcDelete(d *schema.ResourceData, m interface{}) error {

	config := m.(*JDCloudConfig)
	vpcClient := client.NewVpcClient(config.Credential)

	req := apis.NewDeleteVpcRequest(config.Region, d.Id())
	resp, err := vpcClient.DeleteVpc(req)

	if err != nil {
		return fmt.Errorf("[ERROR] resourceVpcDelete failed %s ", err.Error())
	}

	if resp.Error.Code != REQUEST_COMPLETED {
		return fmt.Errorf("[ERROR] resourceVpcDelete failed  code:%d staus:%s message:%s ", resp.Error.Code, resp.Error.Status, resp.Error.Message)
	}

	d.SetId("")
	return nil
}
