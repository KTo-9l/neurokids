
# Юзать у себя в хандлере так:
```go
	perm, err := auth_helpers.GetPerms[auth_helpers.Perms](req)
	if err != nil {
		response = suckhttp.NewResponse(403, "forbidden")
		return
	}
	if p, ok := perm.Perms[0]; !ok || p&auth_helpers.AllPerms == 0 {
		response = suckhttp.NewResponse(403, "forbidden")
		return
	}
```

# nginx

```nginx
upstream auth {
	server 127.0.0.1:9213;
	keepalive 20;
	keepalive_timeout 24h;
}

underscores_in_headers on;
proxy_pass_request_headers      on;

location = /auth {
   	proxy_pass http://auth/;
   	proxy_pass_request_body off;
   	proxy_set_header Content-Length "";
   	proxy_set_header X-Original-URI $request_uri;
	proxy_http_version 1.1;
	proxy_set_header Connection "";
	proxy_set_header Host $host;
	proxy_set_header X-Real-IP $remote_addr;
	proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	proxy_set_header X-Forwarded-Proto $scheme;
	proxy_set_header X-Request-Id $request_id;
}

location / {
	auth_request /auth;
	auth_request_set $perm $upstream_http_x_perm;
	proxy_set_header x-perm $perm;
}
```
