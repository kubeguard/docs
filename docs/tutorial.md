---
title: Tutorial
description: Tutorial
menu:
  product_guard_0.1.0-rc.4:
    identifier: guard-tutorial
    name: Tutorial
    parent: getting-started
    weight: 40
product_name: guard
menu_name: product_guard_0.1.0-rc.4
section_menu_id: getting-started
url: /products/guard/0.1.0-rc.4/getting-started/tutorial/
aliases:
  - /products/guard/0.1.0-rc.4/tutorial/
---
# Tutorial

Guard is a [Kubernetes Webhook Authentication](https://kubernetes.io/docs/admin/authentication/#webhook-token-authentication) server. Guard server requires TLS client certificate for authentication. This certificate is also used to identify whether to use Github or Google to check for user authentication. The `CommonName` and `Organization` fields in the client cert are used for this purpose.

## Github Authenticator
TO use Github, you need a client cert with `CommonName` set to Github organization name and `Organization` set to `Github`. To ease this process, use the Guard cli to issue a client cert/key pair.
```console
$ guard init client {org-name} -o Github
```

![github-webhook-flow](/docs/images/github-webhook-flow.png)

```json
{
  "apiVersion": "authentication.k8s.io/v1beta1",
  "kind": "TokenReview",
  "status": {
    "authenticated": true,
    "user": {
      "username": "<github-login>",
      "uid": "<github-id>",
      "groups": [
        "<team-1>",
        "<team-2>"
      ]
    }
  }
}
```

To use Github authentication, you can use your personal access token with permissions to read `public_repo` and `read:org`. You can use the following command to issue a token:
```
$ guard get token -o github
```
Guard uses the token found in `TokenReview` request object to read user's profile information and list of teams this user is member of. In the `TokenReview` response, `status.user.username` is set to user's Github login, `status.user.groups` is set to teams of the organization in client cert of which this user is a member of.


## Google Authenticator
TO use Google, you need a client cert with `CommonName` set to Google Apps (now G Suite) domain and `Organization` set to `Google`. To ease this process, use the Guard cli to issue a client cert/key pair.
```console
$ guard init client {domain-name} -o Google
```

![google-webhook-flow](/docs/images/google-webhook-flow.png)
```json
{
  "apiVersion": "authentication.k8s.io/v1beta1",
  "kind": "TokenReview",
  "status": {
    "authenticated": true,
    "user": {
      "username": "john@mycompany.com",
      "uid": "<google-id>",
      "groups": [
        "groups-1@mycompany.com",
        "groups-2@mycompany.com"
      ]
    }
  }
}
```
To use Google authentication, you need a token with the following OAuth scopes:
 - https://www.googleapis.com/auth/userinfo.email
 - https://www.googleapis.com/auth/admin.directory.group.readonly

You can use the following command to issue a token:
```
$ guard get token -o google
```
This will run a local HTTP server to issue a token with appropriate OAuth2 scopes. Guard uses the token found in `TokenReview` request object to read user's profile information and list of Google Groups this user is member of. In the `TokenReview` response, `status.user.username` is set to user's Google email, `status.user.groups` is set to email of Google groups under the domain found in client cert of which this user is a member of.

## RBAC Roles

Kubernetes 1.6+ comes with a set of pre-defined set of [user-facing roles](https://kubernetes.io/docs/admin/authorization/rbac/#user-facing-roles). You can create `ClusterRoleBinding`s or `RoleBinding`s to grant permissions to your Github teams or Google groups. Say, you have a Github team called `ops`. You want to make the members of this Github team admin of a cluster. You can do that using the following command:

```console
echo "
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: ops-team
subjects:
- kind: Group
  name: ops
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
" | kubectl apply -f -
```


## Next Steps
- Learn how to install Guard [here](/docs/install.md).
- Want to hack on Guard? Check our [contribution guidelines](/CONTRIBUTING.md).
