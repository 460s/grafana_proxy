# grafana_proxy
Внимание, это демонстрационная конфигурация, предназначенная только для ознакомления.  
Как это работает:
#### Сеть
Создаем docker сеть в рамках которой будут доступны контейнеры с Grafana и proxy.  
`docker network create --driver=bridge --subnet=192.168.0.0/16 gnet`
#### Grafana
Поднимаем контейнер с Grafana, включенным режимом авторизации посредствам заголовка.  
`docker run -d --name=grafana --network=gnet -e "GF_AUTH_PROXY_ENABLED=true" -e "GF_AUTH_PROXY_HEADER_NAME=X-GRAFANA-AUTH" grafana/grafana`
#### Репозиторий
Клонируем репозиторий  
`git clone https://github.com/460s/grafana_proxy.git`  
Переходим в папку  
`cd grafana_proxy`  
Собираем образ proxy  
`docker build -t grafana_proxy .`  
Запускаем контейнер с proxy пробрасывая 4000 порт в хостовую сеть  
`docker run -d --name proxy -p 4000:4000 --network=gnet grafana_proxy:latest`
#### Проверка
Убеждаемся что оба контейнера запущены  
`docker ps`  
Открываем в браузере localhost:4000


