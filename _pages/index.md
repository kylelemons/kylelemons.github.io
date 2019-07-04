---
permalink: /
---

{% for post in site.posts %}
{% include post_header.html post=post %}
{{ post.content | markdownify }}
{% endfor %}
