## git提交规范

## 起因
zzzz...

![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200517171852.png)

>提交log可以看出，很凌乱~

>So,需要一个规范，一个比较固定的提交规则，便于维护和查找

## 规范
```cgo
[源码路径]  版本号  类型 : 具体内容
```
>源码路径：
[参考路径](https://github.com/crab21/Images/blob/master/blog/20200517175159.png)

>版本号：
>
```
X.N.M ---> 
X : 大版本提交 ===> < 10
N : 不同类别(研究) ===> < 100
M : 小版本提交序号 ===> < 100
```
>类型：
>
```
必选：
	fix : 修复
	md : markdown文件修改
	add : 新增
	del : 删除
	refactor : 重构
	test : 测试类修改
	bench : 性能测试
	performance : 性能提升型
	style : 样式风格调整
可选组合：(可作为标签使用)
	skill : 技巧型
	pit : 常见错误
	
	
例：
	md/skill : md文件中小技巧的修改
	md/refactor : md文件重构
```

```cgo
[encoding/json] 0.0.2 md : fix target url
[encoding/json] 0.0.3 md/skill : md skill 
[encoding/json] 0.0.2 add/skill : add json skill 
```
