Hello模板，实现了Partial的方法
{% code
type Hello struct {
Vars map[string]interface{}
}
%}
{% func (h *Hello) Body() %}
<h1>{%v h.Vars["message"] %}</h1>
<div>
    Hello <b>{%v h.Vars["name"] %}!</b>
</div>
{% endfunc %}