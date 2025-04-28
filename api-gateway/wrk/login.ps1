$mountPath = "${PWD}"
$luaScriptPath = "/scripts/login.lua"

docker run -it --rm -v "${mountPath}:/scripts" williamyeh/wrk `
    -t12 -c400 -d30s -s $luaScriptPath `
    http://host.docker.internal:80/project/login