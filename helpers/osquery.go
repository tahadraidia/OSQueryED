package helpers

import (
  "fmt"
  "github.com/osquery/osquery-go/gen/osquery"
)

func IsOSQueryDeamonRunning() bool {
  return  DoesNamedPipeExist("osquery.em")
}

func IsOSQueryInterpreterRunning() bool {
  return  DoesNamedPipeExist("shell.em")
}

func ParseOSQueryResponse(response osquery.ExtensionPluginResponse) {
  for key, element := range response {
    //fmt.Println("Key:", key)
    //fmt.Println(key, ": ")
    for k, e := range element {
      //fmt.Println("Key:", k, "=>", "Element:", e)
      if len(e) > 0 {
        fmt.Println(key, ": ", k, "=>", e)
      }
    }
    //fmt.Println("Key:", key, "=>", "Element:", element)
    fmt.Println(" ")
  }
}
