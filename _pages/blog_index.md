---
layout: default
permalink: /blog/
---
{% for post in site.posts %}
{% unless post.categories contains "Blog" %}{% continue %}{% endunless %}
<div class="row">
    <div class="col-xs-12">
    {% include post_header.html post=post %}
    </div>
    <div class="col-xs-12">
        {%- if post.long -%}
        <blockquote>
        {{ post.excerpt | markdownify }}
        </blockquote>
        <a class="read-more" href="{{ post.url }}">Read more...</a>
        {%- else -%}
        {{ post.content | markdownify }}
        {%- endif -%}
    </div>
</div>
{% endfor %}
