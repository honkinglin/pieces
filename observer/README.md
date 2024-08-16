# Observer API

## IntersectionObserver

### 语法
```javascript
const observer = new IntersectionObserver(callback, options);

observer.observe(targetElement);
observer.unobserve(targetElement); // 停止观察某个元素
observer.disconnect(); // 停止观察所有元素
```
### 用途:
`MutationObserver` 用于监测 DOM 树的变化，如节点的添加、删除、属性变化、或文本内容的变化。它在需要对 DOM 的动态变化做出反应时非常有用，尤其适用于富文本编辑器或需要动态更新页面内容的应用。

### 主要配置选项:
- `childList`: 是否监测子节点的添加或删除。
- `attributes`: 是否监测属性的变化。
- `characterData`: 是否监测节点内容或文本的变化。
- `subtree`: 是否监测整个子树的变化（包括子元素的子元素）。

## MutationObserver

### 语法
```javascript
const observer = new MutationObserver(callback);

observer.observe(targetNode, options);
observer.disconnect(); // 停止观察
observer.takeRecords(); // 获取所有未处理的变化记录
```

### 用途:
`IntersectionObserver` 用于监测元素与其祖先元素（或视口）之间的交集。它广泛用于实现懒加载（lazy loading）、无限滚动（infinite scrolling）、以及其他与滚动事件相关的功能。

### 主要配置选项:
- `root`: 用于指定监视的祖先元素（默认是浏览器视口）。
- `rootMargin`: 用于扩展或缩小 root 元素的边界（类似于 CSS 的 margin 属性）。
- `threshold`: 指定触发回调的阈值（0 到 1 的数组），表示目标元素的可见比例。


## ResizeObserver

### 语法
```javascript
const observer = new ResizeObserver(callback);

observer.observe(targetElement);
observer.unobserve(targetElement); // 停止观察某个元素
observer.disconnect(); // 停止观察所有元素
```

### 用途:
`ResizeObserver` 用于监测元素尺寸的变化。它在响应式设计中非常有用，尤其是当你需要根据元素的大小调整样式或布局时。常见应用场景包括复杂的自适应布局、图表动态调整、以及其他基于元素大小的功能。

### 主要配置选项:
- `box`: 指定观察的尺寸盒子类型（content-box 或 border-box）。

## 总结对比
- `MutationObserver` 监测 DOM 结构和内容的变化，适用于需要动态响应 DOM 修改的场景。
- `IntersectionObserver` 监测元素是否进入或离开视口或某个祖先元素的可视区域，适用于滚动加载和懒加载等场景。
- `ResizeObserver` 监测元素尺寸的变化，适用于响应式设计和布局调整。
