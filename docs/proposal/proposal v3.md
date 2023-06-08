- 项目名称：[为 Envoy Go 扩展建设插件市场](https://summer-ospp.ac.cn/org/prodetail/23f080259?lang=zh&list=pro)
- 申请人：徐皓
- 联系邮箱：nx@nickxu.me

---

[TOC]

---

# 项目背景

## 项目描述

Envoy是当前最流行的网络代理之一，Go扩展是MOSN社区为Envoy增加的Go生态基础，也是MoE框架的基础。

受益于Golang生态系统，研发可以轻松在Envoy实现插件用于更多的长尾场景，其中很多场景都是通用的。

本项目是为Envoy GO扩展构建插件市场。在插件市场中，人们可以在插件市场中贡献他们已经建设好并已经在生产环境中使用的优质插件，并由其他人来使用。通过插件市场，可以让Envoy生态变得更加开放、共享、丰富

## 产出要求

1. 提供一个Envoy GO插件的内容平台，在这里可以发布经过社区review的优秀插件，需要拥有服务端与前端页面
2. 不自建账号体系，通过GitHub OAuth2.0完成用户认证与授权
3. 对接GitHub OpenAPI，支持动态获取插件所在仓库信息，包括README，分支版本以及star数

## 相关链接

- [开源软件供应链点亮计划-开源之夏2023](https://summer-ospp.ac.cn/org/prodetail/23f080259?lang=zh&list=pro)
- [为 Envoy Go 扩展建设插件市场 · Issue #1 · mosn/envoy-golang-filter-hub](https://github.com/mosn/envoy-golang-filter-hub/issues/1)

---

# 问题分析

首先我们的目标是搭建一个插件市场，并分发编译后的 Docker Image ，为了达到这个目标，我们需要解决几个子问题

- 提交流程该如何设计
  - 上架一个插件，用户需要提交哪些内容？
- 审核流程该如何设计
  - 哪些行为需要审核？
  - 如何审核，自建审核系统，还是依托其他服务？
- 如何进行版本管理
  - 如何上架新版本
  - 如何下架一个版本
- 元信息（名称描述分类等）如何存储（或插件的仓库中需要存哪些内容）
  - 是全部都存在插件本体中，然后平台后端全量缓存
  - 还是一部分打包在插件中，另一部分单独在后端保存
- 插件本体（Docker镜像）该如何构建，存储与分发
  - 是我们负责构建存储分发一条龙
  - 还是交给用户存储在第三方（如 Docker Hub），我们只保存地址？

自己摸索无异于闭门造车，我认为可以参考一下其他平台的做法

因此，我研究了 8 个类似的或者可以提高参考价值的平台，观察他们是如何解决这些问题的，总结为下表：

| 平台/市场                                                    | 提交流程                                                     | 审核机制                                                     | 版本管理                                                     | 验证官方发布                                                 | 元信息保存                                                   | 官方文档                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [Visual Studio Code Marketplace](https://marketplace.visualstudio.com/vscode) | 1. 将插件打包为 `.vsix` 文件<br>2. 在 [Azure DevOps](https://azure.microsoft.com/services/devops/) 使用Microsoft 帐户创建账号并获取 Token<br />3. 使用同一Microsoft 帐户在 [Visual Studio Marketplace](https://marketplace.visualstudio.com/manage) 创建发布者 <br />4. 使用 `vsce login` 与  `vsce publish` 命令发布插件 | 在 Visual Studio Marketplace 上由 Marketplace 团队进行审核   | 1. 更新 `package.json` 的版本号<br>2. 使用 `vsce publish` 发布新版本<br />3. 无法直接删除特定版本，需要联系官方支持团队进行处理 | [通过验证是否持有该公司/组织的域名](https://code.visualstudio.com/api/working-with-extensions/publishing-extension#verify-a-publisher) | 保存在项目的 `package.json` 中<br />包含唯一标识、名称、描述、版本号、发布者名称、兼容的 VSC 版本、分类、关键词等 | [Publishing Extensions &#124; Visual Studio Code Extension API](https://code.visualstudio.com/api/working-with-extensions/publishing-extension) |
| [JetBrains Plugin Repository](https://plugins.jetbrains.com/) | 1. 在官网上创建账号<br>2. 在创建插件页面上传插件 JAR/ZIP 文件并发布 | 由 JetBrains 团队进行审核                                    | 1. 上传新的插件文件<br>2. 在“Versions”选项卡中删除旧版本     | 暂无明确的官方验证机制                                       | 大部分保存在项目的 `plugin.xml` 中，但分类是在发布的页面手动指定 | [Uploading a new plugin &#124; JetBrains Marketplace Documentation](https://plugins.jetbrains.com/docs/marketplace/uploading-a-new-plugin.html) |
| [GitHub Actions](https://github.com/marketplace?type=actions) | 1. 在 GitHub 仓库中完成开发，并编写 `action.yml`<br />2. 创建一个 release ，勾上发布到市场，同时填写分类等信息 | 无审核机制，发布后即可使用                                   | 1. 创建新的 release， 并勾上发布到市场<br />2. 如要删除发布，取消勾选并保存即可 | 仓库属于哪个组织，就是由哪个组织发布的                       | 大部分保存在项目的 `action.yml` 中，但分类是在发布的页面手动指定 | [在 GitHub Marketplace 中发布操作 - GitHub 文档](https://docs.github.com/zh/actions/creating-actions/publishing-actions-in-github-marketplace) |
| [Chrome Web Store](https://chrome.google.com/webstore/category/extensions) | 1. 在 Chrome 开发者仪表盘创建新项目<br>2. 上传 `.zip` 文件包含扩展的所有代码<br>3. 填写项目详情，如名称、描述、图标、预览图等<br>4. 提交审核并支付开发者注册费用 | 由 Google 团队进行人工审核，内容涵盖性能、安全、隐私等方面   | 1. 在开发者仪表盘提交新的 `.zip` 文件<br>2. 旧版本不会被自动删除，用户可以在商店中查看所有版本 | 通过 Google 账户验证                                         | 一部分保存在项目的 `manifest.json` 中，如名称、版本号、描述等<br />其余信息如图标、预览图、详细描述在开发者仪表盘填写 | [创建和发布自定义 Chrome 应用和扩展程序 - Chrome Enterprise and Education帮助](https://support.google.com/chrome/a/answer/2714278?hl=zh-Hans) |
| [Apple App Store](https://www.apple.com/app-store/)          | 1. 注册成为 Apple 开发者并支付年费<br>2. 使用 Xcode 开发应用并配置相关信息<br>3. 在 App Store Connect 上创建应用并上传<br>4. 提交审核请求 | 由 Apple 团队进行严格审核，包括功能、安全性、隐私、设计等方面 | 1. 在 Xcode 中更新版本号和构建号<br>2. 在 App Store Connect 上上传新版本并提交审核<br>3. 旧版本自动下架 | 通过 Apple Developer Program 验证身份                        | 最基础的部分保存在项目的  `Info.plist` 中，如版本号，构建号，唯一标识，设备上显示的名称<br />另一部分是在 App Store Connect 上，包括在App Store上的应用名称、描述、版本号、类别、预览截图等 | [将 iOS App 提交至 App Store - Apple Developer](https://developer.apple.com/cn/ios/submit/) |
| [WordPress Plugin Repository](https://cn.wordpress.org/plugins/) | 1. 在官网上注册账号<br>2. 在 SVN 仓库中添加插件代码<br>3. 使用 Readme Validator 验证 `readme.txt`<br>4. 在官网上提交插件并等待审核 | 由 WordPress 团队进行审核，主要关注插件的功能和安全性        | 1. 在 SVN 仓库中更新插件和 `readme.txt` 的版本号<br>2. 在官网上标记新版本的发布<br>3. 旧版本仍然可用 | 没有明确的官方验证机制                                       | 保存在项目的 `readme.txt` 中，包括名称、描述、版本号、作者、标签等 | [zh-cn:开发一个插件 « WordPress Codex](https://codex.wordpress.org/zh-cn:%E5%BC%80%E5%8F%91%E4%B8%80%E4%B8%AA%E6%8F%92%E4%BB%B6) |
| [Docker Extension](https://hub.docker.com/search?q=&type=extension)<br /> | 1. 在Docker Hub上注册账号<br>2. 构建好你的扩展镜像，并提交到 Docker Hub<br />3. 选择一种发布方式并等待审核 | 可以选择[自行发布](https://github.com/docker/extensions-submissions/issues/new?assignees=&labels=&template=1_automatic_review.yaml&title=%5BSubmission%5D%3A+)或者[请求官方审核](https://www.docker.com/products/extensions/submissions/) | 1. 推送新版本的扩展 Docker Image ，并带有递增的版本标记<br />2. 像管理你的镜像版本一样管理你的扩展版本 | [加入 Docker Verified Publisher Program](https://docs.docker.com/docker-hub/publish/) | 保存在扩展镜像中的 `metadata.json` 中，当然官方肯定也会缓存一部分 | [Publish your extension to the marketplace](https://docs.docker.com/desktop/extensions-sdk/extensions/publish/) |
| [ChatGPT Plugins](https://openai.com/blog/chatgpt-plugins)   | 1. 搭建好你的 API 服务<br />2. 以官方的格式创建 JSON/YAML 文件描述你的插件，并保存在域名下 | 目前是在官网通过机器人递交表单，然后人工审核                 | 1. 更新你部署即可<br />2. 不需要维护多版本，访问到的就是你部署的最新版本 | 暂无明确的官方验证机制                                       | 大部分保存在你部署的 API 的域名下的 YAML/JSON 文件中，官方服务器仅保存名称，描述及域名等基本信息 | [Getting Started - OpenAI API](https://platform.openai.com/docs/plugins/getting-started) |

结论如下：

- 在提交流程这一步来看，我感觉 GitHub Actions 的流程和我们的项目是最贴切的，毕竟我们希望尽可能地利用 GitHub 的基础设施，我认为我们也可以将仓库的一个 Release 关联到插件的一个版本

- 审核机制来看，大部分都是平台自建审核功能，而我也注意到了Docker Extension 的 [自行发布](https://github.com/docker/extensions-submissions/issues/new?assignees=&labels=&template=1_automatic_review.yaml&title=%5BSubmission%5D%3A+) 的做法，我认为可以借鉴他的做法：他是一个 issue form 对应一个申请，然后，用 tag 标记状态，并由 GitHub Actions 自动检查是否符合条件，这种做法让我想起来社团里有学长加友链也是[这么](https://github.com/aFlyBird0/blog-friends/issues/new/choose)搞的， issue form 真的可玩性挺高的

  而对于我们来说，可以根据审核的事件类型自定义要不要加人工审核，比如说上架，机器人检查通过之后（这个 repo 的确按照我们规定的格式编写好了插件，可以编译成 Docker Image），可以自动 @ 管理员来人工审核并通过（更改 tag 并关闭 issue），如果有问题则可以在这个 issue 下面继续交流

- 版本管理来看，可以使用与 GitHub Actions 一样的关联 Release 的做法，然后上架或者下架新版本都需要提交审核申请，添加/解除与一个 Release 的关联

- 验证官方发布来看，也可以和 GitHub Actions 一样，仓库在谁手里就是谁发布的，当然还可以以提交审核的方式认证一些别的 tag

- 关于元信息存储，我的看法是与 Visual Studio Code 一样在本体中全量存储（比如在根目录的 `metadata.json` ） ，然后后端数据库缓存一份，并且始终缓存最新版的信息，这种做法对后端应该最方便，但是对用户来说可能有点麻烦，毕竟你要改描述或者分类这种信息也需要再发布一个新版本

  也就是说，只要插件的仓库中有一个能编译出来镜像的 `Dockerfile` ，以及一个符合规范的 `metadata.yaml` ，就够了

- 插件本体存储来看，除去 ChatGPT Plugins 提交的是 API 之外，其他的都是提交并分发一个能离线运行的实体，有些实体是不需要编译的，直接提交源码即可，有些是提交了编译后的产物（如 JAR 包），而在这些案例中大多数都是直接提交编译后的二进制，但是 Docker Extension 有所不同，他是让用户自行将 Docker Image 提交到 Docker Hub ， 然后提交扩展的时候就上交一个链接就好了，如果我们也这样做的话就是把存储成本转嫁给用户，但是从稳定性来看感觉不妥，当然你也可以说 Docker Hub 是他们自家的存储设施，我的结论就是由我们自己负责编译和存储

另外，作为一个市场还可以有评分和评论的功能，但是我感觉没什么必要，评分的话看仓库的 star 应该就可以了，如果对插件有什么看法的话也可以直接去提一个 issue，当然如果要做的话也可以用 [giscus](https://giscus.app/zh-CN) 这种解决方法，直接依托 GitHub 的基础设施


---

# 实现方案

 经过上面的思考，我提出下面的设计方案

## 全局设计

~~（图画的不满意暂时没配）~~

项目整体由三个部分组成：GitHub，Filter Hub，和 Docker Registry

### GitHub

本插件市场可以使用 GitHub 作为基础设施来实现部分功能，不必重复造轮子

- 作为代码仓库

  开发者在 GitHub 上维护自己的项目

- 用户认证与授权

  在组织中的用户即为管理员，使用 GitHub OAuth 登录便可在后端拿到用户身份

- 作为讨论与审核平台

  可以直接在 GitHub 上发起对插件的 Issue 与 PR ， 所有人也都可以对插件给出自己的建议

  若要上架或更新插件，需要使用 issue form 递交申请


### Filter Hub

市场本体的功能可以分为用户相关与插件相关

- 用户相关

  基于GitHub OAuth 2.0的用户认证与授权

- 插件相关

  列出和搜索插件

  接收 GitHub 传来的事件

  管理已发布的插件（管理员手动隐藏）

  为每个发布版本构建并推送 Docker 镜像

### Docker Registry

本部分用于保存与分发每个插件版本的 Dokcer 镜像，供用户拉取使用

应当设置为仅后端服务与能推送与删除镜像，而普通用户只读

## 流程设计

### 上架插件/发布新版本

1. 在自己的 GitHub 仓库中开发好插件，包含 `Dockerfile` 与 `metadata.yaml`

2. 使用 issue form 提交申请，包括自己的仓库地址，同意服务条款

3. 使用 GitHub Actions 检查是否接收了服务条款

4. 使用 GitHub Actions 找到该项目的 Latest Release ，检查是否合规

   1. 元数据是否符合规范、完整

      1. 询问后端是否重名
      2. 是否填写了分类、分类是否在预定义的种类中
      3. 如果定义了 icon、color ，定义是否合规

   2. 是否能够编译

      使用 `docker build` 试编译

5. 自动检测通过，打上 tag ， 等待人工审核

   若未通过，告知原因，并告知需要发布新版本并使用 `/validate` 重新检测

6. 人工审核并更新 tag

7. GitHub Actions 识别到通过的 tag， 向后端上报要关联到的 release 版本

8. 上架完成后可人为关闭 issue

### 更新元信息（如名称、描述等）

在仓库中发布新 Release， 并按照上面的流程申请发布新版本

（ `Plugin Hub` 始终缓存最新版本的元数据）

### 下架插件/下架版本

1. 使用 issue form 发起申请
2. 管理员打上通过的 tag，GitHub Actions 向后端上报

### 验证官方发布

（认证某个组织是官方组织，之后这个组织下的插件都会有 tag）

1. 使用 issue form 发起申请，说明理由
2. 在 issue 中或私下交流
3. 管理员打上通过的 tag，GitHub Actions 向后端上报

### 更新仓库信息

后端会缓存插件的部分仓库信息（如 star 数）并在列表页面显示，显示规则和 GitHub 相同（小于 1k 显示精确，大于 1k 小于 10k 精确到 0.1k，大于 10k 精确到 1k）

前端进入详情页后会拉取仓库精确信息，如果 star 数差别大到会改变显示信息的话，就提醒后端更新 star 数

### 搜索功能

暂时只使用基于前端的简单搜索

后端每次上架新插件或更新仓库信息后，将所有仓库信息打包生成一个索引文件，前端根据这个文件进行搜索

大概与静态博客的搜索功能一样

## 后端架构

~~（图画的不满意暂时没配）~~

大概分为下面几部分

- user

  用户 OAuth 登陆 

  维护官方组织名单

- plugin

  提供插件列出与查询服务

- event

  配合 mq 接收 GitHub 以及前端传来的事件，调用其他服务

- log

  记录与浏览日志

---

# 开发规划

- 第一个月（七月）开发后端与 GitHub Actions，打通各种业务流程与接口
- 第二个月（八月）开发前端页面，整体项目测试
- 第三个月（九月）项目收尾，发现并处理可能的 bug
