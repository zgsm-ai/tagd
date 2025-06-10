# 标记管理

## 背景

AI codereview返回的结果，是一系列Issue，包括问题描述、上下文代码片段，修改建议，等。

用户对review结果存在几种可能：

1. 解决
2. 未处理
3. 标记为待解决；
4. 标记为忽略
5. 标记为误报、不解决

对同一段代码，多次review，可能会发现问题已经解决，或者问题在上一次review中已经发现，是重复问题。

对于重复问题、尽量不再报告给用户，避免消耗用户额外处理精力。

所以，需要有一个模块(tagd)，记录用户对issue的处理结果，避免报告重复issue。

## 技术原理

由于用户对代码的修改，并不存在某种简单标识方法，能够标记issue。

比如：

1. 位置：代码可能增删改查，导致issue关联的代码位置，可能出现漂移。所以用位置标识一个issue，不可行。
2. UUID: issue并非用户提交，不能在提交过程生成uuid来唯一标识一个issue。
3. 内容: 一个issue关联的代码片段有大有小，甚至可能发生不影响issue生成的微小代码变化(如变量更名)，所以无法简单用issue关联的代码片段来唯一标识一个issue。

为了解决上述问题，本模块采用`<代码范围, 问题分类标签, 关键代码片段>`三元组来标识issue。

`<Scope1, Subject1, Code1>`表示，在Scope1的范围内，关键代码片段为Code1的Subject1类型问题。这三个条件是and的关系，如果某个字段未设置，就表明不判断这个条件。

代码范围分三级: 项目、文件、类/函数/方法。Subject1为一个预设的问题分类列表中的成员，由AI选择用来表示问题的分类。关键代码片段是AI输出的引发问题的最小代码片段。

## 数据结构

标记位置：

```go
type TagPostion struct {
    Scope string `json:"scope"`             //代码范围名称
    ScopeType string `json:"scope_type"`    //代码范围的类型: project/file/function/class
    Subject string `json:"subject"`         //问题分类
    KeyCode string `json:"key_code"`        //关键代码片段
}
```

标记:

```go
type Tag struct {
    ID int `json:"id"`                      //标记的ID
    Position TagPosition `json:"position"`  //打标记的代码位置
    Pairs map[string]string `json:"pairs"`  //标记项构成的键值表
}
```

## 接口

标记管理包括以下RESTful API接口：

| 接口 | URL | 说明 |
|------|-----|-----|
| 查询标记列表 | GET /tags | 指定若干键值对，包括: scope, scope_type, key_code, subject，按与的方式查找标记，带有分页机制 |
| 查询标记详情 | GET /tags/{tag_id} | 获取tag_id这个标记的所有内容 |
| 添加标记 | POST /tags | 用户添加的标记，包括项目名，文件名，行号，函数名，函数体，标记值(label)，扩展标记项(键值对)|
| 更新标记 | PUT /tags/{tag_id} | 更新tag_id这个标记的所有内容 |
| 更新标记的键值对 | PUT /tags/{tag_id}/{key} | 更新tag_id标记，把以Key为名的标记项的值更改为新的值 |
| 删除标记 | DELETE /tags/{tag_id} | 删除tag_id标记|

