<script lang="ts">
  import CodeEditor from "$components/ui/CodeEditor.svelte";
  import DropdownSelect from "$components/ui/DropdownSelect.svelte";
  import KeyValue from "$components/ui/KeyValue.svelte";
  import MultipartFormData from "$components/ui/MultipartFormData.svelte";
  import FileSelector from "$components/ui/FileSelector.svelte";
  import type {
    BodyType,
    KeyValueData,
    FormData,
    BodyTypeItem,
    LanguageItem,
    Language,
  } from "$types/custom";

  export let bodyTypes: BodyTypeItem[] = [
    {
      title: "None",
      value: "none",
    },
    {
      title: "form-data",
      value: "form-data",
    },
    {
      title: "x-www-formdata",
      value: "x-www-formdata",
    },
    {
      title: "Raw",
      value: "raw",
    },
    {
      title: "Binary",
      value: "binary",
    },
  ];
  export let activeBodyType: BodyType = "binary";
  export let languageList: LanguageItem[] = [
    {
      title: "Plain Text",
      value: "plain",
    },
    {
      title: "JSON",
      value: "json",
    },
    {
      title: "JavaScript",
      value: "javascript",
    },
    {
      title: "XML",
      value: "xml",
    },
  ];
  export let activeLanguage: Language = "javascript";
  export let rawBodyValue: string = "";
  export let xwwwformdataValue: KeyValueData[] = [];
  export let formdataValue: FormData[] = [];
  export let binaryValue: string[] = [];
</script>

<div class="my-3 grid grid-cols-[11rem,8rem] gap-x-3 items-center">
  <DropdownSelect items={bodyTypes} bind:value={activeBodyType} />
  {#if activeBodyType === "raw"}
    <DropdownSelect items={languageList} bind:value={activeLanguage} />
  {/if}
</div>
{#if activeBodyType === "raw"}
  <CodeEditor bind:value={rawBodyValue} language={activeLanguage} />
{:else if activeBodyType === "form-data"}
  <MultipartFormData bind:rows={formdataValue} />
{:else if activeBodyType === "x-www-formdata"}
  <KeyValue bind:rows={xwwwformdataValue} />
{:else if activeBodyType === "binary"}
  <div class="h-full flex justify-center items-center">
    <div class="border border-white/20 rounded w-96 flex flex-col h-10">
      <FileSelector bind:files={binaryValue} is_multiple={false} />
    </div>
  </div>
{/if}
