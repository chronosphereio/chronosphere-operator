# ArgoCD Health Checks for Chronosphere Operator
This folder contains the configuration file that enables health checking customer resources created by Chronosphere Operator.

## Description
Argo CD supports custom health checks via scripts written in Lua. The script is supplied in resource.customization field of argocd-cm. 

## Getting Started
To deploy the customization, apply the configmap file and restart the operator pod:

`kubectl apply -n argocd -f argocd-cm.yaml`

`kubectl rollout restart StatefulSet argocd-application-controller -n argocd`