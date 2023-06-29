# OSQueryED
OSquery Enumeration and Discovery (OSqueryED) is a PoC that demonstrate how to leverage Osquery interactive shell named pipe implementation to enumerate windows machines. 

# Build Instructions

```sh
go get && go build
```

# Add/Custom Queries

Feel free to edit `WindowsQueries` in `helpers/windows.go`.

```go
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
```

Happy hacking!
