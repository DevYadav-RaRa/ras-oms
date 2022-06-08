package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/helpers"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

func Testing() (string, bool) {
	b, err := ioutil.ReadFile("./src/conf/api.json")
	if err != nil {
		fmt.Print("Bootstrap error: ", err)
		return err.Error(), false
	}

	var demoApi models.ApiPayload
	demoApi.FromJSONString(string(b))

	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------------------------------------")

	fmt.Println("Calling Iam for Authentication")
	var req models.IamRequest
	req.TenantToken = demoApi.TenantToken
	req.BusinessDetails = demoApi.BusinessDetails
	IamAuth := req.GetIamAuthentication("BusinessHeader")

	fmt.Println("Iam Response: ", IamAuth)
	fmt.Println("Authenticated from Iam")

	for i := range demoApi.Orders {
		var temp models.OrderObject
		temp.TenantToken = demoApi.TenantToken
		temp.BusinessDetails = demoApi.BusinessDetails
		temp.Order = demoApi.Orders[i]

		status, resp := helpers.PostOrder(temp)
		if !resp {
			return status, resp
		}

		fmt.Println(status, " :: ", resp)
		fmt.Println("-------------------------------------------")
		fmt.Println("-------------------------------------------")
	}

	return "Success: Processing", true
}

func Bootstrap(appCtx framework.Framework) {
	fmt.Println("Running Bootstrap...")
	Testing()
	fmt.Println("App is ready!")
}
