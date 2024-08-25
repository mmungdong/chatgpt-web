# build: docker buildx build -f Dockerfile -t tokeninzer:0.1 .
# run: docker run --name tokenizer -p 5000:3002 -d tokeninzer:0.1
FROM python:3.10-alpine
WORKDIR app
ENV PORT 3002

ADD requirements.txt /app/requirements.txt
RUN pip install -i https://mirrors.aliyun.com/pypi/simple --upgrade pip
RUN pip install --root-user-action=ignore -i https://mirrors.aliyun.com/pypi/simple -r requirements.txt

# 安装 curl 工具（用于健康检查）
RUN apk add --no-cache curl

# health check
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD curl --fail http://localhost:${PORT} || exit 1

ADD tokenizer.py /app/tokenizer.py
CMD ["sh", "-c", "nuxt --port ${PORT} --module tokenizer.py --workers 2"]