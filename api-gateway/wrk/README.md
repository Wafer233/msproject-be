since wrk is not supprted in windows, we need to use docker

docker pull williamyeh/wrk
docker run -it --rm williamyeh/wrk -t12 -c400 -d30s http://host.docker.internal:8080

docker run -it --rm -v "${PWD}:/scripts" williamyeh/wrk -t12 -c400 -d30s -s
/scripts/login.lua http://host.docker.internal:80/project/login



