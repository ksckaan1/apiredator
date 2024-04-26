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
    GetStats,
    SendRequest,
    WaitWork,
  } from "$lib/wailsjs/go/service/AppService";
  import { domain } from "$lib/wailsjs/go/models";
  import { goto } from "$app/navigation";

  // let urlValue: string = "{{host}}/api/v2/user/{{user_id}}";
  let urlValue: string = "https://jsonplaceholder.typicode.com/todos/1";
  let activeRequestMethod: string = "get";

  let activeTab: RequestTab = "body";

  let numberOfRequestsValue: number = 1;
  let numberOfClientsValue: number = 1;

  let isDurationActive: boolean = false;
  let durationHoursValue: number = 0;
  let durationMinutesValue: number = 0;
  let durationSecondsValue: number = 0;

  let activeBodyType: BodyType = "raw";
  let activeLanguage: Language = "json";
  let rawBodyValue: string = `{
    "username": "me@kaanksc.com",
    "password": "asdf1234"
}`;
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

  let headerRows: KeyValueData[] = [
    {
      key: "Content-Type",
      value: "application/json",
      is_active: true,
    },
    {
      key: "Authorization",
      value: "Bearer {{token}}",
      is_active: true,
    },
    {
      key: "user-id",
      value: "{{user_id}}",
      is_active: true,
    },
    {
      key: "unnecessary",
      value: "example",
      is_active: false,
    },
  ];

  let stats: any;

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
        duration: {
          is_duration_active: isDurationActive,
          hours: Number(durationHoursValue),
          minutes: Number(durationMinutesValue),
          seconds: Number(durationSecondsValue),
        },
      },
    });
    try {
      await SendRequest(data);
      goto("/status");
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
    <div class="grid w-full grid-cols-2 my-5 gap-x-5">
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
    <div class="grid w-full grid-cols-2 mb-5 gap-x-5">
      <Section
        bind:isActive={isDurationActive}
        title="Duration"
        subtitle="Should requests be made over a period of time?"
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

  <pre>
{JSON.stringify(stats, null, "\t")}
  </pre>
</div>
