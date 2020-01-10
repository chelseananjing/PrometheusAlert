package vod

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

// SetEditingProjectMaterials invokes the vod.SetEditingProjectMaterials API synchronously
// api document: https://help.aliyun.com/api/vod/seteditingprojectmaterials.html
func (client *Client) SetEditingProjectMaterials(request *SetEditingProjectMaterialsRequest) (response *SetEditingProjectMaterialsResponse, err error) {
	response = CreateSetEditingProjectMaterialsResponse()
	err = client.DoAction(request, response)
	return
}

// SetEditingProjectMaterialsWithChan invokes the vod.SetEditingProjectMaterials API asynchronously
// api document: https://help.aliyun.com/api/vod/seteditingprojectmaterials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetEditingProjectMaterialsWithChan(request *SetEditingProjectMaterialsRequest) (<-chan *SetEditingProjectMaterialsResponse, <-chan error) {
	responseChan := make(chan *SetEditingProjectMaterialsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetEditingProjectMaterials(request)
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

// SetEditingProjectMaterialsWithCallback invokes the vod.SetEditingProjectMaterials API asynchronously
// api document: https://help.aliyun.com/api/vod/seteditingprojectmaterials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetEditingProjectMaterialsWithCallback(request *SetEditingProjectMaterialsRequest, callback func(response *SetEditingProjectMaterialsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetEditingProjectMaterialsResponse
		var err error
		defer close(result)
		response, err = client.SetEditingProjectMaterials(request)
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

// SetEditingProjectMaterialsRequest is the request struct for api SetEditingProjectMaterials
type SetEditingProjectMaterialsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ProjectId            string `position:"Query" name:"ProjectId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MaterialIds          string `position:"Query" name:"MaterialIds"`
}

// SetEditingProjectMaterialsResponse is the response struct for api SetEditingProjectMaterials
type SetEditingProjectMaterialsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetEditingProjectMaterialsRequest creates a request to invoke SetEditingProjectMaterials API
func CreateSetEditingProjectMaterialsRequest() (request *SetEditingProjectMaterialsRequest) {
	request = &SetEditingProjectMaterialsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SetEditingProjectMaterials", "vod", "openAPI")
	return
}

// CreateSetEditingProjectMaterialsResponse creates a response to parse from SetEditingProjectMaterials response
func CreateSetEditingProjectMaterialsResponse() (response *SetEditingProjectMaterialsResponse) {
	response = &SetEditingProjectMaterialsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}