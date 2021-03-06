这是我们模板的基础实现
{% interface
Partial {
Body()
}
%}
模板编写实现Partial接口的模板
{% func Template(p Partial) %}
<html>
<head>
    <title>Quicktemplate integration with Iris</title>
</head>
<body>
<div>
    Header contents here...
</div>
<div style="margin:10px;">
    {%= p.Body() %}
</div>
</body>
<footer>
    Footer contents here...
</footer>
</html>
{% endfunc %}
基本模板实现。 如果需要，其他页面可以继承
仅覆盖某些部分方法。
{% code type Base struct {} %}
{% func (b *Base) Body() %}This is the base body{% endfunc %}