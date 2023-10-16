## Устройство репозитория 
* Есть единственная стабильная ветка main
* Любая фича или исправление бага делается в отдельной ветке, которая ветвится от main
* Как только фича или багфикс готов, прошёл ревью и тестирование, соответствующая ветка мержится в main. Потом ветка обязательно удаляется
* **backend/** - продуктовый код бекенда; **frontend/** - продуктовый код фронта; 


## Правила работы с репозиторием
#### используется Gitlab Flow
<img width="200" alt="image" src="http://er.jinr.ru/git/help/workflow/release_branches.png">

#### Выпуск релиза
Когда вы готовы сделать релиз, просто возьмите нужный мерж-коммит и повесьте на него тег. Используйте аннотированные (annotated) теги. В них сохраняется дата и автор, как в коммите. Таким образом сохранится информация о том, кто именно и когда принял решение о выпуске релиза. После создания тега запускется пайплайн сборки переходящий в gitlab ci инфраструктурного репозитория (**используется Multi-Project**). 

```bash
git tag -a 1.0.0 -m "Version 1.0.0"
git push origin 1.0.0
```

## Momo Store Helm chart
url адрес магазина - https://momo.commerce-store.ru/

```bash
helm repo add momo --insecure-skip-tls-verify https://nexus.commerce-store.ru/repository/momo-store/
helm repo update
helm upgrade --insecure-skip-tls-verify --install momo-store --create-namespace -n momo-store momo/momo-store --set frontend.dockerconfigjson=<dockerconfigjson> --set backend.dockerconfigjson=<dockerconfigjson> --set frontend.tag=<latest> --set backend.tag=<latest> --atomic --timeout 15m
```
## Инфраструктура.
#### Репозиторий
https://gitlab.praktikum-services.ru/std-016-026/infra 
#### Устройство репозитория
* Есть единственная стабильная ветка main
* Любая фича или исправление бага делается в отдельной ветке, которая ветвится от main
* Как только фича или багфикс готов, прошёл ревью и тестирование, соответствующая ветка мержится в main. Потом ветка обязательно удаляется
* **charts/** - helm chart продуктового кода; **helm/** - values.yaml для helm chart инфраструктуры; **terraform/** - описаная на terraform облачная инфаструктура 

#### Описание
Облачная инфраструктура описана согласно IaC  (за исключением создание внешнего домена). Остальные kubernetes-based ресурсы описаны на базе helm-чартов.  
| **Название** 	| **Тип** 	|  **IaC** 	|   **Развертывание** 	|  **внешний URL** 	|**назначение** 	|
|-------	|-------	|---	|---	|---	|---	|
| Vitrual Private Network     	|   Yandex Cloud    	|   Terraform	|   Flow (Gitlab CI)	|   --/--	|   внутренний связность 	| 	
|      Managed Kubernetes  	|   Yandex Cloud     	|  Terraform 	|   Flow (Gitlab CI)	|   --/--	|  	k8s мощности 	|
|    Ingress controller   	|     k8s-based ресурс  	|   Helm	|   Flow (Gitlab CI)	|   --/--	|  	балансировка и внешний доступ к сервисам (nexus и тп) 	|
|      Prometheus 	|     k8s-based ресурс  	|   Helm	|   Flow (Gitlab CI)	|   	  --/--	| 	сбор метрик 	|
|       Nexus	|      k8s-based ресурс 	|   Helm	|   Flow (Gitlab CI)	|   	 https://nexus.commerce-store.ru	| хранение helm чартов 	|
|     Minio  	|    k8s-based ресурс   	|   Helm	|   Flow (Gitlab CI)	|    https://minio.commerce-store.ru 	| 	хранение больших файлов 	|
|      loki 	|     k8s-based ресурс  	|   Helm	|   Flow (Gitlab CI)	|   --/--	| 		сбор логов 	|
|       promtail	|   k8s-based ресурс    	|   Helm	|   Flow (Gitlab CI)	|   --/--	| 	сбор логов 	|
|      grafana 	|     k8s-based ресурс  	|   Helm	|   Flow (Gitlab CI)	|     https://grafana.commerce-store.ru	| графики	|	
#### Развертывание
В репозитории используется переменные которые необходимо определить до начала развертывания.
| **Название** 	| **Описание** 	|
|-------	|-------	|
| dockerconfigjson     	|  json для доступа container registry gitlab 	| 	
|     kubeconfig  	| kubeconfig для доступа и выкатки релизов на k8s	|
|    minio_access_string   	| строка формата http://< ID >:< КEY >@minio.minio.svc.cluster.local:9000  для записи логов в бакет minio	|  
|      NEXUS_PASSWORD 	|   пароль к https://nexus.commerce-store.ru	|
|       NEXUS_USER	|    логин к https://nexus.commerce-store.ru 	|
|     rootPassword  	|   пароль к https://minio.commerce-store.ru	|
|      rootUser 	|   логин к https://minio.commerce-store.ru	|
|       SSH_PRIVATE_KEY	|  приватный ключ ssh в манифестах terraform	|
|      TF_VAR_ACCESS_KEY 	|   публичный ключ бакета в манифестах terraform	|
|      TF_VAR_CLOUD_ID	|   ID облака yandex cloud	|
|      TF_VAR_SECRET_KEY 	|    приватный ключ бакета в манифестах terraform	|
|      TF_VAR_SSH_LOGIN 	|   логин для доступа по ssh на worker-ноды k8s	|
|      TF_VAR_SSH_PUBLIC_KEY 	|   публичный ключ ssh в манифестах terraform	|
|      TF_VAR_TOKEN 	|   Oath токен для доступа к yandex cloud	|

#### Доступы
Для доступа в https://grafana.commerce-store.ru можно использовать логин/пароль

```bash
user
Password
```
