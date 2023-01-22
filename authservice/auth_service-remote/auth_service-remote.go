// Code generated by Thrift Compiler (0.18.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"talkservice"
	"authservice"
)

var _ = talkservice.GoUnusedProtection__
var _ = authservice.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  LoginResult loginZ(LoginRequest loginRequest)")
  fmt.Fprintln(os.Stderr, "  string confirmE2EELogin(string verifier, string deviceSecret)")
  fmt.Fprintln(os.Stderr, "  void respondE2EELoginRequest(string verifier, E2EEPublicKey publicKey, string encryptedKeyChain, string hashKeyChain, ErrorCode errorCode)")
  fmt.Fprintln(os.Stderr, "  string openAuthSession(AuthSessionRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse updatePassword(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  void logoutZ()")
  fmt.Fprintln(os.Stderr, "  string verifyQrcodeWithE2EE(string verifier, string pinCode, ErrorCode errorCode, E2EEPublicKey publicKey, string encryptedKeyChain, string hashKeyChain)")
  fmt.Fprintln(os.Stderr, "  RSAKey getAuthRSAKey(string authSessionId, IdentityProvider identityProvider)")
  fmt.Fprintln(os.Stderr, "  SecurityCenterResult issueTokenForAccountMigrationSettings(bool enforce)")
  fmt.Fprintln(os.Stderr, "  SetPasswordResponse setPassword(string authSessionId, EncryptedPassword encryptedPassword)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse confirmIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse setIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse setIdentifierAndPassword(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse updateIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse resendIdentifierConfirmation(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr, "  IdentityCredentialResponse removeIdentifier(string authSessionId, IdentityCredentialRequest request)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  var cfg *thrift.TConfiguration = nil
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransportConf(trans, cfg)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := authservice.NewAuthServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "loginZ":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "LoginZ requires 1 args")
      flag.Usage()
    }
    arg105 := flag.Arg(1)
    mbTrans106 := thrift.NewTMemoryBufferLen(len(arg105))
    defer mbTrans106.Close()
    _, err107 := mbTrans106.WriteString(arg105)
    if err107 != nil {
      Usage()
      return
    }
    factory108 := thrift.NewTJSONProtocolFactory()
    jsProt109 := factory108.GetProtocol(mbTrans106)
    argvalue0 := authservice.NewLoginRequest()
    err110 := argvalue0.Read(context.Background(), jsProt109)
    if err110 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.LoginZ(context.Background(), value0))
    fmt.Print("\n")
    break
  case "confirmE2EELogin":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ConfirmE2EELogin requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ConfirmE2EELogin(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "respondE2EELoginRequest":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "RespondE2EELoginRequest requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg114 := flag.Arg(2)
    mbTrans115 := thrift.NewTMemoryBufferLen(len(arg114))
    defer mbTrans115.Close()
    _, err116 := mbTrans115.WriteString(arg114)
    if err116 != nil {
      Usage()
      return
    }
    factory117 := thrift.NewTJSONProtocolFactory()
    jsProt118 := factory117.GetProtocol(mbTrans115)
    argvalue1 := talkservice.NewE2EEPublicKey()
    err119 := argvalue1.Read(context.Background(), jsProt118)
    if err119 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    tmp4, err := (strconv.Atoi(flag.Arg(5)))
    if err != nil {
      Usage()
     return
    }
    argvalue4 := authservice.ErrorCode(tmp4)
    value4 := argvalue4
    fmt.Print(client.RespondE2EELoginRequest(context.Background(), value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "openAuthSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenAuthSession requires 1 args")
      flag.Usage()
    }
    arg122 := flag.Arg(1)
    mbTrans123 := thrift.NewTMemoryBufferLen(len(arg122))
    defer mbTrans123.Close()
    _, err124 := mbTrans123.WriteString(arg122)
    if err124 != nil {
      Usage()
      return
    }
    factory125 := thrift.NewTJSONProtocolFactory()
    jsProt126 := factory125.GetProtocol(mbTrans123)
    argvalue0 := authservice.NewAuthSessionRequest()
    err127 := argvalue0.Read(context.Background(), jsProt126)
    if err127 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenAuthSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "updatePassword":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdatePassword requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg129 := flag.Arg(2)
    mbTrans130 := thrift.NewTMemoryBufferLen(len(arg129))
    defer mbTrans130.Close()
    _, err131 := mbTrans130.WriteString(arg129)
    if err131 != nil {
      Usage()
      return
    }
    factory132 := thrift.NewTJSONProtocolFactory()
    jsProt133 := factory132.GetProtocol(mbTrans130)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err134 := argvalue1.Read(context.Background(), jsProt133)
    if err134 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdatePassword(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "logoutZ":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "LogoutZ requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.LogoutZ(context.Background()))
    fmt.Print("\n")
    break
  case "verifyQrcodeWithE2EE":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "VerifyQrcodeWithE2EE requires 6 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err := (strconv.Atoi(flag.Arg(3)))
    if err != nil {
      Usage()
     return
    }
    argvalue2 := authservice.ErrorCode(tmp2)
    value2 := argvalue2
    arg137 := flag.Arg(4)
    mbTrans138 := thrift.NewTMemoryBufferLen(len(arg137))
    defer mbTrans138.Close()
    _, err139 := mbTrans138.WriteString(arg137)
    if err139 != nil {
      Usage()
      return
    }
    factory140 := thrift.NewTJSONProtocolFactory()
    jsProt141 := factory140.GetProtocol(mbTrans138)
    argvalue3 := talkservice.NewE2EEPublicKey()
    err142 := argvalue3.Read(context.Background(), jsProt141)
    if err142 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.VerifyQrcodeWithE2EE(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "getAuthRSAKey":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetAuthRSAKey requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err := (strconv.Atoi(flag.Arg(2)))
    if err != nil {
      Usage()
     return
    }
    argvalue1 := authservice.IdentityProvider(tmp1)
    value1 := argvalue1
    fmt.Print(client.GetAuthRSAKey(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "issueTokenForAccountMigrationSettings":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "IssueTokenForAccountMigrationSettings requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1) == "true"
    value0 := argvalue0
    fmt.Print(client.IssueTokenForAccountMigrationSettings(context.Background(), value0))
    fmt.Print("\n")
    break
  case "setPassword":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetPassword requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg148 := flag.Arg(2)
    mbTrans149 := thrift.NewTMemoryBufferLen(len(arg148))
    defer mbTrans149.Close()
    _, err150 := mbTrans149.WriteString(arg148)
    if err150 != nil {
      Usage()
      return
    }
    factory151 := thrift.NewTJSONProtocolFactory()
    jsProt152 := factory151.GetProtocol(mbTrans149)
    argvalue1 := authservice.NewEncryptedPassword()
    err153 := argvalue1.Read(context.Background(), jsProt152)
    if err153 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.SetPassword(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "confirmIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ConfirmIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg155 := flag.Arg(2)
    mbTrans156 := thrift.NewTMemoryBufferLen(len(arg155))
    defer mbTrans156.Close()
    _, err157 := mbTrans156.WriteString(arg155)
    if err157 != nil {
      Usage()
      return
    }
    factory158 := thrift.NewTJSONProtocolFactory()
    jsProt159 := factory158.GetProtocol(mbTrans156)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err160 := argvalue1.Read(context.Background(), jsProt159)
    if err160 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.ConfirmIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "setIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg162 := flag.Arg(2)
    mbTrans163 := thrift.NewTMemoryBufferLen(len(arg162))
    defer mbTrans163.Close()
    _, err164 := mbTrans163.WriteString(arg162)
    if err164 != nil {
      Usage()
      return
    }
    factory165 := thrift.NewTJSONProtocolFactory()
    jsProt166 := factory165.GetProtocol(mbTrans163)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err167 := argvalue1.Read(context.Background(), jsProt166)
    if err167 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.SetIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "setIdentifierAndPassword":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetIdentifierAndPassword requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg169 := flag.Arg(2)
    mbTrans170 := thrift.NewTMemoryBufferLen(len(arg169))
    defer mbTrans170.Close()
    _, err171 := mbTrans170.WriteString(arg169)
    if err171 != nil {
      Usage()
      return
    }
    factory172 := thrift.NewTJSONProtocolFactory()
    jsProt173 := factory172.GetProtocol(mbTrans170)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err174 := argvalue1.Read(context.Background(), jsProt173)
    if err174 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.SetIdentifierAndPassword(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "updateIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg176 := flag.Arg(2)
    mbTrans177 := thrift.NewTMemoryBufferLen(len(arg176))
    defer mbTrans177.Close()
    _, err178 := mbTrans177.WriteString(arg176)
    if err178 != nil {
      Usage()
      return
    }
    factory179 := thrift.NewTJSONProtocolFactory()
    jsProt180 := factory179.GetProtocol(mbTrans177)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err181 := argvalue1.Read(context.Background(), jsProt180)
    if err181 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdateIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "resendIdentifierConfirmation":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ResendIdentifierConfirmation requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg183 := flag.Arg(2)
    mbTrans184 := thrift.NewTMemoryBufferLen(len(arg183))
    defer mbTrans184.Close()
    _, err185 := mbTrans184.WriteString(arg183)
    if err185 != nil {
      Usage()
      return
    }
    factory186 := thrift.NewTJSONProtocolFactory()
    jsProt187 := factory186.GetProtocol(mbTrans184)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err188 := argvalue1.Read(context.Background(), jsProt187)
    if err188 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.ResendIdentifierConfirmation(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "removeIdentifier":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "RemoveIdentifier requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg190 := flag.Arg(2)
    mbTrans191 := thrift.NewTMemoryBufferLen(len(arg190))
    defer mbTrans191.Close()
    _, err192 := mbTrans191.WriteString(arg190)
    if err192 != nil {
      Usage()
      return
    }
    factory193 := thrift.NewTJSONProtocolFactory()
    jsProt194 := factory193.GetProtocol(mbTrans191)
    argvalue1 := authservice.NewIdentityCredentialRequest()
    err195 := argvalue1.Read(context.Background(), jsProt194)
    if err195 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.RemoveIdentifier(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}