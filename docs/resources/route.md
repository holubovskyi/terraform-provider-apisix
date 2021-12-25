---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "apisix_route Resource - terraform-provider-apisix"
subcategory: ""
description: |-
  
---

# apisix_route (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String)

### Optional

- **desc** (String)
- **enable_websocket** (Boolean)
- **filter_func** (String)
- **host** (String)
- **hosts** (List of String)
- **is_enabled** (Boolean)
- **labels** (Map of String)
- **methods** (List of String)
- **plugin_config_id** (String)
- **plugins** (Attributes) (see [below for nested schema](#nestedatt--plugins))
- **priority** (Number)
- **remote_addr** (String)
- **remote_addrs** (List of String)
- **script** (String)
- **service_id** (String)
- **timeout** (Attributes) (see [below for nested schema](#nestedatt--timeout))
- **upstream** (Attributes) (see [below for nested schema](#nestedatt--upstream))
- **upstream_id** (String)
- **uri** (String)
- **uris** (List of String)

### Read-Only

- **id** (String) The ID of this resource.

<a id="nestedatt--plugins"></a>
### Nested Schema for `plugins`

Optional:

- **cors** (Attributes) (see [below for nested schema](#nestedatt--plugins--cors))
- **ip_restriction** (Attributes) (see [below for nested schema](#nestedatt--plugins--ip_restriction))
- **prometheus** (Attributes) (see [below for nested schema](#nestedatt--plugins--prometheus))
- **proxy_rewrite** (Attributes) (see [below for nested schema](#nestedatt--plugins--proxy_rewrite))
- **redirect** (Attributes) (see [below for nested schema](#nestedatt--plugins--redirect))
- **request_id** (Attributes) (see [below for nested schema](#nestedatt--plugins--request_id))
- **serverless_post_function** (Attributes) (see [below for nested schema](#nestedatt--plugins--serverless_post_function))
- **serverless_pre_function** (Attributes) (see [below for nested schema](#nestedatt--plugins--serverless_pre_function))

<a id="nestedatt--plugins--cors"></a>
### Nested Schema for `plugins.cors`

Optional:

- **allow_credential** (Boolean)
- **allow_headers** (List of String) Which headers are allowed to set in request when access cross-origin resource. Multiple value use , to split. When allow_credential is false, you can use * to indicate allow all request headers. You also can allow any header forcefully using ** even already enable allow_credential, but it will bring some security risks.
- **allow_methods** (List of String) Which Method is allowed to enable CORS, such as: GET, POST etc. Multiple method use , to split. When allow_credential is false, you can use * to indicate allow all any method. You also can allow any method forcefully using ** even already enable allow_credential, but it will bring some security risks.
- **allow_origins** (List of String) Which Origins is allowed to enable CORS, format as: scheme://host:port, for example: https://somehost.com:8081. Multiple origin use , to split. When allow_credential is false, you can use * to indicate allow any origin. you also can allow all any origins forcefully using ** even already enable allow_credential, but it will bring some security risks.
- **allow_origins_by_regex** (List of String) Use regex expressions to match which origin is allowed to enable CORS, for example, [".*.test.com"] can use to match all subdomain of test.com
- **disable** (Boolean)
- **expose_headers** (List of String) Which headers are allowed to set in response when access cross-origin resource. Multiple value use , to split. When allow_credential is false, you can use * to indicate allow any header. You also can allow any header forcefully using ** even already enable allow_credential, but it will bring some security risks.
- **max_age** (Number) Maximum number of seconds the results can be cached. Within this time range, the browser will reuse the last check result. -1 means no cache. Please note that the maximum value is depended on browser, please refer to MDN for details.


<a id="nestedatt--plugins--ip_restriction"></a>
### Nested Schema for `plugins.ip_restriction`

Optional:

- **blacklist** (List of String)
- **disable** (Boolean)
- **message** (String)
- **whitelist** (List of String)


<a id="nestedatt--plugins--prometheus"></a>
### Nested Schema for `plugins.prometheus`

Optional:

- **disable** (Boolean)
- **prefer_name** (Boolean)


<a id="nestedatt--plugins--proxy_rewrite"></a>
### Nested Schema for `plugins.proxy_rewrite`

Optional:

- **disable** (Boolean)
- **headers** (Map of String)
- **host** (String)
- **method** (String)
- **regex_uri** (Attributes) (see [below for nested schema](#nestedatt--plugins--proxy_rewrite--regex_uri))
- **scheme** (String)
- **uri** (String)

<a id="nestedatt--plugins--proxy_rewrite--regex_uri"></a>
### Nested Schema for `plugins.proxy_rewrite.regex_uri`

Optional:

- **regex** (String)
- **replacement** (String)



<a id="nestedatt--plugins--redirect"></a>
### Nested Schema for `plugins.redirect`

Optional:

- **append_query_string** (Boolean) When set to true, add the query string from the original request to the location header. If the configured uri / regex_uri already contains a query string, the query string from request will be appended to that after an &. Caution: don't use this if you've already handled the query string, e.g. with nginx variable $request_uri, to avoid duplicates.
- **disable** (Boolean)
- **encode_uri** (Boolean) When set to true the uri in Location header will be encoded as per RFC3986
- **http_to_https** (Boolean) When it is set to true and the request is HTTP, will be automatically redirected to HTTPS with 301 response code, and the URI will keep the same as client request
- **regex_uri** (Attributes) (see [below for nested schema](#nestedatt--plugins--redirect--regex_uri))
- **ret_code** (Number) Response code
- **uri** (String) New URL which can contain Nginx variable, eg: /test/index.html, $uri/index.html. You can refer to variables in a way similar to ${xxx} to avoid ambiguity, eg: ${uri}foo/index.html. If you just need the original $ character, add \ in front of it, like this one: /\$foo/index.html. If you refer to a variable name that does not exist, this will not produce an error, and it will be used as an empty string

<a id="nestedatt--plugins--redirect--regex_uri"></a>
### Nested Schema for `plugins.redirect.regex_uri`

Optional:

- **regex** (String)
- **replacement** (String)



<a id="nestedatt--plugins--request_id"></a>
### Nested Schema for `plugins.request_id`

Optional:

- **algorithm** (String)
- **disable** (Boolean)
- **header_name** (String)
- **include_in_response** (Boolean)


<a id="nestedatt--plugins--serverless_post_function"></a>
### Nested Schema for `plugins.serverless_post_function`

Optional:

- **disable** (Boolean)
- **functions** (List of String)
- **phase** (String)


<a id="nestedatt--plugins--serverless_pre_function"></a>
### Nested Schema for `plugins.serverless_pre_function`

Optional:

- **disable** (Boolean)
- **functions** (List of String)
- **phase** (String)



<a id="nestedatt--timeout"></a>
### Nested Schema for `timeout`

Optional:

- **connect** (Number)
- **read** (Number)
- **send** (Number)


<a id="nestedatt--upstream"></a>
### Nested Schema for `upstream`

Optional:

- **checks** (Attributes) (see [below for nested schema](#nestedatt--upstream--checks))
- **desc** (String)
- **discovery_type** (String)
- **hash_on** (String)
- **id** (String) The ID of this resource.
- **keepalive_pool** (Attributes) (see [below for nested schema](#nestedatt--upstream--keepalive_pool))
- **labels** (Map of String)
- **name** (String)
- **nodes** (Attributes List) (see [below for nested schema](#nestedatt--upstream--nodes))
- **pass_host** (String)
- **retries** (Number)
- **retry_timeout** (Number)
- **scheme** (String)
- **service_name** (String)
- **timeout** (Attributes) (see [below for nested schema](#nestedatt--upstream--timeout))
- **tls** (Attributes) (see [below for nested schema](#nestedatt--upstream--tls))
- **type** (String)
- **upstream_host** (String)

<a id="nestedatt--upstream--checks"></a>
### Nested Schema for `upstream.checks`

Optional:

- **active** (Attributes) (see [below for nested schema](#nestedatt--upstream--checks--active))
- **passive** (Attributes) (see [below for nested schema](#nestedatt--upstream--checks--passive))

<a id="nestedatt--upstream--checks--active"></a>
### Nested Schema for `upstream.checks.active`

Optional:

- **concurrency** (Number) The number of targets to be checked at the same time during the active check
- **healthy** (Attributes) (see [below for nested schema](#nestedatt--upstream--checks--active--healthy))
- **host** (String) The hostname of the HTTP request actively checked
- **http_path** (String) The HTTP request path that is actively checked
- **https_verify_certificate** (Boolean) Active check whether to check the SSL certificate of the remote host when HTTPS type checking is used
- **port** (Number) The host port of the HTTP request that is actively checked
- **req_headers** (List of String) Active check When using HTTP or HTTPS type checking, set additional request header information
- **timeout** (Number) The timeout period of the active check (unit: second)
- **type** (String) The type of active check
- **unhealthy** (Attributes) (see [below for nested schema](#nestedatt--upstream--checks--active--unhealthy))

<a id="nestedatt--upstream--checks--active--healthy"></a>
### Nested Schema for `upstream.checks.active.unhealthy`

Optional:

- **http_statuses** (List of Number) Active check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node
- **interval** (Number) Active check (healthy node) check interval (unit: second)
- **successes** (Number) Active check (healthy node) check interval (unit: second)


<a id="nestedatt--upstream--checks--active--unhealthy"></a>
### Nested Schema for `upstream.checks.active.unhealthy`

Optional:

- **http_failures** (Number) Active check (unhealthy node) HTTP or HTTPS type check, determine the number of times that the node is not healthy
- **http_statuses** (List of Number) Active check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node
- **interval** (Number) Active check (unhealthy node) check interval (unit: second)
- **tcp_failures** (Number) Active check (unhealthy node) TCP type check, determine the number of times that the node is not healthy
- **timeouts** (Number) Active check (unhealthy node) to determine the number of timeouts for unhealthy nodes



<a id="nestedatt--upstream--checks--passive"></a>
### Nested Schema for `upstream.checks.passive`

Optional:

- **healthy** (Attributes) (see [below for nested schema](#nestedatt--upstream--checks--passive--healthy))
- **unhealthy** (Attributes) (see [below for nested schema](#nestedatt--upstream--checks--passive--unhealthy))

<a id="nestedatt--upstream--checks--passive--healthy"></a>
### Nested Schema for `upstream.checks.passive.unhealthy`

Optional:

- **http_statuses** (List of Number) Passive check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node
- **successes** (Number) Passive checks (healthy node) determine the number of times a node is healthy


<a id="nestedatt--upstream--checks--passive--unhealthy"></a>
### Nested Schema for `upstream.checks.passive.unhealthy`

Optional:

- **http_failures** (Number) Passive check (unhealthy node) The number of times that the node is not healthy during HTTP or HTTPS type checking
- **http_statuses** (List of Number) Passive check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node
- **tcp_failures** (Number) Passive check (unhealthy node) When TCP type is checked, determine the number of times that the node is not healthy
- **timeouts** (Number) Passive checks (unhealthy node) determine the number of timeouts for unhealthy nodes




<a id="nestedatt--upstream--keepalive_pool"></a>
### Nested Schema for `upstream.keepalive_pool`

Optional:

- **idle_timeout** (Number)
- **requests** (Number)
- **size** (Number)


<a id="nestedatt--upstream--nodes"></a>
### Nested Schema for `upstream.nodes`

Optional:

- **host** (String)
- **port** (Number)
- **weight** (Number)


<a id="nestedatt--upstream--timeout"></a>
### Nested Schema for `upstream.timeout`

Optional:

- **connect** (Number)
- **read** (Number)
- **send** (Number)


<a id="nestedatt--upstream--tls"></a>
### Nested Schema for `upstream.tls`

Optional:

- **client_cert** (String)
- **client_key** (String)

