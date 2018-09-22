# NodeBalancerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | This config&#39;s unique ID | [optional] 
**Port** | **int32** | The port this Config is for. These values must be unique across configs on a single NodeBalancer (you can&#39;t have two configs for port 80, for example).  While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.  | [optional] 
**Protocol** | **string** | The protocol this port is configured to serve. * If this is set to &#x60;https&#x60; you must include an &#x60;ssl_cert&#x60; and an &#x60;ssl_key&#x60;.  | [optional] 
**Algorithm** | **string** | What algorithm this NodeBalancer should use for routing traffic to backends.  | [optional] 
**Stickiness** | **string** | Controls how session stickiness is handled on this port. * If set to &#x60;none&#x60; connections will always be assigned a backend based on the algorithm configured. * If set to &#x60;table&#x60; sessions from the same remote address will be routed to the same   backend.  * For HTTP or HTTPS clients, &#x60;http_cookie&#x60; allows sessions to be   routed to the same backend based on a cookie set by the NodeBalancer.  | [optional] 
**Check** | **string** | The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. * If &#x60;none&#x60; no check is performed. * &#x60;connection&#x60; requires only a connection to the backend to succeed. * &#x60;http&#x60; and &#x60;http_body&#x60; rely on the backend serving HTTP, and that   the response returned matches what is expected.  | [optional] 
**CheckInterval** | **int32** | How often, in seconds, to check that backends are up and serving requests.  | [optional] 
**CheckTimeout** | **int32** | How long, in seconds, to wait for a check attempt before considering it failed.  | [optional] 
**CheckAttempts** | **int32** | How many times to attempt a check before considering a backend to be down.  | [optional] 
**CheckPath** | **string** | The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.  | [optional] 
**CheckBody** | **string** | This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down.  | [optional] 
**CheckPassive** | **bool** | If true, any response from this backend with a &#x60;5xx&#x60; status code will be enough for it to be considered unhealthy and taken out of rotation.  | [optional] 
**CipherSuite** | **string** | What ciphers to use for SSL connections served by this NodeBalancer.  * &#x60;legacy&#x60; is considered insecure and should only be used if necessary.  | [optional] 
**NodebalancerId** | **int32** | The ID for the NodeBalancer this config belongs to.  | [optional] 
**SslCommonname** | **string** | The common name for the SSL certification this port is serving if this port is not configured to use SSL.  | [optional] 
**SslFingerprint** | **string** | The fingerprint for the SSL certification this port is serving if this port is not configured to use SSL.  | [optional] 
**SslCert** | **string** | The certificate this port is serving. This is not returned. If set, this field will come back as \&quot;&amp;lt;REDACTED&amp;gt;\&quot;. Please use the &#x60;ssl_commonname&#x60; and &#x60;ssl_fingerprint&#x60; to identify the certificate.  | [optional] 
**SslKey** | **string** | The private key corresponding to this port&#39;s certificate.  This is not returned. If set, this field will come back as \&quot;&amp;lt;REDACTED&amp;gt;\&quot;. Please use the &#x60;ssl_commonname&#x60; and &#x60;ssl_fingerprint&#x60; to identify the certificate.  | [optional] 
**NodesStatus** | [**NodeBalancerConfigNodesStatus**](NodeBalancerConfig_nodes_status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


