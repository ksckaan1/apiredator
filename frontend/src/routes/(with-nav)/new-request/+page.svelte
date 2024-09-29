<script lang="ts">
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
  import { models } from "$lib/wailsjs/go/models";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import Options from "./parts/Options.svelte";

  let urlValue: string = $state("https://jsonplaceholder.typicode.com/todos/1");
  let activeRequestMethod: string = $state("GET");

  let activeTab: RequestTab = $state("options");

  let activeTestType: string = $state("count");

  let numberOfRequestsValue: number = $state(100);
  let numberOfClientsValue: number = $state(1);

  let testDuration = $state("1m");
  let requestTimeout = $state("");
  let keepAlive = $state(true);

  let activeBodyType: BodyType = $state("none");
  let activeLanguage: Language = $state("json");
  let rawBodyValue: string = $state(``);
  let binaryValue: string[] = $state([]);
  let formdataValue: FormData[] = $state([]);
  let xwwwformdataValue: KeyValueData[] = $state([]);

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

  let headerRows: KeyValueData[] = $state([]);

  $effect(() => {
    (async () => {
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
          keepAlive = cr.options.keep_alive;
        }
      } catch (err) {
        console.error(err);
      }
    })();
  });

  const onSendBtnClicked = async () => {
    let data: models.Data = new models.Data({
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
        keep_alive: keepAlive,
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
    <Button onclick={onSendBtnClicked} icon="la:flag-checkered">Start</Button>
  </div>
  <hr class="my-3" />
  <Tabs bind:value={activeTab} />
  {#if activeTab === "options"}
    <Options
      bind:activeTestType
      bind:testDuration
      bind:requestTimeout
      bind:numberOfClientsValue
      bind:numberOfRequestsValue
      bind:keepAlive
    />
  {:else if activeTab === "header"}
    <KeyValue bind:rows={headerRows} />
  {:else if activeTab === "body"}
    <RequestBody
      bind:activeBodyType
      bind:activeLanguage
      bind:binaryValue
      bind:formdataValue
      bind:rawBodyValue
      bind:xwwwformdataValue
    />
  {/if}
</div>
