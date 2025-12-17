bomba fast attack jak korzystać z kubernetesa absolutnie go nie rozumiejąc

wystawianie jakiegoś kontenera na światło dzienne \
pierw chcemy stworzyć deployment tego kontenera \
następnie tworzymy jego service


pierw ładujemy obraz do minikube

minikube image load <imageName>:<tag> \
minikube nie załąduje nam nowego obrazu pod tym samym tagiem więc nie polecam latest

następnie

kubectl apply -f <plik>.yaml
i tak dla deploymentu oraz service


po czym aby dostać url do któego możemy pukać to
```bash
minikube service <service-name>
```
i dostajemy wszystkie ważne info



docker run --rm -it \
-v "$PWD":/workspace \
-w /workspace \
oeciteam/openenclave-base-ubuntu-20.04 \
/bin/bash
