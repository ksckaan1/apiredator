<script lang="ts">
  import { slide } from "svelte/transition";
  import KeyValue from "$components/ui/KeyValue.svelte";
  import Tabs from "$components/ui/Tabs.svelte";
  import Button from "$components/ui/Button.svelte";
  import DropdownSelect from "$components/ui/DropdownSelect.svelte";
  import UrlInput from "$components/ui/URLInput.svelte";
  import RequestBody from "./parts/RequestBody.svelte";
  import type {
    BodyType,
    KeyValueData,
    Language,
    RequestTab,
    FormData,
  } from "$types/custom";
  import {
    GetCurrentRequest,
    SetCurrentRequest,
    StartCurrentRequest,
  } from "$lib/wailsjs/go/service/AppService";
  import { domain } from "$lib/wailsjs/go/models";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import Options from "./parts/Options.svelte";

  let urlValue: string = "https://jsonplaceholder.typicode.com/todos/1";
  let activeRequestMethod: string = "GET";

  let activeTab: RequestTab = "options";

  let activeTestType: string = "count";

  let numberOfRequestsValue: number = 100;
  let numberOfClientsValue: number = 1;

  let testDuration = "1m";
  let requestTimeout = "";

  let activeBodyType: BodyType = "none";
  let activeLanguage: Language = "json";
  let rawBodyValue: string = ``;
  let binaryValue: string[] = [];
  let formdataValue: FormData[] = [];
  let xwwwformdataValue: KeyValueData[] = [];

  let requestMethods = [
    {
      value: "GET",
      title: "GET",
      color: "text-green-400",
    },
    {
      value: "POST",
      title: "POST",
      color: "text-amber-400",
    },
    {
      value: "PUT",
      title: "PUT",
      color: "text-sky-400",
    },
    {
      value: "PATCH",
      title: "PATCH",
      color: "text-purple-400",
    },
    {
      value: "DELETE",
      title: "DELETE",
      color: "text-red-400",
    },
    {
      value: "HEAD",
      title: "HEAD",
      color: "text-pink-400",
    },
    {
      value: "TRACE",
      title: "TRACE",
      color: "text-teal-400",
    },
    {
      value: "OPTIONS",
      title: "OPTIONS",
      color: "text-emerald-400",
    },
  ];

  let headerRows: KeyValueData[] = [];

  onMount(async () => {
    try {
      let cr = await GetCurrentRequest();

      if (cr) {
        urlValue = cr.request.url;
        activeRequestMethod = cr.request.method;
        numberOfRequestsValue = cr.options.number_of_requests;
        numberOfClientsValue = cr.options.number_of_clients;
        activeTestType = cr.options.test_type;
        testDuration = cr.options.test_duration;
        requestTimeout = cr.options.request_timeout;
        activeBodyType = cr.request.body.type as BodyType;
        activeLanguage = cr.request.body.language as Language;
        rawBodyValue = cr.request.body.raw_value;
        binaryValue = cr.request.body.binary;
        formdataValue = cr.request.body.formdata as FormData[];
        xwwwformdataValue = cr.request.body.xwwwformdata;
      }
    } catch (err) {
      console.error(err);
    }
  });

  const onSendBtnClicked = async () => {
    let data: domain.Data = new domain.Data({
      request: {
        method: activeRequestMethod,
        url: urlValue,
        body: {
          type: activeBodyType,
          language: activeLanguage,
          raw_value: rawBodyValue,
          binary: binaryValue,
          formdata: formdataValue,
          xwwwformdata: xwwwformdataValue,
        },
        header: headerRows,
      },
      options: {
        number_of_requests: Number(numberOfRequestsValue),
        number_of_clients: Number(numberOfClientsValue),
        test_type: activeTestType,
        test_duration: testDuration,
        request_timeout: requestTimeout,
      },
    });
    try {
      await SetCurrentRequest(data);
      await StartCurrentRequest();
      goto("/status", {
        replaceState: true,
      });
    } catch (err) {
      console.error(err);
    }
  };
</script>

<div class="w-full p-5 mx-auto max-w-7xl flex-1 flex flex-col">
  <h1 class="mb-3 text-3xl">New Request</h1>
  <div class="w-full grid grid-cols-[8rem,1fr,7rem] gap-x-2">
    <DropdownSelect items={requestMethods} bind:value={activeRequestMethod} />
    <UrlInput bind:value={urlValue} label="URL"></UrlInput>
    <Button on:click={onSendBtnClicked}>Send</Button>
  </div>
  <hr class="my-3" />
  <Tabs bind:value={activeTab} />
  {#if activeTab === "options"}
    <div transition:slide>
      <Options
        bind:activeTestType
        bind:testDuration
        bind:requestTimeout
        bind:numberOfClientsValue
        bind:numberOfRequestsValue
      />
    </div>
  {:else if activeTab === "header"}
    <div transition:slide>
      <KeyValue bind:rows={headerRows} />
    </div>
  {:else if activeTab === "body"}
    <div transition:slide>
      <RequestBody
        bind:activeBodyType
        bind:activeLanguage
        bind:binaryValue
        bind:formdataValue
        bind:rawBodyValue
        bind:xwwwformdataValue
      />
    </div>
  {/if}
</div>
