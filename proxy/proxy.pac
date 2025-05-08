function FindProxyForURL(url, host) {
  // our local URLs from the domains below example.com don't need a proxy:
  if (shExpMatch(host, "*.vz.bos2.lab")) {
    return "PROXY 192.168.13.11:3128";
  }

  // URLs within this network are accessed through
  // port 8080 on fastproxy.example.com:
  if (isInNet(host, "192.168.13.0.", "255.255.255.128")) {
    return "PROXY 192.168.13.11:3128";
  }

  // All other requests go through port 8080 of proxy.example.com.
  // should that fail to respond, go directly to the WWW:
 if (isInNet(host, "192.168.33.61")){
     return "PROXY proxy.example.com:8080; DIRECT";
 }

  return "PROXY proxy.example.com:8080; DIRECT";
}
