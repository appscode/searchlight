---
title: Check Pod Exec
menu:
  product_searchlight_{{ .version }}:
    identifier: hyperalert-check-pod-exec
    name: Check Pod Exec
    parent: hyperalert-cli
product_name: searchlight
section_menu_id: reference
menu_name: product_searchlight_{{ .version }}
---
## hyperalert check_pod_exec

Check exit code of exec command on Kubernetes container

### Synopsis

Check exit code of exec command on Kubernetes container

```
hyperalert check_pod_exec [flags]
```

### Options

```
  -a, --argv string        Arguments for exec command. [Format: 'arg; arg; arg']
  -c, --cmd string         Exec command. [Default: /bin/sh] (default "/bin/sh")
  -C, --container string   Container name in specified pod
  -h, --help               help for check_pod_exec
  -H, --host string        Icinga host name
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --bypass-validating-webhook-xray   if true, bypasses validating webhook xray checks
      --context string                   Use the context in kubeconfig
      --icinga.checkInterval int         Icinga check_interval in second. [Format: 30, 300] (default 30)
      --kubeconfig string                Path to kubeconfig file with authorization information (the master location is set by the master flag).
      --log-flush-frequency duration     Maximum number of seconds between log flushes (default 5s)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr
      --use-kubeapiserver-fqdn-for-aks   if true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [hyperalert](/docs/reference/hyperalert/hyperalert.md)	 - AppsCode Icinga2 plugin


