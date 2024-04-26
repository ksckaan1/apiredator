<script>
  import CodeEditor from "$components/ui/CodeEditor.svelte";

  let rawBodyValue = `{
    "type": "reportviewerTest  ClickHere ClickHere ClickHere",
    "description": "Click Test",
    "name": "Report By Click",
    "captionBar": true
}`;
</script>

<div class="h-full">
  <CodeEditor bind:value={rawBodyValue} />
</div>
