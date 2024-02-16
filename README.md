[![Go](https://github.com/rthier/study-newhttprouting/actions/workflows/go.yml/badge.svg)](https://github.com/rthier/study-newhttprouting/actions/workflows/go.yml)

> ### Enhanced routing patterns
> HTTP routing in the standard library is now more expressive. The patterns used by net/http.ServeMux have been enhanced to accept methods and wildcards.
> 
> Registering a handler with a method, like "POST /items/create", restricts invocations of the handler to requests with the given method. A pattern with a method takes precedence over a matching pattern without one. As a special case, registering a handler with "GET" also registers it with "HEAD".
> 
> Wildcards in patterns, like /items/{id}, match segments of the URL path. The actual segment value may be accessed by calling the Request.PathValue method. A wildcard ending in "...", like /files/{path...}, must occur at the end of a pattern and matches all the remaining segments.
> 
> A pattern that ends in "/" matches all paths that have it as a prefix, as always. To match the exact pattern including the trailing slash, end it with {$}, as in /exact/match/{$}.
> 
> If two patterns overlap in the requests that they match, then the more specific pattern takes precedence. If neither is more specific, the patterns conflict. This rule generalizes the original precedence rules and maintains the property that the order in which patterns are registered does not matter.
> 
> This change breaks backwards compatibility in small ways, some obvious—patterns with "{" and "}" behave differently— and some less so—treatment of escaped paths has been improved. The change is controlled by a GODEBUG field named httpmuxgo121. Set httpmuxgo121=1 to restore the old behavior.

Source: https://go.dev/doc/go1.22#enhanced_routing_patterns
