# mvc_for_gin
> 这个仅仅是[Gin](https://gin-gonic.com/)框架的MVC结构，orm使用的是[gorm](https://gorm.io/)，配置文件使用的是[viper](https://github.com/spf13/viper)，所有的配置和api均可以在各自的官网查询到。
## 目录结构
mvc_for_gin  
  ├─configs  
  ├─controller  
  ├─database  
  ├─libs  
  │  ├─response  
  │  └─server  
  ├─middleware  
  ├─models  
  ├─router  
  ├─service  
  ├─setting  
  ├─test  
  │  └─User  
  └─views  

### configs目录
该目录下目前仅有一个config.yaml文件，文件名称请勿更改。这是项目的配置文件。当不存在时会自动生成。
### controller目录
控制器目录
### database目录
修改的mysql驱动器，，一般来说不需要修改这个目录
### libs目录
存放一些功能的封装  
response封装了一些响应，对于业务有需求的可以自定义。  
server封装了对于gin.Engine的使用，一般也不需要更改
### middleware目录
存放中间件，里面有中间件的示例，也可以在Gin官网查询到中间件的写法  
### models目录
存在数据库模型的地方，不涉及数据库操作
### router目录
写路由的地方
### service目录
存放操作数据库service的地方，建议文件名对应controller
### setting目录
- configs.go  
调用配置文件，封装了DbConfigs，PjtConfigs，分别用于获取数据库配置和项目配置，注意，，请勿直接通过config.yaml获取配置，而是是使用setting包封装好的这两个配置对象获取配置。  
当您需要添加自己的配置的时候，，请在该目录下的configs.go文件中添加对应的结构体。以保证可以访问到对应配置。
- tables.go  
创建数据库模型后，，在此处进行迁移
### test目录
测试用目录，，使用了vscode中的REST Client插件进行的api测试，如果不用这个插件可以删除这个目录
### views目录
存放html类的地方
### main.go
项目入口，如果您的项目不需要使用数据库，请修改为
```go
http.Default().NoDataBase().Run()
```
这样您就不需要对数据库进行配置（有数据库配置项也不会生效！）  
Default()中有两个默认中间件，，即gin框架本身的logger()和recovery()  
不需要使用这两个中间件的话请使用New()  
如果您使用了html模板请在后面加上.LoadHTML，并传入指定的模板目录  
"views/*"表示加载views下的所有html模板，，详细请参见gin的html模板

