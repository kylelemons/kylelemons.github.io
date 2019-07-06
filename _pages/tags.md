---
permalink: /tags/
title: Tags
---


<div id="archives">
{% for tag in site.tags %}
  <div class="archive-group">
    {% capture tag_name %}{{ tag | first }}{% endcapture %}
    <div id="#{{ tag_name | slugify }}"></div>
    <p></p>

    <h3 class="tag-head">Tag: {{ tag_name }}</h3>
    <a name="{{ tag_name | slugify }}"></a>
    {% for post in site.tags[tag_name] %}
    <article class="archive-item">
        <h4>
            <a href="{{ site.baseurl }}{{ post.url }}">{{post.title}}</a>
            <small>{{ post.date | date: "%-d %B %Y" }}</small>
        </h4>
    </article>
    {% endfor %}
  </div>
{% endfor %}
</div>
