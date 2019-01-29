package main

import (
    "fmt"
    "crypto/tls"
    "net/http"
    "strings"	
)

const (
    PORT       = ":8443"
    PRIV_KEY   = "./private.pem"
    PUBLIC_KEY = "./public.pem"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Nobody should read this.")
    if r.TLS != nil && len(r.TLS.PeerCertificates) > 0 {
      cn := strings.ToLower(r.TLS.PeerCertificates[0].Subject.CommonName)
      fmt.Println("CN: ", cn)
      fmt.Println("DNSNames:")
      fmt.Println(r.TLS.PeerCertificates[0].DNSNames)
      fmt.Println("EmailAddresses:")
      fmt.Println(r.TLS.PeerCertificates[0].EmailAddresses)
      fmt.Println("IPAddresses:")
      fmt.Println(r.TLS.PeerCertificates[0].IPAddresses)
      fmt.Println("URIs:")
      fmt.Println(r.TLS.PeerCertificates[0].URIs)
    }
}

func main() {

    server := &http.Server{
        TLSConfig: &tls.Config{
            ClientAuth: tls.RequireAnyClientCert,
            MinVersion: tls.VersionTLS12,
        },
        Addr: "127.0.0.1:8443",
    }

    http.HandleFunc("/", rootHandler)
    err := server.ListenAndServeTLS(PUBLIC_KEY, PRIV_KEY)
    if err != nil {
        fmt.Printf("main(): %s\n", err)
    }
}
