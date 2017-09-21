# Introduction

Re-usable jumpstart golang template for your next mega (web) app

Includes:

- Dependency management with `dep`
- Live reload with `realize`
- Development and production ready _Dockerfile_s

## Editor config

### Emacs

I _strictly IMHO_ like `go-mode` and `go-company-mode` with autoformat on
save. If you find something better, put it here. Make sure `goimports`
is installed and in your path ([check it out here](https://godoc.org/golang.org/x/tools/cmd/goimports)). Remember that `goimports` might need `go` installed locally.

```elisp
(defun conf--go-mode ()
  (require 'company-go)
  (add-hook 'go-mode-hook
            (lambda ()
              (set (make-local-variable 'company-backends) '(company-go))
              (company-mode)
              (setq gofmt-command "goimports")
              (add-hook 'before-save-hook #'gofmt-before-save))))
(conf--go-mode)
```
