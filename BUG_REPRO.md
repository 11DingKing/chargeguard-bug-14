# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

旧版本数据库升级后服务可以正常启动，但历史站点第一次分配整改负责人会 panic，新建站点没有问题。先不要修改代码，请对比升级恢复与新建路径，定位未准备好的状态如何一路触发崩溃；诊断需给出具体文件、符号和重启复现证据。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-14
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-14.git
- parent SHA：52354ac7319eec6dc372088e216fc82174352494

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-14.git bug-repro
cd bug-repro
git checkout --detach 52354ac7319eec6dc372088e216fc82174352494
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
panic: assignment to entry in nil map [recovered, repanicked]

goroutine 6 [running]:
testing.tRunner.func1.2({0x8cd380, 0xd1d7c0})
	/usr/local/go/src/testing/testing.go:1974 +0x232
testing.tRunner.func1()
	/usr/local/go/src/testing/testing.go:1977 +0x349
panic({0x8cd380?, 0xd1d7c0?})
	/usr/local/go/src/runtime/panic.go:860 +0x13a
chargeguard/internal/charging.(*LegacyStation).Assign(...)
	/app/internal/charging/task_behavior.go:12
chargeguard/internal/httpapi.TaskHTTPHandler({0x960550, 0x284cd96f0c80}, 0x92c58f?)
	/app/internal/httpapi/task_behavior.go:11 +0x6d
chargeguard/internal/httpapi.TestTaskBehavior(0x284cd97aa6c8)
	/app/internal/httpapi/task_behavior_test.go:11 +0xd9
testing.tRunner(0x284cd97aa6c8, 0x957fb0)
	/usr/local/go/src/testing/testing.go:2036 +0xea
created by testing.(*T).Run in goroutine 1
	/usr/local/go/src/testing/testing.go:2101 +0x4c5
FAIL	chargeguard/internal/httpapi	0.067s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
panic: assignment to entry in nil map [recovered, repanicked]

goroutine 8 [running]:
testing.tRunner.func1.2({0x461820, 0x8aec40})
	/usr/local/go/src/testing/testing.go:1974 +0x1a0
testing.tRunner.func1()
	/usr/local/go/src/testing/testing.go:1977 +0x318
panic({0x461820?, 0x8aec40?})
	/usr/local/go/src/runtime/panic.go:860 +0x12c
chargeguard/internal/charging.(*LegacyStation).Assign(...)
	/app/internal/charging/task_behavior.go:12
chargeguard/internal/httpapi.TaskHTTPHandler({0x4f7878, 0x136606220d00}, 0x4c2be6?)
	/app/internal/httpapi/task_behavior.go:11 +0x5c
chargeguard/internal/httpapi.TestTaskBehavior(0x13660626e908)
	/app/internal/httpapi/task_behavior_test.go:11 +0xc8
testing.tRunner(0x13660626e908, 0x4ee410)
	/usr/local/go/src/testing/testing.go:2036 +0xc4
created by testing.(*T).Run in goroutine 1
	/usr/local/go/src/testing/testing.go:2101 +0x3a8
FAIL	chargeguard/internal/httpapi	0.010s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

根因结论必须准确写明出问题的 Go 文件、具体符号和完整失效机制，并由实际复现、源码调查和验证证据支撑；调查结束时目标仓库代码、测试和配置零改动。
