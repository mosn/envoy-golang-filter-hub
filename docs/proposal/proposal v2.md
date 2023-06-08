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

首先我们的目标是搭建一个插件市场，插件本体作为 Docker Image 的形式分发，为了达到这个目标，我们需要解决几个子问题

- 提交流程该如何设计
  - 上架一个插件，用户需要提交哪些内容？
- 审核流程该如何设计
  - 哪些行为需要审核？
  - 如何审核，自建审核系统，还是依托其他服务？
- 如何进行版本管理
  - 如何上架新版本
  - 如何下架一个版本
- 元信息（名称描述分类等）如何存储
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



总结如下：

在提交流程这一步来看，我感觉 GitHub Actions 的流程和我们的项目是最贴切的，毕竟我们希望尽可能地利用 GitHub 的基础设施，我认为我们也可以将仓库的一个 Release 关联到插件的一个版本

审核机制来看，大部分都是平台自建审核功能，而我也注意到了Docker Extension 的 [自行发布](https://github.com/docker/extensions-submissions/issues/new?assignees=&labels=&template=1_automatic_review.yaml&title=%5BSubmission%5D%3A+) 的做法，我认为可以借鉴他的做法：他是一个 issue form 对应一个申请，然后，用 tag 标记状态，并由 GitHub Actions 自动检查是否符合条件，这种做法让我想起来社团里有学长加友链也是[这么](https://github.com/aFlyBird0/blog-friends/issues/new/choose)搞的， issue form 真的可玩性挺高的

而对于我们来说，可以根据审核的事件类型自定义要不要加人工审核，比如说上架，机器人检查通过之后（这个 repo 的确按照我们规定的格式编写好了插件，可以编译成 Docker Image），可以自动 @ 管理员来人工审核并通过（更改 tag 并关闭 issue），如果有问题则可以在这个 issue 下面继续交流

版本管理来看，可以使用与 GitHub Actions 一样的关联 Release 的做法，然后上架或者下架新版本都需要提交审核申请，添加/解除与一个 Release 的关联

验证官方发布来看，也可以和 GitHub Actions 一样，仓库在谁手里就是谁发布的，当然还可以以提交审核的方式认证一些别的 tag

关于元信息存储，我的看法是与 Visual Studio Code 一样在本体中全量存储（比如在根目录的 `metadata.json` ） ，然后后端数据库缓存一份，并且始终缓存最新版的信息，这种做法对后端应该最方便，但是对用户来说可能有点麻烦，毕竟你要改描述或者分类这种信息也需要再发布一个新版本

插件本体存储来看，除去 ChatGPT Plugins 提交的是 API 之外，其他的都是提交并分发一个能离线运行的实体，有些实体是不需要编译的，直接提交源码即可，有些是提交了编译后的产物（如 JAR 包），而在这些案例中大多数都是直接提交编译后的二进制，但是 Docker Extension 有所不同，他是让用户自行将 Docker Image 提交到 Docker Hub ， 然后提交扩展的时候就上交一个链接就好了，如果我们也这样做的话就是把存储成本转嫁给用户，但是从稳定性来看感觉不妥，而 Docker Hub 也可以说是他们自家的存储设施，我的结论就是由我们自己负责编译和存储

另外，作为一个市场还可以有评分和评论的功能，但是我感觉没什么必要，评分的话看仓库的 star 应该就可以了，如果对插件有什么看法的话也可以直接去提一个 issue，当然如果要做的话也可以用 [giscus](https://giscus.app/zh-CN) 这种解决方法，直接依托 GitHub 的基础设施


---

# 实现方案

经过上面的分析已经有一个大概的轮廓了，在正式的 Proposal 中我会以全局设计、使用流程、后端架构三个方面来介绍

## 全局设计

![](./pic/v1.全局设计.png)

全局设计和第一版应该差不多，还是分为这三个部分，这个图还需要重新细化一下，然后会重新写几段介绍文字

## 使用流程

（在流程敲定之后，下一版这里会是一些流程图）

重新梳理一下审核流程

1. 在要进行需要审核的操作时，会触发对应的审核流程，并被重定向到一个 GitHub 仓库的 issue form 中
2. 用户填写 issue form 并提出 issue
3. bot 会进行预检查，这是否是符合规定的申请，如有问题会自动指出，用户可评论某些关键字命令申请重新检查（如 `/validate` )
4. 通过预检查后，如果有必要会 @ 管理员人工审核，否则会自动打上通过的 tag ，关闭 issue 并通知服务器触发后续流程





如果没什么问题的话，下一版会细化成一些小标题



## 后端架构

后端使用单体服务，服务模块划分会有所修改，具体需要等流程确定后再设计



---

# 开发规划

目前的大致开发规划：

- 第一个月（或者从现在开始）完成后端开发，调试好与 GitHub 的调用流程
- 第二个月编写前端，同时发现并处理可能的 bug