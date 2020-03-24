<p align="center">
  <b>
    <span style="font-size:larger;">prefixtree-go</span>
  </b>
  <br />
   <a href="https://travis-ci.org/detailyang/prefixtree-go"><img src="https://travis-ci.org/detailyang/prefixtree-go.svg?branch=master" /></a>
   <a href="https://ci.appveyor.com/project/detailyang/prefixtree-go"><img src="https://ci.appveyor.com/api/projects/status/e5h8gx4rxosernmt?svg=true" /></a>
   <a href="https://godoc.org/github.com/detailyang/prefixtree-go">
      <img src="https://godoc.org/github.com/detailyang/prefixtree-go?status.svg"/>
   </a>
   <br />
   <b>prefixtree-go implements the radix tree which follows the longest match rule.</b>
</p>

```golang
trie.Add("/", 0)
trie.Add("/1", 1)
trie.Add("/2", 2)
trie.Add("/3", 3)
trie.Add("/1/0", 10)
trie.Add("/1/1", 11)
trie.Add("/2/0", 20)
trie.Add("/2/1", 21)
trie.Add("/3/0", 30)
trie.Add("/3/1", 31)

=> longest match

"/" => 0
"/1" =>  1
"/2" =>  2
"/3" =>  3
"/1/01" => 10
"/1/12" => 11
"/2/01" => 20
"/2/12" => 21
"/3/01" => 30
"/3/12" => 31
"/abcd" => 0
```
