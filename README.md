# prometheus_http_sd

An "extension" of Prometheus _file_sd_, listening for HTTP requests on port 9091.

* **PUT /target?host=xxx** 
Add given host as a Prometheus target to scrape

* **DELETE /target?host=xxx** 
Remove given host from Prometheus scrape targets

* Usage:

In _prometheus.yml_, file_sd must be configured, for example:

```yml
  - job_name: 'somejob'
    file_sd_configs:
    - files:
      - /path/for/file_sd.json
```

```bash
  # Run the listener and write targets to /path/for/file_sd.json
  prometheus_http_sd /path/for/file_sd.json
```
