use 'godoc cmd/github.com/GreenRaccoon23/slices' for documentation on the github.com/GreenRaccoon23/slices command 

<!--
Copyright 2009 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->
<!--
Note: Static (i.e., not template-generated) href and id
attributes start with "pkg-" to make it impossible for
them to conflict with generated attributes (some of which
correspond to Go identifiers).
-->

<script type='text/javascript'>
document.ANALYSIS_DATA = ;
document.CALLGRAPH = ;
</script>



<div id="short-nav">
<dl>
<dd><code>import "github.com/GreenRaccoon23/slices"</code></dd>
</dl>
<dl>
<dd><a href="#pkg-overview" class="overviewLink">Overview</a></dd>
<dd><a href="#pkg-index" class="indexLink">Index</a></dd>


</dl>
</div>
<!-- The package's Name is printed as title by the top-level template -->
<div id="pkg-overview" class="toggleVisible">
<!-- <div class="collapsed">
	<h2 class="toggleButton" title="Click to show Overview section">Overview ▹</h2>
</div> -->
<div class="expanded">
	<h2 class="toggleButton" title="Click to hide Overview section">Overview ▾</h2>
	<p>
Package slices is a collection of handy methods for string slices.
</p>

</div>
</div>


<div id="pkg-index" class="toggleVisible">
<!-- <div class="collapsed">
    <h2 class="toggleButton" title="Click to show Index section">Index ▹</h2>
</div> -->
<div class="expanded">
    <h2 class="toggleButton" title="Click to hide Index section">Index ▾</h2>

<!-- Table of contents for API; must be named manual-nav to turn off auto nav. -->
<div id="manual-nav">
<dl>
	<dd><a href="#Compact">func Compact(bloated []string) (compacted []string)</a></dd>
	<dd><a href="#Concat">func Concat(slc []string) string</a></dd>
	<dd><a href="#Contains">func Contains(slc []string, s string) bool</a></dd>
	<dd><a href="#Copy">func Copy(slc []string) []string</a></dd>
	<dd><a href="#Cut">func Cut(slc []string, start int, stop int) []string</a></dd>
	<dd><a href="#Equals">func Equals(slc1 []string, slc2 []string) bool</a></dd>
	<dd><a href="#Extract">func Extract(excess []string, wanted ...string) (extracted []string)</a></dd>
	<dd><a href="#Filter">func Filter(unfiltered []string, unwanted ...string) (filtered []string)</a></dd>
	<dd><a href="#IsEmpty">func IsEmpty(slc []string) bool</a></dd>
	<dd><a href="#IsSameArray">func IsSameArray(slc1 []string, slc2 []string) bool</a></dd>
	<dd><a href="#Join">func Join(slc []string, by string) string</a></dd>
	<dd><a href="#Pop">func Pop(slc []string) (string, []string)</a></dd>
	<dd><a href="#Push">func Push(slc []string, args ...string) []string</a></dd>
	<dd><a href="#Shift">func Shift(slc []string) (string, []string)</a></dd>
	<dd><a href="#Unshift">func Unshift(slc []string, s string) []string</a></dd>
</dl>
</div><!-- #manual-nav -->




<h4>Package files</h4>
<p>
<span style="font-size:90%">

	<a href="/src/github.com/GreenRaccoon23/slices/string.go">string.go</a>

</span>
</p>

</div><!-- .expanded -->
</div><!-- #pkg-index -->

<div id="pkg-callgraph" class="toggle" style="display: none">
<div class="collapsed">
<h2 class="toggleButton" title="Click to show Internal Call Graph section">Internal call graph ▹</h2>
</div> <!-- .expanded -->
<div class="expanded">
<h2 class="toggleButton" title="Click to hide Internal Call Graph section">Internal call graph ▾</h2>
<p>
  In the call graph viewer below, each node
  is a function belonging to this package
  and its children are the functions it
  calls&mdash;perhaps dynamically.
</p>
<p>
  The root nodes are the entry points of the
  package: functions that may be called from
  outside the package.
  There may be non-exported or anonymous
  functions among them if they are called
  dynamically from another package.
</p>
<p>
  Click a node to visit that function's source code.
  From there you can visit its callers by
  clicking its declaring <code>func</code>
  token.
</p>
<p>
  Functions may be omitted if they were
  determined to be unreachable in the
  particular programs or tests that were
  analyzed.
</p>
<!-- Zero means show all package entry points. -->
<ul style="margin-left: 0.5in" id="callgraph-0" class="treeview"></ul>
</div>
</div> <!-- #pkg-callgraph -->




<h2 id="Compact">func <a href="/src/target/string.go?s=2725:2776#L110">Compact</a>
	<a class="permalink" href="#Compact">&#xb6;</a>
</h2>
<pre>func Compact(bloated []<a href="/pkg/builtin/#string">string</a>) (compacted []<a href="/pkg/builtin/#string">string</a>)</pre>
<p>
Compact generates a copy of a slice with any empty strings removed.
The slice is not modified in place; the original will be unchanged.
</p>


<h2 id="Concat">func <a href="/src/target/string.go?s=1214:1246#L46">Concat</a>
	<a class="permalink" href="#Concat">&#xb6;</a>
</h2>
<pre>func Concat(slc []<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
<p>
Concat concatenates/joins all elements of a slice into a single string.
</p>


<h2 id="Contains">func <a href="/src/target/string.go?s=161:203#L1">Contains</a>
	<a class="permalink" href="#Contains">&#xb6;</a>
</h2>
<pre>func Contains(slc []<a href="/pkg/builtin/#string">string</a>, s <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
<p>
Return true if any element in a slice matches a string.
</p>


<h2 id="Copy">func <a href="/src/target/string.go?s=2475:2507#L102">Copy</a>
	<a class="permalink" href="#Copy">&#xb6;</a>
</h2>
<pre>func Copy(slc []<a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
<p>
Copy generates a full copy of a slice,
i.e., one which points to a different underlying array.
</p>


<h2 id="Cut">func <a href="/src/target/string.go?s=2139:2191#L87">Cut</a>
	<a class="permalink" href="#Cut">&#xb6;</a>
</h2>
<pre>func Cut(slc []<a href="/pkg/builtin/#string">string</a>, start <a href="/pkg/builtin/#int">int</a>, stop <a href="/pkg/builtin/#int">int</a>) []<a href="/pkg/builtin/#string">string</a></pre>
<p>
Cut gets a slice of a slice.
It gets the elements of a slice from index &#39;start&#39; to index &#39;stop&#39;.
&#39;start&#39; is inclusive (will include the element at that index).
&#39;stop&#39; is exclusive (will not include the element at that index).
If &#39;start&#39; is 0 and &#39;stop&#39; is -1, it generates a new copy of the slice.
</p>


<h2 id="Equals">func <a href="/src/target/string.go?s=705:751#L22">Equals</a>
	<a class="permalink" href="#Equals">&#xb6;</a>
</h2>
<pre>func Equals(slc1 []<a href="/pkg/builtin/#string">string</a>, slc2 []<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
<p>
Equals tests whether all the elements of two slices are equal.
These elements do NOT need to point to the same memory location,
only to hold an equal value.
</p>


<h2 id="Extract">func <a href="/src/target/string.go?s=3599:3667#L143">Extract</a>
	<a class="permalink" href="#Extract">&#xb6;</a>
</h2>
<pre>func Extract(excess []<a href="/pkg/builtin/#string">string</a>, wanted ...<a href="/pkg/builtin/#string">string</a>) (extracted []<a href="/pkg/builtin/#string">string</a>)</pre>
<p>
Extract gets elements from a slice.
It return a new slice of the elements pulled from the original.
The new slice contains only the &#39;wanted&#39; elements
which the original &#39;excess&#39; slice contains.
</p>


<h2 id="Filter">func <a href="/src/target/string.go?s=3103:3175#L123">Filter</a>
	<a class="permalink" href="#Filter">&#xb6;</a>
</h2>
<pre>func Filter(unfiltered []<a href="/pkg/builtin/#string">string</a>, unwanted ...<a href="/pkg/builtin/#string">string</a>) (filtered []<a href="/pkg/builtin/#string">string</a>)</pre>
<p>
Filter removes elements from a slice.
It returns a copy of a slice with unwanted strings removed.
The slice is not modified in place; the original will be unchanged.
</p>


<h2 id="IsEmpty">func <a href="/src/target/string.go?s=395:426#L9">IsEmpty</a>
	<a class="permalink" href="#IsEmpty">&#xb6;</a>
</h2>
<pre>func IsEmpty(slc []<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
<p>
IsEmpty tests whether a slice has 0 elements
or is full of empty strings.
</p>


<h2 id="IsSameArray">func <a href="/src/target/string.go?s=1059:1110#L41">IsSameArray</a>
	<a class="permalink" href="#IsSameArray">&#xb6;</a>
</h2>
<pre>func IsSameArray(slc1 []<a href="/pkg/builtin/#string">string</a>, slc2 []<a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#bool">bool</a></pre>
<p>
IsSameArray tests whether two slices point to the same array.
</p>


<h2 id="Join">func <a href="/src/target/string.go?s=1521:1562#L61">Join</a>
	<a class="permalink" href="#Join">&#xb6;</a>
</h2>
<pre>func Join(slc []<a href="/pkg/builtin/#string">string</a>, by <a href="/pkg/builtin/#string">string</a>) <a href="/pkg/builtin/#string">string</a></pre>
<p>
Join concatenates/joins all elements of a slice into a single string
and inserts a common string between each joined element.
</p>


<h2 id="Pop">func <a href="/src/target/string.go?s=4236:4277#L166">Pop</a>
	<a class="permalink" href="#Pop">&#xb6;</a>
</h2>
<pre>func Pop(slc []<a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, []<a href="/pkg/builtin/#string">string</a>)</pre>
<p>
Pop removes the last element of a slice.
It return the removed element along with the modified slice.
</p>


<h2 id="Push">func <a href="/src/target/string.go?s=4045:4093#L160">Push</a>
	<a class="permalink" href="#Push">&#xb6;</a>
</h2>
<pre>func Push(slc []<a href="/pkg/builtin/#string">string</a>, args ...<a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
<p>
Push adds elements to a slice and returns the modified slice.
It is a direct call to the built-in &#39;append()&#39; func.
It is meant to be clear, readable method for stack implementations.
</p>


<h2 id="Shift">func <a href="/src/target/string.go?s=4755:4798#L186">Shift</a>
	<a class="permalink" href="#Shift">&#xb6;</a>
</h2>
<pre>func Shift(slc []<a href="/pkg/builtin/#string">string</a>) (<a href="/pkg/builtin/#string">string</a>, []<a href="/pkg/builtin/#string">string</a>)</pre>
<p>
Shift removes the first element of a slice.
It returns the removed element along with the modified slice.
</p>


<h2 id="Unshift">func <a href="/src/target/string.go?s=4556:4601#L180">Unshift</a>
	<a class="permalink" href="#Unshift">&#xb6;</a>
</h2>
<pre>func Unshift(slc []<a href="/pkg/builtin/#string">string</a>, s <a href="/pkg/builtin/#string">string</a>) []<a href="/pkg/builtin/#string">string</a></pre>
<p>
Unshift inserts an element at the beginning of a slice,
and moves the rest of the elements up an index.
It does not overwrite the first element.
It returns the modified slice.
</p>
