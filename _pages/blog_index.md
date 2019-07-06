---
permalink: /blog/
---

{% for post in site.posts -%}
{% unless post.categories contains "Blog" %}{% continue %}{% endunless %}
{% include post_header.html post=post -%}
{{ post.content | markdownify -}}
{% endfor -%}
