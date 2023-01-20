
# Change Log
All notable changes to this project will be documented in this file.

## [0.2.1] - 2023-01-20
### Added
- Run only for hosts in playbook

### Changed
- Repo file format

## [0.1.2] - 2023-01-18
### Added
- Clear table button
- List all files in repo, add Run button
- Short file names on Index page
- Binary and Deb release

### Changed
- In play.Exec implement Quit before timeout ends

## [0.1.1] - 2023-01-17
### Added
- Color HEADs on Index page
- Docker workflows

### Fixed
- Repo update

## [0.1.0] - 2023-01-11
- [x] Execute Ansible playbook for 1 host
- [x] Get host list from inventory
- [x] Get repo HEAD (`git rev-parse --short HEAD`)
- [x] Check that each playbook file exists and yaml/yml
- [x] Get a list of chnaged playbooks (`git diff --name-only 9fc2add HEAD`)
- [x] Check if repo was updated
- [x] For each playbook and for each host run Ansible
