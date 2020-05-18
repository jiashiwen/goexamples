

### /api/v1/createtask

#### POST
##### Summary:

Create Task

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| CreateRequest | body | json for createtask | Yes | [httpgin.CreateRequest](#httpgin.createrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [httpgin.Response](#httpgin.response) |

### Models


#### httpgin.CreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| autoStart | boolean |  | No |
| taskName | string | TaskName define your taskname | No |
| taskType | string |  | No |

#### httpgin.Response

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | string |  | No |
| error | object |  | No |
| message | object |  | No |