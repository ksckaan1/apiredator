<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import * as monaco from "monaco-editor";
  import jsonWorker from "monaco-editor/esm/vs/language/json/json.worker?worker";
  import htmlWorker from "monaco-editor/esm/vs/language/html/html.worker?worker";
  import tsWorker from "monaco-editor/esm/vs/language/typescript/ts.worker?worker";
  import editorWorker from "monaco-editor/esm/vs/editor/editor.worker?worker";
  import { ayuTheme } from "$lib/theme/ayu-dark";

  export let value = "";
  export let language = "";

  let editorElement: HTMLDivElement;
  let editor: monaco.editor.IStandaloneCodeEditor;
  let model: monaco.editor.ITextModel;

  const loadCode = (code: string, language: string) => {
    model = monaco.editor.createModel(code, language);
    model.onDidChangeContent((e) => {
      value = model.getValue();
    });
    editor.setModel(model);
  };

  onMount(async () => {
    self.MonacoEnvironment = {
      getWorker: function (_: any, label: string) {
        if (label === "json") return new jsonWorker();
        if (label === "xml") return new htmlWorker();
        if (label === "javascript") return new tsWorker();
        return new editorWorker();
      },
    };

    monaco.languages.typescript.typescriptDefaults.setEagerModelSync(true);

    monaco.editor.defineTheme("ayu-dark", ayuTheme);

    editor = monaco.editor.create(editorElement, {
      automaticLayout: true,
      theme: "ayu-dark",
      fontSize: 16,
    });

    loadCode(value, language);
  });

  $: {
    if (language && model) loadCode(model.getValue(), language);
  }

  onDestroy(() => {
    monaco?.editor.getModels().forEach((model) => model.dispose());
    editor?.dispose();
  });
</script>

<div
  class="h-96 border border-white/20 rounded overflow-hidden"
  bind:this={editorElement}
/>
