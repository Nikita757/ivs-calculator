{% if obj.type == 'func' %}
    {# Creating the parameters line #}
    {% set ns = namespace(tmpstring='') %}
    {% set argjoin = joiner(', ') %}
    {% for param in obj.parameters %}
        {% set ns.tmpstring = ns.tmpstring ~ argjoin() ~ param.name ~ ' ' ~ param.type %}
    {% endfor %}
.. {{ obj.ref_type }}:: {{ obj.name }}({{ ns.tmpstring }})
{% else %}
.. go:{{ obj.ref_type }}:: {{ obj.name }}
{% endif %}

{% macro render() %}{{ obj.docstring }}{% endmacro %}
{{ render()|format_docstring|indent(4) }}

{% if obj.children %}
    {% for child in obj.children|sort %}
{% macro render_child() %}{{ child.render() }}{% endmacro %}
{{ render_child()|indent(4) }}
    {% endfor %}
{% endif %}
