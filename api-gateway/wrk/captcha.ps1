$mountPath = "${PWD}"
$luaScriptPath = "/scripts/captcha.lua"

# 12 threads; 400 connections; 30 seconds duration
docker run -it --rm -v "${mountPath}:/scripts" williamyeh/wrk `
    -t12 -c400 -d30s -s $luaScriptPath `
    http://host.docker.internal:80/project/login/getCaptcha