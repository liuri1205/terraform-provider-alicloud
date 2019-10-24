package smartag

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyACL invokes the smartag.ModifyACL API synchronously
// api document: https://help.aliyun.com/api/smartag/modifyacl.html
func (client *Client) ModifyACL(request *ModifyACLRequest) (response *ModifyACLResponse, err error) {
	response = CreateModifyACLResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyACLWithChan invokes the smartag.ModifyACL API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyacl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyACLWithChan(request *ModifyACLRequest) (<-chan *ModifyACLResponse, <-chan error) {
	responseChan := make(chan *ModifyACLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyACL(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyACLWithCallback invokes the smartag.ModifyACL API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyacl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyACLWithCallback(request *ModifyACLRequest, callback func(response *ModifyACLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyACLResponse
		var err error
		defer close(result)
		response, err = client.ModifyACL(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyACLRequest is the request struct for api ModifyACL
type ModifyACLRequest struct {
	*requests.RpcRequest
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyACLResponse is the response struct for api ModifyACL
type ModifyACLResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyACLRequest creates a request to invoke ModifyACL API
func CreateModifyACLRequest() (request *ModifyACLRequest) {
	request = &ModifyACLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifyACL", "smartag", "openAPI")
	return
}

// CreateModifyACLResponse creates a response to parse from ModifyACL response
func CreateModifyACLResponse() (response *ModifyACLResponse) {
	response = &ModifyACLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
