# FreezeLib 示例集合

**Language / 语言**: [English](README.md) | [中文](README_CN.md)

本目录包含了展示 FreezeLib 各种功能的综合示例。

## 📁 示例分类

### [01-basic/](01-basic/) - 基础用法
- 简单代码截图生成
- 基本配置
- 入门示例

### [02-formats/](02-formats/) - 输出格式
- SVG 输出示例
- PNG 输出示例
- 格式对比
- 质量设置

### [03-themes/](03-themes/) - 主题展示
- 流行主题演示
- 主题对比
- 自定义主题创建

### [04-languages/](04-languages/) - 编程语言
- 不同语言的语法高亮
- 语言特定优化
- 多语言项目

### [05-terminal/](05-terminal/) - 终端输出
- ANSI 颜色支持
- 终端样式
- 命令输出截图

### [06-advanced/](06-advanced/) - 高级配置
- 复杂样式选项
- 性能优化
- 自定义字体和布局

### [07-batch/](07-batch/) - 批量处理
- 多文件处理
- 自动化工作流
- 批量操作

### [08-auto-language-detection/](08-auto-language-detection/) - 自动语言检测
- 智能语言检测
- 内容分析检测
- 文件名检测
- 自定义检测器配置
- 检测 API 使用

### [09-supported-options/](09-supported-options/) - 支持选项列表
- 获取所有支持的语言列表
- 获取所有支持的主题列表
- 获取所有可用的预设列表
- 验证语言、主题、预设支持
- 简单高效的 API

## 🚀 快速开始

运行所有示例：

```bash
cd examples
go run run_all_examples.go
```

运行特定分类：

```bash
# 基础示例
cd examples/01-basic
go run main.go

# 自动语言检测
cd examples/08-auto-language-detection
go run main.go

# 支持选项列表
cd examples/09-supported-options
go run main.go
```



## 🤝 贡献示例

添加新示例：

1. 选择合适的分类或创建新分类
2. 遵循命名约定：`example_name.go`
3. 包含代码和生成的输出
4. 在 README.md 中添加文档
5. 使用 `go run main.go` 测试

