# 构建最小运行时镜像
FROM scratch

COPY --chown=65534:0  cfgedit /cfgedit
USER 65534
WORKDIR /

# ENTRYPOINT ["/cfgedit --incluster=false"]