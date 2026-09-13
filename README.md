# 档案安全记录服务

这是档案库房异常记录的后端起点。接口契约和样例事件放在 `contracts/` 与 `fixtures/`，数据目录默认使用 SQLite 文件。运行时配置通过环境变量提供，敏感配置不要提交到仓库。

本地启动：

```bash
docker compose up --build
```

服务监听 `8080`，`/healthz` 仅用于检查进程是否可用。
