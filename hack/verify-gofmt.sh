#!/bin/bash

# Copyright 2022 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

echo "Verifying gofmt"

diff=$(find . -name "*.go" | grep -v "\/vendor\/" | grep -v "\/cmd\/cyclonus\/" | xargs gofmt -s -d 2>&1)
if [[ -n "${diff}" ]]; then
  echo "${diff}"
  echo
  echo "Please run make fmt to fix the issue(s)"
  exit 1
fi
echo "No issue found"
