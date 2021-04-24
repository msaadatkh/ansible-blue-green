# Blue/Green deployment with Ansible and Docker Swarm

### Stack info
- Remote node OS: `Ubuntu 20.04`
- Docker version: `20.10.6`
- Ansible version: `2.9.6`
- Python versoin: `3.8.5`

### Architecture
Docker Swarm used for app deployment and standalone Nginx serivce as Reverse Proxy. I don't change Nginx config for re-route and just using `http://127.0.0.1:3000` as endpoint. Because of using Docker Swarm Service we don't need to deploy the new version of app on new port and then modify Nginx to re-route the traffic to the new instance.

App deploy by stack file that include one service (`app`). In stack file we have `2` replicas for current version of app (`v1`). After we changed the version of app in `vars` folder to new version and re-apply it with playbook-stack, the progress of updating current instance of app (which has two replica) begins. Swarm updates current replicas of app service one by one. It stops one replica of current version and deploy new version, If the new version becomes healthy and ready to receive traffic, then imeadiately stop another replica of old version and deploy another replica of new version.

In this process we don't have any downtime or traffic disruption, because at worst case we have one replica of current version ( which is working good). If the new version that deploys not working at all (`v4`), Service (the first replica that was killed for the update process) automatically rollback to previous stable version.

In Dockerfile `Healthceck` is determined to help Service for making decision of keeping new version or rollback it to previous version.
For rollback from current healthy version to another healthy version (Bonus point), you just need to change `version` variable in `vars` folder and apply the playbook just like before.

### Order of playbooks:
- `playbook-nginx-docker.yml` (which installs Nginx and Docker)
- `playbook-build-push-images.yml` (which build and push images from remote node to DockerHub)
- `playbook-deploy-stack.yml` (which deploy/rollback between versions) 

DockerHub blue-green repo:  https://hub.docker.com/repository/docker/msaadat/blue-green
