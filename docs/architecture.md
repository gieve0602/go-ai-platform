# architecture design

## Day 1

### 为什么未来要将 LLM Gateway、RAG、Agent 拆成独立模块？

```plain
拆成独立模块的原因是如果三个模块内容都方在一起，会导致项目臃肿，杂乱，结构不清晰，不能复用,
如果拆为独立的模块，LLM Gateway就可以对接不同的大模型，对LLM的修改不会影响其它的模块；各模块
通过事先规定号的接口进行对接，同理其它模块也是。
```

### 为什么不应该让 HTTP Handler 直接调用大模型？

```plain
如果通过http handler直接调用大模型， 会导致handler代码过于繁杂，不利于修改和复用，调用不同的模型
就得写不同的handler
```

### 如果未来接入 OpenAI、Claude 和本地模型，应该如何设计统一调用接口？

```plain
应该通过统一的gateway通过不同的路由去条用不同的模型
```
