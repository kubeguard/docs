[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/guard)](https://goreportcard.com/report/github.com/appscode/guard)

# Guard
 Guard by AppsCode is a [Kubernetes Webhook Authentication](https://kubernetes.io/docs/admin/authentication/#webhook-token-authentication) server. Using guard, you can log into your Kubernetes cluster using your Github or Google authentication token. Guard also sets authenticated user's groups to his Github teams or Google groups. This allows cluster administrator to setup RBAC rules based on membership in Github teams or Google groups.

## Supported Versions
Kubernetes 1.6+

## Installation
To install Guard, please follow the guide [here](/docs/install.md).

## Contribution guidelines
Want to help improve Guard? Please start [here](/CONTRIBUTING.md).

---

**The guard server collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support
If you have any questions, you can reach out to us.
* [Slack](https://slack.appscode.com)
* [Twitter](https://twitter.com/AppsCodeHQ)
