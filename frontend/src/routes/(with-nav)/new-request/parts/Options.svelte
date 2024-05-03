<script lang="ts">
  import { fade } from "svelte/transition";
  import Section from "$components/ui/Section.svelte";
  import NumberInput from "$components/ui/NumberInput.svelte";
  import SegmentedSelect from "$components/ui/SegmentedSelect.svelte";

  export let activeTestType = "count";
  export let numberOfRequestsValue = 1;
  export let numberOfClientsValue = 1;
  export let durationHoursValue = 0;
  export let durationMinutesValue = 0;
  export let durationSecondsValue = 0;

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
</script>

<div class="grid grid-cols-1 md:grid-cols-2 w-full mt-5 gap-5">
  <Section title="Test type" subtitle="Which test type you are using?">
    <div class="w-min">
      <SegmentedSelect bind:activeItem={activeTestType} items={testTypes} />
    </div>
  </Section>
  <Section
    title="Number of Clients"
    subtitle="How many clients will make requests at the same time?"
  >
    <NumberInput bind:value={numberOfClientsValue} min={1} max={100000000000} />
  </Section>
  {#if activeTestType === "count"}
    <div
      in:fade={{
        delay: 200,
        duration: 200,
      }}
      out:fade={{
        duration: 200,
      }}
    >
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
    </div>
  {:else if activeTestType === "duration"}
    <div
      in:fade={{
        delay: 200,
        duration: 200,
      }}
      out:fade={{
        duration: 200,
      }}
    >
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
    </div>
  {/if}
</div>
