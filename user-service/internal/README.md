# internal
个人感觉理解这个才是重点
## application
放service, dto
### service
其实就是我以前写过的webook里头的service，但是把一部分用到dto的给分离了，这个是主要的服务
### dto
其实就是把request放出来，之前的话是写在函数里头了
## domain
### model
这个很好理解，就是之前的domain底下的不带tag的struct
### repository
就是把之前repo里头的interface单独抽取出来
### service
就是把之前service里面用不到实体的抽取出来
## infrastructure
### repository
这个就是以前的repo和dao的结合实现
### persistence
主要放mysql的初始化和初始化
## interface
