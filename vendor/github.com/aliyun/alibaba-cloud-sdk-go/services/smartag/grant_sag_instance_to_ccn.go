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

// GrantSagInstanceToCcn invokes the smartag.GrantSagInstanceToCcn API synchronously
// api document: https://help.aliyun.com/api/smartag/grantsaginstancetoccn.html
func (client *Client) GrantSagInstanceToCcn(request *GrantSagInstanceToCcnRequest) (response *GrantSagInstanceToCcnResponse, err error) {
	response = CreateGrantSagInstanceToCcnResponse()
	err = client.DoAction(request, response)
	return
}

// GrantSagInstanceToCcnWithChan invokes the smartag.GrantSagInstanceToCcn API asynchronously
// api document: https://help.aliyun.com/api/smartag/grantsaginstancetoccn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantSagInstanceToCcnWithChan(request *GrantSagInstanceToCcnRequest) (<-chan *GrantSagInstanceToCcnResponse, <-chan error) {
	responseChan := make(chan *GrantSagInstanceToCcnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantSagInstanceToCcn(request)
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

// GrantSagInstanceToCcnWithCallback invokes the smartag.GrantSagInstanceToCcn API asynchronously
// api document: https://help.aliyun.com/api/smartag/grantsaginstancetoccn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantSagInstanceToCcnWithCallback(request *GrantSagInstanceToCcnRequest, callback func(response *GrantSagInstanceToCcnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantSagInstanceToCcnResponse
		var err error
		defer close(result)
		response, err = client.GrantSagInstanceToCcn(request)
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

// GrantSagInstanceToCcnRequest is the request struct for api GrantSagInstanceToCcn
type GrantSagInstanceToCcnRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CcnUid               string           `position:"Query" name:"CcnUid"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	CcnInstanceId        string           `position:"Query" name:"CcnInstanceId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GrantSagInstanceToCcnResponse is the response struct for api GrantSagInstanceToCcn
type GrantSagInstanceToCcnResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGrantSagInstanceToCcnRequest creates a request to invoke GrantSagInstanceToCcn API
func CreateGrantSagInstanceToCcnRequest() (request *GrantSagInstanceToCcnRequest) {
	request = &GrantSagInstanceToCcnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "GrantSagInstanceToCcn", "smartag", "openAPI")
	return
}

// CreateGrantSagInstanceToCcnResponse creates a response to parse from GrantSagInstanceToCcn response
func CreateGrantSagInstanceToCcnResponse() (response *GrantSagInstanceToCcnResponse) {
	response = &GrantSagInstanceToCcnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
