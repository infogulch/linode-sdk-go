/*
 * Linode API
 *
 * # Introduction The Linode API provides the ability to programmatically manage the full range of Linode products and services.  This reference is designed to assist application developers and system administrators.  Each endpoint includes descriptions, request syntax, and examples using standard HTTP requests. Response data is returned in JSON format.   This document was generated from our OpenAPI Specification.  See the [OpenAPI website](https://www.openapis.org) for more information.  [Download the Linode OpenAPI Specification](/openapi.yaml)   # Changelog  [View our Changelog](/changelog) to see release notes on all changes made to our API.  # Access and Authentication  Some endpoints are publicly accessible without requiring authentication. All endpoints affecting your Account, however, require either a Personal Access Token or OAuth authentication (when using third-party applications).  ## Personal Access Token  The easiest way to access the API is with a Personal Access Token (PAT) generated from the [Linode Cloud Manager](https://cloud.linode.com/profile/tokens).  All scopes for the OAuth security model (defined below) apply to this security model as well.  ### Authentication  | Security Scheme Type: | HTTP | |-----------------------|------| | **HTTP Authorization Scheme** | bearer |  ## OAuth  The OAuth workflow is a three-step process to authenticate a User before an application can start making API calls on the User's behalf. If all you need is a Personal Access Token, feel free to skip ahead to the next section.  First, the User visits the application's website and is directed to log with Linode. The User is then redirected to Linode's authentication server and presented the scope levels the application is requesting. Once the User accepts the request for access, we redirect them back to the application's specified redirect URI with an access code.  Once the User has logged in to Linode and you have received an exchange code, you will need to exchange that access code for an Authorization token. You do this by making an HTTP POST request to the following address:  ``` https://login.linode.com/oauth/token ```  Make this request as `application/x-www-form-urlencoded` or as `multipart/form-data` and include the following parameters in the POST body:  | PARAMETER | DESCRIPTION | |-----------|-------------| | client_id | Your app's client ID | | client_secret | Your app's client secret | | code | The code you just received from the redirect |  You'll get a reponse like this:  ```json {   \"scope\": \"linodes:create\",   \"access_token\": \"03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c\"   \"token_type\": \"bearer\",   \"expires_in\": 7200, } ```  Included in the reponse is `access_token`. With this token, you can proceed to make authenticated HTTP requests to the API by adding this header to each request:  ``` Authorization: Bearer 03d084436a6c91fbafd5c4b20c82e5056a2e9ce1635920c30dc8d81dc7a6665c ```  ### Authentication  | Security Scheme Type: | Oauth2 | |-----------------------|--------| | **AuthorizationCode Oauth Flow** |  **Authorization URL:** https://login.linode.com/oauth/authorize<br />**Token URL:** https://login.linode.com/oauth/token<br />**Scopes:**<br /><ul><li>`account:read_only` - Allows access to GET information about your Account.</li><li>`account:read_write` - Allows access to all endpoints related to your Account.</li><li>`domains:read_only` - Allows access to GET Domains on your Account.</li><li>`domains:read_write` - Allows access to all Domain endpoints.</li><li>`events:read_only` - Allows access to GET your Events.</li><li>`events:read_write` - Allows access to all endpoints related to your Events.</li><li>`images:read_only` - Allows access to GET your Images.</li><li>`images:read_write` - Allows access to all endpoints related to your Images.</li><li>`ips:read_only` - Allows access to GET your ips.</li><li>`ips:read_write` - Allows access to all endpoints related to your ips.</li><li>`linodes:read_only` - Allows access to GET Linodes on your Account.</li><li>`linodes:read_write` - Allow access to all endpoints related to your Linodes.</li><li>`longview:read_only` - Allows access to GET your Longview Clients.</li><li>`longview:read_write` - Allows access to all endpoints related to your Longview Clients.</li><li>`nodebalancers:read_only` - Allows access to GET NodeBalancers on your Account.</li><li>`nodebalancers:read_write` - Allows access to all NodeBalancer endpoints.</li><li>`stackscripts:read_only` - Allows access to GET your StackScripts.</li><li>`stackscripts:read_write` - Allows access to all endpoints related to your StackScripts.</li><li>`volumes:read_only` - Allows access to GET your Volumes.</li><li>`volumes:read_write` - Allows access to all endpoints related to your Volumes.</li></ul><br />|  # Requests  Requests must be made over HTTPS to ensure transactions are encrypted. The following Request methods are supported:  | METHOD | USAGE | |--------|-------| | GET    | Retrieves data about collections and individual resources. | | POST   | For collections, creates a new resource of that type. Also used to perform actions on action endpoints. | | PUT    | Updates an existing resource. | | DELETE | Deletes a resource. This is a destructive action. |   # Responses  Actions will return one following HTTP response status codes:  | STATUS  | DESCRIPTION | |---------|-------------| | 200 OK  | The request was successful. | | 204 No Content | The server successfully fulfilled the request and there is no additional content to send. | | 400 Bad Request | You submitted an invalid request (missing parameters, etc.). | | 401 Unauthorized | You failed to authenticate for this resource. | | 403 Forbidden | You are authenticated, but don't have permission to do this. | | 404 Not Found | The resource you're requesting does not exist. | | 429 Too Many Requests | You've hit a rate limit. | | 500 Internal Server Error | Please [open a Support Ticket](/#operation/createTicket). |  # Errors  Success is indicated via [Standard HTTP status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes). `2xx` codes indicate success, `4xx` codes indicate a request error, and `5xx` errors indicate a server error. A request error might be an invalid input, a required parameter being omitted, or a malformed request. A server error means something went wrong processing your request. If this occurs, please [open a Support Ticket](/#operation/createTicket) and let us know. Though errors are logged and we work quickly to resolve issues, opening a ticket and providing us with reproducable steps and data is always helpful.  The `errors` field is an array of the things that went wrong with your request. We will try to include as many of the problems in the response as possible, but it's conceivable that fixing these errors and resubmitting may result in new errors coming back once we are able to get further along in the process of handling your request.   Within each error object, the `field` parameter will be included if the error pertains to a specific field in the JSON you've submitted. This will be omitted if there is no relevant field. The `reason` is a human-readable explanation of the error, and will always be included.  # Pagination  Resource lists are always paginated. The response will look similar to this:  ```json {     \"data\": [ ... ],     \"page\": 1,     \"pages\": 3,     \"results\": 300 } ```  Pages start at 1. You may retrieve a specific page of results by adding `?page=x` to your URL (for example, `?page=4`). Page sizes default to 100, and can be set to return between 25 and 100. Page size can be set using `?page_size=x`.  # Filtering and Sorting  Collections are searchable by fields they include, marked in the spec as `x-linode-filterable: true`. Filters are passed in the `X-Filter` header and are formatted as JSON objects. Here is a request call for Linode Types in our \"standard\" class:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H 'X-Filter: { \\     \"class\": \"standard\"   }' ```  The filter object's keys are the keys of the object you're filtering, and the values are accepted values. You can add multiple filters by including more than one key. For example, filtering for \"standard\" Linode Types that offer one vcpu:  ```Shell  curl \"https://api.linode.com/v4/linode/types\" \\   -H 'X-Filter: { \\     \"class\": \"standard\",     \"vcpus\": 1   }' ```  In the above example, both filters are combined with an \"and\" operation. However, if you wanted either Types with one vcpu or Types in our \"standard\" class, you can add an operator:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H 'X-Filter: {     \"+or\": [       { \"vcpus\": 1 },       { \"class\": \"standard\" }     ]   }' ```  Each filter in the `+or` array is its own filter object, and all conditions in it are combined with an \"and\" operation as they were in the previous example.  Other operators are also available. Operators are keys of a Filter JSON object. Their value must be of the appropriate type, and they are evaluated as described below:  | OPERATOR | TYPE   | DESCRIPTION                       | |----------|--------|-----------------------------------| | +and     | array  | All conditions must be true.       | | +or      | array  | One condition must be true.        | | +gt      | number | Value must be greater than number. | | +gte     | number | Value must be greater than or equal to number. | | +lt      | number | Value must be less than number. | | +lte     | number | Value must be less than or equal to number. | | +contains | string | Given string must be in the value. | | +neq      | string | Does not equal the value.          | | +order_by | string | Attribute to order the results by - must be filterable. | | +order    | string | Either \"asc\" or \"desc\". Defaults to \"asc\". Requires `+order_by`. |  For example, filtering for [Linode Types](/#operation/getLinodeTypes) that offer memory equal to or higher than 61440:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H 'X-Filter: {     \"memory\": {       \"+gte\": 61440     }   }' ```  You can combine and nest operators to construct arbitrarily-complex queries. For example, give me all [Linode Types](/#operation/getLinodeTypes) which are either `standard` or `highmem` class, and have between 12 and 20 vcpus:  ```Shell curl \"https://api.linode.com/v4/linode/types\" \\   -H 'X-Filter: {     \"+or\": [       {         \"+or\": [           {             \"class\": \"standard\"           },           {             \"class\": \"highmem\"           }         ]       },       {         \"+and\": [           {             \"vcpus\": {               \"+gte\": 12             }           },           {             \"vcpus\": {               \"+lte\": 20             }           }         ]       }     ]   }' ```  # CLI (Command Line Interface)  The [Linode CLI](https://github.com/linode/linode-cli) allows you to easily work with the API using intuitive and simple syntax. It requires a [Personal Access Token](/#section/Personal-Access-Token) for authentication, and gives you access to all of the features and functionality of the Linode API that are documented here with CLI examples.  Endpoints that do not have CLI examples are currently unavailable through the CLI, but can be accessed via other methods such as Shell commands and other third-party applications. 
 *
 * API version: 4.0.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// A Personal Access Token is a token generated manually to access the API without going through an OAuth login.  Personal Access Tokens can have scopes just like OAuth tokens do, and are commonly used to give access to command-line tools like the Linode CLI, or when writing your own integrations. 
type PersonalAccessToken struct {
	// This token's unique ID, which can be used to revoke it. 
	Id int32 `json:"id,omitempty"`
	// The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the [Linode CLI](https://github.com/linode/linode-cli), require tokens with access to `*`. Tokens with more restrictive scopes are generally more secure. 
	Scopes string `json:"scopes,omitempty"`
	// The date and time this token was created. 
	Created time.Time `json:"created,omitempty"`
	// This token's label.  This is for display purposes only, but can be used to more easily track what you're using each token for. 
	Label string `json:"label,omitempty"`
	// The token used to access the API.  When the token is created, the full token is returned here.  Otherwise, only the first 16 characters are returned. 
	Token string `json:"token,omitempty"`
	// When this token will expire.  Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated.  Tokens may be created with \"null\" as their expiry and will never expire unless revoked. 
	Expiry time.Time `json:"expiry,omitempty"`
}
