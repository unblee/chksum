#!/bin/bash
set -e

ver=v$(gobump show -r)
make crossbuild
ghr -username unblee -replace ${ver} dist/${ver}
