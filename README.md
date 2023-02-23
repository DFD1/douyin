# 字节跳动青训营项目

## 仿抖音项目  
- 运行环境  
golang1.18  
gin1.7.7  
mysql5.7  
ffmpeg  

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/public/video_name 即可   
- 基础接口  
/douyin/feed/ - 视频流接口  
/douyin/user/register/ - 用户注册接口  
/douyin/user/login/ - 用户登录接口  
/douyin/user/ - 用户信息  
/douyin/publish/action/ - 视频投稿  
/douyin/publish/list/ - 发布列表  
- 互动接口  
/douyin/favorite/action/ - 赞操作  
/douyin/favorite/list/ - 喜欢列表  
/douyin/comment/action/ - 评论操作  
/douyin/comment/list/ - 视频评论列表  
