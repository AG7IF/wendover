# Deployment Checklist

- [ ] Open PR from `develop`
- [ ] Verify PR builds and speculative plans pass
- [ ] Merge into `main`
- [ ] Ensure Terraform changes to ECR are in place
- [ ] Approve `build-hold-wendsrv`
- [ ] Ensure containers are built and pushed
- [ ] Apply Terraform changes
- [ ] Approve `deploy-hold-web`
- [ ] Approve `deploy-hold-wendsrv`