package main

import (
  "fmt"
  "log"
  "time"
  "os"
  "os/exec"
  helper "github.com/tahadraidia/OSQueryED/helpers"
  "github.com/osquery/osquery-go"
)

func main() {
  if(false == helper.IsOSQueryInterpreterRunning()) {
    var osqueryi_binary = helper.FindOSQueryIWindowsLocation()
    if len(osqueryi_binary) > 0 {
      cmd := exec.Command(osqueryi_binary)
      if err := cmd.Start(); err != nil { 
	fmt.Println("Error: ", err)
      }   
      time.Sleep(1 * time.Second)
    } else {
      fmt.Println("Failed finding osquery interpreter")
      os.Exit(0)
    }
  }

  client, err := osquery.NewClient(helper.OSQUERYI_SOCKET, 10*time.Second)
  if err != nil {
    log.Fatalf("Error creating Thrift client: %v", err)
  }
  defer client.Close()

  for _, q := range helper.WindowsQueries {

    fmt.Println(" ")
    fmt.Println(" ")
    fmt.Println("-------- ", q, "---------")
    resp, err := client.Query(q)

    if err != nil {
      log.Fatalf("Error communicating with osqueryd: %v", err)
    }

    if resp.Status.Code != 0 {
      log.Fatalf("osqueryd returned error: %s", resp.Status.Code)
    }

    helper.ParseOSQueryResponse(resp.Response)
    fmt.Println("-------------------------")
    fmt.Println(" ")
    fmt.Println(" ")
  }
}
