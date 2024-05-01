<script lang="ts">
  import KeyValue from "$components/ui/KeyValue.svelte";
  import NumberInput from "$components/ui/NumberInput.svelte";
  import Section from "$components/ui/Section.svelte";
  import Tabs from "$components/ui/Tabs.svelte";
  import Button from "$components/ui/Button.svelte";
  import DropdownSelect from "$components/ui/DropdownSelect.svelte";
  import UrlInput from "$components/ui/URLInput.svelte";
  import RequestBody from "$components/parts/RequestBody.svelte";
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
  import SegmentedSelect from "$components/ui/SegmentedSelect.svelte";

  let urlValue: string = "https://jsonplaceholder.typicode.com/todos/1";
  let activeRequestMethod: string = "GET";

  let activeTab: RequestTab = "options";

  let activeTestType: string = "count";

  let numberOfRequestsValue: number = 1;
  let numberOfClientsValue: number = 1;

  let durationHoursValue: number = 0;
  let durationMinutesValue: number = 0;
  let durationSecondsValue: number = 0;

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

  let testTypes = [
    {
      title: "Count",
      value: "count",
    },
    {
      title: "Duration",
      value: "duration",
    },
  ];

  let headerRows: KeyValueData[] = [];

  onMount(async () => {
    // try {
    let cr = await GetCurrentRequest();

    if (cr) {
      urlValue = cr.request.url;
      activeRequestMethod = cr.request.method;
      numberOfRequestsValue = cr.options.number_of_requests;
      numberOfClientsValue = cr.options.number_of_clients;
      activeTestType = cr.options.test_type;
      durationHoursValue = cr.options.duration.hours;
      durationMinutesValue = cr.options.duration.minutes;
      durationSecondsValue = cr.options.duration.seconds;
      activeBodyType = cr.request.body.type as BodyType;
      activeLanguage = cr.request.body.language as Language;
      rawBodyValue = cr.request.body.raw_value;
      binaryValue = cr.request.body.binary;
      formdataValue = cr.request.body.formdata as FormData[];
      xwwwformdataValue = cr.request.body.xwwwformdata;
    }
    // }
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
        duration: {
          hours: Number(durationHoursValue),
          minutes: Number(durationMinutesValue),
          seconds: Number(durationSecondsValue),
        },
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
    <div class="w-full my-5">
      <Section title="Test type" subtitle="Which test type you are using?">
        <div class="w-min">
          <SegmentedSelect bind:activeItem={activeTestType} items={testTypes} />
        </div>
      </Section>
    </div>
    <div class="grid w-full grid-cols-2 mb-5 gap-x-5">
      {#if activeTestType === "count"}
        <Section
          title="Number of Requests"
          subtitle="How many requests per client?"
        >
          <NumberInput
            bind:value={numberOfRequestsValue}
            min={1}
            max={100000000000}
          />
        </Section>
      {:else if activeTestType === "duration"}
        <Section
          title="Duration"
          subtitle="How long will requests be sent in total?"
        >
          <div class="flex gap-x-3">
            <NumberInput
              label="Hours"
              bind:value={durationHoursValue}
              min={0}
              max={100000000000}
            />
            <NumberInput
              label="Minutes"
              bind:value={durationMinutesValue}
              min={0}
              max={100000000000}
            />
            <NumberInput
              label="Seconds"
              bind:value={durationSecondsValue}
              min={0}
              max={100000000000}
            />
          </div>
        </Section>
      {/if}
      <Section
        title="Number of Clients"
        subtitle="How many clients will make requests at the same time?"
      >
        <NumberInput
          bind:value={numberOfClientsValue}
          min={1}
          max={100000000000}
        />
      </Section>
    </div>
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
