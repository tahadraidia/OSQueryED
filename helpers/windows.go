package helpers

import (
  //"fmt"
  "io/ioutil"
  "log"
  "os"
)

const OSQUERYI_SOCKET = "\\\\.\\pipe\\shell.em"
const OSQUERYD_SOCKET = "\\\\.\\pipe\\osquery.em" // for future use.

var WindowsQueries = []string {
  "select * from os_version", // OS Version.
  "select * from patches", // Windows OS Patches.
  "select device_name, image from drivers where image != ''", // List unused drivers.
  "select groupname, group_sid from groups", // list local groups.
  "select type, user, host, pid from logged_in_users", // Logged in users;
  "select user, logon_domain, authentication_package from logon_sessions", // logon sessions.
  "select * from ntdomains", // Active Directory.
  "select * from pipes", // listing pipes.
  }

var osqueryi_windows_locations = []string {
  // TODO: added more potential locations,
  // kolide and fleet for example.
  "C:\\Program Files\\osquery\\osqueryi.exe",
}

func FindOSQueryIWindowsLocation() string {
  for _, f := range osqueryi_windows_locations {
    if _, err := os.Stat(f); err == nil {
      return f
    }
  }
  return " "
}

func DoesNamedPipeExist(namedpipe string) bool {
  files, err := ioutil.ReadDir("//./pipe/")
  if err != nil {
    log.Fatal(err)
    return false
  }

  for _, file := range files {
    if file.Name() == namedpipe {
      return true
    }
  }
  return false
}
