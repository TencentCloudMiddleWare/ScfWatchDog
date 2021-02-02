# 设计

用户的服务如果不进行改造那么一定是server结构。但是scf是一个server->client结构，用户实际上要做的是逻辑部分，故而无法和各种web框架一起使用

watchdog需要做的就是作为client的地位，将从scf拉取的请求再次请求到用户的server，避免用户改造。后续还将针对不同的语言添加更多功能

# 版本1

1.执行customruntime的bootstarp功能，用于拉起客户server
2.从scf拉取请求并请求到用户server 作为一个proxy作用
3.记录结构化日志，用于方便分析

# 环境变量
1. WATCHDOG_RUN_PATH 拉起的server路径及命令
2. WATCHDOG_DSIPATCH_MODE 转发模式，当前只有api,plain
3. WATCHDOG_DSIPATCH_PATH 转发路径，如果为api则为http//127.0.0.1:port即可，如果为plain需要加上所需路径
4. WATCHDOG_DEBUG 如果为true则设置log为debug等级