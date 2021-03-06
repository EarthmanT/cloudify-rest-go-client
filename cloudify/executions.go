/*
Copyright (c) 2017 GigaSpaces Technologies Ltd. All rights reserved

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudify

import (
	"encoding/json"
	rest "github.com/0lvin-cfy/cloudify-rest-go-client/cloudifyrest"
	"log"
	"net/url"
)

type CloudifyExecutionPost struct {
	WorkflowId   string                 `json:"workflow_id"`
	DeploymentId string                 `json:"deployment_id"`
	Parameters   map[string]interface{} `json:"parameters"`
}

func (exec *CloudifyExecutionPost) SetJsonParameters(parameters string) {
	if len(parameters) == 0 {
		exec.Parameters = map[string]interface{}{}
		return
	}

	err := json.Unmarshal([]byte(parameters), &exec.Parameters)
	if err != nil {
		log.Fatal(err)
	}
}

type CloudifyExecution struct {
	// have id, owner information
	rest.CloudifyResource
	// contain information from post
	CloudifyExecutionPost
	IsSystemWorkflow bool   `json:"is_system_workflow"`
	ErrorMessage     string `json:"error"`
	BlueprintId      string `json:"blueprint_id"`
	Status           string `json:"status"`
}

type CloudifyExecutionGet struct {
	// can be response from api
	rest.CloudifyBaseMessage
	CloudifyExecution
}

type CloudifyExecutions struct {
	rest.CloudifyBaseMessage
	Metadata rest.CloudifyMetadata `json:"metadata"`
	Items    []CloudifyExecution   `json:"items"`
}

// change params type if you want use non uniq values in params
func (cl *CloudifyClient) GetExecutions(params map[string]string) CloudifyExecutions {
	var executions CloudifyExecutions

	values := url.Values{}
	for key, value := range params {
		values.Set(key, value)
	}

	err := cl.Get("executions?"+values.Encode(), &executions)
	if err != nil {
		log.Fatal(err)
	}

	return executions
}

func (cl *CloudifyClient) PostExecution(exec CloudifyExecutionPost) CloudifyExecutionGet {
	var execution CloudifyExecutionGet

	var err error

	err = cl.Post("executions", exec, &execution)
	if err != nil {
		log.Fatal(err)
	}

	return execution
}
