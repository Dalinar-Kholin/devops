a) imagePullNever: powoduje że obrazy nigdy nie będą pobierane z repozytorium,a tylko pobierane z lokalnego noda

b

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml // pobranie

kubectl get pods -n kubernetes-dashboard // sprawdzenie czy działa

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
``` 
stworzneie usera

kubectl apply -f dashboard-admin.yaml

kubectl -n kubernetes-dashboard create token admin-user // pobranie tokenu do uwierzytelnienia na http 

kubectl proxy

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/





c)
- skomilowany frontend --> such pliki http i js /css
- jakieś walidatory requestów
- reverse proxy

