本课程的实战项目是一个针对第三方网站的爬虫，课程开设以来，也的确遇到了一些常见的对于第三方依赖的问题：

返回的网页格式变了，导致视频里的代码不好用
对方网站对爬虫的过滤更严了，我们可以采用轮换request header中的user-agent的方法，比如这里(https://github.com/EDDYCJY/fake-useragent)，但效果仍然有限。

因此本课程现在推出模拟相亲网站，用来返回和当时的第三方网站极其类似的数据结构，最大限度上的让视屏里的代码不作更改就可运行。它的特点有：

* 无限数据生成
* 资源占用小，不吃内存
* 响应极快，高并发
* 没有任何反爬限制，想爬多快都行

具体用法是：
首先，git clone本课程最新代码
```shell
git clone https://git.imooc.com/coding-180/coding-180.git
```

这里需要同学的慕课网用户名和密码登陆。已经拉过这个代码的同学，git pull就行。

然后就可以起服务器了
一定要确保自己在clone下来的代码的“根”目录下哦，我上面的例子，也就是这个叫coding-180目录下。

```shell
coding-180$ go run mockserver/main.go
```

### 访问首页
去http://localhost:8080看看吧。大致长这样：
![img.png](img.png)

第一步就是点击“进入”，看看里面都有些什么用户。

### 放宽解析器对URL匹配的要求
在视频中，我们会在抓取用户相关URL的时候，精确匹配 http://www.zhenai.com/… 等，我们现在当然不能这样，
由于在模拟相亲网站上，所有的地址都被映射为：http://localhost:8080/mock/www.zhenai.com/… 
其实我们只需匹配 www.zhenai.com 这样的特征串即可。
具体的改动请参考 https://git.imooc.com/coding-180/coding-180/commit/1c9e644e901c6c84de99ad20c14e73c45abc06ec
不清楚也没有关系，跟着视频走我就会讲解正则表达式，并且在视频中边打边讲解。

### 运行爬虫
视频中，我们会把http://www.zhenai.com/zhenghun
页面作为爬虫的起始页面，也需要相应的改成 http://localhost:8080/mock/www.zhenai.com/zhenghun，
由此我们的爬虫就可以运行了。我们上述git clone下来的是本课程上完后的最终代码，当然也可以直接运行。
我们的模拟网站背后的数据是无限的（仅受制于int64的大小），所以想要爬下来多少都行。

想要爬快一点的话，记得修改Config.Qps的值哦。

### 模拟相亲网站的设计和实现
那么，我是如何做到无限数据生成，资源占用小，响应极快，高并发这些看似不可能同时实现的目标的呢？可以看一下我的源码。
网站采用gin框架进行开发，利用了伪随机数的概念进行数据生成，运用了单元测试来保证生成数据的正确性，并且注重代码规范性。
可以当作这门课程的复习使用。