## 搭建基本编程环境

> 编程环境截图

![调试中击中断点的截图](https://raw.githubusercontent.com/Zhang-Yue0621/work1-qsc/main/pic1.png)


---

> git截图

![命令行形式Git截图](https://raw.githubusercontent.com/Zhang-Yue0621/work1-qsc/main/pic2.png)

---
---

## git:版本控制工具


> 1.如果不借助 git，即没有这一样工具，只使用您常用的即时通信工具来进行代码协作，您认为会遇到哪些困难？

1.无法回滚版本，每次更改后只保留了最新的版本，想要混滚版本时较为困难。

2.团队协作困难，无法及时获得别人的最新代码，需要大量的文件收发工作。


---

>2.github仓库截图

![总算上传成功了](https://raw.githubusercontent.com/Zhang-Yue0621/work1-qsc/main/pic3.png)

>上传代码步骤

github:
```
    git init
    git add ***.*
    git commit -m "***"
    git branch -M main
    git remote add origin git@github.com:***    //此处需用SSH
    git push -u origin main
```
gitee:
```
    git init
    touch ***.*
    git add ***.*
    git commit -m "***"
    git remote add origin https://gitee.com/***
    git push -u origin master
```
---

>结合本节的第一个问题，您现在认为 git 是如何解决您前面提出的问题的？

1.git可以记录每次更改，回滚版本

2.利用git可以将团队联系起来，将所有代码统一放入代码仓库，统一管理。

---
---
## Final

>请给这篇作业评分（满分 10 分），并提出相关的修改意见。

10分

本次作业内容是最基础但也是最实用的一部分，由浅入深地引导小白完成任务。我觉得挺好的（没啥要改的