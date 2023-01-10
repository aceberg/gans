
# Change Log
All notable changes to this project will be documented in this file.

## [0.1.0] - 2023-01-11
- [x] Execute Ansible playbook for 1 host
- [x] Get host list from inventory
- [x] Get repo HEAD (`git rev-parse --short HEAD`)
- [x] Check that each playbook file exists and yaml/yml
- [x] Get a list of chnaged playbooks (`git diff --name-only 9fc2add HEAD`)
- [x] Check if repo was updated
- [x] For each playbook and for each host run Ansible
