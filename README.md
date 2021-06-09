# ⚠️注意
目前云函数已经支持web函数和镜像函数，意味着你的代码可以零改造上云，因此这个项目要退休了。

- [镜像函数](https://cloud.tencent.com/document/product/583/56051)
- [web函数](https://cloud.tencent.com/document/product/583/56124)

# ScfWatchDog
scf云函数自定义运行时的守护进程，能够让你的服务端应用花最小的改造就能部署到scf云函数


## 注意
请不要使用初始化过于耗时的框架，这会导致scf函数性能受到影响
