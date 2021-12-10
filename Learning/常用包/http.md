# net

net：为网络I/O提供了一个可移植的接口，包括TCP/IP、UDP、域名解析和Unix域套接字。

# net/http

http是net的子包

http包：提供HTTP客户端和服务器实现

    resp, err := http.Get("http://example.com/")
    ...
    resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
    ...
    resp, err := http.PostForm("http://example.com/form",
        url.Values{"key": {"Value"}, "id": {"123"}})

用完这个clien必须关掉response的body：

    resp, err := http.Get("http://example.com/")
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    // ...

For control over HTTP client headers, redirect policy, and other settings, create a Client:

    client := &http.Client{
        CheckRedirect: redirectPolicyFunc,
    }

    resp, err := client.Get("http://example.com")
    // ...

    req, err := http.NewRequest("GET", "http://example.com", nil)
    // ...
    req.Header.Add("If-None-Match", `W/"wyzzy"`)
    resp, err := client.Do(req)
    // ...

For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:

    tr := &http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    30 * time.Second,
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://example.com")


# net/http/httputil

httputil是http的子包

httputil提供了HTTP实用功能，补充了 net/http 包中比较常用的功能。

