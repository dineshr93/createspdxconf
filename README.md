<h1 id="yocto-config-for-create-spdx">Yocto config for create spdx</h1>

Configure the create-spdx bbclass parameters in the `workdir/conf/local.conf` 
check : [create-spdx-2.2.bbclass](https://git.yoctoproject.org/poky/tree/meta/classes/create-spdx-2.2.bbclass)

<h2 id="build">docker binary</h2>
<h3 id="binary">binary</h3>
<pre class="codeblock language-sh">docker pull dineshr93/yconfcreatespdx:1.0</pre>
<h3 id="binary">Usage</h3>
<pre class="codeblock language-sh">docker run -v ${PWD}:${PWD} yconfcreatespdx:1.0 ${PWD}/reposync_dir insert_position</pre>


<h2 id="build">build</h2>
<h3 id="binary">binary</h3>
<pre class="codeblock language-sh">make build
</pre>
<h3 id="docker">docker</h3>
<pre class="codeblock language-sh">make dbuild
</pre>
<h2 id="run">Run</h2>
<pre class="codeblock language-sh">
<span class="hljs-built_in">cd</span> Dir_where_reposync_folder_is_present
 
docker run -v <span class="hljs-variable">${PWD}</span>:<span class="hljs-variable">${PWD}</span> yconfcreatespdx:1.0 reposync_dir insert_position

Please provide reposync folder to process that contains workdir , insert position (starts from 0) to <span class="hljs-built_in">local</span> conf file as an argument
check https://git.yoctoproject.org/poky/tree/meta/classes/create-spdx-2.2.bbclass


docker run -v <span class="hljs-variable">${PWD}</span>:<span class="hljs-variable">${PWD}</span> yconfcreatespdx:1.0 <span class="hljs-variable">${PWD}</span>/reposync 3

<span class="hljs-built_in">cat</span> /home/dinesh/dev/createspdxconf/reposync/workdir/conf/local.conf
INHERIT += <span class="hljs-string">&quot;create-spdx-2.2&quot;</span>
SPDX_PRETTY = <span class="hljs-string">&quot;1&quot;</span>
SPDX_ARCHIVE_PACKAGED = <span class="hljs-string">&quot;1&quot;</span>
SPDX_INCLUDE_SOURCES = <span class="hljs-string">&quot;1&quot;</span>
SPDX_ARCHIVE_SOURCES = <span class="hljs-string">&quot;1&quot;</span>
</pre>
