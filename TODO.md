# Roadmap

## ToDo
- [ ] Use Repo timeout for playbook repeat
- [ ] In play.Exec implement Quit before timeout ends

## Done
- [x] Execute Ansible playbook for 1 host
- [x] Get host list from inventory
- [x] Get repo HEAD (`git rev-parse --short HEAD`)
- [x] Check that each playbook file exists and yaml/yml
- [x] Get a list of chnaged playbooks (`git diff --name-only 9fc2add HEAD`)
- [x] Check if repo was updated
- [x] For each playbook and for each host run Ansible
- [x] Rename conf.Timeout to Interval
- [x] Store results in DB