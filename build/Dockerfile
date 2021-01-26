# 构建最小运行时镜像
FROM scratch

COPY --chown=65534:0  cfgedit /cfgedit
COPY --chown=65534:0  ca/ /ca/
USER 65534
WORKDIR /

#  ["/cfgedit --incluster=false"]